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

// DeleteCmd encapsulates the commands for deleting Organizations
func DeleteCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		deleteCmd := &cobra.Command{
			Use:   "delete <name>",
			Args:  cobra.ExactArgs(1),
			Short: "delete an Organization with the given name",
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				name := args[0]

				end := ch.Printer.PrintProgress(fmt.Sprintf("Deleting Organization %s...", name))
				_, err := client.Access.DeleteAccessOrganizationName(access.NewDeleteAccessOrganizationNameParamsWithContext(ctx).WithName(name))
				end()
				if err != nil {
					return err
				}

				analytics.Event("delete-organization")

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("%s %s %s\n", printer.BoldRed("Organization"), printer.BoldGreen(name), printer.BoldRed("deleted"))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"deleted": name,
				})
			},
		}

		cmd.AddCommand(deleteCmd)
	}
}

//
//func DeleteCmd(ch *cmdutil.Helper) *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "delete <id>",
//		Args:  cobra.ExactArgs(1),
//		Short: "delete an API Key with the given ID",
//		RunE: func(cmd *cobra.Command, args []string) error {
//			ctx := cmd.Context()
//			client, err := ch.Client()
//			if err != nil {
//				return err
//			}
//
//			id := args[0]
//
//			end := ch.Printer.PrintProgress(fmt.Sprintf("Deleting API Key %s...", id))
//			_, err = client.Access.DeleteAccessApikeyID(access.NewDeleteAccessApikeyIDParamsWithContext(ctx).WithID(id))
//			end()
//			if err != nil {
//				return err
//			}
//
//			if ch.Printer.Format() == printer.Human {
//				ch.Printer.Printf("API Key %s %s\n", printer.BoldGreen(id), printer.BoldRed("deleted"))
//				return nil
//			}
//
//			return ch.Printer.PrintResource(map[string]string{
//				"deleted": id,
//			})
//		},
//	}
//
//	return cmd
//}
