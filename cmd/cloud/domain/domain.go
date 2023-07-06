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

package domain

import (
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"
	"github.com/loopholelabs/scale-cli/cmd/utils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/spf13/cobra"
)

// Cmd encapsulates the commands for functions.
func Cmd(hidden bool) command.SetupCommand[*config.Config] {
	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		domainCmd := &cobra.Command{
			Use:                "domain <command>",
			Aliases:            []string{"domain"},
			Short:              "Create, list, and manage domains",
			Hidden:             hidden,
			PersistentPostRunE: utils.PostRunAnalytics(ch),
		}

		addSetup := AddCmd(false)
		addSetup(domainCmd, ch)

		verifySetup := VerifyCmd(false)
		verifySetup(domainCmd, ch)

		removeSetup := RemoveCmd(false)
		removeSetup(domainCmd, ch)

		attachSetup := AttachCmd(false)
		attachSetup(domainCmd, ch)

		detachSetup := DetachCmd(false)
		detachSetup(domainCmd, ch)

		listSetup := ListCmd(false)
		listSetup(domainCmd, ch)

		cmd.AddCommand(domainCmd)
	}
}
