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

package template

var (
	LUT = map[string]func() []byte{
		"go": Go,
	}
)

const (
	GoTemplate = `module scale

go 1.18
{{range .}}
require {{.Name}} {{.Version}}
{{end}}
`
)

func Go() []byte {
	return []byte(`package scale

import (
	"github.com/loopholelabs/scale/go/context"
)

func Scale(ctx *context.Context) *context.Context {
	ctx.Response().SetBody("Hello, World!")
	return ctx
}`)
}
