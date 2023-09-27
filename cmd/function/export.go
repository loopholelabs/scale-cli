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
	"github.com/loopholelabs/scale"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
	"os"
	"path"
)

// ExportCmd encapsulates the commands for exporting Functions
func ExportCmd() command.SetupCommand[*config.Config] {
	var outputName string
	var raw bool
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		exportCmd := &cobra.Command{
			Use:      "export <org>/<name>:<tag> <output_path>",
			Args:     cobra.ExactArgs(2),
			Short:    "export a compiled scale function to the given output path",
			Long:     "Export a compiled scale function to the given output path. The output path must always be a directory and the function will be exported to a file with the name <org>-<name>-<tag>.scale by default. This can be overridden using the --output-name flag. If the org is not specified or the function is not associated with an org, no organization will be used.",
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

				parsed := scale.Parse(args[0])
				if parsed.Organization == "" && !scalefunc.ValidString(parsed.Organization) {
					return utils.InvalidStringError("organization name", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return utils.InvalidStringError("function name", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return utils.InvalidStringError("function tag", parsed.Tag)
				}

				e, err := st.Get(parsed.Name, parsed.Tag, parsed.Organization, "")
				if err != nil {
					return fmt.Errorf("failed to export function %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}
				if e == nil {
					return fmt.Errorf("function %s/%s:%s does not exist", parsed.Organization, parsed.Name, parsed.Tag)
				}

				analytics.Event("export-function", map[string]string{"language": string(e.Schema.Language)})

				output := args[1]

				outputPath := output
				if !path.IsAbs(outputPath) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					outputPath = path.Join(wd, outputPath)
				}

				oInfo, err := os.Stat(outputPath)
				if err != nil {
					ch.Printer.Printf("Creating output directory %s\n", printer.BoldBlue(outputPath))
					err = os.MkdirAll(outputPath, 0755)
					if err != nil {
						return fmt.Errorf("failed to create output directory %s: %w", outputPath, err)
					}
				} else if !oInfo.IsDir() {
					return fmt.Errorf("output path %s is not a directory", outputPath)
				}

				if outputName == "" {
					suffix := "scale"
					if raw {
						suffix = "wasm"
					}
					outputPath = path.Join(outputPath, fmt.Sprintf("%s-%s-%s.%s", parsed.Organization, parsed.Name, parsed.Tag, suffix))
				} else {
					outputPath = path.Join(outputPath, outputName)
				}

				if raw {
					err = os.WriteFile(outputPath, e.Schema.Function, 0644)
				} else {
					err = os.WriteFile(outputPath, e.Schema.Encode(), 0644)
				}
				if err != nil {
					return fmt.Errorf("failed to write function to %s: %w", outputPath, err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Exported scale function %s to %s\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", parsed.Organization, parsed.Name, parsed.Tag)), printer.BoldBlue(outputPath))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"destination": outputPath,
					"org":         parsed.Organization,
					"name":        parsed.Name,
					"tag":         parsed.Tag,
					"hash":        e.Hash,
				})
			},
		}

		exportCmd.Flags().StringVarP(&outputName, "output-name", "o", "", "the (optional) output name of the function to export")
		exportCmd.Flags().BoolVar(&raw, "raw", false, "export the raw wasm module instead of the compiled scale function")
		cmd.AddCommand(exportCmd)
	}
}
