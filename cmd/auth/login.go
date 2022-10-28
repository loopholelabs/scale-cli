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

package auth

import (
	"encoding/json"
	"fmt"
	"github.com/loopholelabs/auth/pkg/client"
	"github.com/loopholelabs/auth/pkg/client/discover"
	"github.com/loopholelabs/auth/pkg/token/tokenKind"
	"github.com/loopholelabs/scale-cli/internal/auth"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// LoginCmd is the command for logging into a Scale account.
func LoginCmd(ch *cmdutil.Helper) *cobra.Command {
	var clientID string
	var authEndpoint string
	var authBasePath string

	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.ExactArgs(0),
		Short: "Authenticate with the Scale API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !printer.IsTTY {
				return errors.New("The 'login' command requires an interactive shell")
			}

			_, c := client.UnauthenticatedClient(authEndpoint, authBasePath, auth.DefaultScheme, nil)

			d, err := discover.Discover(c.Transport, fmt.Sprintf("https://%s", path.Join(authEndpoint, authBasePath)))
			if err != nil {
				return fmt.Errorf("error discovering auth endpoint: %w", err)
			}

			var end func()
			go func() {
				<-cmd.Context().Done()
				if end != nil {
					end()
				}
				os.Exit(0)
			}()

			flow := client.DeviceFlow(d.GetHosts(), client.NewCompatibleClient(c.Transport), d.GetScopes(), clientID, func(userCode string, verificationURI string) error {
				bold := color.New(color.Bold)
				_, _ = bold.Printf("\nConfirmation Code: ")
				boldGreen := bold.Add(color.FgGreen)
				_, _ = boldGreen.Fprintln(color.Output, userCode)

				ch.Printer.Printf("\nIf something goes wrong, copy and paste this URL into your browser: %s\n\n", printer.Bold(verificationURI))
				end = ch.Printer.PrintProgress("Waiting for confirmation...")
				return nil
			}, nil)

			token, err := client.GetToken(flow)
			if end != nil {
				end()
			}
			if err != nil {
				return fmt.Errorf("error getting token: %w", err)
			}

			err = writeToken(config.FromClientToken(token, tokenKind.OAuthKind, authEndpoint, authBasePath, clientID))
			if err != nil {
				return errors.Wrap(err, "error logging in")
			}

			// We explicitly stop here so we can replace the spinner with our success
			// message.
			end()
			ch.Printer.Println("Successfully logged in.")

			return nil
		},
	}

	cmd.Flags().StringVar(&clientID, "client-id", auth.OAuthClientID, "The client ID for the Scale Auth Service.")
	cmd.Flags().StringVar(&authEndpoint, "auth", auth.DefaultEndpoint, "The Scale Auth Service Endpoint")
	cmd.Flags().StringVar(&authBasePath, "base-path", auth.DefaultBasePath, "The Scale Auth Service Base Path")

	return cmd
}

func writeToken(token *config.Token) error {
	configDir, err := config.Dir()
	if err != nil {
		return err
	}

	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(configDir, 0771)
		if err != nil {
			return errors.Wrap(err, "error creating config directory")
		}
	} else if err != nil {
		return err
	}

	tokenPath, err := config.TokenPath()
	if err != nil {
		return err
	}

	tokenData, err := json.Marshal(token)
	if err != nil {
		return err
	}

	err = os.WriteFile(tokenPath, tokenData, config.TokenFileMode)
	if err != nil {
		return errors.Wrap(err, "error writing token")
	}

	return nil
}
