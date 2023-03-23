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
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/posthog/posthog-go"
	"github.com/spf13/cobra"
	"os"
	"path"
	textTemplate "text/template"
	"time"
)

const (
	defaultSignature = "http@v0.3.8"
)

var (
	extensionLUT = map[string]string{
		string(scalefunc.Go):   "go",
		string(scalefunc.Rust): "rs",
	}
)

// NewCmd encapsulates the commands for creating new Functions
func NewCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string
	var language string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		newCmd := &cobra.Command{
			Use:      "new <name> [flags]",
			Args:     cobra.ExactArgs(1),
			Short:    "generate a new scale function with the given name",
			Hidden:   hidden,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				name := args[0]
				if name == "" || !scalefunc.ValidString(name) {
					return utils.InvalidStringError("function name", name)
				}

				extension, ok := extensionLUT[language]
				if !ok {
					return fmt.Errorf("language %s is not supported", language)
				}

				scaleFile := &scalefile.ScaleFile{
					Version:   scalefile.V1Alpha,
					Name:      name,
					Tag:       utils.DefaultTag,
					Signature: defaultSignature,
					Language:  scalefunc.Language(language),
					Source:    fmt.Sprintf("scale.%s", extension),
				}

				if _, err := os.Stat(directory); os.IsNotExist(err) {
					err = os.MkdirAll(directory, 0755)
					if err != nil {
						return fmt.Errorf("error creating directory %s: %w", directory, err)
					}
				}

				scaleFilePath := path.Join(directory, "scalefile")

				switch language {
				case "go":
					if analytics.Client != nil {
						_ = analytics.Client.Enqueue(posthog.Capture{
							DistinctId: analytics.MachineID,
							Event:      "new-function",
							Timestamp:  time.Now(),
							Properties: posthog.NewProperties().Set("language", "go"),
						})
					}
					scaleFile.Dependencies = []scalefile.Dependency{
						{
							Name:    "github.com/loopholelabs/scale-signature",
							Version: "v0.2.11",
						},
						{
							Name:    "github.com/loopholelabs/scale-signature-http",
							Version: "v0.3.8",
						},
					}

					tmpl, err := textTemplate.New("dependencies").Parse(template.GoTemplate)
					if err != nil {
						return fmt.Errorf("error parsing dependency template: %w", err)
					}

					dependencyFile, err := os.Create(fmt.Sprintf("%s/go.mod", directory))
					if err != nil {
						return fmt.Errorf("error creating dependencies file: %w", err)
					}

					err = tmpl.Execute(dependencyFile, scaleFile.Dependencies)
					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing dependencies file: %w", err)
					}
				case "rust":
					if analytics.Client != nil {
						_ = analytics.Client.Enqueue(posthog.Capture{
							DistinctId: analytics.MachineID,
							Event:      "new-function",
							Timestamp:  time.Now(),
							Properties: posthog.NewProperties().Set("language", "rust"),
						})
					}
					scaleFile.Dependencies = []scalefile.Dependency{
						{
							Name:    "scale_signature_http",
							Version: "0.3.8",
						},
						{
							Name:    "scale_signature",
							Version: "0.2.11",
						},
					}

					tmpl, err := textTemplate.New("dependencies").Parse(template.RustTemplate)
					if err != nil {
						return fmt.Errorf("error parsing dependency template: %w", err)
					}

					dependencyFile, err := os.Create(fmt.Sprintf("%s/Cargo.toml", directory))
					if err != nil {
						return fmt.Errorf("error creating dependencies file: %w", err)
					}

					err = tmpl.Execute(dependencyFile, scaleFile.Dependencies)

					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing dependencies file: %w", err)
					}
				default:
					return fmt.Errorf("language %s is not supported", language)
				}

				err := scalefile.Write(scaleFilePath, scaleFile)
				if err != nil {
					return fmt.Errorf("error writing scalefile: %w", err)
				}

				err = os.WriteFile(fmt.Sprintf("%s/%s", directory, scaleFile.Source), template.LUT[language](), 0644)
				if err != nil {
					return fmt.Errorf("error writing source file: %w", err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully created new %s scale function %s\n", printer.BoldGreen(language), printer.BoldGreen(name))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"path":     scaleFilePath,
					"name":     name,
					"language": language,
				})
			},
		}

		newCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale function in")
		newCmd.Flags().StringVarP(&language, "language", "l", string(scalefunc.Go), "the language to create the new scale function in (go, rust)")

		cmd.AddCommand(newCmd)
	}
}
