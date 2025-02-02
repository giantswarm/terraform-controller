/*
Copyright 2019 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	time "time"

	terraformcontrollercattleiov1 "github.com/giantswarm/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	versioned "github.com/giantswarm/terraform-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/giantswarm/terraform-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/giantswarm/terraform-controller/pkg/generated/listers/terraformcontroller.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExecutionInformer provides access to a shared informer and lister for
// Executions.
type ExecutionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ExecutionLister
}

type executionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExecutionInformer constructs a new informer for Execution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExecutionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExecutionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExecutionInformer constructs a new informer for Execution type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExecutionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TerraformcontrollerV1().Executions(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TerraformcontrollerV1().Executions(namespace).Watch(options)
			},
		},
		&terraformcontrollercattleiov1.Execution{},
		resyncPeriod,
		indexers,
	)
}

func (f *executionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExecutionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *executionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&terraformcontrollercattleiov1.Execution{}, f.defaultInformer)
}

func (f *executionInformer) Lister() v1.ExecutionLister {
	return v1.NewExecutionLister(f.Informer().GetIndexer())
}
