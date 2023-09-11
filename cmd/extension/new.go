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

package extension

import (
	"fmt"
	"os"
	"path"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/template"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/spf13/cobra"
)

// NewCmd encapsulates the commands for creating new Extensions
func NewCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		newCmd := &cobra.Command{
			Use:      "new [flags]",
			Args:     cobra.ExactArgs(0),
			Short:    "create a new scale extension",
			Hidden:   hidden,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				analytics.Event("new-extension")
				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}
				err := os.WriteFile(path.Join(sourceDir, fmt.Sprintf("scale.extension")), []byte(template.ExtensionFile), 0644)
				if err != nil {
					return fmt.Errorf("error writing extension: %w", err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully created new scale extension\n")
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"path": path.Join(directory, fmt.Sprintf("scale.extension")),
				})
			},
		}

		newCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale extension in")

		cmd.AddCommand(newCmd)
	}
}
