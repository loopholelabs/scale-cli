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

package registry

import (
	"fmt"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/client/registry"
	"github.com/loopholelabs/scale/go/storage"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/spf13/cobra"
)

// PushCmd encapsulates the commands for pushing Functions
func PushCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		var name string
		var tag string
		var org string
		var public bool
		pushCmd := &cobra.Command{
			Use:   "push [<name>:<tag> | [<org>/<name>:<tag>]",
			Short: "push a locally available scale function to the registry",
			Long:  "Push a locally available scale function to the registry. The function must be available in the local cache directory. If no cache directory is specified, the default cache directory will be used. If the org is not specified, it will default to the local organization.",
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.Default
				if ch.Config.CacheDirectory != "" {
					var err error
					st, err = storage.New(ch.Config.CacheDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.CacheDirectory, err)
					}
				}

				parsed := utils.ParseFunction(args[0])
				if parsed.Organization == "" {
					parsed.Organization = utils.DefaultOrganization
				}

				if parsed.Organization == "" || !scalefunc.ValidString(parsed.Organization) {
					return fmt.Errorf("invalid organization name: %s", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return fmt.Errorf("invalid function name: %s", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return fmt.Errorf("invalid tag: %s", parsed.Tag)
				}

				e, err := st.Get(parsed.Name, parsed.Tag, parsed.Organization, "")
				if err != nil {
					return fmt.Errorf("failed to push function %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}
				if e == nil {
					return fmt.Errorf("function %s/%s:%s does not exist", parsed.Organization, parsed.Name, parsed.Tag)
				}

				end := ch.Printer.PrintProgress(fmt.Sprintf("Pushing %s/%s:%s to Scale Registry...", parsed.Organization, parsed.Name, parsed.Tag))

				ctx := cmd.Context()
				client := ch.Config.APIClient()

				if org != "" {
					parsed.Organization = org
				}

				if name != "" {
					parsed.Name = name
				}

				if tag != "" {
					parsed.Tag = tag
				}

				if parsed.Organization == "" || !scalefunc.ValidString(parsed.Organization) {
					return fmt.Errorf("invalid organization name: %s", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return fmt.Errorf("invalid function name: %s", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return fmt.Errorf("invalid tag: %s", parsed.Tag)
				}

				e.ScaleFunc.Name = parsed.Name
				e.ScaleFunc.Tag = parsed.Tag

				params := registry.NewPostRegistryFunctionParamsWithContext(ctx).WithFunction(utils.NewScaleFunctionNamedReadCloser(e.ScaleFunc)).WithPublic(&public)
				if parsed.Organization != utils.DefaultOrganization {
					params = params.WithOrganization(&parsed.Organization)
				}

				res, err := client.Registry.PostRegistryFunction(params)
				end()
				if err != nil {
					return err
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Pushed %s to the Scale Registry\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", res.GetPayload().Organization, res.GetPayload().Name, res.GetPayload().Tag)))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"name":   res.GetPayload().Name,
					"tag":    res.GetPayload().Tag,
					"org":    res.GetPayload().Organization,
					"public": fmt.Sprintf("%t", res.GetPayload().Public),
					"hash":   res.GetPayload().Hash,
				})
			},
		}

		pushCmd.Flags().StringVarP(&name, "override-name", "n", "", "the name of the function")
		pushCmd.Flags().StringVarP(&tag, "override-tag", "t", "", "the tag of the function")
		pushCmd.Flags().StringVarP(&org, "override-org", "o", "", "the organization of the function")
		pushCmd.Flags().BoolVarP(&public, "public", "p", false, "make the function public")

		cmd.AddCommand(pushCmd)
	}
}
