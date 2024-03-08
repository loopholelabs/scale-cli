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
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/signature"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

// GenerateCmd encapsulates the commands for generating a Signature from a Signature File
func GenerateCmd(hidden bool) command.SetupCommand[*config.Config] {
	var directory string

	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		generateCmd := &cobra.Command{
			Use:      "generate <name>:<tag> [flags]",
			Args:     cobra.ExactArgs(1),
			Short:    "generate a scale signature from a signature file",
			Hidden:   hidden,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				sourceDir := directory
				if !path.IsAbs(sourceDir) {
					wd, err := os.Getwd()
					if err != nil {
						return fmt.Errorf("failed to get working directory: %w", err)
					}
					sourceDir = path.Join(wd, sourceDir)
				}

				signaturePath := path.Join(sourceDir, "scale.signature")
				signatureFile, err := signature.ReadSchema(signaturePath)
				if err != nil {
					return fmt.Errorf("failed to read signature file at %s: %w", signaturePath, err)
				}

				nametag := strings.Split(args[0], ":")
				if len(nametag) != 2 {
					return fmt.Errorf("invalid name or tag %s", args[0])
				}
				name := nametag[0]
				tag := nametag[1]

				if name == "" || !scalefunc.ValidString(name) {
					return utils.InvalidStringError("name", name)
				}

				if strings.Contains(name, "-") {
					return utils.InvalidStringError("signature name cannot have a '-' found in:", name)
				}

				if tag == "" || !scalefunc.ValidString(tag) {
					return utils.InvalidStringError("tag", tag)
				}

				end := ch.Printer.PrintProgress(fmt.Sprintf("Generating scale signature local/%s:%s...", name, tag))

				st := storage.DefaultSignature
				if ch.Config.StorageDirectory != "" {
					st, err = storage.NewSignature(ch.Config.StorageDirectory)
					if err != nil {
						end()
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				oldEntry, err := st.Get(name, tag, "local", "")
				if err != nil {
					end()
					return fmt.Errorf("failed to check if scale signature already exists: %w", err)
				}

				if oldEntry != nil {
					err = st.Delete(name, tag, oldEntry.Organization, oldEntry.Hash)
					if err != nil {
						end()
						return fmt.Errorf("failed to delete existing scale signature %s:%s: %w", name, tag, err)
					}
				}

				err = st.Put(name, tag, "local", signatureFile)
				if err != nil {
					end()
					return fmt.Errorf("failed to store scale signature: %w", err)
				}

				end()

				if ch.Printer.Format() == printer.Human {
					ch.Printer.Printf("Successfully generated scale signature %s\n", printer.BoldGreen(fmt.Sprintf("local/%s:%s", name, tag)))
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

		generateCmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory containing the signature file")

		cmd.AddCommand(generateCmd)
	}
}
