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
	"fmt"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	adapter "github.com/loopholelabs/scale-go/adapters/fasthttp"
	"github.com/loopholelabs/scale-go/runtime"
	"github.com/loopholelabs/scale-go/scalefunc"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
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
		logger.Debug().Msgf("run called with name '%s'", name)
		names := strings.Split(name, ":")
		if len(names) != 2 {
			name = fmt.Sprintf("%s:latest", name)
		}
		scaleFunc, err := storage.Default.Get(name)
		if err != nil {
			logger.Fatal().Err(err).Msgf("error getting scale function '%s'", name)
		}
		listen := cmd.Flag("listen").Value.String()
		if listen == "" {
			logger.Fatal().Msg("listen address must be specified")
		}
		r, err := runtime.New(context.Background(), []scalefunc.ScaleFunc{*scaleFunc})
		if err != nil {
			logger.Fatal().Err(err).Msg("error creating runtime")
		}

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		server := fasthttp.Server{
			Handler:         adapter.New(nil, r).Handle,
			CloseOnShutdown: true,
			IdleTimeout:     time.Second,
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Info().Msgf("scale function %s listening on %s", name, listen)
			err = server.ListenAndServe(listen)
			if err != nil {
				logger.Fatal().Err(err).Msg("error starting server")
			}
		}()
		<-stop
		logger.Debug().Msg("shutting down server")
		err = server.Shutdown()
		if err != nil {
			logger.Fatal().Err(err).Msg("error shutting down server")
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("listen", "l", "127.0.0.1:8080", "the address the scale function should listen on")
}
