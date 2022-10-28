/*
	Copyright 2022 Loophole Labs

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

package apikey

import (
	"fmt"
	"github.com/loopholelabs/auth/pkg/utils"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/pkg/client/access"
	"github.com/spf13/cobra"
)

func GetCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <name>",
		Args:  cobra.ExactArgs(1),
		Short: "get information about an API Key with the given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			name := args[0]

			end := ch.Printer.PrintProgress(fmt.Sprintf("Retrieving API Key %s...", name))
			res, err := client.Access.GetAccessApikeyName(access.NewGetAccessApikeyNameParamsWithContext(ctx).WithName(name))
			end()
			if err != nil {
				return err
			}

			return ch.Printer.PrintResource(apiKeyRedacted{
				Name:    res.Payload.Name,
				Created: utils.Int64ToTime(res.Payload.CreatedAt).String(),
				ID:      res.Payload.ID,
			})
		},
	}

	return cmd
}
