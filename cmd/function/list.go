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
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
)

// ListCmd encapsulates the commands for listing Functions
func ListCmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:      "list",
			Short:    "list locally available scale functions",
			Args:     cobra.NoArgs,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.DefaultFunction
				if ch.Config.StorageDirectory != "" {
					var err error
					st, err = storage.NewFunction(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}
				scaleFuncEntries, err := st.List()
				if err != nil {
					return fmt.Errorf("failed to list scale functions: %w", err)
				}

				if len(scaleFuncEntries) == 0 && ch.Printer.Format() == printer.Human {
					ch.Printer.Println("No Scale Functions available yet.")
					return nil
				}

				funcs := make([]functionModel, len(scaleFuncEntries))
				for i, entry := range scaleFuncEntries {
					funcs[i] = functionModel{
						Name:      entry.Schema.Name,
						Tag:       entry.Schema.Tag,
						Signature: entry.Schema.Signature.Name,
						Language:  string(entry.Schema.Language),
						Version:   string(entry.Schema.Tag), // FIXME
						Hash:      entry.Hash,
						Org:       entry.Organization,
					}
				}

				return ch.Printer.PrintResource(funcs)
			},
		}

		cmd.AddCommand(listCmd)
	}
}
