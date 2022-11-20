/*
	Copyright 2022 Loophole Labs

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
	"errors"
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	remoteSignature "github.com/loopholelabs/scale-cli/internal/signature"
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/signature"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"os"
	"path"
	textTemplate "text/template"
)

var (
	ManifestLUT = map[scalefile.Language]string{
		scalefile.Go: "go.mod",
	}
)

func AddCmd(ch *cmdutil.Helper) *cobra.Command {
	var directory string
	var local string
	var scaleFilePath string
	var manifest string
	var language string

	cmd := &cobra.Command{
		Use:     "add <signature> [flags]",
		Args:    cobra.ExactArgs(1),
		Short:   "add a scale signature to an existing scale function",
		PreRunE: cmdutil.CheckAuthentication(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			signatureString := args[0]
			signatureNamespace, signatureName, signatureVersion := signature.ParseSignature(signatureString)

			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			var scaleFileLanguage scalefile.Language
			if manifest != "" {
				if language == "" {
					return errors.New("language must be specified when using the --manifest flag")
				}
				invalid := true
				for _, l := range scalefile.AcceptedLanguages {
					if string(l) == language {
						invalid = false
						break
					}
				}
				if invalid {
					return fmt.Errorf("language %s is not supported", language)
				}
				scaleFileLanguage = scalefile.Language(language)
			} else {
				sc, err := scalefile.Read(scaleFilePath)
				if err != nil {
					return fmt.Errorf("error reading scale file: %w", err)
				}
				scaleFileLanguage = sc.Language
				manifest = path.Join(path.Dir(scaleFilePath), ManifestLUT[scaleFileLanguage])
			}
			manifestData, err := os.ReadFile(manifest)
			if err != nil {
				return fmt.Errorf("error reading manifest file: %w", err)
			}

			scaleFile, err := scalefile.Read(scaleFilePath)
			if err != nil {
				return fmt.Errorf("error reading scale file: %w", err)
			}

			switch scaleFileLanguage {
			case scalefile.Go:
				sourcePath := modfile.ModulePath(manifestData)
				if sourcePath == "" {
					return errors.New("failed to find module path in go.mod")
				}
				if local != "" {
					sourcePath = path.Join(sourcePath, local, signatureName)
				} else {
					dependency, err := remoteSignature.GetRemoteGoSignature(client, ctx, signatureNamespace, signatureName, signatureVersion)
					if err != nil {
						return err
					}

					dependencyFile, err := os.Create(fmt.Sprintf("%s/go.mod", directory))
					if err != nil {
						return fmt.Errorf("error creating dependencies file: %w", err)
					}

					tmpl, err := textTemplate.New("dependencies").Parse(template.GoTemplate)
					if err != nil {
						return fmt.Errorf("error parsing dependency template: %w", err)
					}

					sourcePath = dependency.Name
					dependencies := make([]scalefile.Dependency, len(scaleFile.Dependencies)+1)
					copy(dependencies, scaleFile.Dependencies)
					dependencies[len(dependencies)-1] = *dependency
					err = tmpl.Execute(dependencyFile, dependencies)
					if err != nil {
						_ = dependencyFile.Close()
						return fmt.Errorf("error writing dependencies file: %w", err)
					}
				}

				err = signature.CreateGoSignature(scaleFilePath, directory, sourcePath)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("language %s is not supported", scaleFileLanguage)
			}

			if signatureNamespace != "" {
				scaleFile.Signature = fmt.Sprintf("%s/%s@%s", signatureNamespace, signatureName, signatureVersion)
			} else {
				scaleFile.Signature = fmt.Sprintf("%s@%s", signatureName, signatureVersion)
			}
			err = scalefile.Write(scaleFilePath, scaleFile)
			if err != nil {
				return fmt.Errorf("error writing scale file: %w", err)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully added scale signature %s\n", printer.BoldGreen(signatureName))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name":      signatureName,
				"Directory": directory,
				"Local":     local,
				"ScaleFile": scaleFilePath,
				"Manifest":  manifest,
				"Language":  string(scaleFileLanguage),
			})
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", "signature", "the directory to register the scale signature relative to the scalefile")
	cmd.Flags().StringVarP(&local, "local", "l", "", "the path to the local signature relative to the scalefile")
	cmd.Flags().StringVar(&scaleFilePath, "scalefile", "scalefile", "the path to the scale file for the scale function")
	cmd.Flags().StringVar(&manifest, "manifest", "", "the path to the manifest file for the guest language for the scale function")
	cmd.Flags().StringVar(&language, "language", "", "the language of the scale function if the --manifest flag is used (go, rust, etc.)")
	return cmd
}
