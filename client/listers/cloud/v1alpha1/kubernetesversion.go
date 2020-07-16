/*
Copyright 2020 AppsCode Inc.

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
	v1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubernetesVersionLister helps list KubernetesVersions.
type KubernetesVersionLister interface {
	// List lists all KubernetesVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KubernetesVersion, err error)
	// Get retrieves the KubernetesVersion from the index for a given name.
	Get(name string) (*v1alpha1.KubernetesVersion, error)
	KubernetesVersionListerExpansion
}

// kubernetesVersionLister implements the KubernetesVersionLister interface.
type kubernetesVersionLister struct {
	indexer cache.Indexer
}

// NewKubernetesVersionLister returns a new KubernetesVersionLister.
func NewKubernetesVersionLister(indexer cache.Indexer) KubernetesVersionLister {
	return &kubernetesVersionLister{indexer: indexer}
}

// List lists all KubernetesVersions in the indexer.
func (s *kubernetesVersionLister) List(selector labels.Selector) (ret []*v1alpha1.KubernetesVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubernetesVersion))
	})
	return ret, err
}

// Get retrieves the KubernetesVersion from the index for a given name.
func (s *kubernetesVersionLister) Get(name string) (*v1alpha1.KubernetesVersion, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubernetesversion"), name)
	}
	return obj.(*v1alpha1.KubernetesVersion), nil
}
