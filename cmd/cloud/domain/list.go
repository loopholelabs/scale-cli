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
	"github.com/loopholelabs/scale/go/client/domain"
	"github.com/spf13/cobra"
)

// ListCmd encapsulates the command for listing all domains.
func ListCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:      "list",
			Short:    "lists all your domains",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := domain.NewGetDomainParamsWithContext(ctx)
				res, err := client.Domain.GetDomain(params)
				if err != nil {
					return err
				}
				domains := res.GetPayload()
				domainRows := make([]domainTableRow, len(domains))
				for i, domain := range domains {
					attached := ""
					if domain.Deployment != nil {
						attached = fmt.Sprintf("%s.%s", domain.Deployment.Subdomain, domain.Deployment.RootDomain)
					}

					domainRows[i] = domainTableRow{
						Domain:   domain.Domain,
						State:    formattedState(domain.State),
						Attached: attached,
					}
				}
				if len(domains) == 0 {
					ch.Printer.Print("No domains found\n")
					return nil
				}
				ch.Printer.PrintResource(domainRows)

				return nil
			},
		}

		cmd.AddCommand(listCmd)
	}
}

func formattedState(state string) string {
	switch state {
	case "pending":
		return printer.BoldYellow("Pending")
	case "issuing":
		return printer.BoldBlue("Issuing")
	case "ready":
		return printer.BoldGreen("Ready")
	case "attached":
		return printer.BoldGreen("Attached")
	default:
		return "Unknown"
	}
}

type domainTableRow struct {
	Domain   string `header:"Domain"`
	State    string `header:"State"`
	Attached string `header:"Attached to"`
}
