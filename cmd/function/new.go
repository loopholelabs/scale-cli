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
	"os"
	"path"
	"strings"
	textTemplate "text/template"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/registry"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/template"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/extension"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
)

// NewCmd encapsulates the commands for creating new Functions
func NewCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string
	var language string
	var signature string
	var extensions []string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		newCmd := &cobra.Command{
			Use:      "new <name>:<tag> [flags]",
			Args:     cobra.ExactArgs(1),
			Short:    "generate a new scale function with the given name",
			Hidden:   hidden,
			PreRunE:  utils.PreRunOptionalAuthenticatedAPI(ch),
			PostRunE: utils.PostRunOptionalAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.DefaultSignature
				et := storage.DefaultExtension
				if ch.Config.StorageDirectory != "" {
					var err error
					st, err = storage.NewSignature(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
					et, err = storage.NewExtension(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				nametag := strings.Split(args[0], ":")
				if len(nametag) != 2 {
					return fmt.Errorf("invalid name or tag %s", args[0])
				}
				functionName := nametag[0]
				functionTag := nametag[1]

				if functionName == "" || !scalefunc.ValidString(functionName) {
					return utils.InvalidStringError("function name", functionName)
				}

				if functionTag == "" || !scalefunc.ValidString(functionTag) {
					return utils.InvalidStringError("function tag", functionTag)
				}

				if signature == "" {
					return fmt.Errorf("signature is required")
				}

				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}

				_, err := os.Stat(sourceDir)
				if err != nil && os.IsNotExist(err) {
					err = os.MkdirAll(sourceDir, 0755)
					if err != nil {
						return fmt.Errorf("error creating directory %s: %w", sourceDir, err)
					}
				}

				var extensionData = make([]extension.ExtensionInfo, 0)
				for _, e := range extensions {
					var extensionPath string
					parsedExtension := scale.Parse(e)
					if parsedExtension.Organization == "local" {
						extensionPath, err = et.Path(parsedExtension.Name, parsedExtension.Tag, parsedExtension.Organization, "")
						if err != nil {
							return fmt.Errorf("error while getting extension %s: %w", parsedExtension.Name, err)
						}

						ext, err := et.Get(parsedExtension.Name, parsedExtension.Tag, parsedExtension.Organization, "")
						if err != nil {
							return fmt.Errorf("error while getting extension %s: %w", parsedExtension.Name, err)
						}

						switch scalefunc.Language(language) {
						case scalefunc.Go:
							extensionData = append(extensionData, extension.ExtensionInfo{
								Name:    ext.Schema.Name,
								Path:    path.Join(extensionPath, "golang", "guest"),
								Version: "v0.1.0",
							})
						default:
							panic("Only go extension for now")
						}
					} else {
						panic("Only local extension for now")
					}
				}

				var signaturePath string
				var signatureVersion string
				var signatureContext string
				parsedSignature := scale.Parse(signature)
				if parsedSignature.Organization == "local" {
					signaturePath, err = st.Path(parsedSignature.Name, parsedSignature.Tag, parsedSignature.Organization, "")
					if err != nil {
						return fmt.Errorf("error while getting signature %s: %w", parsedSignature.Name, err)
					}
					sig, err := st.Get(parsedSignature.Name, parsedSignature.Tag, parsedSignature.Organization, "")
					if err != nil {
						return fmt.Errorf("error while getting signature %s: %w", parsedSignature.Name, err)
					}
					switch scalefunc.Language(language) {
					case scalefunc.Go:
						signatureVersion = ""
						signaturePath = path.Join(signaturePath, "golang", "guest")
					case scalefunc.Rust:
						signatureVersion = ""
						signaturePath = path.Join(signaturePath, "rust", "guest")
					case scalefunc.TypeScript:
						signatureVersion = ""
						signaturePath = path.Join(signaturePath, "typescript", "guest")
					default:
						return fmt.Errorf("language %s is not supported", language)
					}
					signatureContext = sig.Schema.Context
				} else {
					ctx := cmd.Context()
					client := ch.Config.APIClient()

					end := ch.Printer.PrintProgress(fmt.Sprintf("Fetching signature %s/%s:%s...", parsedSignature.Organization, parsedSignature.Name, parsedSignature.Tag))
					res, err := client.Registry.GetRegistrySignatureOrgNameTag(registry.NewGetRegistrySignatureOrgNameTagParamsWithContext(ctx).WithOrg(parsedSignature.Organization).WithName(parsedSignature.Name).WithTag(parsedSignature.Tag))
					end()
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsedSignature.Organization, parsedSignature.Name, parsedSignature.Tag, err)
					}

					switch scalefunc.Language(language) {
					case scalefunc.Go:
						signatureVersion = ""
						signaturePath = res.GetPayload().GolangImportPathGuest
					case scalefunc.Rust:
						signatureVersion = "0.1.0"
						signaturePath = ""
					case scalefunc.TypeScript:
						return fmt.Errorf("typescript functions are not currently supported via the registry")
					default:
						return fmt.Errorf("language %s is not supported", language)
					}
					signatureContext = res.GetPayload().Context
				}

				scaleFile := &scalefile.Schema{
					Version: scalefile.V1AlphaVersion,
					Name:    functionName,
					Tag:     functionTag,
					Signature: scalefile.SignatureSchema{
						Organization: parsedSignature.Organization,
						Name:         parsedSignature.Name,
						Tag:          parsedSignature.Tag,
					},
				}

				scaleFile.Extensions = make([]scalefile.ExtensionSchema, 0)

				for _, ex := range extensions {
					parsedExtension := scale.Parse(ex)

					scaleFile.Extensions = append(scaleFile.Extensions, scalefile.ExtensionSchema{
						Organization: parsedExtension.Organization,
						Name:         parsedExtension.Name,
						Tag:          parsedExtension.Tag,
					})
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
						"package_name":      functionName,
						"signature_path":    signaturePath,
						"signature_version": signatureVersion,
						"extensions":        extensionData,
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
						"package_name": functionName,
						"context_name": signatureContext,
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
						"package_name":      functionName,
						"signature_package": fmt.Sprintf("%s_%s_%s_guest", parsedSignature.Organization, parsedSignature.Name, parsedSignature.Tag),
						"signature_version": signatureVersion,
						"signature_path":    signaturePath,
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
						"context_name": signatureContext,
					})
					if err != nil {
						_ = funcFile.Close()
						return fmt.Errorf("error writing lib.rs file: %w", err)
					}

					err = funcFile.Close()
					if err != nil {
						return fmt.Errorf("error closing lib.rs file: %w", err)
					}
				case scalefunc.TypeScript:
					scaleFile.Language = string(scalefunc.TypeScript)
					scaleFile.Function = "scale"
					analytics.Event("new-function", map[string]string{"language": "typescript"})

					packageTempl, err := textTemplate.New("dependencies").Parse(template.TypescriptPackageTemplate)
					if err != nil {
						return fmt.Errorf("error parsing package.json template: %w", err)
					}

					dependencyFile, err := os.Create(path.Join(sourceDir, "package.json"))
					if err != nil {
						return fmt.Errorf("error creating package.json file: %w", err)
					}

					err = packageTempl.Execute(dependencyFile, map[string]interface{}{
						"package_name":      functionName,
						"signature_path":    signaturePath,
						"signature_version": signatureVersion,
					})
					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing package.json file: %w", err)
					}

					err = dependencyFile.Close()
					if err != nil {
						return fmt.Errorf("error closing package.json file: %w", err)
					}

					funcTempl, err := textTemplate.New("function").Parse(template.TypeScriptFunctionTemplate)
					if err != nil {
						return fmt.Errorf("error parsing function template: %w", err)
					}

					funcFile, err := os.Create(path.Join(sourceDir, "index.ts"))
					if err != nil {
						return fmt.Errorf("error creating index.ts file: %w", err)
					}

					err = funcTempl.Execute(funcFile, map[string]interface{}{
						"context_name": signatureContext,
					})
					if err != nil {
						_ = funcFile.Close()
						return fmt.Errorf("error writing index.ts file: %w", err)
					}

					err = funcFile.Close()
					if err != nil {
						return fmt.Errorf("error closing index.ts file: %w", err)
					}
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
					ch.Printer.Printf("Successfully created new %s scale function %s:%s\n", printer.BoldGreen(language), printer.BoldGreen(functionName), printer.BoldGreen(functionTag))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"path":     scaleFilePath,
					"name":     functionName,
					"tag":      functionTag,
					"language": language,
				})
			},
		}

		newCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale function in")
		newCmd.Flags().StringVarP(&language, "language", "l", string(scalefunc.Go), "the language for the new scale function (go, rust, ts)")
		newCmd.Flags().StringVarP(&signature, "signature", "s", "", "the signature to use with the new scale function")
		newCmd.Flags().StringArrayVarP(&extensions, "extension", "e", []string{}, "the extensions to use in the scale function")
		cmd.AddCommand(newCmd)
	}
}
