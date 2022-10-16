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
	"context"
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	adapter "github.com/loopholelabs/scale-go/adapters/http"
	"github.com/loopholelabs/scale-go/runtime"
	"github.com/loopholelabs/scale-go/scalefunc"
	"github.com/spf13/cobra"
	"net/http"
)

var runCmd = &cobra.Command{
	Use:   "run [function name]",
	Args:  cobra.ExactArgs(1),
	Short: "run a scale function that has already been built",
	Long: `Run a scale function locally that has already been built. This command can
also be used to temporarily expose your scale function to the internet using lynk.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := config.Init(cmd, false)
		name := args[0]
		tag := cmd.Flag("tag").Value.String()
		logger.Debug().Msgf("run called with name '%s' and tag '%s'", name, tag)
		var scaleFunc *scalefunc.ScaleFunc
		var err error
		if tag == "" {
			scaleFunc, err = storage.Default.Get(name)
			if err != nil {
				logger.Fatal().Err(err).Msgf("error getting scale function '%s'", name)
			}
		} else {
			scaleFunc, err = storage.Default.Get(name, tag)
			if err != nil {
				logger.Fatal().Err(err).Msgf("error getting scale function '%s' with tag '%s'", name, tag)
			}
		}
		listen := cmd.Flag("listen").Value.String()
		if listen == "" {
			logger.Fatal().Msg("listen address must be specified")
		}
		r, err := runtime.New(context.Background(), []scalefunc.ScaleFunc{*scaleFunc})
		if err != nil {
			logger.Fatal().Err(err).Msg("error creating runtime")
		}

		logger.Info().Msgf("scale function %s listening on %s", name, listen)
		err = http.ListenAndServe(listen, adapter.New(nil, r))
		if err != nil {
			logger.Fatal().Err(err).Msg("error starting http server")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("tag", "t", "", "the (optional) tag to use for this module")
	runCmd.Flags().StringP("listen", "l", "127.0.0.1:8080", "the address the scale function should listen on")
}
