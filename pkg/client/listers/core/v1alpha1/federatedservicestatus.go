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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/kubefed/pkg/apis/core/v1alpha1"
)

// FederatedServiceStatusLister helps list FederatedServiceStatuses.
type FederatedServiceStatusLister interface {
	// List lists all FederatedServiceStatuses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedServiceStatus, err error)
	// FederatedServiceStatuses returns an object that can list and get FederatedServiceStatuses.
	FederatedServiceStatuses(namespace string) FederatedServiceStatusNamespaceLister
	FederatedServiceStatusListerExpansion
}

// federatedServiceStatusLister implements the FederatedServiceStatusLister interface.
type federatedServiceStatusLister struct {
	indexer cache.Indexer
}

// NewFederatedServiceStatusLister returns a new FederatedServiceStatusLister.
func NewFederatedServiceStatusLister(indexer cache.Indexer) FederatedServiceStatusLister {
	return &federatedServiceStatusLister{indexer: indexer}
}

// List lists all FederatedServiceStatuses in the indexer.
func (s *federatedServiceStatusLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedServiceStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedServiceStatus))
	})
	return ret, err
}

// FederatedServiceStatuses returns an object that can list and get FederatedServiceStatuses.
func (s *federatedServiceStatusLister) FederatedServiceStatuses(namespace string) FederatedServiceStatusNamespaceLister {
	return federatedServiceStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedServiceStatusNamespaceLister helps list and get FederatedServiceStatuses.
type FederatedServiceStatusNamespaceLister interface {
	// List lists all FederatedServiceStatuses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedServiceStatus, err error)
	// Get retrieves the FederatedServiceStatus from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedServiceStatus, error)
	FederatedServiceStatusNamespaceListerExpansion
}

// federatedServiceStatusNamespaceLister implements the FederatedServiceStatusNamespaceLister
// interface.
type federatedServiceStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedServiceStatuses in the indexer for a given namespace.
func (s federatedServiceStatusNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedServiceStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedServiceStatus))
	})
	return ret, err
}

// Get retrieves the FederatedServiceStatus from the indexer for a given namespace and name.
func (s federatedServiceStatusNamespaceLister) Get(name string) (*v1alpha1.FederatedServiceStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedservicestatus"), name)
	}
	return obj.(*v1alpha1.FederatedServiceStatus), nil
}
