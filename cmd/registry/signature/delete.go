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

package signature

import (
	"fmt"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/registry"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/spf13/cobra"
)

// DeleteCmd encapsulates the commands for deleting Signatures
func DeleteCmd() command.SetupCommand[*config.Config] {
	var public bool
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		pushCmd := &cobra.Command{
			Use:      "delete <org>/<name>:<tag> [flags]",
			Args:     cobra.ExactArgs(1),
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				parsed := utils.Parse(args[0])
				if parsed.Organization != "" && !scalefunc.ValidString(parsed.Organization) {
					return utils.InvalidStringError("organization name", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return utils.InvalidStringError("function name", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return utils.InvalidStringError("function tag", parsed.Tag)
				}

				ctx := cmd.Context()
				client := ch.Config.APIClient()
				end := ch.Printer.PrintProgress(fmt.Sprintf("Deleting %s/%s:%s from the Scale Registry...", parsed.Organization, parsed.Name, parsed.Tag))

				_, err := client.Registry.DeleteRegistrySignatureOrgNameTag(registry.NewDeleteRegistrySignatureOrgNameTagParamsWithContext(ctx).WithOrg(parsed.Organization).WithName(parsed.Name).WithTag(parsed.Tag))
				end()
				if err != nil {
					return err
				}

				analytics.Event("delete-signature")

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Deleted %s from the Scale Registry\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", parsed.Organization, parsed.Name, parsed.Tag)))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"name": parsed.Name,
					"tag":  parsed.Tag,
					"org":  parsed.Organization,
				})
			},
		}

		pushCmd.Flags().BoolVar(&public, "public", true, "whether the signature is publicly available")
		_ = pushCmd.Flags().MarkHidden("public")

		cmd.AddCommand(pushCmd)
	}
}
