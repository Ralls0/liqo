// Copyright 2019-2022 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	virtualkubeletv1alpha1 "github.com/liqotech/liqo/apis/virtualkubelet/v1alpha1"
	versioned "github.com/liqotech/liqo/pkg/client/clientset/versioned"
	internalinterfaces "github.com/liqotech/liqo/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/liqotech/liqo/pkg/client/listers/virtualkubelet/v1alpha1"
)

// ShadowPodInformer provides access to a shared informer and lister for
// ShadowPods.
type ShadowPodInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ShadowPodLister
}

type shadowPodInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewShadowPodInformer constructs a new informer for ShadowPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewShadowPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredShadowPodInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredShadowPodInformer constructs a new informer for ShadowPod type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredShadowPodInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtualkubeletV1alpha1().ShadowPods(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtualkubeletV1alpha1().ShadowPods(namespace).Watch(context.TODO(), options)
			},
		},
		&virtualkubeletv1alpha1.ShadowPod{},
		resyncPeriod,
		indexers,
	)
}

func (f *shadowPodInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredShadowPodInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *shadowPodInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&virtualkubeletv1alpha1.ShadowPod{}, f.defaultInformer)
}

func (f *shadowPodInformer) Lister() v1alpha1.ShadowPodLister {
	return v1alpha1.NewShadowPodLister(f.Informer().GetIndexer())
}
