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

const (
	GoModfileTemplate = `
module {{ .package_name }}

go 1.20

replace signature v0.1.0 => {{ .signature_path }} {{ .signature_version }}

require signature v0.1.0

{{ range $extension := .extensions -}}
		replace {{ $extension.Name }} => {{ $extension.Path }}
{{end -}}

{{ range $extension := .extensions -}}
		require {{ $extension.Name }} {{ $extension.Version }}
{{end -}}

`

	GoFunctionTemplate = `
package {{ .package_name }}

import (
	"signature"
)

func Scale(ctx *signature.{{ .context_name }}) (*signature.{{ .context_name }}, error) {
	return signature.Next(ctx)
}
`

	RustCargofileTemplate = `
[package]
name = "{{ .package_name }}"
version = "0.1.0"
edition = "2021"

[lib]
path = "lib.rs"

[dependencies]
{{ if .signature_path }}
signature = { package = "{{ .signature_package }}", path = "{{ .signature_path }}" }
{{ end }}

{{ if .signature_version }}
signature = { package = "{{ .signature_package }}", version = "{{ .signature_version }}", registry = "scale" }
{{ end }}

{{ range $extension := .extensions -}}
{{ $extension.Name }} = { package = "{{ $extension.Package }}", path = "{{ $extension.Path }}" }
{{end -}}

[profile.release]
opt-level = 3
lto = true
codegen-units = 1
`
	RustFunctionTemplate = `
use signature::types;

pub fn scale(ctx: Option<types::{{ .context_name }}>) -> Result<Option<types::{{ .context_name }}>, Box<dyn std::error::Error>> {
    return signature::next(ctx);
}
`

	TypescriptPackageTemplate = `
{
  "name": "{{ .package_name }}",
  "version": "0.1.0",
  "main": "index.ts",
  "dependencies": {
    "signature": "file:{{ .signature_path }}"

		{{ range $extension := .extensions -}}
		,"{{ $extension.Name }}": "file:{{ $extension.Path }}"
		{{end -}}
	}
}
`
	TypeScriptFunctionTemplate = `
import * as signature from "signature";

export function scale(ctx?: signature.{{ .context_name }}): signature.{{ .context_name }} | undefined {
    return signature.Next(ctx);
}
`

	SignatureFile = `
version = "v1alpha"
context = "context"
model Context {
  string MyString {
    default = "DefaultValue"
  }
}`

	ExtensionFile = `
	version = "v1alpha"
	name = "HttpFetch"
	tag = "alpha"
	
	function New {
		params = "HttpConfig"
		return = "HttpConnector"	
	}
	
	model HttpConfig {
		int32 Timeout {
			default = 60
		}
	}
	
	model HttpResponse {
		string_map Headers {
			value = "StringList"
		}
		int32 StatusCode {
			default = 0
		}
		bytes Body {
			initial_size = 0
		}
	}
	
	model StringList {
		string_array Values {
			initial_size = 0
		}
	}
	
	model ConnectionDetails {
		string Url {
			default = "https://google.com"
		}
	}
	
	interface HttpConnector {
		function Fetch {
			params = "ConnectionDetails"
			return = "HttpResponse"
		}
	}`
)
