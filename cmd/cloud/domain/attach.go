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

// AttachCmd encapsulates the commands for deploying Functions
func AttachCmd(hidden bool) command.SetupCommand[*config.Config] {
	var force bool

	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		attachCmd := &cobra.Command{
			Use:      "attach [domain] [function ID]",
			Args:     cobra.ExactArgs(2),
			Short:    "list your deployed Scale functions",
			Long:     "list your deployed Scale function",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := cloud.NewPostCloudAttachParamsWithContext(ctx).WithRequest(&models.ModelsAttachDomainRequest{
					Domain:   args[0],
					Function: args[1],
					Force:    force,
				})

				var end = ch.Printer.PrintProgress("Attaching domain to function...")
				res, err := client.Cloud.PostCloudAttach(params)
				end()
				if err != nil {
					return err
				}
				domainInfo := res.GetPayload()
				customDomain := domainInfo.CustomDomain[len(domainInfo.CustomDomain)-1]
				defaultDomain := fmt.Sprintf("%s.%s", domainInfo.Subdomain, domainInfo.RootDomain)
				ch.Printer.Printf("Attached domain %s to %s\n", printer.Bold(customDomain), printer.Bold(defaultDomain))
				return nil
			},
		}

		attachCmd.Flags().BoolVarP(&hidden, "force", "f", false, "forcibly detach the domain from its current function if it's already attached")

		cmd.AddCommand(attachCmd)
	}
}
