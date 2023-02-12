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
	Path      string
	Name      string
	Signature string
}

func LocalBuild(ctx context.Context, scaleFile *scalefile.ScaleFile, c chan<- *scalefunc.ScaleFunc) error {

	scaleFunc := &scalefunc.ScaleFunc{
		Version:   scalefunc.V1Alpha,
		Name:      fmt.Sprintf("%s:latest", scaleFile.Name),
		Signature: scaleFile.Signature,
		Language:  scalefunc.Language(scaleFile.Language),
	}

	isErr := true

	if scaleFunc.Language == "go" {

		module := &Module{
			Path:      fmt.Sprintf("%s.go", scaleFile.Name),
			Name:      scaleFile.Name,
			Signature: scaleFile.Signature,
		}

		tinygo, err := exec.LookPath("tinygo")
		if err != nil {
		}

		g := compile.NewGenerator()

		_, err = os.Stat(module.Path)
		if err != nil {
		}

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name)), 0755)
		if !os.IsExist(err) {
		}

		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "main.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
		}

		importPath := "github.com/loopholelabs/scale-cli"
		err = g.GenerateGoMain(file, module.Signature, fmt.Sprintf("%s/%s-build/scale", importPath, module.Name))
		if err != nil {
		}

		err = file.Close()
		if err != nil {
		}

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale"), 0755)
		if !os.IsExist(err) {
		}

		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "scale.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
		}

		file, err = os.Open(fmt.Sprintf("%s.go", scaleFile.Name))
		if err != nil {
		}

		_, err = io.Copy(scale, file)
		if err != nil {
		}

		err = scale.Close()
		if err != nil {
		}

		err = file.Close()
		if err != nil {
		}

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature"), 0755)
		if !os.IsExist(err) {
		}

		copiedSignature, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature", "signature.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
		}

		signature, err := os.Open(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature", "signature.go"))
		if err != nil {
		}

		_, err = io.Copy(copiedSignature, signature)
		if err != nil {
		}

		err = signature.Close()
		if err != nil {
		}

		err = copiedSignature.Close()
		if err != nil {
		}

		wd, err := os.Getwd()
		if err != nil {
		}

		cmd := exec.Command(tinygo, "build", "-o", fmt.Sprintf("%s.wasm", module.Name), "-scheduler=none", "-target=wasi", "--no-debug", "main.go")
		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-build", module.Name))

		err = cmd.Run()
		if err != nil {
		}

		data, err := os.ReadFile(path.Join(cmd.Dir, fmt.Sprintf("%s.wasm", module.Name)))
		if err == nil {
			scaleFunc.Function = data
			isErr = false
		}
	}
	if scaleFunc.Language == "rust" {

		module := &Module{
			Path:      fmt.Sprintf("%s.rs", scaleFile.Name),
			Name:      scaleFile.Name,
			Signature: scaleFile.Signature,
		}

		cargo, err := exec.LookPath("cargo")
		if err != nil {
			fmt.Println(err)
		}

		g := rustCompile.NewGenerator()

		_, err = os.Stat(module.Path)
		if err != nil {
			fmt.Println(err)
		}

		moduleDir := path.Dir(module.Path)

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name)), 0755)
		if !os.IsExist(err) {
			fmt.Println(err)
		}

		//thing, err := os.Create(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "lib.rs"))
		file, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "lib.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		//file, err := os.Open(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "lib.rs"))
		fmt.Println(file, "file 1")
		if err != nil {
			fmt.Println(err)
		}

		importPath := "scale/scale.rs"
		err = g.GenerateLibRs(file, module.Signature, importPath)

		//os.Create(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "Cargo.toml"))
		cargoFile, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "Cargo.toml"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		fmt.Println(cargoFile, "file 2")
		dependencies := []*scalefile.Dependency{}
		err = g.GenerateCargoTomlfile(cargoFile, dependencies)
		if err != nil {
			fmt.Println(err, "1")
		}

		err = file.Close()
		if err != nil {
			fmt.Println(err, "2")
		}

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale"), 0755)
		if err != nil {
			fmt.Println(err, "3")
		}

		os.Create(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "scale.rs"))
		scale, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "scale.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}

		file, err = os.Open(fmt.Sprintf("%s.rs", scaleFile.Name))
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.Copy(scale, file)
		if err != nil {
			fmt.Println(err)
		}

		err = scale.Close()
		if err != nil {
			fmt.Println(err)
		}

		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}

		err = os.Mkdir(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature"), 0755)
		if !os.IsExist(err) {
			fmt.Println(err)
		}

		copiedSignature, err := os.OpenFile(path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature", "signature.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}

		signature, err := os.Open((path.Join(moduleDir, fmt.Sprintf("%s-build", module.Name), "scale", "signature", "signature.rs")))
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.Copy(copiedSignature, signature)
		if err != nil {
			fmt.Println(err)
		}

		err = signature.Close()
		if err != nil {
			fmt.Println(err)
		}

		err = copiedSignature.Close()
		if err != nil {
			fmt.Println(err)
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		cmd := exec.Command(cargo, "build", "--target", "wasm32-unknown-unknown", "--manifest-path", "Cargo.toml")

		cmd.Dir = path.Join(wd, moduleDir, fmt.Sprintf("%s-build", module.Name))
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		data, err := os.ReadFile(path.Join(cmd.Dir, "target/wasm32-unknown-unknown/debug/compile.wasm"))

		if err == nil {
			scaleFunc.Function = data
			isErr = false
		}
	}

	if scaleFunc.Language == "typescript" {
		errors.New("TypeScript support not implemented")
	}

	if isErr {
		errors.New("problem with build")
	}

	c <- scaleFunc
	return nil
}
