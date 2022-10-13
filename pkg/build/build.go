/*
 * Copyright 2022 Loophole Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package build

import (
	"context"
	"errors"
	"github.com/loopholelabs/frisbee-go"
	"github.com/loopholelabs/frisbee-go/pkg/packet"
	"github.com/loopholelabs/scale-cli/pkg/scalefile"
	"github.com/loopholelabs/scale-cli/rpc/build"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
	"time"
)

func Build(input []byte, outputPath string, token string, scaleFile scalefile.ScaleFile, logger zerolog.Logger) {
	start := time.Now()
	client, err := build.NewClient(nil, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("error while creating builder fRPC client")
	}

	builderServer := viper.GetString("builder")

	isErr := true
	streamDone := make(chan struct{})
	err = client.Connect(builderServer, func(stream *frisbee.Stream) {
		streamPacket := build.NewBuildStreamPacket()
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
			case build.BuildSTDOUT:
				logger.Info().Msg(string(streamPacket.Data))
			case build.BuildSTDERR:
				logger.Warn().Msg(string(streamPacket.Data))
			case build.BuildOUTPUT:
				logger.Info().Msgf("Writing Compiled Scale Function to %s", outputPath)
				err = os.WriteFile(outputPath, streamPacket.Data, 0644)
				if err != nil {
					logger.Error().Err(err).Msgf("error while writing compiled scale function to %s", outputPath)
				}
				isErr = false
			case build.BuildCLOSE:
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
	req := build.NewBuildRequest()
	req.Token = token
	req.ScaleFile.Name = scaleFile.Name
	req.ScaleFile.Input = input
	req.ScaleFile.BuildConfig.Language = scaleFile.Builder.Language
	for _, dependency := range scaleFile.Builder.Dependencies {
		req.ScaleFile.BuildConfig.Dependencies = append(req.ScaleFile.BuildConfig.Dependencies, &build.BuildDependency{
			Name:    dependency.Name,
			Version: dependency.Version,
		})
	}
	logger.Info().Msgf("Compiling %s...", scaleFile.Name)
	start = time.Now()
	res, err := client.Service.Build(context.Background(), req)
	if err != nil {
		logger.Fatal().Err(err).Msg("error while sending compilation request to wabuild server")
	}
	logger.Debug().Msgf("stream ID: %d", res.StreamID)
	<-streamDone
	logger.Debug().Msgf("completed remote build in %s", time.Since(start))
	if isErr {
		logger.Fatal().Msg("Error Occurred while Compiling Module")
	}
}
