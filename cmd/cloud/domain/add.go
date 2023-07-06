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
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/client/domain"
	"github.com/loopholelabs/scale/go/client/models"
	"github.com/spf13/cobra"
)

// DNSRecordTableRow is a row in the DNS record table
type DNSRecordTableRow struct {
	Type  string `header:"Type"`
	Name  string `header:"Name"`
	Value string `header:"Value"`
}

// AddCmd encapsulates the command for adding a domain
func AddCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		addCmd := &cobra.Command{
			Use:      "add [domain]",
			Args:     cobra.ExactArgs(1),
			Short:    "add a custom domain",
			Long:     "add a custom domain to your user or organization",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := domain.NewPostDomainParamsWithContext(ctx).WithRequest(&models.ModelsCreateDomainRequest{
					Domain: args[0],
				})

				var end = ch.Printer.PrintProgress("Starting domain verification...")
				res, err := client.Domain.PostDomain(params)
				if err != nil {
					end()
					return err
				}
				end()
				domainInfo := res.GetPayload()
				ch.Printer.Print("Please add the following records to your domain's DNS provider:\n\n")
				ch.Printer.PrintResource([]DNSRecordTableRow{
					{Type: "CNAME", Name: domainInfo.Domain, Value: domainInfo.Cname},
					{Type: "CNAME", Name: "_acme-challenge." + domainInfo.Domain, Value: domainInfo.TxtCname},
				})
				ch.Printer.Printf(
					"Scale Cloud will periodically check for the records to be added, you can check your domain's status with %s\n",
					printer.Bold("scale cloud domain verify "+domainInfo.Domain),
				)
				return nil
			},
		}

		cmd.AddCommand(addCmd)
	}
}
