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
	"github.com/spf13/cobra"
)

// RemoveCmd encapsulates the command for adding a domain
func RemoveCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		removeCmd := &cobra.Command{
			Use:      "remove [domain]",
			Args:     cobra.ExactArgs(1),
			Short:    "removes a custom domain",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := domain.NewDeleteDomainDomainParamsWithContext(ctx).WithDomain(args[0])

				var end = ch.Printer.PrintProgress("Removing domain..")
				_, err := client.Domain.DeleteDomainDomain(params)
				if err != nil {
					end()
					return err
				}
				end()
				ch.Printer.Printf("Domain %s has been removed\n", printer.Bold(args[0]))
				return nil
			},
		}

		cmd.AddCommand(removeCmd)
	}
}
