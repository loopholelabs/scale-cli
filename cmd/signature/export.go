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

package signature

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
	"path/filepath"
)

// ExportCmd encapsulates the commands for exporting Signatures
func ExportCmd() command.SetupCommand[*config.Config] {
	var manifest bool
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		exportCmd := &cobra.Command{
			Use:      "export <org>/<name>:<tag> <language> <guest|host> <output_path>",
			Args:     cobra.ExactArgs(4),
			Short:    "export a generated scale signature to the given output path",
			Long:     "Export a generated scale signature to the given output path. The output path must always be a directory.",
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.DefaultSignature
				if ch.Config.StorageDirectory != "" {
					var err error
					st, err = storage.NewSignature(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate signature storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				parsed := scale.Parse(args[0])
				if parsed.Organization == "" && !scalefunc.ValidString(parsed.Organization) {
					return utils.InvalidStringError("organization name", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return utils.InvalidStringError("signature name", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return utils.InvalidStringError("signature tag", parsed.Tag)
				}

				language := args[1]
				kind := args[2]
				output := args[3]

				kindString := "guest"
				switch kind {
				case "guest":
					switch scalefunc.Language(language) {
					case scalefunc.Go, scalefunc.Rust:
					default:
						return fmt.Errorf("invalid signature language %s for guest: must be go or rust", language)
					}
					kindString = "guest"
				case "host":
					switch scalefunc.Language(language) {
					case scalefunc.Go:
					default:
						return fmt.Errorf("invalid signature language %s for guest: must be go", language)
					}
					kindString = "host"
				default:
					return fmt.Errorf("invalid signature kind %s: must be guest or host", kind)
				}

				analytics.Event("export-signature", map[string]string{"language": language})

				switch scalefunc.Language(language) {
				case scalefunc.Go:
					language = "golang"
				case scalefunc.Rust:
					language = "rust"
				default:
					return fmt.Errorf("invalid signature language %s: must be go, rust, or typescript", language)
				}

				signaturePath, err := st.Path(parsed.Name, parsed.Tag, parsed.Organization, "")
				if err != nil {
					return fmt.Errorf("failed to export signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}

				outputPath := output
				if !path.IsAbs(outputPath) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					outputPath = path.Join(wd, outputPath)
				}

				oInfo, err := os.Stat(output)
				if err != nil {
					return fmt.Errorf("failed to stat output path %s: %w", output, err)
				}

				if !oInfo.IsDir() {
					return fmt.Errorf("output path %s is not a directory", output)
				}

				files, err := filepath.Glob(path.Join(signaturePath, language, kindString, "*"))
				if err != nil {
					return fmt.Errorf("failed to glob signature files: %w", err)
				}

				for _, file := range files {
					name := filepath.Base(file)
					if manifest || (name != "go.mod" && name != "Cargo.toml") {
						data, err := os.ReadFile(file)
						if err != nil {
							return fmt.Errorf("failed to read file %s: %w", file, err)
						}
						err = os.WriteFile(path.Join(outputPath, filepath.Base(file)), data, 0644)
						if err != nil {
							return fmt.Errorf("failed to write file %s: %w", file, err)
						}
					}
				}
				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Exported scale function %s to %s\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", parsed.Organization, parsed.Name, parsed.Tag)), printer.BoldBlue(output))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"destination": output,
					"org":         parsed.Organization,
					"name":        parsed.Name,
					"tag":         parsed.Tag,
				})
			},
		}

		exportCmd.Flags().BoolVar(&manifest, "manifest", true, "export the manifest file with the compiled signature")

		cmd.AddCommand(exportCmd)
	}
}
