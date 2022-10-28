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

package function

import (
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	"github.com/spf13/cobra"
)

func ListCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "list compiled scale functions",
		RunE: func(cmd *cobra.Command, args []string) error {

			scaleFuncEntries, err := storage.Default.List()
			if err != nil {
				return fmt.Errorf("failed to list scale functions: %w", err)
			}

			if len(scaleFuncEntries) == 0 && ch.Printer.Format() == printer.Human {
				ch.Printer.Println("No Scale Functions have been compiled yet.")
				return nil
			}

			funcs := make([]scaleFunction, len(scaleFuncEntries))
			for i, entry := range scaleFuncEntries {
				funcs[i] = scaleFunction{
					Name:       entry.ScaleFile.Name,
					Tag:        entry.Tag,
					Language:   entry.ScaleFile.Build.Language,
					Middleware: entry.ScaleFile.Middleware,
				}
			}

			return ch.Printer.PrintResource(funcs)
		},
	}

	return cmd
}
