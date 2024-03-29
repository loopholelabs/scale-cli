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

package auth

import (
	"fmt"
	"github.com/go-openapi/runtime/client"
	"github.com/loopholelabs/auth/pkg/client/device"
	"github.com/loopholelabs/auth/pkg/client/session"
	"github.com/loopholelabs/auth/pkg/client/userinfo"
	"github.com/loopholelabs/auth/pkg/kind"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/cmdutils/pkg/printer"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/utils"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	ErrInteractive = errors.New("The 'login' command requires an interactive shell when the output format is not 'json'")
	ErrNoSession   = errors.New("No session found")
)

// LoginCmd encapsulates the commands for logging in
func LoginCmd(hidden bool) command.SetupCommand[*config.Config] {
	var apiKey string
	var organization string
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		loginCmd := &cobra.Command{
			Use:      "login [flags]",
			Short:    "Authenticate with the Scale Authentication API",
			Hidden:   hidden,
			PreRunE:  utils.PreRunUpdateCheck(ch),
			PostRunE: utils.PostRunAnalytics(ch),
			RunE: func(cmd *cobra.Command, args []string) error {
				if !printer.IsTTY {
					if ch.Printer.Format() == printer.Human {
						return ErrInteractive
					}
				}

				ctx := cmd.Context()

				c := ch.Config.NewUnauthenticatedAuthClient()

				var end func()
				if cmd.Flags().Changed("api-key") {
					end = ch.Printer.PrintProgress("Authenticating... (press Ctrl+C to cancel)")
					go func() {
						<-ctx.Done()
						if end != nil {
							end()
						}
						os.Exit(0)
					}()
					ch.Config.Session = session.New(kind.APIKey, apiKey, time.Time{})
				} else {
					flow, err := c.Device.PostDeviceFlow(device.NewPostDeviceFlowParamsWithContext(ctx))
					if err != nil {
						return fmt.Errorf("error getting device flow: %w", err)
					}

					browserURL := fmt.Sprintf("https://%s/device-auth?code=%s&organization=%s", ch.Config.UIEndpoint, flow.GetPayload().DeviceCode, organization)
					switch ch.Printer.Format() {
					case printer.Human:
						ch.Printer.Printf("\n%s%s\n", printer.Bold("Confirmation Code: "), printer.BoldGreen(flow.GetPayload().DeviceCode))
						ch.Printer.Printf("Opening browser to %s\n", printer.Bold(browserURL))
						err = browser.OpenURL(browserURL)
						if err != nil {
							ch.Printer.Printf("Failed to open browser: %s\n", err)
						}

						ch.Printer.Printf("\nIf something goes wrong, copy and paste this URL into your browser: %s\n\n", printer.Bold(browserURL))
						end = ch.Printer.PrintProgress("Waiting for confirmation... (press Ctrl+C to cancel)")
						go func() {
							<-ctx.Done()
							if end != nil {
								end()
							}
							os.Exit(0)
						}()
					case printer.JSON, printer.CSV:
						err = ch.Printer.PrintJSON(map[string]string{
							"code": flow.GetPayload().DeviceCode,
							"url":  browserURL,
						})
						if err != nil {
							return fmt.Errorf("error printing JSON: %w", err)
						}
					}

					ticker := time.NewTicker(time.Duration(flow.GetPayload().PollingRate)*time.Second + time.Millisecond*500)
					for {
						select {
						case <-ctx.Done():
							return fmt.Errorf("error while waiting for confirmation: %w", cmd.Context().Err())
						case <-ticker.C:
							_, err := c.Device.PostDevicePoll(device.NewPostDevicePollParamsWithContext(ctx).WithCode(flow.GetPayload().UserCode))
							if err != nil {
								if _, ok := err.(*device.PostDevicePollForbidden); ok {
									continue
								}
								return fmt.Errorf("error polling for confirmation: %w", err)
							}
							cookies := c.Transport.(*client.Runtime).Jar.Cookies(ch.Config.SessionCookieURL())
							if len(cookies) == 0 {
								return ErrNoSession
							}
							ch.Config.Session = session.New(kind.Session, cookies[0].Value, cookies[0].Expires)
							goto DONE
						}
					}
				}
			DONE:
				c, err := ch.Config.NewAuthenticatedAuthClient()
				if err != nil {
					return fmt.Errorf("error creating authenticated auth client: %w", err)
				}

				info, err := c.Userinfo.PostUserinfo(userinfo.NewPostUserinfoParamsWithContext(ctx))
				if err != nil {
					return fmt.Errorf("error getting user info: %w", err)
				}

				analytics.AssociateUser(info.GetPayload().Identifier, info.GetPayload().Organization)
				analytics.Event("login", map[string]string{
					"kind": string(info.GetPayload().Kind),
				})

				err = ch.Config.WriteSession()
				if err != nil {
					return fmt.Errorf("error writing session: %w", err)
				}

				if end != nil {
					end()
					end = nil
				}

				switch ch.Printer.Format() {
				case printer.JSON, printer.CSV:
					return ch.Printer.PrintJSON(map[string]string{
						"identifier":   info.GetPayload().Identifier,
						"kind":         string(info.GetPayload().Kind),
						"organization": info.GetPayload().Organization,
					})
				case printer.Human:
					ch.Printer.Printf("Logged in as %s (using organization %s)\n", printer.Bold(info.GetPayload().Identifier), printer.Bold(info.GetPayload().Organization))
				}
				return nil
			},
		}

		loginCmd.Flags().StringVarP(&apiKey, "api-key", "a", "", "The API Key to authenticate with the Scale API")
		loginCmd.Flags().StringVarP(&organization, "organization", "o", "", "The organization to authenticate with")
		cmd.AddCommand(loginCmd)
	}
}
