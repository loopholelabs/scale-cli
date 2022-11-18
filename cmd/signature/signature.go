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

package signature

import (
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

// Cmd returns the base command for signature.
func Cmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "signature <command>",
		Aliases: []string{"sig"},
		Short:   "Create, list, and manage Scale Signatures",
	}

	cmd.AddCommand(NewCmd(ch))
	cmd.AddCommand(GenerateCmd(ch))

	return cmd
}
