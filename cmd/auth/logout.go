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
	"bufio"
	"io"
	"os"

	"github.com/loopholelabs/scale-cli/internal/auth"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/printer"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func LogoutCmd(ch *cmdutil.Helper) *cobra.Command {
	var clientID string
	var apiURL string

	cmd := &cobra.Command{
		Use:   "logout",
		Args:  cobra.NoArgs,
		Short: "Log out of the Scale API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if ch.Config.IsAuthenticated() != nil {
				ch.Printer.Println("Already logged out. Exiting...")
				return nil
			}

			if printer.IsTTY {
				ch.Printer.Println("Press Enter to log out of the Scale API.")
				_ = waitForEnter(cmd.InOrStdin())
			}

			end := ch.Printer.PrintProgress("Logging out...")
			defer end()

			err := deleteToken()
			if err != nil {
				return err
			}
			end()
			ch.Printer.Println("Successfully logged out.")

			return nil
		},
	}

	cmd.Flags().StringVar(&clientID, "client-id", auth.OAuthClientID, "The client ID for the Scale CLI application.")
	cmd.Flags().StringVar(&apiURL, "api", auth.DefaultEndpoint, "The Scale API URL.")
	return cmd
}

func deleteToken() error {
	tokenPath, err := config.TokenPath()
	if err != nil {
		return err
	}

	err = os.Remove(tokenPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "error removing access token file")
		}
	}

	configFile, err := config.DefaultConfigPath()
	if err != nil {
		return err
	}

	err = os.Remove(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "error removing default config file")
		}
	}

	return nil
}

func waitForEnter(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Err()
}
