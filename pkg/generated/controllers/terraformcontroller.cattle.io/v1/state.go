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
	"context"

	v1 "github.com/giantswarm/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	clientset "github.com/giantswarm/terraform-controller/pkg/generated/clientset/versioned/typed/terraformcontroller.cattle.io/v1"
	informers "github.com/giantswarm/terraform-controller/pkg/generated/informers/externalversions/terraformcontroller.cattle.io/v1"
	listers "github.com/giantswarm/terraform-controller/pkg/generated/listers/terraformcontroller.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type StateHandler func(string, *v1.State) (*v1.State, error)

type StateController interface {
	StateClient

	OnChange(ctx context.Context, name string, sync StateHandler)
	OnRemove(ctx context.Context, name string, sync StateHandler)
	Enqueue(namespace, name string)

	Cache() StateCache

	Informer() cache.SharedIndexInformer
	GroupVersionKind() schema.GroupVersionKind

	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
	AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler)
	Updater() generic.Updater
}

type StateClient interface {
	Create(*v1.State) (*v1.State, error)
	Update(*v1.State) (*v1.State, error)
	UpdateStatus(*v1.State) (*v1.State, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.State, error)
	List(namespace string, opts metav1.ListOptions) (*v1.StateList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.State, err error)
}

type StateCache interface {
	Get(namespace, name string) (*v1.State, error)
	List(namespace string, selector labels.Selector) ([]*v1.State, error)

	AddIndexer(indexName string, indexer StateIndexer)
	GetByIndex(indexName, key string) ([]*v1.State, error)
}

type StateIndexer func(obj *v1.State) ([]string, error)

type stateController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.StatesGetter
	informer          informers.StateInformer
	gvk               schema.GroupVersionKind
}

func NewStateController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.StatesGetter, informer informers.StateInformer) StateController {
	return &stateController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromStateHandlerToHandler(sync StateHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.State
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.State))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *stateController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.State))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateStateOnChange(updater generic.Updater, handler StateHandler) StateHandler {
	return func(key string, obj *v1.State) (*v1.State, error) {
		if obj == nil {
			return handler(key, nil)
		}

		copyObj := obj.DeepCopy()
		newObj, err := handler(key, copyObj)
		if newObj != nil {
			copyObj = newObj
		}
		if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
			newObj, err := updater(copyObj)
			if newObj != nil && err == nil {
				copyObj = newObj.(*v1.State)
			}
		}

		return copyObj, err
	}
}

func (c *stateController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *stateController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *stateController) OnChange(ctx context.Context, name string, sync StateHandler) {
	c.AddGenericHandler(ctx, name, FromStateHandlerToHandler(sync))
}

func (c *stateController) OnRemove(ctx context.Context, name string, sync StateHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromStateHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *stateController) Enqueue(namespace, name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), namespace, name)
}

func (c *stateController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *stateController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *stateController) Cache() StateCache {
	return &stateCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *stateController) Create(obj *v1.State) (*v1.State, error) {
	return c.clientGetter.States(obj.Namespace).Create(obj)
}

func (c *stateController) Update(obj *v1.State) (*v1.State, error) {
	return c.clientGetter.States(obj.Namespace).Update(obj)
}

func (c *stateController) UpdateStatus(obj *v1.State) (*v1.State, error) {
	return c.clientGetter.States(obj.Namespace).UpdateStatus(obj)
}

func (c *stateController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.States(namespace).Delete(name, options)
}

func (c *stateController) Get(namespace, name string, options metav1.GetOptions) (*v1.State, error) {
	return c.clientGetter.States(namespace).Get(name, options)
}

func (c *stateController) List(namespace string, opts metav1.ListOptions) (*v1.StateList, error) {
	return c.clientGetter.States(namespace).List(opts)
}

func (c *stateController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.States(namespace).Watch(opts)
}

func (c *stateController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.State, err error) {
	return c.clientGetter.States(namespace).Patch(name, pt, data, subresources...)
}

type stateCache struct {
	lister  listers.StateLister
	indexer cache.Indexer
}

func (c *stateCache) Get(namespace, name string) (*v1.State, error) {
	return c.lister.States(namespace).Get(name)
}

func (c *stateCache) List(namespace string, selector labels.Selector) ([]*v1.State, error) {
	return c.lister.States(namespace).List(selector)
}

func (c *stateCache) AddIndexer(indexName string, indexer StateIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.State))
		},
	}))
}

func (c *stateCache) GetByIndex(indexName, key string) (result []*v1.State, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.State))
	}
	return result, nil
}
