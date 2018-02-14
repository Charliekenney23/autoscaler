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

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
)

// ClusterRoleLister helps list ClusterRoles.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	List(selector labels.Selector) (ret []*rbac.ClusterRole, err error)
	// Get retrieves the ClusterRole from the index for a given name.
	Get(name string) (*rbac.ClusterRole, error)
	ClusterRoleListerExpansion
}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	indexer cache.Indexer
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(indexer cache.Indexer) ClusterRoleLister {
	return &clusterRoleLister{indexer: indexer}
}

// List lists all ClusterRoles in the indexer.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*rbac.ClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbac.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the index for a given name.
func (s *clusterRoleLister) Get(name string) (*rbac.ClusterRole, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbac.Resource("clusterrole"), name)
	}
	return obj.(*rbac.ClusterRole), nil
}