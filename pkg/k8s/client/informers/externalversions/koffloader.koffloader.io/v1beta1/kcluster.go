// Copyright 2024 Authors of koffloader-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	koffloaderkoffloaderiov1beta1 "github.com/koffloader-io/koffloader/pkg/k8s/apis/koffloader.koffloader.io/v1beta1"
	versioned "github.com/koffloader-io/koffloader/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/koffloader-io/koffloader/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/koffloader-io/koffloader/pkg/k8s/client/listers/koffloader.koffloader.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KClusterInformer provides access to a shared informer and lister for
// KClusters.
type KClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KClusterLister
}

type kClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKClusterInformer constructs a new informer for KCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKClusterInformer constructs a new informer for KCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KoffloaderV1beta1().KClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KoffloaderV1beta1().KClusters().Watch(context.TODO(), options)
			},
		},
		&koffloaderkoffloaderiov1beta1.KCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *kClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKClusterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&koffloaderkoffloaderiov1beta1.KCluster{}, f.defaultInformer)
}

func (f *kClusterInformer) Lister() v1beta1.KClusterLister {
	return v1beta1.NewKClusterLister(f.Informer().GetIndexer())
}
