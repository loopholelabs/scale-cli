/*
	Copyright 2023 Loophole Labs

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

package domain

import (
	"fmt"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/client/cloud"
	"github.com/loopholelabs/scale/go/client/models"
	"github.com/spf13/cobra"
)

// DetachCmd encapsulates the commands for deploying Functions
func DetachCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		detachCmd := &cobra.Command{
			Use:      "detach [domain]",
			Short:    "detach the domain from its Scale Cloud function",
			Long:     "detach the domain from its Scale Cloud function",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := cloud.NewPostCloudDetachParamsWithContext(ctx).WithRequest(&models.ModelsDetachDomainRequest{
					Domain: args[0],
				})

				var end = ch.Printer.PrintProgress("Detaching domain from function...")
				res, err := client.Cloud.PostCloudDetach(params)
				end()
				if err != nil {
					return err
				}
				domainInfo := res.GetPayload()
				domain := fmt.Sprintf("%s.%s", domainInfo.Subdomain, domainInfo.RootDomain)
				if len(domainInfo.CustomDomain) > 0 {
					domain = domainInfo.CustomDomain[len(domainInfo.CustomDomain)-1]
				}

				ch.Printer.Printf(
					"Detached domain. The function is now available at %s\n",
					printer.Bold(domain),
				)
				return nil
			},
		}

		cmd.AddCommand(detachCmd)
	}
}
