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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1 "pharmer.dev/cloud/client/clientset/versioned/typed/cloud/v1"
)

type FakeCloudV1 struct {
	*testing.Fake
}

func (c *FakeCloudV1) CloudProviders() v1.CloudProviderInterface {
	return &FakeCloudProviders{c}
}

func (c *FakeCloudV1) Credentials() v1.CredentialInterface {
	return &FakeCredentials{c}
}

func (c *FakeCloudV1) CredentialFormats() v1.CredentialFormatInterface {
	return &FakeCredentialFormats{c}
}

func (c *FakeCloudV1) KubernetesVersions() v1.KubernetesVersionInterface {
	return &FakeKubernetesVersions{c}
}

func (c *FakeCloudV1) MachineTypes() v1.MachineTypeInterface {
	return &FakeMachineTypes{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCloudV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}