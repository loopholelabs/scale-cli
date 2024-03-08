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
	"github.com/loopholelabs/scale-cli/client/models"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/spf13/cobra"
)

// CreateCmd encapsulates the commands for creating Organizations
func CreateCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		createCmd := &cobra.Command{
			Use:   "create <name>",
			Args:  cobra.ExactArgs(1),
			Short: "Create an Organization with the given name",
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()

				name := args[0]

				end := ch.Printer.PrintProgress(fmt.Sprintf("Creating Organization %s...", name))
				req := &models.ModelsCreateOrganizationRequest{
					Name: name,
				}
				res, err := client.Access.PostAccessOrganization(access.NewPostAccessOrganizationParamsWithContext(ctx).WithRequest(req))
				end()
				if err != nil {
					return err
				}

				analytics.Event("create-organization")

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Created Organization '%s'\n", printer.Bold(res.Payload.ID))
					return nil
				}

				return ch.Printer.PrintResource(organization{
					Created: res.GetPayload().CreatedAt,
					ID:      res.GetPayload().ID,
					Default: fmt.Sprintf("%t", res.GetPayload().Default),
				})
			},
		}

		cmd.AddCommand(createCmd)
	}
}
