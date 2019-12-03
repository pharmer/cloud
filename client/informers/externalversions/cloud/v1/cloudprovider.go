/*
Copyright The Pharmer Authors.

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

package v1

import (
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	cloudv1 "pharmer.dev/cloud/apis/cloud/v1"
	versioned "pharmer.dev/cloud/client/clientset/versioned"
	internalinterfaces "pharmer.dev/cloud/client/informers/externalversions/internalinterfaces"
	v1 "pharmer.dev/cloud/client/listers/cloud/v1"
)

// CloudProviderInformer provides access to a shared informer and lister for
// CloudProviders.
type CloudProviderInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CloudProviderLister
}

type cloudProviderInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCloudProviderInformer constructs a new informer for CloudProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudProviderInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCloudProviderInformer constructs a new informer for CloudProvider type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudProviderInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1().CloudProviders().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1().CloudProviders().Watch(options)
			},
		},
		&cloudv1.CloudProvider{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudProviderInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudProviderInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudProviderInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1.CloudProvider{}, f.defaultInformer)
}

func (f *cloudProviderInformer) Lister() v1.CloudProviderLister {
	return v1.NewCloudProviderLister(f.Informer().GetIndexer())
}
