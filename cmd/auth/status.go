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

package auth

import (
	"fmt"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/userinfo"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/spf13/cobra"
)

// StatusCmd encapsulates the commands for the authentication status
func StatusCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		statusCmd := &cobra.Command{
			Use:      "status [flags]",
			Short:    "Retrieve the current Authentication Status using the Scale Authentication API",
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				end := ch.Printer.PrintProgress("Retrieving authentication status...")
				ctx := cmd.Context()
				c := ch.Config.APIClient()
				res, err := c.Userinfo.PostUserinfo(userinfo.NewPostUserinfoParamsWithContext(ctx))
				end()
				if err != nil {
					return utils.ErrNotAuthenticated
				}

				analytics.Event("status")

				return ch.Printer.PrintResource(map[string]string{
					"email":         res.GetPayload().Email,
					"org":           res.GetPayload().Organization,
					"organizations": fmt.Sprintf("%s", res.GetPayload().Organizations),
				})
			},
		}

		cmd.AddCommand(statusCmd)
	}
}
