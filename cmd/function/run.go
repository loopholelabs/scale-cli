/*
	Copyright 2023 Loophole Labs

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
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	adapter "github.com/loopholelabs/scale-http-adapters/fasthttp"
	"github.com/loopholelabs/scale/go"
	"github.com/loopholelabs/scale/go/storage"
	"github.com/loopholelabs/scalefile/scalefunc"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// RunCmd encapsulates the commands for running Functions
func RunCmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		var listen string
		runCmd := &cobra.Command{
			Use:    "run [ ...[ <name>:<tag> ] | [ <org>/<name>:<tag> ] ] [flags]",
			Args:   cobra.MinimumNArgs(1),
			Short:  "run a compiled scale function",
			Long:   "Run a compiled scale function by starting an HTTP server that will listen for incoming requests and execute the specified functions in a chain. It's possible to specify multiple functions to be executed in a chain. The functions will be executed in the order they are specified. The scalefile must be in the current directory or specified with the --directory flag.",
			Hidden: hidden,
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.Default
				if ch.Config.CacheDirectory != "" {
					var err error
					st, err = storage.New(ch.Config.CacheDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.CacheDirectory, err)
					}
				}

				fns := make([]*scalefunc.ScaleFunc, 0, len(args))
				for _, f := range args {
					parsed := utils.ParseFunction(f)
					if parsed.Organization == "" {
						parsed.Organization = utils.DefaultOrganization
					}

					if parsed.Organization != "" && !scalefunc.ValidString(parsed.Organization) {
						return utils.InvalidStringError("organization name", parsed.Organization)
					}

					if parsed.Name == "" || !scalefunc.ValidString(parsed.Name) {
						return utils.InvalidStringError("function name", parsed.Name)
					}

					if parsed.Tag == "" || !scalefunc.ValidString(parsed.Tag) {
						return utils.InvalidStringError("function tag", parsed.Tag)
					}

					e, err := st.Get(parsed.Name, parsed.Tag, parsed.Organization, "")
					if err != nil {
						return fmt.Errorf("failed to get function %s: %w", f, err)
					}

					if e == nil {
						return fmt.Errorf("function %s not found", f)
					}

					fns = append(fns, e.ScaleFunc)
				}

				ctx := cmd.Context()
				r, err := runtime.New(ctx, fns)
				if err != nil {
					return fmt.Errorf("failed to create runtime: %w", err)
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
					ch.Printer.Printf("scale functions %s listening at %s", printer.BoldGreen(args), printer.BoldGreen(listen))
					err = server.ListenAndServe(listen)
					if err != nil {
						ch.Printer.Printf("error starting server: %v", printer.BoldRed(err))
					}
				}()
				<-stop
				err = server.Shutdown()
				if err != nil {
					return fmt.Errorf("failed to shutdown server: %w", err)
				}
				wg.Wait()
				return nil
			},
		}

		runCmd.Flags().StringVarP(&listen, "listen", "l", ":8080", "the address to listen on")
		cmd.AddCommand(runCmd)
	}
}
