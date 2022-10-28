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
	"os"
	"strings"
)

func DeleteCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <name> [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "delete a compiled scale function",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			names := strings.Split(name, ":")
			if len(names) != 2 {
				name = fmt.Sprintf("%s:latest", name)
			}
			err := storage.Default.Delete(name)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("function %s does not exist", name)
				}
				return fmt.Errorf("failed to delete function %s: %w", name, err)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Scale Function %s %s\n", printer.BoldGreen(args[0]), printer.BoldRed("deleted"))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"deleted": args[0],
			})
		},
	}

	return cmd
}
