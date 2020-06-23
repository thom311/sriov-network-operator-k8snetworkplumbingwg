/*
Copyright The Kubernetes Authors.

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
	"context"
	time "time"

	sriovnetworkv1 "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1"
	versioned "github.com/openshift/sriov-network-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift/sriov-network-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/sriov-network-operator/pkg/client/listers/sriovnetwork/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SriovNetworkNodeStateInformer provides access to a shared informer and lister for
// SriovNetworkNodeStates.
type SriovNetworkNodeStateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SriovNetworkNodeStateLister
}

type sriovNetworkNodeStateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSriovNetworkNodeStateInformer constructs a new informer for SriovNetworkNodeState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSriovNetworkNodeStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSriovNetworkNodeStateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSriovNetworkNodeStateInformer constructs a new informer for SriovNetworkNodeState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSriovNetworkNodeStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovNetworkNodeStates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovNetworkNodeStates(namespace).Watch(context.TODO(), options)
			},
		},
		&sriovnetworkv1.SriovNetworkNodeState{},
		resyncPeriod,
		indexers,
	)
}

func (f *sriovNetworkNodeStateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSriovNetworkNodeStateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sriovNetworkNodeStateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sriovnetworkv1.SriovNetworkNodeState{}, f.defaultInformer)
}

func (f *sriovNetworkNodeStateInformer) Lister() v1.SriovNetworkNodeStateLister {
	return v1.NewSriovNetworkNodeStateLister(f.Informer().GetIndexer())
}
