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

package vultr

import (
	"fmt"
	"testing"

	"pharmer.dev/cloud/pkg/credential"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func TestRegion(t *testing.T) {
	client, err := NewClient(getToken())
	if err != nil {
		t.Error(err)
	}
	regions, err := client.ListRegions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(regions)
}

func TestInstance(t *testing.T) {
	client, err := NewClient(getToken())
	if err != nil {
		t.Error(err)
	}
	instances, err := client.ListMachineTypes()
	if err != nil {
		t.Error(err)
	}
	for _, i := range instances {
		fmt.Println(i.Spec.SKU)
	}
	fmt.Println("total:", len(instances))
}

func getToken() credential.Vultr {
	var v credential.Vultr
	utilruntime.Must(v.LoadFromJSON("/home/ac/Downloads/cred/vultr.json"))
	return v
}
