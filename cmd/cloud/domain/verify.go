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

// VerifyCmd encapsulates the command for showing the verification status of a domain
func VerifyCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		verifyCmd := &cobra.Command{
			Use:      "verify [domain]",
			Args:     cobra.ExactArgs(1),
			Short:    "check the status of a custom domain",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := domain.NewGetDomainNameParamsWithContext(ctx).WithName(args[0])
				res, err := client.Domain.GetDomainName(params)
				if err != nil {
					return err
				}
				domain := res.GetPayload()

				switch domain.State {
				case "pending":
					ch.Printer.Printf(
						"Your domain is %s. Scale Cloud is periodically checking your domain for these required name records:\n\n",
						printer.BoldYellow("pending verification"),
					)
					ch.Printer.PrintResource([]DNSRecordTableRow{
						{Type: "CNAME", Name: domain.Domain, Value: domain.Cname},
						{Type: "CNAME", Name: "_acme-challenge." + domain.Domain, Value: domain.TxtCname},
					})
					break

				case "issuing":
					ch.Printer.Printf(
						"Your domain is being %s and should soon be ready for deployment.\n",
						printer.BoldBlue("issued a certificate"),
					)
					break

				case "ready":
					ch.Printer.Printf(
						"Your domain is %s. You may now attach a scale function to this domain using %s\n",
						printer.BoldGreen("ready"),
						printer.Bold("scale cloud domain attach"),
					)
					break

				case "attached":
					ch.Printer.Printf(
						"Your domain is %s and is attached to %s.\n",
						printer.BoldGreen("ready"),
						fmt.Sprintf("%s.%s", domain.Deployment.Subdomain, domain.Deployment.RootDomain),
					)
					break

				}
				return nil
			},
		}

		cmd.AddCommand(verifyCmd)
	}
}
