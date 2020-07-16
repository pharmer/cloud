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
	"go.bytebuilders.dev/resource-model/apis"
	v1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DigitalOcean struct {
	CommonSpec

	token string
}

func (c DigitalOcean) Token() string { return get(c.Data, DigitalOceanToken, c.token) }

func (c *DigitalOcean) LoadFromEnv() {
	c.CommonSpec.LoadFromEnv(c.Format())
}

func (c DigitalOcean) IsValid() (bool, error) {
	return c.CommonSpec.IsValid(c.Format())
}

func (c *DigitalOcean) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.token, apis.DigitalOcean+"."+DigitalOceanToken, c.token, "DigitalOcean token")
}

func (c DigitalOcean) RequiredFlags() []string {
	return []string{apis.DigitalOcean + "." + DigitalOceanToken}
}

func (_ DigitalOcean) Format() v1.CredentialFormat {
	return v1.CredentialFormat{
		ObjectMeta: metav1.ObjectMeta{
			Name: apis.DigitalOcean,
			Labels: map[string]string{
				apis.KeyCloudProvider: apis.DigitalOcean,
			},
			Annotations: map[string]string{
				apis.KeyClusterCredential: "",
				apis.KeyDNSCredential:     "",
			},
		},
		Spec: v1.CredentialFormatSpec{
			Provider:      apis.DigitalOcean,
			DisplayFormat: "field",
			Fields: []v1.CredentialField{
				{

					Envconfig: "DIGITALOCEAN_TOKEN",
					Form:      "digitalocean_token",
					JSON:      DigitalOceanToken,
					Label:     "Personal Access Token",
					Input:     "password",
				},
			},
		},
	}
}
