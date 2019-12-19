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
	"time"

	v1 "github.com/giantswarm/terraform-controller/pkg/apis/terraformcontroller.cattle.io/v1"
	scheme "github.com/giantswarm/terraform-controller/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExecutionsGetter has a method to return a ExecutionInterface.
// A group's client should implement this interface.
type ExecutionsGetter interface {
	Executions(namespace string) ExecutionInterface
}

// ExecutionInterface has methods to work with Execution resources.
type ExecutionInterface interface {
	Create(*v1.Execution) (*v1.Execution, error)
	Update(*v1.Execution) (*v1.Execution, error)
	UpdateStatus(*v1.Execution) (*v1.Execution, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Execution, error)
	List(opts metav1.ListOptions) (*v1.ExecutionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Execution, err error)
	ExecutionExpansion
}

// executions implements ExecutionInterface
type executions struct {
	client rest.Interface
	ns     string
}

// newExecutions returns a Executions
func newExecutions(c *TerraformcontrollerV1Client, namespace string) *executions {
	return &executions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the execution, and returns the corresponding execution object, and an error if there is any.
func (c *executions) Get(name string, options metav1.GetOptions) (result *v1.Execution, err error) {
	result = &v1.Execution{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("executions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Executions that match those selectors.
func (c *executions) List(opts metav1.ListOptions) (result *v1.ExecutionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ExecutionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("executions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested executions.
func (c *executions) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("executions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a execution and creates it.  Returns the server's representation of the execution, and an error, if there is any.
func (c *executions) Create(execution *v1.Execution) (result *v1.Execution, err error) {
	result = &v1.Execution{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("executions").
		Body(execution).
		Do().
		Into(result)
	return
}

// Update takes the representation of a execution and updates it. Returns the server's representation of the execution, and an error, if there is any.
func (c *executions) Update(execution *v1.Execution) (result *v1.Execution, err error) {
	result = &v1.Execution{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("executions").
		Name(execution.Name).
		Body(execution).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *executions) UpdateStatus(execution *v1.Execution) (result *v1.Execution, err error) {
	result = &v1.Execution{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("executions").
		Name(execution.Name).
		SubResource("status").
		Body(execution).
		Do().
		Into(result)
	return
}

// Delete takes name of the execution and deletes it. Returns an error if one occurs.
func (c *executions) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("executions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *executions) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("executions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched execution.
func (c *executions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Execution, err error) {
	result = &v1.Execution{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("executions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
