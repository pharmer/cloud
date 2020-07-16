/*
Copyright 2020 AppsCode Inc.

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
	v1alpha1 "go.bytebuilders.dev/resource-model/client/clientset/versioned/typed/cloud/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCloudV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCloudV1alpha1) CloudProviders() v1alpha1.CloudProviderInterface {
	return &FakeCloudProviders{c}
}

func (c *FakeCloudV1alpha1) Credentials() v1alpha1.CredentialInterface {
	return &FakeCredentials{c}
}

func (c *FakeCloudV1alpha1) CredentialFormats() v1alpha1.CredentialFormatInterface {
	return &FakeCredentialFormats{c}
}

func (c *FakeCloudV1alpha1) KubernetesVersions() v1alpha1.KubernetesVersionInterface {
	return &FakeKubernetesVersions{c}
}

func (c *FakeCloudV1alpha1) MachineTypes() v1alpha1.MachineTypeInterface {
	return &FakeMachineTypes{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCloudV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
