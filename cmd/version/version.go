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
