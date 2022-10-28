package apikey

import (
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client/access"
	"github.com/spf13/cobra"
)

func DeleteCmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <id>",
		Args:  cobra.ExactArgs(1),
		Short: "delete an API Key with the given ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client, err := ch.Client()
			if err != nil {
				return err
			}

			id := args[0]

			end := ch.Printer.PrintProgress(fmt.Sprintf("Deleting API Key %s...", id))
			res, err := client.Access.DeleteAccessApikeyID(access.NewDeleteAccessApikeyIDParamsWithContext(ctx).WithID(id))
			end()
			if err != nil {
				return err
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("API Key %s %s\n", printer.BoldGreen(res.Payload), printer.BoldRed("deleted"))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"deleted": id,
			})
		},
	}

	return cmd
}
