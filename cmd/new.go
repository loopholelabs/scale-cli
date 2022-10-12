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

package cmd

import (
	"fmt"
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/loopholelabs/scale-cli/pkg/scalefile"
	"github.com/loopholelabs/scale-cli/pkg/template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	languageLUT = map[string]string{
		"go": "go",
		//"js":   "js",
		//"rust": "rs",
	}
)

var newCmd = &cobra.Command{
	Use:   "new [language] [name]",
	Short: "new generates a scalefile for a scale function with the given name and language",
	Long:  `New generates a scalefile for a scale function with the given name and language.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := config.Init(cmd)
		language := args[0]
		name := args[1]

		logger.Debug().Msgf("new called with language '%s' and name '%s'", language, name)
		extension, ok := languageLUT[language]
		if !ok {
			logger.Fatal().Msgf("language '%s' is not supported", language)
		}

		scaleFile := scalefile.ScaleFile{
			Name: name,
			Builder: scalefile.Build{
				Language:     language,
				Dependencies: nil,
			},
			File: fmt.Sprintf("%s.%s", name, extension),
		}

		directory := viper.GetString("directory")
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			err = os.MkdirAll(directory, 0755)
			if err != nil {
				logger.Fatal().Err(err).Msgf("error creating directory '%s'", directory)
			}
		}

		err := scalefile.Write(fmt.Sprintf("%s/scalefile", directory), scaleFile)
		if err != nil {
			logger.Fatal().Err(err).Msg("error writing scalefile")
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s", directory, scaleFile.File), template.LUT[language](), 0644)
		if err != nil {
			logger.Fatal().Err(err).Msg("error writing scale function template")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("directory", "d", ".", "the directory to create the scale function in")
	err := viper.BindPFlag("directory", newCmd.Flags().Lookup("directory"))
	if err != nil {
		panic(err)
	}
}
