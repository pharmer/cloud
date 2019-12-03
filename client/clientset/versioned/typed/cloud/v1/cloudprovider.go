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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "pharmer.dev/cloud/apis/cloud/v1"
	scheme "pharmer.dev/cloud/client/clientset/versioned/scheme"
)

// CloudProvidersGetter has a method to return a CloudProviderInterface.
// A group's client should implement this interface.
type CloudProvidersGetter interface {
	CloudProviders() CloudProviderInterface
}

// CloudProviderInterface has methods to work with CloudProvider resources.
type CloudProviderInterface interface {
	Create(*v1.CloudProvider) (*v1.CloudProvider, error)
	Update(*v1.CloudProvider) (*v1.CloudProvider, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.CloudProvider, error)
	List(opts metav1.ListOptions) (*v1.CloudProviderList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CloudProvider, err error)
	CloudProviderExpansion
}

// cloudProviders implements CloudProviderInterface
type cloudProviders struct {
	client rest.Interface
}

// newCloudProviders returns a CloudProviders
func newCloudProviders(c *CloudV1Client) *cloudProviders {
	return &cloudProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the cloudProvider, and returns the corresponding cloudProvider object, and an error if there is any.
func (c *cloudProviders) Get(name string, options metav1.GetOptions) (result *v1.CloudProvider, err error) {
	result = &v1.CloudProvider{}
	err = c.client.Get().
		Resource("cloudproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudProviders that match those selectors.
func (c *cloudProviders) List(opts metav1.ListOptions) (result *v1.CloudProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CloudProviderList{}
	err = c.client.Get().
		Resource("cloudproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudProviders.
func (c *cloudProviders) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cloudproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudProvider and creates it.  Returns the server's representation of the cloudProvider, and an error, if there is any.
func (c *cloudProviders) Create(cloudProvider *v1.CloudProvider) (result *v1.CloudProvider, err error) {
	result = &v1.CloudProvider{}
	err = c.client.Post().
		Resource("cloudproviders").
		Body(cloudProvider).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudProvider and updates it. Returns the server's representation of the cloudProvider, and an error, if there is any.
func (c *cloudProviders) Update(cloudProvider *v1.CloudProvider) (result *v1.CloudProvider, err error) {
	result = &v1.CloudProvider{}
	err = c.client.Put().
		Resource("cloudproviders").
		Name(cloudProvider.Name).
		Body(cloudProvider).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudProvider and deletes it. Returns an error if one occurs.
func (c *cloudProviders) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cloudproviders").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudProviders) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cloudproviders").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudProvider.
func (c *cloudProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.CloudProvider, err error) {
	result = &v1.CloudProvider{}
	err = c.client.Patch(pt).
		Resource("cloudproviders").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
