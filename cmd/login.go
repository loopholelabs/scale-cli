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
	"github.com/go-openapi/strfmt"
	authServer "github.com/loopholelabs/scale-cli/pkg/api/auth"
	"github.com/loopholelabs/scale-cli/pkg/api/client"
	"github.com/loopholelabs/scale-cli/pkg/api/client/auth"
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to your scale account",
	Long: `Login to your scale account using your browser 
or by providing an API key.

You can create API keys at https://app.scale.sh/account/api-keys`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := config.Init(cmd)

		defaultConfig := client.DefaultTransportConfig()
		api := viper.GetString("api")
		apiURL, err := url.Parse(api)
		if err != nil {
			logger.Fatal().Err(err).Msg("Invalid API URL")
		}
		defaultConfig.Schemes = []string{apiURL.Scheme}
		defaultConfig.Host = apiURL.Host
		client := client.NewHTTPClientWithConfig(strfmt.Default, defaultConfig)
		res, err := client.Auth.GetAuthGithub(auth.NewGetAuthGithubParams().WithRedirect("http://localhost:8085/auth/callback"))
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to get github auth url")
		}

		logger.Info().Msgf("Please visit %s to authenticate", res.Payload.AuthURL)
		as, err := authServer.Do("localhost:8085", logger)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to authenticate")
		}

		logger.Info().Msgf("Successfully authenticated %s", as.Username)
		viper.Set("auth", as)
		err = viper.WriteConfig()
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to write config")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringP("api-key", "a", "", "API key to use for authentication")
	err := viper.BindPFlag("api-key", loginCmd.PersistentFlags().Lookup("api-key"))
	if err != nil {
		panic(err)
	}
}
