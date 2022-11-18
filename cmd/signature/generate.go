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
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func GenerateCmd(ch *cmdutil.Helper) *cobra.Command {
	var directory string
	var protoc string

	cmd := &cobra.Command{
		Use:     "generate <name> [flags]",
		Args:    cobra.ExactArgs(1),
		Short:   "regenerate a scale signature from its modified proto file",
		PreRunE: cmdutil.CheckAuthentication(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if protoc == "" {
				protoc, err = exec.LookPath("protoc")
				if err != nil {
					return errors.New("failed to find protoc binary in $PATH, you can specify it manually using the --protoc flag")
				}
			}

			name := args[0]

			scaleBinary, err := os.Executable()
			if err != nil {
				return fmt.Errorf("error finding scale binary: %w", err)
			}

			goProtocCmd := exec.Command(protoc, "--plugin", fmt.Sprintf("protoc-gen-go-scale-signature=%s", scaleBinary), fmt.Sprintf("--go-scale-signature_out=%s/%s", directory, name), fmt.Sprintf("%s/%s/signature.proto", directory, name))
			goProtocCmd.Env = append(os.Environ(), "SCALE_PROTOC=true")

			err = goProtocCmd.Run()
			if err != nil {
				return fmt.Errorf("error generating go code from proto file: %w", err)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully regenerated scale signature %s\n", printer.BoldGreen(name))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name":      name,
				"Directory": directory,
			})
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", "signature", "the directory to create the new scale signature in")
	cmd.Flags().StringVar(&protoc, "protoc", "", "the path to the protoc binary")
	return cmd
}
