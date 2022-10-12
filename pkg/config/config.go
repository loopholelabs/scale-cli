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
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

func Init(cmd *cobra.Command) *zerolog.Logger {
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

	if viper.GetBool("debug") {
		logger = logger.Level(zerolog.DebugLevel)
	}

	return &logger
}
