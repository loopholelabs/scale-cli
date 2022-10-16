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
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/loopholelabs/scale-cli/internal/token"
	"github.com/loopholelabs/scale-cli/pkg/build"
	"github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/loopholelabs/scale-cli/pkg/client/auth"
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/loopholelabs/scale-go/scalefile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path"
	"time"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build a scale function from a scalefile using the scale build service",
	Long: `Build a scale function from a scalefile using the scale build service. 

The scalefile is a YAML file that describes the scale function and its dependencies. 
The scale build service will build the scale function and return the compiled module.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := config.Init(cmd)

		t := viper.GetStringMapString("auth")
		if t["access_token"] == "" {
			logger.Fatal().Msg("You must be logged in to use the scale build service. Please run `scale login` to login.")
		}

		expired, err := token.Expired(t["access_token"])
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to check if access token is expired")
		}
		if expired {
			logger.Debug().Msg("access token is expired, refreshing")
			defaultConfig := client.DefaultTransportConfig()
			api := viper.GetString("api")
			apiURL, err := url.Parse(api)
			if err != nil {
				logger.Fatal().Err(err).Msg("Invalid API URL")
			}
			defaultConfig.Schemes = []string{apiURL.Scheme}
			defaultConfig.Host = apiURL.Host
			client := client.NewHTTPClientWithConfig(strfmt.Default, defaultConfig)
			res, err := client.Auth.PostAuthRefresh(auth.NewPostAuthRefreshParams().WithGrantType("refresh_token").WithRefreshToken(t["refresh_token"]))
			if err != nil {
				logger.Fatal().Err(err).Msg("You must be logged in to use the scale build service. If you were already logged in, your access token may have expired. Please run `scale login` again.")
			}
			t["refresh_token"] = res.Payload.RefreshToken
			t["access_token"] = res.Payload.AccessToken
			t["token_type"] = res.Payload.TokenType
			err = viper.WriteConfig()
			if err != nil {
				logger.Fatal().Err(err).Msg("failed to update config")
			}
		}

		scalefilePath := cmd.Flag("scalefile").Value.String()
		if scalefilePath == "" {
			logger.Fatal().Msg("scalefile path is required")
		}
		logger.Debug().Msgf("build called with scalefile '%s'", scalefilePath)
		scaleFile, err := scalefile.Read(scalefilePath)
		if err != nil {
			logger.Fatal().Err(err).Msg("error reading scalefile")
		}
		directory := path.Dir(scalefilePath)
		inputPath := path.Join(directory, scaleFile.File)

		start := time.Now()
		input, err := os.ReadFile(inputPath)
		if err != nil {
			logger.Fatal().Err(err).Msgf("error while reading scale function file %s", inputPath)
		}
		logger.Debug().Msgf("read scale function file %s in %s", inputPath, time.Since(start))

		outputPath := fmt.Sprintf("%s.wasm", path.Join(directory, scaleFile.Name))

		build.Build(input, outputPath, t["access_token"], scaleFile, new(tls.Config), logger)

		logger.Info().Msg("Module Compilation Completed")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().StringP("builder", "b", "build.scale.sh:8192", "Scale Build Service URL")
	buildCmd.Flags().StringP("scalefile", "s", "scalefile", "the scalefile to use")

	err := viper.BindPFlag("builder", buildCmd.PersistentFlags().Lookup("builder"))
	if err != nil {
		panic(err)
	}
}
