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

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	shipper_v1 "github.com/bookingcom/shipper/pkg/apis/shipper/v1"
	versioned "github.com/bookingcom/shipper/pkg/client/clientset/versioned"
	internalinterfaces "github.com/bookingcom/shipper/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/bookingcom/shipper/pkg/client/listers/shipper/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StrategyInformer provides access to a shared informer and lister for
// Strategies.
type StrategyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.StrategyLister
}

type strategyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStrategyInformer constructs a new informer for Strategy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStrategyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStrategyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStrategyInformer constructs a new informer for Strategy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStrategyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipperV1().Strategies(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipperV1().Strategies(namespace).Watch(options)
			},
		},
		&shipper_v1.Strategy{},
		resyncPeriod,
		indexers,
	)
}

func (f *strategyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStrategyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *strategyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&shipper_v1.Strategy{}, f.defaultInformer)
}

func (f *strategyInformer) Lister() v1.StrategyLister {
	return v1.NewStrategyLister(f.Informer().GetIndexer())
}