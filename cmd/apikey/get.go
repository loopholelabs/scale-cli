package apikey

import (
	"fmt"
	"github.com/loopholelabs/auth/pkg/utils"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/pkg/client/access"
	"github.com/spf13/cobra"
)

func GetCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <name>",
		Args:  cobra.ExactArgs(1),
		Short: "get information about an API Key with the given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			name := args[0]

			end := ch.Printer.PrintProgress(fmt.Sprintf("Retrieving API Key %s...", name))
			res, err := client.Access.GetAccessApikeyName(access.NewGetAccessApikeyNameParamsWithContext(ctx).WithName(name))
			end()
			if err != nil {
				return err
			}

			return ch.Printer.PrintResource(apiKeyRedacted{
				Name:    res.Payload.Name,
				Created: utils.Int64ToTime(res.Payload.CreatedAt).String(),
				ID:      res.Payload.ID,
			})
		},
	}

	return cmd
}
