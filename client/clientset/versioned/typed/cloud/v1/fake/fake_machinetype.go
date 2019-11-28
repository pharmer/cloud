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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	testing "k8s.io/client-go/testing"
	cloudv1 "pharmer.dev/cloud/apis/cloud/v1"
)

// FakeMachineTypes implements MachineTypeInterface
type FakeMachineTypes struct {
	Fake *FakeCloudV1
}

var machinetypesResource = schema.GroupVersionResource{Group: "cloud.pharmer.io", Version: "v1", Resource: "machinetypes"}

var machinetypesKind = schema.GroupVersionKind{Group: "cloud.pharmer.io", Version: "v1", Kind: "MachineType"}

// Get takes name of the machineType, and returns the corresponding machineType object, and an error if there is any.
func (c *FakeMachineTypes) Get(name string, options v1.GetOptions) (result *cloudv1.MachineType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machinetypesResource, name), &cloudv1.MachineType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.MachineType), err
}

// List takes label and field selectors, and returns the list of MachineTypes that match those selectors.
func (c *FakeMachineTypes) List(opts v1.ListOptions) (result *cloudv1.MachineTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machinetypesResource, machinetypesKind, opts), &cloudv1.MachineTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudv1.MachineTypeList{ListMeta: obj.(*cloudv1.MachineTypeList).ListMeta}
	for _, item := range obj.(*cloudv1.MachineTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a machineType and creates it.  Returns the server's representation of the machineType, and an error, if there is any.
func (c *FakeMachineTypes) Create(machineType *cloudv1.MachineType) (result *cloudv1.MachineType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machinetypesResource, machineType), &cloudv1.MachineType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.MachineType), err
}

// Update takes the representation of a machineType and updates it. Returns the server's representation of the machineType, and an error, if there is any.
func (c *FakeMachineTypes) Update(machineType *cloudv1.MachineType) (result *cloudv1.MachineType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machinetypesResource, machineType), &cloudv1.MachineType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.MachineType), err
}

// Delete takes name of the machineType and deletes it. Returns an error if one occurs.
func (c *FakeMachineTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(machinetypesResource, name), &cloudv1.MachineType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(machinetypesResource, listOptions)

	_, err := c.Fake.Invokes(action, &cloudv1.MachineTypeList{})
	return err
}

// Patch applies the patch and returns the patched machineType.
func (c *FakeMachineTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.MachineType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machinetypesResource, name, pt, data, subresources...), &cloudv1.MachineType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.MachineType), err
}