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
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/loopholelabs/auth/pkg/client"
	"github.com/loopholelabs/scale-cli/internal/auth"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/build"
	apiClient "github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/loopholelabs/scale-cli/pkg/storage"
	"github.com/loopholelabs/scalefile"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func BuildCmd(ch *cmdutil.Helper) *cobra.Command {
	var scaleFilePath string
	var name string

	cmd := &cobra.Command{
		Use:      "build [flags]",
		Args:     cobra.ExactArgs(0),
		Short:    "build a scale function",
		PreRunE:  cmdutil.CheckAuthentication(ch.Config),
		PostRunE: cmdutil.UpdateToken(ch.Config),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			if scaleFilePath == "" {
				return errors.New("scalefile path is required")
			}

			scaleFile, err := scalefile.Read(scaleFilePath)
			if err != nil {
				return fmt.Errorf("failed to read scalefile: %w", err)
			}

			directory := path.Dir(scaleFilePath)
			sourcePath := path.Join(directory, scaleFile.Source)

			source, err := os.ReadFile(sourcePath)
			if err != nil {
				return fmt.Errorf("failed to read source file: %w", err)
			}

			expired, err := auth.Expired(ch.Config.Token.AccessToken)
			if err != nil {
				return fmt.Errorf("failed to check token expiration: %w", err)
			}

			if expired {
				ts, _, err := client.AuthenticatedClient(ch.Config.Endpoint, apiClient.DefaultBasePath, apiClient.DefaultSchemes, nil, path.Join(ch.Config.Token.Endpoint, ch.Config.Token.BasePath), ch.Config.Token.ClientID, ch.Config.Token.Kind, client.NewToken(ch.Config.Token.AccessToken, ch.Config.Token.TokenType, ch.Config.Token.RefreshToken, ch.Config.Token.Expiry))
				if err != nil {
					return fmt.Errorf("failed to create authenticated client: %w", err)
				}
				t, err := ts.Token()
				if err != nil {
					return fmt.Errorf("failed to get refreshed token: %w", err)
				}
				ch.Config.Token.AccessToken = t.AccessToken
				ch.Config.Token.TokenType = t.TokenType
				ch.Config.Token.RefreshToken = t.RefreshToken
				ch.Config.Token.Expiry = t.Expiry
			}

			// pending build service
			// scaleFunc, err := build.RemoteBuild(ctx, ch.Config.Build, name, source, ch.Config.Token.AccessToken, scaleFile, new(tls.Config), ch)
			scaleFunc, err := build.LocalBuild(ctx, name, source, scaleFile, ch)

			if err != nil {
				return err
			}
			if name != "" {
				scaleFunc.Name = name
			} else {
				name = scaleFunc.Name
			}
			err = storage.Default.Put(scaleFunc.Name, scaleFunc)
			if err != nil {
				return err
			}

			if ch.Printer.Format() == printer.Human {
				ch.Printer.Printf("Successfully built scale function %s\n", printer.BoldGreen(name))
				return nil
			}

			return ch.Printer.PrintResource(map[string]string{
				"Name": name,
			})
		},
	}

	cmd.Flags().StringVar(&ch.Config.Build, "build-service", "build.scale.sh:8192", "The endpoint for the Scale Build Service.")

	cmd.Flags().StringVarP(&scaleFilePath, "scalefile", "s", "scalefile", "the scalefile to use")
	cmd.Flags().StringVarP(&name, "name", "n", "", "the (optional) name of this scale function")

	return cmd
}
