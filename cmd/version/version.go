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

package version

import (
	"fmt"
	"strings"

	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"

	"github.com/spf13/cobra"
)

// Cmd encapsulates the commands for showing a version
func Cmd(ch *cmdutil.Helper, ver, commit, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version <command>",
		// we can also show the version via `--version`, hence this doesn't
		// need to be displayed.
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if ch.Printer.Format() == printer.Human {
				ch.Printer.Println(Format(ver, commit, buildDate))
				return nil
			}

			v := map[string]string{
				"version":    ver,
				"commit":     commit,
				"build_date": buildDate,
			}
			return ch.Printer.PrintResource(v)
		},
	}

	return cmd
}

// Format formats a version string with the given information.
func Format(ver, commit, buildDate string) string {
	if ver == "" && buildDate == "" && commit == "" {
		return "scale cli version (built from source)"
	}

	ver = strings.TrimPrefix(ver, "v")

	return fmt.Sprintf("scale cli version %s (build date: %s commit: %s)\n", ver, buildDate, commit)
}
