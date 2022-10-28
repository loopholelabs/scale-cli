package apikey

import (
	"github.com/loopholelabs/auth/pkg/utils"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client/access"
	"github.com/spf13/cobra"
)

func ListCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list API Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			end := ch.Printer.PrintProgress("Retrieving API Keys...")
			res, err := client.Access.GetAccessApikey(access.NewGetAccessApikeyParamsWithContext(ctx))
			end()
			if err != nil {
				return err
			}

			if len(res.Payload) == 0 && ch.Printer.Format() == printer.Human {
				ch.Printer.Println("No API Keys have been created yet.")
				return nil
			}

			keys := make([]apiKeyRedacted, len(res.Payload))
			for i, key := range res.Payload {
				keys[i] = apiKeyRedacted{
					Name:    key.Name,
					Created: utils.Int64ToTime(key.CreatedAt).String(),
					ID:      key.ID,
				}
			}
			return ch.Printer.PrintResource(keys)
		},
	}

	return cmd
}
