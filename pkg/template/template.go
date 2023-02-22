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

package template

var (
	LUT = map[string]func() []byte{
		"go":   Go,
		"rust": Rust,
	}
)

const (
	GoTemplate = `module scale

go 1.18
{{range .}}
require {{.Name}} {{.Version}}
{{end}}
`
	RustTemplate = `[package]
name = "scale"
version = "0.1.0"
edition = "2021"

[dependencies]
{{range .}}
{{.Name}} = "{{.Version}}"
{{end}}

[lib]
crate-type = ["cdylib"]
path = "scale.rs"
`
)

func Go() []byte {
	return []byte(`//go:build tinygo || js || wasm
package scale

import (
	signature "github.com/loopholelabs/scale-signature-http"
)

func Scale(ctx *signature.Context) (*signature.Context, error) {
	ctx.Response().SetBody("Hello, World!")
	return ctx.Next()
}`)
}

func Rust() []byte {
	return []byte(`use scale_signature_http::context::Context;

pub fn scale(ctx: &mut Context) -> Result<&mut Context, Box<dyn std::error::Error>> {
    ctx.response().set_body("Hello, World!".to_string());
    Ok(ctx)
}`)
}
