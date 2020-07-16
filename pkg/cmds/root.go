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

	"github.com/appscode/go/flags"
	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"go.bytebuilders.dev/resource-model/apis"
	"k8s.io/client-go/kubernetes/scheme"
	"kmodules.xyz/client-go/logs"
	"kmodules.xyz/client-go/tools/cli"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "pharmer-tools",
		Short:             `Pharmer tools by Appscode`,
		DisableAutoGenTag: true,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			flags.DumpAll(c.Flags())
			cli.SendAnalytics(c, v.Version.Version)
			return apis.AddToScheme(scheme.Scheme)
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	logs.ParseFlags()
	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")

	rootCmd.AddCommand(NewCmdGenData())
	rootCmd.AddCommand(NewCmdKubeSupport())
	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}
