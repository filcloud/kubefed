package kubefedctl

import (
	"time"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apiextv1b1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	ctlutil "sigs.k8s.io/kubefed/pkg/controller/util"
	"sigs.k8s.io/kubefed/pkg/kubefedctl/util"
)

func JoinClusterForNamespacePhaseOne(clusterConfig *rest.Config, kubefedNamespace,
	joiningNamespace, hostClusterName, joiningClusterName, secretName string,
	scope apiextv1b1.ResourceScope, dryRun, errorOnExisting bool) (*corev1.Secret, error) {

	clusterClientset, err := util.ClusterClientset(clusterConfig)
	if err != nil {
		klog.V(2).Infof("Failed to get joining cluster clientset: %v", err)
		return nil, err
	}

	klog.V(2).Infof("Performing preflight checks.")
	err = performPreflightChecks(clusterClientset, joiningClusterName, hostClusterName, joiningNamespace, errorOnExisting)
	if err != nil {
		return nil, err
	}

	klog.V(2).Infof("Creating %s namespace in joining cluster", joiningNamespace)
	_, err = createKubeFedNamespace(clusterClientset, joiningNamespace,
		joiningClusterName, dryRun)
	if err != nil {
		klog.V(2).Infof("Error creating %s namespace in joining cluster: %v",
			joiningNamespace, err)
		return nil, err
	}
	klog.V(2).Infof("Created %s namespace in joining cluster", joiningNamespace)

	saName, err := createAuthorizedServiceAccount(clusterClientset,
		joiningNamespace, joiningClusterName, hostClusterName,
		scope, dryRun, errorOnExisting)
	if err != nil {
		return nil, err
	}

	secret, err := populateSecretInHostClusterPhaseOne(clusterClientset,
		saName, joiningNamespace, secretName, dryRun)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func populateSecretInHostClusterPhaseOne(clusterClientset kubeclient.Interface,
	saName, joiningNamespace, secretName string,
	dryRun bool) (*corev1.Secret, error) {

	if dryRun {
		dryRunSecret := &corev1.Secret{}
		dryRunSecret.Name = secretName
		return dryRunSecret, nil
	}

	// Get the secret from the joining cluster.
	var secret *corev1.Secret
	err := wait.PollImmediate(1*time.Second, serviceAccountSecretTimeout, func() (bool, error) {
		sa, err := clusterClientset.CoreV1().ServiceAccounts(joiningNamespace).Get(saName,
			metav1.GetOptions{})
		if err != nil {
			return false, nil
		}

		for _, objReference := range sa.Secrets {
			saSecretName := objReference.Name
			var err error
			secret, err = clusterClientset.CoreV1().Secrets(joiningNamespace).Get(saSecretName,
				metav1.GetOptions{})
			if err != nil {
				return false, nil
			}
			if secret.Type == corev1.SecretTypeServiceAccountToken {
				klog.V(2).Infof("Using secret named: %s", secret.Name)
				return true, nil
			}
		}
		return false, nil
	})

	if err != nil {
		klog.V(2).Infof("Could not get service account secret from joining cluster: %v", err)
		return nil, err
	}

	return secret, nil
}

func JoinClusterForNamespacePhaseTwo(hostConfig, clusterConfig *rest.Config, kubefedNamespace,
	hostClusterName, joiningClusterName, secretName string,
	dryRun, errorOnExisting bool, secretForCluster *corev1.Secret) (*fedv1b1.KubeFedCluster, error) {

	hostClientset, err := util.HostClientset(hostConfig)
	if err != nil {
		klog.V(2).Infof("Failed to get host cluster clientset: %v", err)
		return nil, err
	}

	client, err := genericclient.New(hostConfig)
	if err != nil {
		klog.V(2).Infof("Failed to get kubefed clientset: %v", err)
		return nil, err
	}

	secret, caBundle, err := populateSecretInHostClusterPhaseTwo(hostClientset,
		kubefedNamespace, joiningClusterName, secretName, dryRun, secretForCluster)
	if err != nil {
		klog.V(2).Infof("Error creating secret in host cluster: %s due to: %v", hostClusterName, err)
		return nil, err
	}

	var disabledTLSValidations []fedv1b1.TLSValidation
	if clusterConfig.TLSClientConfig.Insecure {
		disabledTLSValidations = append(disabledTLSValidations, fedv1b1.TLSAll)
	}

	kubefedCluster, err := createKubeFedCluster(client, joiningClusterName, clusterConfig.Host,
		secret.Name, kubefedNamespace, caBundle, disabledTLSValidations, dryRun, errorOnExisting)
	if err != nil {
		klog.V(2).Infof("Failed to create federated cluster resource: %v", err)
		return nil, err
	}

	klog.V(2).Info("Created federated cluster resource")
	return kubefedCluster, nil
}

func populateSecretInHostClusterPhaseTwo(hostClientset kubeclient.Interface,
	hostNamespace, joiningClusterName, secretName string,
	dryRun bool, secret *corev1.Secret) (*corev1.Secret, []byte, error) {

	klog.V(2).Infof("Creating cluster credentials secret in host cluster")

	if dryRun {
		dryRunSecret := &corev1.Secret{}
		dryRunSecret.Name = secretName
		return dryRunSecret, nil, nil
	}

	token, ok := secret.Data[ctlutil.TokenKey]
	if !ok {
		return nil, nil, errors.Errorf("Key %q not found in service account secret", ctlutil.TokenKey)
	}

	// Create a secret in the host cluster containing the token.
	v1Secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: hostNamespace,
		},
		Data: map[string][]byte{
			ctlutil.TokenKey: token,
		},
	}

	if secretName == "" {
		v1Secret.GenerateName = joiningClusterName + "-"
	} else {
		v1Secret.Name = secretName
	}

	v1SecretResult, err := hostClientset.CoreV1().Secrets(hostNamespace).Create(&v1Secret)
	if err != nil {
		klog.V(2).Infof("Could not create secret in host cluster: %v", err)
		return nil, nil, err
	}

	// caBundle is optional so no error is suggested if it is not
	// found in the secret.
	caBundle := secret.Data["ca.crt"]

	klog.V(2).Infof("Created secret in host cluster named: %s", v1SecretResult.Name)
	return v1SecretResult, caBundle, nil
}
