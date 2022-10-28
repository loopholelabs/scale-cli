package apikey

import (
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

type apiKey struct {
	Name    string `header:"name" json:"name"`
	Created string `header:"created" json:"created"`
	ID      string `header:"id" json:"id"`
	Value   string `header:"value" json:"value"`
}

type apiKeyRedacted struct {
	Name    string `header:"name" json:"name"`
	Created string `header:"created" json:"created"`
	ID      string `header:"id" json:"id"`
}

// Cmd encapsulates the command for interacting with API Keys.
func Cmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "api-key <action>",
		Short:             "Create, list, and manage API Keys",
		PersistentPreRunE: cmdutil.CheckAuthentication(ch.Config),
	}

	cmd.AddCommand(CreateCmd(ch))
	cmd.AddCommand(ListCmd(ch))
	cmd.AddCommand(GetCmd(ch))
	cmd.AddCommand(DeleteCmd(ch))
	return cmd
}
