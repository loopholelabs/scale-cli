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
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/registry"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/build"
	"github.com/loopholelabs/scale/extension"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/signature"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
)

// BuildCmd encapsulates the commands for building Functions
func BuildCmd(hidden bool) command.SetupCommand[*config.Config] {
	var name string
	var tag string
	var directory string

	var release bool
	var target string

	var goBin string
	var tinygoBin string
	var cargoBin string
	var npmBin string

	var tinygoArgs []string
	var cargoArgs []string

	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		buildCmd := &cobra.Command{
			Use:      "build [flags]",
			Args:     cobra.ExactArgs(0),
			Short:    "build a scale function locally and store it in the cache",
			Long:     "Build a scale function locally and store it in the cache. The scalefile must be in the current directory or specified with the --directory flag.",
			Hidden:   hidden,
			PreRunE:  utils.PreRunOptionalAuthenticatedAPI(ch),
			PostRunE: utils.PostRunOptionalAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				sfPath := path.Join(directory, "scalefile")
				sf, err := scalefile.ReadSchema(sfPath)
				if err != nil {
					return fmt.Errorf("failed to read scalefile at %s: %w", sfPath, err)
				}

				if name == "" {
					name = sf.Name
				} else {
					sf.Name = name
				}

				if tag == "" {
					tag = sf.Tag
				} else {
					sf.Tag = tag
				}

				if sf.Name == "" || !scalefunc.ValidString(sf.Name) {
					return utils.InvalidStringError("name", sf.Name)
				}

				if sf.Tag == "" || !scalefunc.ValidString(sf.Tag) {
					return utils.InvalidStringError("tag", sf.Tag)
				}

				buildTarget := build.WASITarget
				switch strings.ToLower(target) {
				case "wasi":
					buildTarget = build.WASITarget
				case "wasm":
					buildTarget = build.WASMTarget
				default:
					return fmt.Errorf("invalid build target %s", target)
				}

				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}
				analytics.Event("build-function", map[string]string{"language": sf.Language})

				stb := storage.DefaultBuild
				if ch.Config.StorageDirectory != "" {
					stb, err = storage.NewBuild(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate build storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				var signatureSchema *signature.Schema
				if sf.Signature.Organization == "local" {
					sts := storage.DefaultSignature
					if ch.Config.StorageDirectory != "" {
						sts, err = storage.NewSignature(ch.Config.StorageDirectory)
						if err != nil {
							return fmt.Errorf("failed to instantiate signature storage for %s: %w", ch.Config.StorageDirectory, err)
						}
					}

					sig, err := sts.Get(sf.Signature.Name, sf.Signature.Tag, sf.Signature.Organization, "")
					if err != nil {
						return fmt.Errorf("failed to get signature local/%s:%s: %w", sf.Signature.Name, sf.Signature.Tag, err)
					}
					if sig == nil {
						return fmt.Errorf("signature local/%s:%s not found", sf.Signature.Name, sf.Signature.Tag)
					}

					signatureSchema = sig.Schema
				} else {
					ctx := cmd.Context()
					client := ch.Config.APIClient()
					end := ch.Printer.PrintProgress(fmt.Sprintf("Fetching signature %s/%s:%s...", sf.Signature.Organization, sf.Signature.Name, sf.Signature.Tag))
					res, err := client.Registry.GetRegistrySignatureOrgNameTag(registry.NewGetRegistrySignatureOrgNameTagParamsWithContext(ctx).WithOrg(sf.Signature.Organization).WithName(sf.Signature.Name).WithTag(sf.Signature.Tag))
					end()
					if err != nil {
						return fmt.Errorf("failed to fetch signature %s/%s:%s: %w", sf.Signature.Organization, sf.Signature.Name, sf.Signature.Tag, err)
					}

					signatureSchemaData, err := base64.StdEncoding.DecodeString(res.GetPayload().Schema)
					if err != nil {
						return fmt.Errorf("failed to decode signature schema: %w", err)
					}
					signatureSchema = new(signature.Schema)
					err = signatureSchema.Decode(signatureSchemaData)
					if err != nil {
						return fmt.Errorf("failed to decode signature schema: %w", err)
					}
				}

				out := ch.Printer.Out()

				// Deal with extensions...
				extensionData := make([]extension.Info, 0)
				extensionSchemas := make([]*extension.Schema, 0)

				ets := storage.DefaultExtension
				if ch.Config.StorageDirectory != "" {
					ets, err = storage.NewExtension(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate extension storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				for _, e := range sf.Extensions {
					if e.Organization == "local" {
						extensionPath, err := ets.Path(e.Name, e.Tag, e.Organization, "")

						ext, err := ets.Get(e.Name, e.Tag, e.Organization, "")
						if err != nil {
							return fmt.Errorf("failed to get extension %s:%s: %w", e.Name, e.Tag, err)
						}

						extensionData = append(extensionData, extension.Info{
							Name:    e.Name,
							Path:    path.Join(extensionPath, "golang", "guest"),
							Version: "v0.1.0",
						})

						extensionSchemas = append(extensionSchemas, ext.Schema)

					} else {
						panic("Only local atm")
					}
				}

				// extensionData is setup for use in generating go.mod...
				ch.Printer.PrintProgress(fmt.Sprintf("Building scale function local/%s:%s...", sf.Name, sf.Tag))
				var scaleFunc *scalefunc.V1BetaSchema
				switch scalefunc.Language(sf.Language) {
				case scalefunc.Go:
					opts := &build.LocalGolangOptions{
						Stdout:           out,
						Scalefile:        sf,
						SourceDirectory:  sourceDir,
						SignatureSchema:  signatureSchema,
						Storage:          stb,
						Release:          release,
						Target:           build.WASITarget,
						GoBin:            goBin,
						TinyGoBin:        tinygoBin,
						Args:             tinygoArgs,
						Extensions:       extensionData,
						ExtensionSchemas: extensionSchemas,
					}
					scaleFunc, err = build.LocalGolang(opts)
				case scalefunc.Rust:
					opts := &build.LocalRustOptions{
						Stdout:          out,
						Scalefile:       sf,
						SourceDirectory: sourceDir,
						SignatureSchema: signatureSchema,
						Storage:         stb,
						Release:         release,
						Target:          buildTarget,
						CargoBin:        cargoBin,
						Args:            cargoArgs,
						//						Extensions:       extensionData,
						ExtensionSchemas: extensionSchemas,
					}
					scaleFunc, err = build.LocalRust(opts)
				case scalefunc.TypeScript:
					opts := &build.LocalTypescriptOptions{
						Stdout:          out,
						Scalefile:       sf,
						SourceDirectory: sourceDir,
						SignatureSchema: signatureSchema,
						Storage:         stb,
						Release:         release,
						Target:          buildTarget,
						NPMBin:          npmBin,
					}
					scaleFunc, err = build.LocalTypescript(opts)
				default:
					return fmt.Errorf("language %s is not supported for local builds", sf.Language)
				}
				if err != nil {
					return fmt.Errorf("failed to build scale function: %w", err)
				}

				fs := storage.DefaultFunction
				if ch.Config.StorageDirectory != "" {
					fs, err = storage.NewFunction(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				oldEntry, err := fs.Get(scaleFunc.Name, scaleFunc.Tag, "local", "")
				if err != nil {
					return fmt.Errorf("failed to check if scale function already exists: %w", err)
				}

				if oldEntry != nil {
					err = fs.Delete(name, tag, oldEntry.Organization, oldEntry.Hash)
					if err != nil {
						return fmt.Errorf("failed to delete existing scale function %s:%s: %w", name, tag, err)
					}
				}

				err = fs.Put(scaleFunc.Name, scaleFunc.Tag, "local", scaleFunc)
				if err != nil {
					return fmt.Errorf("failed to store scale function: %w", err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully built scale function %s\n", printer.BoldGreen(fmt.Sprintf("local/%s:%s", scaleFunc.Name, scaleFunc.Tag)))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"name":      name,
					"tag":       tag,
					"org":       "local",
					"directory": directory,
				})
			},
		}

		buildCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory containing the scalefile")
		buildCmd.Flags().StringVarP(&name, "name", "n", "", "the (optional) name of this scale function")
		buildCmd.Flags().StringVarP(&tag, "tag", "t", "", "the (optional) tag of this scale function")

		buildCmd.Flags().BoolVar(&release, "release", false, "build the function in release mode")
		buildCmd.Flags().StringVar(&target, "target", "wasi", "the compile target for the function")

		buildCmd.Flags().StringVar(&tinygoBin, "tinygo", "", "the (optional) path to the tinygo binary")
		buildCmd.Flags().StringVar(&goBin, "go", "", "the (optional) path to the go binary")
		buildCmd.Flags().StringVar(&cargoBin, "cargo", "", "the (optional) path to the cargo binary")
		buildCmd.Flags().StringVar(&npmBin, "npm", "", "the (optional) path to the npm binary")

		buildCmd.Flags().StringSliceVar(&tinygoArgs, "tinygo-args", []string{}, "list of (optional) tinygo build arguments")
		buildCmd.Flags().StringSliceVar(&cargoArgs, "cargo-args", []string{}, "list of (optional) cargo build arguments")

		cmd.AddCommand(buildCmd)
	}
}
