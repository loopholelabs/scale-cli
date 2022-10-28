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
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client/access"
	"github.com/spf13/cobra"
)

func CreateCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <name>",
		Args:  cobra.ExactArgs(1),
		Short: "create an API Key with the given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			name := args[0]

			end := ch.Printer.PrintProgress(fmt.Sprintf("Creating API Key %s...", name))
			res, err := client.Access.PostAccessApikeyName(access.NewPostAccessApikeyNameParamsWithContext(ctx).WithName(name))
			end()
			if err != nil {
				return err
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Created API Key '%s': %s (this will only be displayed once)\n", printer.Bold(res.Payload.Name), printer.BoldGreen(res.Payload.Apikey))
				return nil
			}

			return ch.Printer.PrintResource(apiKey{
				Created: utils.Int64ToTime(res.Payload.CreatedAt).String(),
				ID:      res.Payload.ID,
				Name:    res.Payload.ID,
				Value:   res.Payload.Apikey,
			})
		},
	}

	return cmd
}
