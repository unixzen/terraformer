// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	hetzner_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/hetzner"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdHetznerImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hetzner",
		Short: "Import current state to Terraform configuration from Hetzner",
		Long:  "Import current state to Terraform configuration from Hetzner",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newHetznerProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newHetznerProvider()))
	//	baseProviderFlags(cmd.PersistentFlags(), &options, "instance", "hetzner_instance=name1:name2:name3")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "hcloud_server")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	return cmd
}

func newHetznerProvider() terraformutils.ProviderGenerator {
	return &hetzner_terraforming.HetznerProvider{}
}
