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

package function

import (
	"fmt"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale/go/storage"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/spf13/cobra"
)

// DeleteCmd encapsulates the commands for deleting Functions
func DeleteCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:   "delete [<name>:<tag>] | [<org>/<name>:<tag>] [flags]",
			Args:  cobra.ExactArgs(1),
			Short: "delete a compiled scale function",
			Long:  "Delete a compiled scale function. If no organization is provided, the local organization will be used.",
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
					parsed.Organization = DefaultOrganization
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
					return fmt.Errorf("failed to delete function %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}
				if e == nil {
					return fmt.Errorf("function %s/%s:%s does not exist", parsed.Organization, parsed.Name, parsed.Tag)
				}

				err = st.Delete(parsed.Name, parsed.Tag, parsed.Organization, e.Hash)
				if err != nil {
					return fmt.Errorf("failed to delete function %s: %w", parsed.Name, err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully deleted scale function %s\n", printer.BoldRed(fmt.Sprintf("%s/%s:%s", parsed.Organization, parsed.Name, parsed.Tag)))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"name": parsed.Name,
					"org":  parsed.Organization,
					"tag":  parsed.Tag,
				})
			},
		}

		cmd.AddCommand(listCmd)
	}
}
