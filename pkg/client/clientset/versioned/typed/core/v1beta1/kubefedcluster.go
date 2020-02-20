/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	scheme "sigs.k8s.io/kubefed/pkg/client/clientset/versioned/scheme"
)

// KubeFedClustersGetter has a method to return a KubeFedClusterInterface.
// A group's client should implement this interface.
type KubeFedClustersGetter interface {
	KubeFedClusters(namespace string) KubeFedClusterInterface
}

// KubeFedClusterInterface has methods to work with KubeFedCluster resources.
type KubeFedClusterInterface interface {
	Create(*v1beta1.KubeFedCluster) (*v1beta1.KubeFedCluster, error)
	Update(*v1beta1.KubeFedCluster) (*v1beta1.KubeFedCluster, error)
	UpdateStatus(*v1beta1.KubeFedCluster) (*v1beta1.KubeFedCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.KubeFedCluster, error)
	List(opts v1.ListOptions) (*v1beta1.KubeFedClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KubeFedCluster, err error)
	KubeFedClusterExpansion
}

// kubeFedClusters implements KubeFedClusterInterface
type kubeFedClusters struct {
	client rest.Interface
	ns     string
}

// newKubeFedClusters returns a KubeFedClusters
func newKubeFedClusters(c *CoreV1beta1Client, namespace string) *kubeFedClusters {
	return &kubeFedClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubeFedCluster, and returns the corresponding kubeFedCluster object, and an error if there is any.
func (c *kubeFedClusters) Get(name string, options v1.GetOptions) (result *v1beta1.KubeFedCluster, err error) {
	result = &v1beta1.KubeFedCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubefedclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubeFedClusters that match those selectors.
func (c *kubeFedClusters) List(opts v1.ListOptions) (result *v1beta1.KubeFedClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KubeFedClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubefedclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubeFedClusters.
func (c *kubeFedClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubefedclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kubeFedCluster and creates it.  Returns the server's representation of the kubeFedCluster, and an error, if there is any.
func (c *kubeFedClusters) Create(kubeFedCluster *v1beta1.KubeFedCluster) (result *v1beta1.KubeFedCluster, err error) {
	result = &v1beta1.KubeFedCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubefedclusters").
		Body(kubeFedCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubeFedCluster and updates it. Returns the server's representation of the kubeFedCluster, and an error, if there is any.
func (c *kubeFedClusters) Update(kubeFedCluster *v1beta1.KubeFedCluster) (result *v1beta1.KubeFedCluster, err error) {
	result = &v1beta1.KubeFedCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubefedclusters").
		Name(kubeFedCluster.Name).
		Body(kubeFedCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kubeFedClusters) UpdateStatus(kubeFedCluster *v1beta1.KubeFedCluster) (result *v1beta1.KubeFedCluster, err error) {
	result = &v1beta1.KubeFedCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubefedclusters").
		Name(kubeFedCluster.Name).
		SubResource("status").
		Body(kubeFedCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubeFedCluster and deletes it. Returns an error if one occurs.
func (c *kubeFedClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubefedclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubeFedClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubefedclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubeFedCluster.
func (c *kubeFedClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KubeFedCluster, err error) {
	result = &v1beta1.KubeFedCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubefedclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
