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
	"github.com/loopholelabs/scale-cli/pkg/storage"
	adapter "github.com/loopholelabs/scale/go/adapters/fasthttp"
	"github.com/loopholelabs/scale/go/runtime"
	"github.com/loopholelabs/scale/go/scalefunc"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func RunCmd(ch *cmdutil.Helper) *cobra.Command {
	var listen string
	cmd := &cobra.Command{
		Use:   "run <function> [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "run a compiled scale function",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			name := args[0]
			names := strings.Split(name, ":")
			if len(names) != 2 {
				name = fmt.Sprintf("%s:latest", name)
			}
			scaleFunc, err := storage.Default.Get(name)
			if err != nil {
				return fmt.Errorf("failed to get function %s: %w", name, err)
			}
			r, err := runtime.New(ctx, []scalefunc.ScaleFunc{*scaleFunc})
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
				ch.Printer.Printf("scale function %s listening at %s", printer.BoldGreen(name), printer.BoldGreen(listen))
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

	cmd.Flags().StringVarP(&listen, "listen", "l", "127.0.0.1:8080", "the address the scale function should listen on")

	return cmd
}
