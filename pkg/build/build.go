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

package build

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/loopholelabs/frisbee-go"
	"github.com/loopholelabs/frisbee-go/pkg/packet"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/rpc/build"
	"github.com/loopholelabs/scale/scalefile"
	"github.com/loopholelabs/scale/scalefunc"
	"time"
)

func Build(ctx context.Context, endpoint string, name string, input []byte, token string, scaleFile *scalefile.ScaleFile, tlsConfig *tls.Config, ch *cmdutil.Helper) (*scalefunc.ScaleFunc, error) {
	client, err := build.NewClient(tlsConfig, nil)
	if err != nil {
		return nil, err
	}

	scaleFunc := &scalefunc.ScaleFunc{
		Version:   scalefunc.V1Alpha,
		Name:      fmt.Sprintf("%s:latest", scaleFile.Name),
		Signature: scaleFile.Signature,
		Language:  scalefunc.Language(scaleFile.Language),
	}

	for _, f := range scaleFile.Extensions {
		scaleFunc.Extensions = append(scaleFunc.Extensions, scalefunc.Extension(f))
	}

	isErr := true
	streamDone := make(chan struct{})
	err = client.Connect(endpoint, func(stream *frisbee.Stream) {
		streamPacket := build.NewBuildStreamPacket()
		for {
			p, err := stream.ReadPacket()
			if err != nil {
				packet.Put(p)
				if !errors.Is(err, frisbee.StreamClosed) {
					ch.Printer.Printf("%s %v\n", printer.Red("Error reading packet from builder stream:"), printer.BoldRed(err))
				}
				break
			}
			err = streamPacket.Decode(p.Content.Bytes())
			packet.Put(p)
			if err != nil {
				ch.Printer.Printf("%s %v\n", printer.Red("Error decoding packet from builder stream:"), printer.BoldRed(err))
				break
			}
			switch streamPacket.Type {
			case build.BuildSTDOUT:
				ch.Printer.Printf("%s", printer.BoldBlue(string(streamPacket.Data))) // Ignoring newline because it's already in the data
			case build.BuildSTDERR:
				ch.Printer.Printf("%s", printer.BoldYellow(string(streamPacket.Data))) // Ignoring newline because it's already in the data
			case build.BuildOUTPUT:
				scaleFunc.Function = streamPacket.Data
				isErr = false
			case build.BuildCLOSE:
				break
			}
		}
		_ = stream.Close()
		streamDone <- struct{}{}
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to build service: %w", err)
	}
	req := build.NewBuildRequest()
	req.Token = token
	req.ScaleFile.Name = scaleFile.Name
	req.ScaleFile.Input = input
	req.ScaleFile.BuildConfig.Language = string(scaleFunc.Language)
	for _, dependency := range scaleFile.Dependencies {
		req.ScaleFile.BuildConfig.Dependencies = append(req.ScaleFile.BuildConfig.Dependencies, &build.BuildDependency{
			Name:    dependency.Name,
			Version: dependency.Version,
		})
	}
	ch.Printer.Printf("%s %s...\n", printer.BoldBlue("Building scale function"), printer.BoldGreen(name))
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
	defer cancel()
	_, err = client.Service.Build(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error sending build request to build service: %w", err)
	}
	select {
	case <-streamDone:
	case <-ctx.Done():
		return nil, fmt.Errorf("error waiting for build stream to close: %w", ctx.Err())
	}
	if isErr {
		return nil, errors.New("error occured during build")
	}

	return scaleFunc, nil
}
