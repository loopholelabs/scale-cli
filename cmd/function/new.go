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
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/spf13/cobra"
	"os"
	"path"
	textTemplate "text/template"
)

const (
	defaultSignature = "http@v0.3.4"
)

var (
	extensionLUT = map[string]string{
		string(scalefile.Go):   "go",
		string(scalefile.Rust): "rs",
	}
)

// NewCmd encapsulates the commands for creating new Functions
func NewCmd() command.SetupCommand[*config.Config] {
	var directory string
	var language string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		listCmd := &cobra.Command{
			Use:   "new <name> [flags]",
			Args:  cobra.ExactArgs(1),
			Short: "generate a new scale function with the given name",
			RunE: func(cmd *cobra.Command, args []string) error {
				name := args[0]
				if name == "" || !scalefunc.ValidString(name) {
					return fmt.Errorf("invalid name %s", name)
				}

				extension, ok := extensionLUT[language]
				if !ok {
					return fmt.Errorf("language %s is not supported", language)
				}

				scaleFile := &scalefile.ScaleFile{
					Version:   scalefile.V1Alpha,
					Name:      name,
					Tag:       "v0.1.0",
					Signature: defaultSignature,
					Language:  scalefile.Language(language),
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
					scaleFile.Dependencies = []scalefile.Dependency{
						{
							Name:    "github.com/loopholelabs/scale-signature",
							Version: "v0.2.9",
						},
						{
							Name:    "github.com/loopholelabs/scale-signature-http",
							Version: "v0.3.4",
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
					scaleFile.Dependencies = []scalefile.Dependency{
						{
							Name:    "scale_signature_http",
							Version: "0.3.4",
						},
						{
							Name:    "scale_signature",
							Version: "0.2.9",
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

		listCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale function in")
		listCmd.Flags().StringVarP(&language, "language", "l", string(scalefile.Go), "the language to create the new scale function in (go, rust)")

		cmd.AddCommand(listCmd)
	}
}
