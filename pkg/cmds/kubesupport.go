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

package cmds

import (
	"flag"

	"github.com/appscode/go/term"
	"github.com/spf13/cobra"
	"go.bytebuilders.dev/resource-model/pkg/cmds/options"
	"go.bytebuilders.dev/resource-model/pkg/providers"
)

func NewCmdKubeSupport() *cobra.Command {
	opts := options.NewKubernetesData()
	cmd := &cobra.Command{
		Use:               "kubesupport",
		Short:             "Add kubernetes version support for a given or all cloud provider",
		Example:           "",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.ValidateFlags(cmd, args); err != nil {
				term.Fatalln(err)
			}
			err := providers.AddKubernetesSupport(opts)
			if err != nil {
				term.Fatalln(err)
			} else {
				term.Successln("Kubernetes support successfully added for ", opts.Provider)
			}
		},
	}
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	opts.AddFlags(cmd.Flags())
	return cmd
}
