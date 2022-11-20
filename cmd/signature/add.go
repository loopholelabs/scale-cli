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
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/signature"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

var (
	ManifestLUT = map[scalefile.Language]string{
		scalefile.Go: "go.mod",
	}
)

func AddCmd(ch *cmdutil.Helper) *cobra.Command {
	var directory string
	var local string
	var scaleFile string
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

			var scalefileLanguage scalefile.Language
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
				scalefileLanguage = scalefile.Language(language)
			} else {
				sc, err := scalefile.Read(scaleFile)
				if err != nil {
					return fmt.Errorf("error reading scale file: %w", err)
				}
				scalefileLanguage = sc.Language
				manifest = path.Join(path.Dir(scaleFile), ManifestLUT[scalefileLanguage])
			}
			manifestData, err := os.ReadFile(manifest)
			if err != nil {
				return fmt.Errorf("error reading manifest file: %w", err)
			}

			switch scalefileLanguage {
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

					sourcePath = dependency.Name
				}

				err = signature.CreateGoSignature(scaleFile, directory, sourcePath)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("language %s is not supported", scalefileLanguage)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully added scale signature %s\n", printer.BoldGreen(signatureName))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name":      signatureName,
				"Directory": directory,
				"Local":     local,
				"ScaleFile": scaleFile,
				"Manifest":  manifest,
				"Language":  string(scalefileLanguage),
			})
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", "signature", "the directory to register the scale signature relative to the scalefile")
	cmd.Flags().StringVarP(&local, "local", "l", "", "the path to the local signature relative to the scalefile")
	cmd.Flags().StringVar(&scaleFile, "scalefile", "scalefile", "the path to the scale file for the scale function")
	cmd.Flags().StringVar(&manifest, "manifest", "", "the path to the manifest file for the guest language for the scale function")
	cmd.Flags().StringVar(&language, "language", "", "the language of the scale function if the --manifest flag is used (go, rust, etc.)")
	return cmd
}
