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

package function

import (
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

type scaleFunction struct {
	Name       string `header:"name" json:"name"`
	Tag        string `header:"tag" json:"tag"`
	Language   string `header:"language" json:"language"`
	Middleware bool   `header:"middleware" json:"middleware"`
	Version    string `header:"version,v1" json:"version"`
}

// Cmd returns the base command for function.
func Cmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "function <command>",
		Aliases: []string{"fn"},
		Short:   "Create, list, and manage Scale Functions",
	}

	cmd.AddCommand(BuildCmd(ch))
	cmd.AddCommand(NewCmd(ch))
	cmd.AddCommand(ListCmd(ch))
	cmd.AddCommand(DeleteCmd(ch))
	cmd.AddCommand(RunCmd(ch))
	return cmd
}
