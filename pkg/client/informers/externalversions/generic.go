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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/kubefed/pkg/apis/core/v1alpha1"
	v1beta1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	multiclusterdnsv1alpha1 "sigs.k8s.io/kubefed/pkg/apis/multiclusterdns/v1alpha1"
	schedulingv1alpha1 "sigs.k8s.io/kubefed/pkg/apis/scheduling/v1alpha1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=core, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusterpropagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterPropagatedVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedservicestatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedServiceStatuses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("propagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().PropagatedVersions().Informer()}, nil

		// Group=core, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("federatedtypeconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().FederatedTypeConfigs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("kubefedclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().KubeFedClusters().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("kubefedconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1beta1().KubeFedConfigs().Informer()}, nil

		// Group=multiclusterdns, Version=v1alpha1
	case multiclusterdnsv1alpha1.SchemeGroupVersion.WithResource("dnsendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().DNSEndpoints().Informer()}, nil
	case multiclusterdnsv1alpha1.SchemeGroupVersion.WithResource("domains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().Domains().Informer()}, nil
	case multiclusterdnsv1alpha1.SchemeGroupVersion.WithResource("ingressdnsrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().IngressDNSRecords().Informer()}, nil
	case multiclusterdnsv1alpha1.SchemeGroupVersion.WithResource("servicednsrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().ServiceDNSRecords().Informer()}, nil

		// Group=scheduling, Version=v1alpha1
	case schedulingv1alpha1.SchemeGroupVersion.WithResource("replicaschedulingpreferences"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().ReplicaSchedulingPreferences().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
