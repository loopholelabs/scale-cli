package function

import (
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/loopholelabs/scale/go/scalefile"
	"github.com/spf13/cobra"
	"os"
)

var (
	extensionLUT = map[string]string{
		"go": "go",
	}
)

func NewCmd(ch *cmdutil.Helper) *cobra.Command {
	var directory string
	var middleware bool

	cmd := &cobra.Command{
		Use:     "new <language> <name>",
		Args:    cobra.ExactArgs(2),
		Short:   "generate a new scale function with the given name and language",
		PreRunE: cmdutil.CheckAuthentication(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			language := args[0]
			name := args[1]

			extension, ok := extensionLUT[language]
			if !ok {
				return fmt.Errorf("language %s is not supported", language)
			}

			scaleFile := scalefile.ScaleFile{
				Version: "v1",
				Name:    name,
				Build: scalefile.Build{
					Language:     language,
					Dependencies: scalefile.DefaultDependencies,
				},
				Source:     fmt.Sprintf("%s.%s", name, extension),
				Middleware: middleware,
			}

			if _, err := os.Stat(directory); os.IsNotExist(err) {
				err = os.MkdirAll(directory, 0755)
				if err != nil {
					return fmt.Errorf("error creating directory %s: %w", directory, err)
				}
			}

			err := scalefile.Write(fmt.Sprintf("%s/scalefile", directory), scaleFile)
			if err != nil {
				return fmt.Errorf("error writing scalefile: %w", err)
			}

			err = os.WriteFile(fmt.Sprintf("%s/%s", directory, scaleFile.Source), template.LUT[language](), 0644)
			if err != nil {
				return fmt.Errorf("error writing source file: %w", err)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully created new %s function %s\n", printer.BoldGreen(language), printer.BoldGreen(name))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name":     name,
				"Language": language,
			})
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale function in")
	cmd.Flags().BoolVarP(&middleware, "middleware", "m", false, "create a middleware function")

	return cmd
}
