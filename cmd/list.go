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
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(0),
	Short: "list lists all the scale functions that are available",
	Long:  `List lists all the scale functions that are available.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := config.Init(cmd, false)
		scaleFuncEntries, err := storage.Default.List()
		if err != nil {
			logger.Fatal().Err(err).Msg("error listing scale functions")
		}
		for _, entry := range scaleFuncEntries {
			if entry.Tag != "" {
				logger.Info().Msgf("%s:%s", entry.ScaleFunc.ScaleFile.Name, entry.Tag)
			} else {
				logger.Info().Msgf("%s", entry.ScaleFunc.ScaleFile.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
