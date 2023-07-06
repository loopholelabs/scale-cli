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

package cloud

import (
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/cmd/cloud/domain"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/spf13/cobra"
)

// Cmd encapsulates the commands for functions.
func Cmd() command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		cloudCmd := &cobra.Command{
			Use:                "cloud <command>",
			Aliases:            []string{"cl"},
			Short:              "Create, list, and manage uploaded Scale Functions",
			PersistentPostRunE: utils.PostRunAnalytics(ch),
		}

		deploySetup := DeployCmd(false)
		deploySetup(cloudCmd, ch)

		listSetup := ListCmd(false)
		listSetup(cloudCmd, ch)

		deleteSetup := DeleteCmd(false)
		deleteSetup(cloudCmd, ch)

		domainSetup := domain.Cmd(false)
		domainSetup(cloudCmd, ch)

		deployAliasSetup := DeployCmd(true)
		deployAliasSetup(cmd, ch)

		listAliasSetup := ListCmd(true)
		listAliasSetup(cmd, ch)

		domainAliasSetup := domain.Cmd(true)
		domainAliasSetup(cmd, ch)

		cmd.AddCommand(cloudCmd)
	}
}
