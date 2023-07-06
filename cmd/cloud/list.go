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
	"fmt"
	"time"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/client/cloud"
	"github.com/spf13/cobra"
)

type DeploymentTableRow struct {
	ID           string `header:"ID"`
	Functions    string `header:"Functions"`
	Created      string `header:"Created"`
	URL          string `header:"URL"`
	CustomDomain string `header:"Custom Domain"`
}

// ListCmd encapsulates the commands for listing Functions
func ListCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:      "list",
			Short:    "list your deployed Scale functions",
			Long:     "list your deployed Scale function",
			Hidden:   hidden,
			PreRunE:  utils.PreRunAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()
				client := ch.Config.APIClient()
				params := cloud.NewGetCloudFunctionParamsWithContext(ctx)

				result, err := client.Cloud.GetCloudFunction(params)
				if err != nil {
					return err
				}
				deployments := make([]DeploymentTableRow, len(result.Payload))
				for i, deployment := range result.Payload {
					names := ""
					for i, fun := range deployment.Functions {
						names += fmt.Sprintf("%s:%s", fun.Name, fun.Tag)
						if i < len(deployment.Functions)-1 {
							names += " → "
						}
					}

					link := fmt.Sprintf("https://%s.%s", deployment.Subdomain, deployment.RootDomain)
					date, err := time.Parse(time.RFC3339, deployment.CreatedAt)
					if err != nil {
						return err
					}
					hasCustomDomain := ""
					if len(deployment.CustomDomains) > 0 || i == 2 {
						hasCustomDomain = "✅ Attached"
					}

					deployments[i] = DeploymentTableRow{
						ID:           deployment.Identifier,
						Functions:    names,
						Created:      date.Local().Format("2006-01-02 15:04:05"),
						URL:          link,
						CustomDomain: hasCustomDomain,
					}
				}

				if len(deployments) == 0 {
					ch.Printer.Print("No deployments found\n")
					return nil
				}
				ch.Printer.PrintResource(deployments)

				return nil
			},
		}

		cmd.AddCommand(listCmd)
	}
}
