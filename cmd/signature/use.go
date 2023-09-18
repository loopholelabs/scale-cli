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
	"os"
	"path"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale"
	"github.com/loopholelabs/scale-cli/client/registry"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/compile/golang"
	"github.com/loopholelabs/scale/compile/rust"
	"github.com/loopholelabs/scale/compile/typescript"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
)

// UseCmd encapsulates the commands for using a Signature
func UseCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		useCmd := &cobra.Command{
			Use:      "use <org>/<name>:<tag> [flags]",
			Args:     cobra.ExactArgs(1),
			Short:    "create a new scale signature with the given name and tag",
			Hidden:   hidden,
			PreRunE:  utils.PreRunOptionalAuthenticatedAPI(ch),
			PostRunE: utils.PostRunOptionalAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				var err error
				parsed := scale.Parse(args[0])
				if parsed.Organization != "" && !scalefunc.ValidString(parsed.Organization) {
					return utils.InvalidStringError("organization name", parsed.Organization)
				}

				if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
					return utils.InvalidStringError("signature name", parsed.Name)
				}

				if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
					return utils.InvalidStringError("signature tag", parsed.Tag)
				}

				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}

				sf, err := scalefile.ReadSchema(path.Join(sourceDir, "scalefile"))
				if err != nil {
					return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}

				var signaturePath string
				var signatureVersion string
				if parsed.Organization == "local" {
					st := storage.DefaultSignature
					if ch.Config.StorageDirectory != "" {
						st, err = storage.NewSignature(ch.Config.StorageDirectory)
						if err != nil {
							return fmt.Errorf("failed to instantiate signature storage for %s: %w", ch.Config.StorageDirectory, err)
						}
					}

					signaturePath, err = st.Path(parsed.Name, parsed.Tag, parsed.Organization, "")
					if err != nil || signaturePath == "" {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
					switch scalefunc.Language(sf.Language) {
					case scalefunc.Go:
						signaturePath = path.Join(signaturePath, "golang", "guest")
					case scalefunc.Rust:
						signaturePath = path.Join(signaturePath, "rust", "guest")
					case scalefunc.TypeScript:
						signaturePath = path.Join(signaturePath, "typescript", "guest")
					default:
						return fmt.Errorf("failed to use signature %s/%s:%s: unknown or unsupported language", parsed.Organization, parsed.Name, parsed.Tag)
					}
				} else {
					ctx := cmd.Context()
					client := ch.Config.APIClient()

					end := ch.Printer.PrintProgress(fmt.Sprintf("Fetching signature %s/%s:%s...", parsed.Organization, parsed.Name, parsed.Tag))
					res, err := client.Registry.GetRegistrySignatureOrgNameTag(registry.NewGetRegistrySignatureOrgNameTagParamsWithContext(ctx).WithOrg(parsed.Organization).WithName(parsed.Name).WithTag(parsed.Tag))
					end()
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					switch scalefunc.Language(sf.Language) {
					case scalefunc.Go:
						signaturePath = res.GetPayload().GolangImportPathGuest
					case scalefunc.Rust:
						signaturePath = res.GetPayload().RustImportPathGuest
					case scalefunc.TypeScript:
						return fmt.Errorf("failed to use signature %s/%s:%s: typescript is not currently support via the registry", parsed.Organization, parsed.Name, parsed.Tag)
					default:
						return fmt.Errorf("failed to use signature %s/%s:%s: unknown or unsupported language", parsed.Organization, parsed.Name, parsed.Tag)
					}
					signatureVersion = "0.1.0"
				}

				switch scalefunc.Language(sf.Language) {
				case scalefunc.Go:
					modfileData, err := os.ReadFile(path.Join(sourceDir, "go.mod"))
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					m, err := golang.ParseManifest(modfileData)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = m.RemoveReplacement("signature", "v0.1.0")
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = m.RemoveReplacement("signature", "")
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = m.AddReplacement("signature", "v0.1.0", signaturePath, signatureVersion)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
					modfileData, err = m.Write()
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = os.WriteFile(path.Join(sourceDir, "go.mod"), modfileData, 0644)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
				case scalefunc.Rust:
					cargofileData, err := os.ReadFile(path.Join(sourceDir, "Cargo.toml"))
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					m, err := rust.ParseManifest(cargofileData)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = m.RemoveDependency("signature")
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					if parsed.Organization == "local" {
						err = m.AddDependencyWithPath("signature", rust.DependencyPath{
							Path:    signaturePath,
							Package: fmt.Sprintf("%s_%s_%s_guest", parsed.Organization, parsed.Name, parsed.Tag),
						})
						if err != nil {
							return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
						}
					} else {
						err = m.AddDependencyWithVersion("signature", rust.DependencyVersion{
							Version:  signatureVersion,
							Package:  fmt.Sprintf("%s_%s_%s_guest", parsed.Organization, parsed.Name, parsed.Tag),
							Registry: "scale",
						})
						if err != nil {
							return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
						}
					}

					cargofileData, err = m.Write()
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = os.WriteFile(path.Join(sourceDir, "Cargo.toml"), cargofileData, 0644)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
				case scalefunc.TypeScript:
					packageFileData, err := os.ReadFile(path.Join(sourceDir, "package.json"))
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					p, err := typescript.ParseManifest(packageFileData)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = p.RemoveDependency("signature")
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = p.AddDependency("signature", signaturePath)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
					packageFileData, err = p.Write()
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}

					err = os.WriteFile(path.Join(sourceDir, "package.json"), packageFileData, 0644)
					if err != nil {
						return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
					}
				default:
					return fmt.Errorf("failed to use signature %s/%s:%s: unknown or unsupported language", parsed.Organization, parsed.Name, parsed.Tag)
				}

				sf.Signature.Name = parsed.Name
				sf.Signature.Tag = parsed.Tag
				sf.Signature.Organization = parsed.Organization

				sfData, err := sf.Encode()
				if err != nil {
					return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}

				err = os.WriteFile(path.Join(sourceDir, "scalefile"), sfData, 0644)
				if err != nil {
					return fmt.Errorf("failed to use signature %s/%s:%s: %w", parsed.Organization, parsed.Name, parsed.Tag, err)
				}

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully using scale signature %s\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", parsed.Organization, parsed.Name, parsed.Tag)))
					return nil
				}

				return ch.Printer.PrintResource(map[string]string{
					"path": signaturePath,
					"name": parsed.Name,
					"org":  parsed.Organization,
					"tag":  parsed.Tag,
				})
			},
		}

		useCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory that contains the scalefile and the function source")

		cmd.AddCommand(useCmd)
	}
}
