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

package registry

import (
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/cmd/registry/function"
	"github.com/loopholelabs/scale-cli/cmd/registry/signature"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/spf13/cobra"
)

// Cmd encapsulates the commands for registry.
func Cmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		registryCmd := &cobra.Command{
			Use:     "registry <command>",
			Aliases: []string{"reg"},
			Short:   "Create, list, and manage Scale Functions in the registry",
		}

		functionSetup := function.Cmd()
		functionSetup(registryCmd, ch)

		signatureSetup := signature.Cmd()
		signatureSetup(registryCmd, ch)

		cmd.AddCommand(registryCmd)
	}
}
