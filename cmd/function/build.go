package function

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/build"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	"github.com/loopholelabs/scale/go/scalefile"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

func BuildCmd(ch *cmdutil.Helper) *cobra.Command {
	var scaleFilePath string
	var name string

	cmd := &cobra.Command{
		Use:     "build [flags]",
		Args:    cobra.ExactArgs(0),
		Short:   "build a scale function",
		PreRunE: cmdutil.CheckAuthentication(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if scaleFilePath == "" {
				return errors.New("scalefile path is required")
			}

			scaleFile, err := scalefile.Read(scaleFilePath)
			if err != nil {
				return err
			}

			directory := path.Dir(scaleFilePath)
			sourcePath := path.Join(directory, scaleFile.Source)

			source, err := os.ReadFile(sourcePath)
			if err != nil {
				return err
			}

			scaleFunc, err := build.Build(ctx, ch.Config.Build, source, ch.Config.Token.AccessToken, scaleFile, new(tls.Config), ch)
			if err != nil {
				return err
			}
			if name != "" {
				scaleFunc.ScaleFile.Name = name
				names := strings.Split(name, ":")
				var fname string
				if len(names) == 2 {
					scaleFunc.ScaleFile.Name = names[0]
					scaleFunc.Tag = names[1]
					fname = fmt.Sprintf("%s:%s", scaleFunc.ScaleFile.Name, scaleFunc.Tag)
				} else {
					scaleFunc.Tag = "latest"
					fname = fmt.Sprintf("%s:%s", name, scaleFunc.Tag)
				}
				err = storage.Default.Put(fname, scaleFunc)
			} else {
				name = scaleFunc.ScaleFile.Name
				err = storage.Default.Put(fmt.Sprintf("%s:%s", name, scaleFunc.Tag), scaleFunc)
			}
			if err != nil {
				return err
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully built function %s\n", printer.BoldGreen(name))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name": name,
				"Tag":  scaleFunc.Tag,
			})
		},
	}

	cmd.Flags().StringVar(&ch.Config.Build, "build-service", "build.scale.sh:8192", "The endpoint for the Scale Build Service.")

	cmd.Flags().StringVarP(&scaleFilePath, "scalefile", "s", "scalefile", "the scalefile to use")
	cmd.Flags().StringVarP(&name, "name", "n", "", "the (optional) name of this scale function")

	return cmd
}
