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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	cloudv1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	versioned "go.bytebuilders.dev/resource-model/client/clientset/versioned"
	internalinterfaces "go.bytebuilders.dev/resource-model/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.bytebuilders.dev/resource-model/client/listers/cloud/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineTypeInformer provides access to a shared informer and lister for
// MachineTypes.
type MachineTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineTypeLister
}

type machineTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMachineTypeInformer constructs a new informer for MachineType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineTypeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineTypeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMachineTypeInformer constructs a new informer for MachineType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineTypeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().MachineTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().MachineTypes().Watch(context.TODO(), options)
			},
		},
		&cloudv1alpha1.MachineType{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineTypeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineTypeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1alpha1.MachineType{}, f.defaultInformer)
}

func (f *machineTypeInformer) Lister() v1alpha1.MachineTypeLister {
	return v1alpha1.NewMachineTypeLister(f.Informer().GetIndexer())
}
