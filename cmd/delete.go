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
	"github.com/loopholelabs/scale-cli/pkg/storage"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [scale function name]",
	Args:  cobra.ExactArgs(1),
	Short: "delete deletes a specific scale function",
	Long:  `Delete deletes a specific scale function given the name.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := config.Init(cmd, false)
		name := args[0]
		names := strings.Split(name, ":")
		if len(names) != 2 {
			name = fmt.Sprintf("%s:latest", name)
		}
		err := storage.Default.Delete(name)
		if err != nil {
			if os.IsNotExist(err) {
				logger.Fatal().Msgf("scale function '%s' does not exist", name)
			}
			logger.Fatal().Err(err).Msg("error deleting scale function")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
