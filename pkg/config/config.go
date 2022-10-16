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

package config

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/loopholelabs/scale-cli/internal/token"
	"github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/loopholelabs/scale-cli/pkg/client/auth"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path"
)

func Init(cmd *cobra.Command, isAuth bool) (zerolog.Logger, map[string]string) {
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
	if cmd.Flag("config").Value.String() != "" {
		viper.SetConfigFile(cmd.Flag("config").Value.String())
	}
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok || errors.Is(err, os.ErrNotExist) {
			logger.Info().Msg("config file not found, creating new config")
			err = os.MkdirAll(path.Dir(viper.ConfigFileUsed()), 0755)
			if err != nil {
				logger.Fatal().Err(err).Msg("failed to create config directory")
			}
			err = viper.WriteConfig()
			if err != nil {
				logger.Fatal().Err(err).Msg("failed to create config file")
			}
		} else {
			logger.Fatal().Err(err).Msg("failed to read config")
		}
	}

	if cmd.Flag("debug").Value.String() == "true" {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	if isAuth {
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
			c := client.NewHTTPClientWithConfig(strfmt.Default, defaultConfig)
			res, err := c.Auth.PostAuthRefresh(auth.NewPostAuthRefreshParams().WithGrantType("refresh_token").WithRefreshToken(t["refresh_token"]))
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

		return logger, t
	}

	return logger, nil
}
