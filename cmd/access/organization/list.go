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

package organization

import (
	"fmt"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/access"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/spf13/cobra"
)

// ListCmd encapsulates the commands for listing Organizations
func ListCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:   "list",
			Short: "list Organizations",
			Args:  cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				end := ch.Printer.PrintProgress("Retrieving Organizations...")
				res, err := client.Access.GetAccessOrganization(access.NewGetAccessOrganizationParamsWithContext(ctx))
				end()
				if err != nil {
					return err
				}

				analytics.Event("list-organization")

				if len(res.Payload) == 0 && ch.Printer.Format() == printer.Human {
					ch.Printer.Println("No Organizations have been created yet.")
					return nil
				}

				orgs := make([]organization, 0, len(res.Payload))
				for _, org := range res.Payload {
					orgs = append(orgs, organization{
						Created: org.CreatedAt,
						ID:      org.ID,
						Default: fmt.Sprintf("%t", org.Default),
					})
				}

				return ch.Printer.PrintResource(orgs)
			},
		}

		cmd.AddCommand(listCmd)
	}
}
