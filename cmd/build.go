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
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/loopholelabs/frisbee-go"
	"github.com/loopholelabs/frisbee-go/pkg/packet"
	"github.com/loopholelabs/scale-cli/pkg/api/client"
	"github.com/loopholelabs/scale-cli/pkg/api/client/auth"
	"github.com/loopholelabs/scale-cli/pkg/builder"
	"github.com/loopholelabs/scale-cli/pkg/config"
	"github.com/loopholelabs/scale-cli/pkg/scalefile"
	"github.com/loopholelabs/scale-cli/pkg/token"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path"
	"time"
)

const (
	BuilderPort = 8080
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

		scalefilePath := viper.GetString("scalefile")
		logger.Debug().Msgf("build called with scalefile '%s'", scalefilePath)
		scaleFile, err := scalefile.Read(scalefilePath)
		if err != nil {
			logger.Fatal().Err(err).Msg("error reading scalefile")
		}
		directory := path.Dir(scalefilePath)
		filePath := path.Join(directory, scaleFile.File)

		output := fmt.Sprintf("%s/%s.wasm", directory, scaleFile.Name)

		start := time.Now()
		input, err := os.ReadFile(filePath)
		if err != nil {
			logger.Fatal().Err(err).Msgf("error while reading scale function file %s", filePath)
		}
		logger.Debug().Msgf("read scale function file %s in %s", filePath, time.Since(start))

		var clientLogger *zerolog.Logger
		if viper.GetBool("debug") {
			zlogger := logger.With().Str("COMPONENT", "BUILD").Logger().Level(zerolog.InfoLevel)
			clientLogger = &zlogger
		}

		start = time.Now()
		client, err := builder.NewClient(nil, clientLogger)
		if err != nil {
			logger.Fatal().Err(err).Msg("error while creating builder fRPC client")
		}

		builderServer := fmt.Sprintf("%s:%d", viper.GetString("builder"), BuilderPort)

		isErr := true
		streamDone := make(chan struct{})
		err = client.Connect(builderServer, func(stream *frisbee.Stream) {
			streamPacket := builder.NewStreamPacket()
			for {
				p, err := stream.ReadPacket()
				if err != nil {
					packet.Put(p)
					if !errors.Is(err, frisbee.StreamClosed) {
						logger.Error().Err(err).Msg("error while reading packet from builder stream")
					}
					break
				}
				err = streamPacket.Decode(p.Content.Bytes())
				packet.Put(p)
				if err != nil {
					logger.Error().Err(err).Msg("error while decoding packet from builder stream")
					break
				}
				switch streamPacket.Type {
				case builder.STDOUT:
					logger.Info().Msg(string(streamPacket.Data))
				case builder.STDERR:
					logger.Warn().Msg(string(streamPacket.Data))
				case builder.OUTPUT:
					logger.Info().Msgf("Writing Compiled Scale Function to %s", output)
					err = os.WriteFile(output, streamPacket.Data, 0644)
					if err != nil {
						logger.Error().Err(err).Msgf("error while writing compiled scale function to %s", output)
					}
					isErr = false
				case builder.CLOSE:
					break
				}
			}
			_ = stream.Close()
			streamDone <- struct{}{}
		})
		if err != nil {
			logger.Fatal().Err(err).Msg("error while connecting to scale build server")
		}
		logger.Debug().Msgf("connected to scale build server in %s", time.Since(start))
		req := builder.NewBuildRequest()
		req.Authorization = t["access_token"]
		req.ScaleFile.Name = scaleFile.Name
		req.ScaleFile.Input = input
		req.ScaleFile.Build.Language = scaleFile.Builder.Language
		for _, dependency := range scaleFile.Builder.Dependencies {
			req.ScaleFile.Build.Dependencies = append(req.ScaleFile.Build.Dependencies, &builder.Dependency{
				Name:    dependency.Name,
				Version: dependency.Version,
			})
		}
		logger.Info().Msgf("Compiling %s...", scaleFile.Name)
		start = time.Now()
		res, err := client.BuildService.Build(context.Background(), req)
		if err != nil {
			logger.Fatal().Err(err).Msg("error while sending compilation request to wabuild server")
		}
		logger.Debug().Msgf("stream ID: %d", res.StreamID)
		<-streamDone
		logger.Debug().Msgf("completed remote build in %s", time.Since(start))
		if isErr {
			logger.Fatal().Msg("Error Occurred while Compiling Module")
		}
		logger.Info().Msg("Module Compilation Completed")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringP("scalefile", "s", "scalefile", "the scalefile to use")
	buildCmd.Flags().StringP("builder", "b", "build.scale.com", "Scale Build Service URL")

	err := viper.BindPFlag("scalefile", buildCmd.Flags().Lookup("scalefile"))
	if err != nil {
		panic(err)
	}

	err = viper.BindPFlag("builder", buildCmd.Flags().Lookup("builder"))
	if err != nil {
		panic(err)
	}
}
