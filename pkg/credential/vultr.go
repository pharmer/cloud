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

package credential

import (
	"github.com/spf13/pflag"
	"go.bytebuilders.dev/resource-model/apis"
	v1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Vultr struct {
	CommonSpec

	token string
}

func (c Vultr) Token() string { return get(c.Data, VultrAPIToken, c.token) }

func (c *Vultr) LoadFromEnv() {
	c.CommonSpec.LoadFromEnv(c.Format())
}

func (c Vultr) IsValid() (bool, error) {
	return c.CommonSpec.IsValid(c.Format())
}

func (c *Vultr) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.token, apis.Vultr+"."+VultrAPIToken, c.token, "Vultr token")
}

func (_ Vultr) RequiredFlags() []string {
	return []string{apis.Vultr + "." + VultrAPIToken}
}

func (_ Vultr) Format() v1.CredentialFormat {
	return v1.CredentialFormat{
		ObjectMeta: metav1.ObjectMeta{
			Name: apis.Vultr,
			Labels: map[string]string{
				apis.KeyCloudProvider: apis.Vultr,
			},
			Annotations: map[string]string{
				apis.KeyClusterCredential: "",
				apis.KeyDNSCredential:     "",
			},
		},
		Spec: v1.CredentialFormatSpec{
			Provider:      apis.Vultr,
			DisplayFormat: "field",
			Fields: []v1.CredentialField{
				{
					Envconfig: "VULTR_TOKEN",
					Form:      "vultr_token",
					JSON:      VultrAPIToken,
					Label:     "Personal Access Token",
					Input:     "password",
				},
			},
		},
	}
}
