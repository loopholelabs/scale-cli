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

type Dependency struct {
	Name    string
	Version string
}

const (
	GoModfileTemplate = `
module {{ .package }}

go 1.20

replace {{ .old_signature_dependency }} {{ .old_signature_version }} => {{ .new_signature_dependency }} {{ .new_signature_version }}

{{ range $dependency := .dependencies -}}
    require {{ $dependency.Name }} {{ $dependency.Version }}
{{end -}}
`

	GoFunctionTemplate = `
package {{ .package }}

import (
	"signature"
)

func Scale(ctx *signature.{{ .context }}) (*signature.{{ .context }}, error) {
	return signature.Next(ctx)
}
`

	RustCargofileTemplate = `
[package]
name = "{{ .package }}"
version = "{{ .version }}"
edition = "2021"

[lib]
path = "lib.rs"

[dependencies]
{{ range $dependency := .dependencies -}}
{{ $dependency.Name }} = "{{ $dependency.Version }}"
{{end -}}

{{ if .signature_path }}
{{ .signature_dependency }} = { package = "{{ .signature_package }}", path = "{{ .signature_path }}" }
{{ end }}

{{ if .signature_version }}
{{ .signature_dependency }} = { package = "{{ .signature_package }}", version = "{{ .signature_version }}", registry = "scale" }
{{ end }}

[profile.release]
opt-level = 3
lto = true
codegen-units = 1
`
	RustFunctionTemplate = `
use signature::types;

pub fn scale(ctx: Option<&mut types::{{ .context }}>) -> Result<Option<types::{{ .context }}>, Box<dyn std::error::Error>> {
    return signature::next(ctx);
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
		int32 timeout {
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
		string url {
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
