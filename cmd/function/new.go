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
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/spf13/cobra"
	"os"
	textTemplate "text/template"
)

var (
	extensionLUT = map[string]string{
		"go": "go",
	}
)

func NewCmd(ch *cmdutil.Helper) *cobra.Command {
	var directory string

	cmd := &cobra.Command{
		Use:     "new <language> <name> [flags]",
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

			scaleFile := &scalefile.ScaleFile{
				Version:   scalefile.V1Alpha,
				Name:      name,
				Signature: "http",
				Language:  scalefile.Language(language),
				Dependencies: []scalefile.Dependency{
					{
						Name:    "github.com/loopholelabs/scale",
						Version: "v0.0.10",
					},
				},
				Extensions: nil,
				Source:     fmt.Sprintf("%s.%s", name, extension),
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

			tmpl, err := textTemplate.New("dependencies").Parse(template.GoTemplate)
			if err != nil {
				return fmt.Errorf("error parsing dependency template: %w", err)
			}

			switch language {
			case "go":
				dependencyFile, err := os.Create(fmt.Sprintf("%s/go.mod", directory))
				if err != nil {
					return fmt.Errorf("error creating dependencies file: %w", err)
				}
				err = tmpl.Execute(dependencyFile, scaleFile.Dependencies)
				if err != nil {
					_ = dependencyFile.Close()
					return fmt.Errorf("error writing dependencies file: %w", err)
				}
			default:
				return fmt.Errorf("language %s is not supported", language)
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully created new %s scale function %s\n", printer.BoldGreen(language), printer.BoldGreen(name))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name":     name,
				"Language": language,
			})
		},
	}

	cmd.Flags().StringVarP(&directory, "directory", "d", ".", "the directory to create the new scale function in")

	return cmd
}