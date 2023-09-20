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
	"github.com/loopholelabs/scale"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/client/registry"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/loopholelabs/scale/signature/converter"
	"github.com/loopholelabs/scale/storage"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
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
			Use:      "run [ ...[ <org>/<name>:<tag> ] ] [flags]",
			Args:     cobra.MinimumNArgs(1),
			Short:    "run a compiled scale function",
			Long:     "Run a compiled scale function by starting an HTTP server that will listen for incoming requests and execute the specified functions in a chain. It's possible to specify multiple functions to be executed in a chain. The functions will be executed in the order they are specified.",
			Hidden:   hidden,
			PreRunE:  utils.PreRunOptionalAuthenticatedAPI(ch),
			PostRunE: utils.PostRunAuthenticatedAPI(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				st := storage.DefaultFunction
				if ch.Config.StorageDirectory != "" {
					var err error
					st, err = storage.NewFunction(ch.Config.StorageDirectory)
					if err != nil {
						return fmt.Errorf("failed to instantiate function storage for %s: %w", ch.Config.StorageDirectory, err)
					}
				}

				fns := make([]*scalefunc.Schema, 0, len(args))
				for _, f := range args {
					parsed := scale.Parse(f)
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
						analytics.Event("pull-registry")

						ctx := cmd.Context()
						client := ch.Config.APIClient()

						end := ch.Printer.PrintProgress(fmt.Sprintf("Function %s was not found not found, pulling from the registry...", printer.BoldGreen(f)))

						res, err := client.Registry.GetRegistryFunctionOrgNameTag(registry.NewGetRegistryFunctionOrgNameTagParamsWithContext(ctx).WithOrg(parsed.Organization).WithName(parsed.Name).WithTag(parsed.Tag))
						if err != nil {
							end()
							return fmt.Errorf("failed to get function %s: %w", f, err)
						}

						resp, err := http.DefaultClient.Get(res.GetPayload().PresignedURL)
						if err != nil {
							end()
							return fmt.Errorf("failed to download function: %w", err)
						}

						data, err := io.ReadAll(resp.Body)
						if err != nil {
							end()
							return fmt.Errorf("failed to read response body: %w", err)
						}

						err = resp.Body.Close()
						if err != nil {
							end()
							return fmt.Errorf("failed to close response body: %w", err)
						}

						s := new(scalefunc.Schema)
						err = s.Decode(data)
						if err != nil {
							end()
							return fmt.Errorf("failed to decode function: %w", err)
						}

						err = st.Put(res.GetPayload().Function.Name, res.GetPayload().Function.Tag, res.GetPayload().Function.Organization, s)
						if err != nil {
							end()
							return fmt.Errorf("failed to save function: %w", err)
						}

						if ch.Printer.Format() == printer.Human {
							ch.Printer.Printf("Pulled %s from the Scale Registry\n", printer.BoldGreen(fmt.Sprintf("%s/%s:%s", parsed.Organization, s.Name, s.Tag)))
						}
						fns = append(fns, s)
					} else {
						fns = append(fns, e.Schema)
					}
				}

				analytics.Event("run-function", map[string]string{"chain-size": fmt.Sprintf("%d", len(fns))})

				ctx := cmd.Context()
				typecheckSignature, err := converter.NewSignature(fns[0].SignatureSchema)
				if err != nil {
					return fmt.Errorf("failed to create type check signature: %w", err)
				}

				writer := ch.Printer.Out()
				s, err := scale.New(scale.NewConfig(typecheckSignature.Signature).WithContext(ctx).WithFunctions(fns).WithStdout(writer).WithStderr(writer))
				if err != nil {
					return fmt.Errorf("failed to create scale: %w", err)
				}

				stop := make(chan os.Signal, 1)
				signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

				server := fasthttp.Server{
					Handler: func(ctx *fasthttp.RequestCtx) {
						if !ctx.IsPost() {
							ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
							return
						}
						instance, err := s.Instance()
						if err != nil {
							ctx.Error(fmt.Sprintf("Failed to create instance: %v", err), fasthttp.StatusInternalServerError)
							return
						}
						sig, err := converter.NewSignature(fns[0].SignatureSchema)
						if err != nil {
							ctx.Error(fmt.Sprintf("Failed to create signature: %v", err), fasthttp.StatusInternalServerError)
							return
						}
						err = sig.FromJSON(ctx.PostBody())
						if err != nil {
							ctx.Error(fmt.Sprintf("Failed to parse signature: %v", err), fasthttp.StatusInternalServerError)
							return
						}
						err = instance.Run(ctx, sig)
						if err != nil {
							ctx.Error(fmt.Sprintf("Failed to run function: %v", err), fasthttp.StatusInternalServerError)
							return
						}

						body, err := sig.ToJSON()
						if err != nil {
							ctx.Error(fmt.Sprintf("Failed to encode signature: %v", err), fasthttp.StatusInternalServerError)
							return
						}

						ctx.SetContentType("application/json")
						ctx.SetStatusCode(fasthttp.StatusOK)
						ctx.SetBody(body)
					},
					CloseOnShutdown: true,
					IdleTimeout:     time.Second,
				}

				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					ch.Printer.Printf("Scale Functions %s listening at %s", printer.BoldGreen(args), printer.BoldGreen(listen))
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
