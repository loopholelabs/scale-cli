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
	"github.com/loopholelabs/scale/go/compile"
	rustCompile "github.com/loopholelabs/scale/rust/compile"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"io"
	"os"
	"os/exec"
	"path"
	"time"
)

func RemoteBuild(ctx context.Context, endpoint string, name string, input []byte, token string, scaleFile *scalefile.ScaleFile, tlsConfig *tls.Config, ch *cmdutil.Helper) (*scalefunc.ScaleFunc, error) {
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

type Module struct {
	Path         string
	Name         string
	Signature    string
	Dependencies []*scalefile.Dependency
}

func LocalBuild(ctx context.Context, endpoint string, name string, input []byte, token string, scaleFile *scalefile.ScaleFile, tlsConfig *tls.Config, ch *cmdutil.Helper) (*scalefunc.ScaleFunc, error) {

	importPath := "./" //TODO

	scaleFunc := &scalefunc.ScaleFunc{
		Version:   scalefunc.V1Alpha,
		Name:      fmt.Sprintf("%s:latest", scaleFile.Name),
		Signature: scaleFile.Signature,
		Language:  scalefunc.Language(scaleFile.Language),
	}

	module := &Module{
		Path:         name,
		Name:         scaleFile.Name,
		Signature:    scaleFile.Signature,
		Dependencies: []*scalefile.Dependency{}, //TODO
	}

	isErr := true

	if scaleFunc.Language == "go" {

		tinygo, err := exec.LookPath("tinygo")

		g := compile.NewGenerator()

		_, err = os.Stat(module.Path)

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name)), 0755)
		if !os.IsExist(err) {
		}

		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "main.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

		err = g.GenerateGoMain(file, module.Signature, fmt.Sprintf("%s/%s/%s-build/scale", importPath, module.Name, module.Name))

		err = file.Close()

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale"), 0755)
		if !os.IsExist(err) {
		}

		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "scale.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

		file, err = os.Open(module.Path)

		_, err = io.Copy(scale, file)

		err = scale.Close()

		err = file.Close()

		wd, err := os.Getwd()

		cmd := exec.Command(tinygo, "build", "-o", fmt.Sprintf("%s.wasm", module.Name), "-scheduler=none", "-target=wasi", "--no-debug", "main.go")
		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-build", module.Name))

		err = cmd.Run()

		data, err := os.ReadFile(path.Join(cmd.Dir, fmt.Sprintf("%s.wasm", module.Name)))
		scaleFunc.Function = data
		isErr = false
	}
	if scaleFunc.Language == "rust" {

		cargo, err := exec.LookPath("cargo")

		g := rustCompile.NewGenerator()

		_, err = os.Stat(module.Path)

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name)), 0755)
		if !os.IsExist(err) {
		}

		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "lib.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

		err = g.GenerateLibRs(file, module.Signature, importPath)

		cargoFile, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "Cargo.toml"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		dependencies := []*scalefile.Dependency{}
		err = g.GenerateCargoTomlfile(cargoFile, dependencies)

		err = file.Close()

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale"), 0755)
		if !os.IsExist(err) {
		}

		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "scale.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

		file, err = os.Open(module.Path)

		_, err = io.Copy(scale, file)

		err = scale.Close()

		err = file.Close()

		wd, err := os.Getwd()

		cmd := exec.Command(cargo, "build", "--target", "wasm32-unknown-unknown", "--manifest-path", "Cargo.toml")

		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-build", module.Name))

		err = cmd.Run()

		data, err := os.ReadFile(path.Join(cmd.Dir, "target/wasm32-unknown-unknown/debug/compile.wasm"))
		scaleFunc.Function = data
		isErr = false
	}

	//if scaleFunc.Language == "typescript" {

	//}

	if isErr {
		return nil, errors.New("error occured during build")
	}

	return scaleFunc, nil
}
