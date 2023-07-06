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

package cloud

import (
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/client/cloud"
	"github.com/spf13/cobra"
)

// DeleteCmd encapsulates the commands for deleting Functions
func DeleteCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		deleteCmd := &cobra.Command{
			Use:      "delete [function identifier]",
			Args:     cobra.ExactArgs(1),
			Short:    "delete a deployed Scale function",
			Long:     "delete a deployed Scale function",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := cloud.NewDeleteCloudFunctionIDParamsWithContext(ctx).WithID(args[0])

				var end = ch.Printer.PrintProgress("Deleting function...\n")
				_, err := client.Cloud.DeleteCloudFunctionID(params)
				if err != nil {
					end()
					return err
				}
				end()
				ch.Printer.Printf("The function has been deleted from Scale Cloud\n")
				return nil
			},
		}

		cmd.AddCommand(deleteCmd)
	}
}
