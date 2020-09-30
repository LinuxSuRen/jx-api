/*
Copyright 2020 The Jenkins X Authors.

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

package v1

import (
	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EnvironmentRoleBindingLister helps list EnvironmentRoleBindings.
// All objects returned here must be treated as read-only.
type EnvironmentRoleBindingLister interface {
	// List lists all EnvironmentRoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EnvironmentRoleBinding, err error)
	// EnvironmentRoleBindings returns an object that can list and get EnvironmentRoleBindings.
	EnvironmentRoleBindings(namespace string) EnvironmentRoleBindingNamespaceLister
	EnvironmentRoleBindingListerExpansion
}

// environmentRoleBindingLister implements the EnvironmentRoleBindingLister interface.
type environmentRoleBindingLister struct {
	indexer cache.Indexer
}

// NewEnvironmentRoleBindingLister returns a new EnvironmentRoleBindingLister.
func NewEnvironmentRoleBindingLister(indexer cache.Indexer) EnvironmentRoleBindingLister {
	return &environmentRoleBindingLister{indexer: indexer}
}

// List lists all EnvironmentRoleBindings in the indexer.
func (s *environmentRoleBindingLister) List(selector labels.Selector) (ret []*v1.EnvironmentRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EnvironmentRoleBinding))
	})
	return ret, err
}

// EnvironmentRoleBindings returns an object that can list and get EnvironmentRoleBindings.
func (s *environmentRoleBindingLister) EnvironmentRoleBindings(namespace string) EnvironmentRoleBindingNamespaceLister {
	return environmentRoleBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EnvironmentRoleBindingNamespaceLister helps list and get EnvironmentRoleBindings.
// All objects returned here must be treated as read-only.
type EnvironmentRoleBindingNamespaceLister interface {
	// List lists all EnvironmentRoleBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EnvironmentRoleBinding, err error)
	// Get retrieves the EnvironmentRoleBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EnvironmentRoleBinding, error)
	EnvironmentRoleBindingNamespaceListerExpansion
}

// environmentRoleBindingNamespaceLister implements the EnvironmentRoleBindingNamespaceLister
// interface.
type environmentRoleBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EnvironmentRoleBindings in the indexer for a given namespace.
func (s environmentRoleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1.EnvironmentRoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EnvironmentRoleBinding))
	})
	return ret, err
}

// Get retrieves the EnvironmentRoleBinding from the indexer for a given namespace and name.
func (s environmentRoleBindingNamespaceLister) Get(name string) (*v1.EnvironmentRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("environmentrolebinding"), name)
	}
	return obj.(*v1.EnvironmentRoleBinding), nil
}
