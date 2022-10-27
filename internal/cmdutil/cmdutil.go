package cmdutil

import (
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/spf13/cobra"
	exec "golang.org/x/sys/execabs"
	"path/filepath"
	"strings"
)

const WarnAuthMessage = "not authenticated yet. Please run 'scale auth login'"

// Helper is passed to every single command and is used by individual
// subcommands.
type Helper struct {
	// Config contains globally sourced configuration
	Config *config.Config

	// Client returns the Scale API client
	Client func() (*client.ScaleAPIV1, error)

	// Printer is used to print output of a command to stdout.
	Printer *printer.Printer

	// bebug defines the debug mode
	debug *bool
}

func (h *Helper) SetDebug(debug *bool) {
	h.debug = debug
}

func (h *Helper) Debug() bool { return *h.debug }

// RequiredArgs - required arguments are not available.
func RequiredArgs(reqArgs ...string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		n := len(reqArgs)
		if len(args) >= n {
			return nil
		}

		missing := reqArgs[len(args):]

		a := fmt.Sprintf("arguments <%s>", strings.Join(missing, ", "))
		if len(missing) == 1 {
			a = fmt.Sprintf("argument <%s>", missing[0])
		}

		return fmt.Errorf("missing %s \n\n%s", a, cmd.UsageString())
	}
}

// CheckAuthentication checks whether the user is authenticated and returns an
// actionable error message.
func CheckAuthentication(cfg *config.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := cfg.IsAuthenticated(); err != nil {
			return fmt.Errorf("%s\nError: %s", WarnAuthMessage, err.Error())
		}

		return nil
	}
}

// IsUnderHomebrew checks whether the given binary is under the homebrew path.
// copied from: https://github.com/cli/cli/blob/trunk/cmd/gh/main.go#L298
func IsUnderHomebrew(binpath string) bool {
	if binpath == "" {
		return false
	}

	brewExe, err := exec.LookPath("brew")
	if err != nil {
		return false
	}

	brewPrefixBytes, err := exec.Command(brewExe, "--prefix").Output()
	if err != nil {
		return false
	}

	brewBinPrefix := filepath.Join(strings.TrimSpace(string(brewPrefixBytes)), "bin") + string(filepath.Separator)
	return strings.HasPrefix(binpath, brewBinPrefix)
}

// HasHomebrew check whether the user has installed brew
func HasHomebrew() bool {
	_, err := exec.LookPath("brew")
	return err == nil
}
