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
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/template"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	textTemplate "text/template"
)

// NewCmd encapsulates the commands for creating new Functions
func NewCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string
	var language string
	var signature string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		newCmd := &cobra.Command{
			Use:      "new <name>:<tag> [flags]",
			Args:     cobra.ExactArgs(1),
			Short:    "generate a new scale function with the given name",
			Hidden:   hidden,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.DefaultSignature
				if ch.Config.StorageDirectory != "" {
					var err error
					st, err = storage.NewSignature(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				nametag := strings.Split(args[0], ":")
				if len(nametag) != 2 {
					return fmt.Errorf("invalid name or tag %s", args[0])
				}
				name := nametag[0]
				tag := nametag[1]

				if name == "" || !scalefunc.ValidString(name) {
					return utils.InvalidStringError("function name", name)
				}

				if tag == "" || !scalefunc.ValidString(tag) {
					return utils.InvalidStringError("function tag", tag)
				}

				if signature == "" {
					return fmt.Errorf("signature is required")
				}

				org, name, tag := scalefunc.ParseFunctionName(signature)
				signaturePath, err := st.Path(name, tag, org, "")
				if err != nil {
					return fmt.Errorf("error while getting signature %s: %w", signature, err)
				}
				sig, err := st.Get(name, tag, org, "")
				if err != nil {
					return fmt.Errorf("error while getting signature %s: %w", signature, err)
				}

				scaleFile := &scalefile.Schema{
					Version: scalefile.V1AlphaVersion,
					Name:    name,
					Tag:     "latest",
					Signature: scalefile.SignatureSchema{
						Organization: org,
						Name:         name,
						Tag:          tag,
					},
				}

				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}

				if _, err = os.Stat(sourceDir); os.IsNotExist(err) {
					err = os.MkdirAll(sourceDir, 0755)
					if err != nil {
						return fmt.Errorf("error creating directory %s: %w", sourceDir, err)
					}
				}

				scaleFilePath := path.Join(sourceDir, "scalefile")
				switch scalefunc.Language(language) {
				case scalefunc.Go:
					scaleFile.Language = string(scalefunc.Go)
					scaleFile.Function = "Scale"
					analytics.Event("new-function", map[string]string{"language": "go"})

					modfileTempl, err := textTemplate.New("dependencies").Parse(template.GoModfileTemplate)
					if err != nil {
						return fmt.Errorf("error parsing go.mod template: %w", err)
					}

					dependencyFile, err := os.Create(path.Join(sourceDir, "go.mod"))
					if err != nil {
						return fmt.Errorf("error creating go.mod file: %w", err)
					}

					err = modfileTempl.Execute(dependencyFile, map[string]interface{}{
						"package":                  name,
						"old_signature_dependency": "signature",
						"old_signature_version":    "",
						"new_signature_dependency": path.Join(signaturePath, "golang", "guest"),
						"new_signature_version":    "",
						"dependencies": []template.Dependency{
							{
								Name:    "signature",
								Version: "v0.1.0",
							},
						},
					})
					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing go.mod file: %w", err)
					}

					err = dependencyFile.Close()
					if err != nil {
						return fmt.Errorf("error closing go.mod file: %w", err)
					}

					funcTempl, err := textTemplate.New("function").Parse(template.GoFunctionTemplate)
					if err != nil {
						return fmt.Errorf("error parsing function template: %w", err)
					}

					funcFile, err := os.Create(path.Join(sourceDir, "main.go"))
					if err != nil {
						return fmt.Errorf("error creating main.go file: %w", err)
					}

					err = funcTempl.Execute(funcFile, map[string]interface{}{
						"package": name,
						"context": sig.Schema.Context,
					})
					if err != nil {
						_ = funcFile.Close()
						return fmt.Errorf("error writing main.go file: %w", err)
					}

					err = funcFile.Close()
					if err != nil {
						return fmt.Errorf("error closing main.go file: %w", err)
					}
				case scalefunc.Rust:
					scaleFile.Language = string(scalefunc.Rust)
					scaleFile.Function = "scale"
					analytics.Event("new-function", map[string]string{"language": "rust"})

					cargofileTempl, err := textTemplate.New("dependencies").Parse(template.RustCargofileTemplate)
					if err != nil {
						return fmt.Errorf("error parsing Cargo.toml template: %w", err)
					}

					dependencyFile, err := os.Create(path.Join(sourceDir, "Cargo.toml"))
					if err != nil {
						return fmt.Errorf("error creating Cargo.toml file: %w", err)
					}

					err = cargofileTempl.Execute(dependencyFile, map[string]interface{}{
						"package":              name,
						"version":              "0.1.0",
						"signature_dependency": "signature",
						"signature_path":       path.Join(signaturePath, "rust", "guest"),
						"signature_package":    fmt.Sprintf("%s_%s_%s_guest", scaleFile.Signature.Organization, scaleFile.Signature.Name, scaleFile.Signature.Tag),
					})
					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing Cargo.toml file: %w", err)
					}

					err = dependencyFile.Close()
					if err != nil {
						return fmt.Errorf("error closing Cargo.toml file: %w", err)
					}

					funcTempl, err := textTemplate.New("function").Parse(template.RustFunctionTemplate)
					if err != nil {
						return fmt.Errorf("error parsing function template: %w", err)
					}

					funcFile, err := os.Create(path.Join(sourceDir, "lib.rs"))
					if err != nil {
						return fmt.Errorf("error creating lib.rs file: %w", err)
					}

					err = funcTempl.Execute(funcFile, map[string]interface{}{
						"context": sig.Schema.Context,
					})
					if err != nil {
						_ = funcFile.Close()
						return fmt.Errorf("error writing lib.rs file: %w", err)
					}

					err = funcFile.Close()
					if err != nil {
						return fmt.Errorf("error closing lib.rs file: %w", err)
					}
				//case scalefunc.TypeScript:
				//	scaleFile.Language = scalefunc.TypeScript
				//	scaleFile.Dependencies = []scalefile.Dependency{
				//		{
				//			Name:    "@loopholelabs/scale-signature-http",
				//			Version: "0.3.8",
				//		},
				//		{
				//			Name:    "@loopholelabs/scale-signature",
				//			Version: "0.2.11",
				//		},
				//	}
				//
				//	tmpl, err := textTemplate.New("dependencies").Parse(template.TypeScriptTemplate)
				//	if err != nil {
				//		return fmt.Errorf("error parsing dependency template: %w", err)
				//	}
				//
				//	dependencyFile, err := os.Create(fmt.Sprintf("%s/package.json", directory))
				//	if err != nil {
				//		return fmt.Errorf("error creating dependencies file: %w", err)
				//	}
				//
				//	err = tmpl.Execute(dependencyFile, scaleFile.Dependencies)
				//
				//	if err != nil {
				//		_ = dependencyFile.Close()
				//		return fmt.Errorf("error writing dependencies file: %w", err)
				//	}
				default:
					return fmt.Errorf("language %s is not supported", language)
				}

				scaleFileData, err := scaleFile.Encode()
				if err != nil {
					return fmt.Errorf("error encoding scalefile: %w", err)
				}
				err = os.WriteFile(scaleFilePath, scaleFileData, 0644)
				if err != nil {
					return fmt.Errorf("error writing scalefile: %w", err)
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
		newCmd.Flags().StringVarP(&language, "language", "l", string(scalefunc.Go), "the language for the new scale function (go, rust, ts)")
		newCmd.Flags().StringVarP(&signature, "signature", "s", "", "the signature to use with the new scale function")

		cmd.AddCommand(newCmd)
	}
}
