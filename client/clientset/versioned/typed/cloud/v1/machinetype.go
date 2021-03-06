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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "pharmer.dev/cloud/apis/cloud/v1"
	scheme "pharmer.dev/cloud/client/clientset/versioned/scheme"
)

// MachineTypesGetter has a method to return a MachineTypeInterface.
// A group's client should implement this interface.
type MachineTypesGetter interface {
	MachineTypes() MachineTypeInterface
}

// MachineTypeInterface has methods to work with MachineType resources.
type MachineTypeInterface interface {
	Create(ctx context.Context, machineType *v1.MachineType, opts metav1.CreateOptions) (*v1.MachineType, error)
	Update(ctx context.Context, machineType *v1.MachineType, opts metav1.UpdateOptions) (*v1.MachineType, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MachineType, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MachineTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineType, err error)
	MachineTypeExpansion
}

// machineTypes implements MachineTypeInterface
type machineTypes struct {
	client rest.Interface
}

// newMachineTypes returns a MachineTypes
func newMachineTypes(c *CloudV1Client) *machineTypes {
	return &machineTypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineType, and returns the corresponding machineType object, and an error if there is any.
func (c *machineTypes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MachineType, err error) {
	result = &v1.MachineType{}
	err = c.client.Get().
		Resource("machinetypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineTypes that match those selectors.
func (c *machineTypes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MachineTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MachineTypeList{}
	err = c.client.Get().
		Resource("machinetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineTypes.
func (c *machineTypes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("machinetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineType and creates it.  Returns the server's representation of the machineType, and an error, if there is any.
func (c *machineTypes) Create(ctx context.Context, machineType *v1.MachineType, opts metav1.CreateOptions) (result *v1.MachineType, err error) {
	result = &v1.MachineType{}
	err = c.client.Post().
		Resource("machinetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineType).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineType and updates it. Returns the server's representation of the machineType, and an error, if there is any.
func (c *machineTypes) Update(ctx context.Context, machineType *v1.MachineType, opts metav1.UpdateOptions) (result *v1.MachineType, err error) {
	result = &v1.MachineType{}
	err = c.client.Put().
		Resource("machinetypes").
		Name(machineType.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineType).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineType and deletes it. Returns an error if one occurs.
func (c *machineTypes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machinetypes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineTypes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("machinetypes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineType.
func (c *machineTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MachineType, err error) {
	result = &v1.MachineType{}
	err = c.client.Patch(pt).
		Resource("machinetypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
