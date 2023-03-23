/*
	Copyright 2023 Loophole Labs

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
	"errors"
	"fmt"
	"github.com/loopholelabs/scale/go/compile"
	rustCompile "github.com/loopholelabs/scale/rust/compile"
	"github.com/loopholelabs/scalefile"
	"github.com/loopholelabs/scalefile/scalefunc"
	"io"
	"os"
	"os/exec"
	"path"
)

var (
	ErrNoGo     = errors.New("go not found in PATH. Please install go: https://golang.org/doc/install")
	ErrNoTinyGo = errors.New("tinygo not found in PATH. Please install tinygo: https://tinygo.org/getting-started/")
	ErrNoCargo  = errors.New("cargo not found in PATH. Please install cargo: https://doc.rust-lang.org/cargo/getting-started/installation.html")
)

type Module struct {
	Source    string
	Name      string
	Signature string
}

func LocalBuild(scaleFile *scalefile.ScaleFile, goBin string, tinygo string, cargo string, baseDir string) (*scalefunc.ScaleFunc, error) {
	scaleFunc := &scalefunc.ScaleFunc{
		Version:   scalefunc.V1Alpha,
		Name:      scaleFile.Name,
		Tag:       scaleFile.Tag,
		Signature: scaleFile.Signature,
		Language:  scaleFile.Language,
	}

	switch scaleFunc.Language {
	case scalefunc.Go:
		module := &Module{
			Name:      scaleFile.Name,
			Source:    scaleFile.Source,
			Signature: "github.com/loopholelabs/scale-signature-http",
		}

		if goBin != "" {
			stat, err := os.Stat(goBin)
			if err != nil {
				return nil, fmt.Errorf("unable to find go binary %s: %w", goBin, err)
			}
			if !(stat.Mode()&0111 != 0) {
				return nil, fmt.Errorf("go binary %s is not executable", goBin)
			}
		} else {
			var err error
			goBin, err = exec.LookPath("go")
			if err != nil {
				return nil, ErrNoGo
			}
		}

		if tinygo != "" {
			stat, err := os.Stat(tinygo)
			if err != nil {
				return nil, fmt.Errorf("unable to find tinygo binary %s: %w", tinygo, err)
			}
			if !(stat.Mode()&0111 != 0) {
				return nil, fmt.Errorf("tinygo binary %s is not executable", tinygo)
			}
		} else {
			var err error
			tinygo, err = exec.LookPath("tinygo")
			if err != nil {
				return nil, ErrNoTinyGo
			}
		}

		g := compile.NewGenerator()

		moduleSourcePath := path.Join(baseDir, module.Source)
		_, err := os.Stat(moduleSourcePath)
		if err != nil {
			return nil, fmt.Errorf("unable to find module %s: %w", moduleSourcePath, err)
		}

		buildDir := path.Join(baseDir, "build")
		defer func() {
			_ = os.RemoveAll(buildDir)
		}()

		err = os.Mkdir(buildDir, 0755)
		if !os.IsExist(err) && err != nil {
			return nil, fmt.Errorf("unable to create build %s directory: %w", buildDir, err)
		}

		file, err := os.OpenFile(path.Join(buildDir, "main.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create main.go file: %w", err)
		}

		err = g.GenerateGoMain(file, "compile/scale", module.Signature)
		if err != nil {
			return nil, fmt.Errorf("unable to generate main.go file: %w", err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close main.go file: %w", err)
		}

		file, err = os.OpenFile(path.Join(buildDir, "go.mod"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create go.mod file: %w", err)
		}

		deps := make([]*scalefile.Dependency, 0, len(scaleFile.Dependencies))
		for _, dep := range scaleFile.Dependencies {
			var d = dep
			deps = append(deps, &d)
		}
		err = g.GenerateGoModfile(file, deps)
		if err != nil {
			return nil, fmt.Errorf("unable to generate go.mod file: %w", err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close go.mod file: %w", err)
		}

		scalePath := path.Join(buildDir, "scale")
		err = os.Mkdir(scalePath, 0755)
		if !os.IsExist(err) && err != nil {
			return nil, fmt.Errorf("unable to create scale source directory %s: %w", scalePath, err)
		}

		scale, err := os.OpenFile(path.Join(scalePath, "scale.go"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create scale.go file: %w", err)
		}

		file, err = os.Open(moduleSourcePath)
		if err != nil {
			return nil, fmt.Errorf("unable to open scale source file: %w", err)
		}

		_, err = io.Copy(scale, file)
		if err != nil {
			return nil, fmt.Errorf("unable to copy scale source file: %w", err)
		}

		err = scale.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close scale.go file: %w", err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close scale source file: %w", err)
		}

		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("unable to get working directory: %w", err)
		}

		cmd := exec.Command(goBin, "mod", "tidy")
		cmd.Dir = path.Join(wd, buildDir)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return nil, fmt.Errorf("unable to compile scale function: %s", output)
			}
			return nil, fmt.Errorf("unable to compile scale function: %w", err)
		}

		cmd = exec.Command(tinygo, "build", "-o", "scale.wasm", "-scheduler=none", "-target=wasi", "--no-debug", "main.go")
		cmd.Dir = path.Join(wd, buildDir)

		output, err = cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return nil, fmt.Errorf("unable to compile scale function: %s", output)
			}
			return nil, fmt.Errorf("unable to compile scale function: %w", err)
		}

		data, err := os.ReadFile(path.Join(cmd.Dir, "scale.wasm"))
		if err != nil {
			return nil, fmt.Errorf("unable to read compiled wasm file: %w", err)
		}
		scaleFunc.Function = data
	case scalefunc.Rust:
		module := &Module{
			Name:      scaleFile.Name,
			Source:    scaleFile.Source,
			Signature: "scale_signature_http",
		}

		if cargo != "" {
			stat, err := os.Stat(cargo)
			if err != nil {
				return nil, fmt.Errorf("unable to find cargo binary %s: %w", cargo, err)
			}
			if !(stat.Mode()&0111 != 0) {
				return nil, fmt.Errorf("cargo binary %s is not executable", cargo)
			}
		} else {
			var err error
			cargo, err = exec.LookPath("cargo")
			if err != nil {
				return nil, ErrNoCargo
			}
		}

		g := rustCompile.NewGenerator()

		moduleSourcePath := path.Join(baseDir, module.Source)
		_, err := os.Stat(moduleSourcePath)
		if err != nil {
			return nil, fmt.Errorf("unable to find module %s: %w", moduleSourcePath, err)
		}

		buildDir := path.Join(baseDir, "build")
		defer func() {
			_ = os.RemoveAll(buildDir)
		}()

		err = os.Mkdir(buildDir, 0755)
		if !os.IsExist(err) && err != nil {
			return nil, fmt.Errorf("unable to create build %s directory: %w", buildDir, err)
		}

		file, err := os.OpenFile(path.Join(buildDir, "lib.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create lib.rs file: %w", err)
		}

		err = g.GenerateRsLib(file, "scale/scale.rs", module.Signature)
		if err != nil {
			return nil, fmt.Errorf("unable to generate lib.rs file: %w", err)
		}

		cargoFile, err := os.OpenFile(path.Join(buildDir, "Cargo.toml"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create Cargo.toml file: %w", err)
		}

		deps := make([]*scalefile.Dependency, 0, len(scaleFile.Dependencies))
		for _, dep := range scaleFile.Dependencies {
			var d = dep
			deps = append(deps, &d)
		}
		err = g.GenerateRsCargo(cargoFile, deps, module.Signature, "")
		if err != nil {
			return nil, fmt.Errorf("unable to generate Cargo.toml file: %w", err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close lib.rs file: %w", err)
		}

		scalePath := path.Join(buildDir, "scale")
		err = os.Mkdir(scalePath, 0755)
		if !os.IsExist(err) && err != nil {
			return nil, fmt.Errorf("unable to create scale source directory %s: %w", scalePath, err)
		}

		scale, err := os.OpenFile(path.Join(scalePath, "scale.rs"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return nil, fmt.Errorf("unable to create scale.rs file: %w", err)
		}

		file, err = os.Open(moduleSourcePath)
		if err != nil {
			return nil, fmt.Errorf("unable to open scale source file: %w", err)
		}

		_, err = io.Copy(scale, file)
		if err != nil {
			return nil, fmt.Errorf("unable to copy scale source file: %w", err)
		}

		err = scale.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close scale.rs file: %w", err)
		}

		err = file.Close()
		if err != nil {
			return nil, fmt.Errorf("unable to close scale source file: %w", err)
		}

		cmd := exec.Command(cargo, "build", "--target", "wasm32-unknown-unknown", "--manifest-path", "Cargo.toml", "--release")
		cmd.Dir = buildDir

		output, err := cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return nil, fmt.Errorf("unable to compile scale function: %s", output)
			}
			return nil, fmt.Errorf("unable to compile scale function: %w", err)
		}

		data, err := os.ReadFile(path.Join(cmd.Dir, "target/wasm32-unknown-unknown/release/compile.wasm"))
		if err != nil {
			return nil, fmt.Errorf("unable to read compiled wasm file: %w", err)
		}
		scaleFunc.Function = data
	default:
		return nil, fmt.Errorf("%s support not implemented", scaleFile.Language)
	}

	return scaleFunc, nil
}
