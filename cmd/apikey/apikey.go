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

package apikey

import (
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

type apiKey struct {
	Name    string `header:"name" json:"name"`
	ID      string `header:"id" json:"id"`
	Value   string `header:"value" json:"value"`
	Created string `header:"created" json:"created"`
}

type apiKeyRedacted struct {
	Name    string `header:"name" json:"name"`
	ID      string `header:"id" json:"id"`
	Created string `header:"created" json:"created"`
}

// Cmd encapsulates the command for interacting with API Keys.
func Cmd(ch *cmdutil.Helper) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "api-key <action>",
		Short:              "Create, list, and manage API Keys",
		PersistentPreRunE:  cmdutil.CheckAuthentication(ch.Config),
		PersistentPostRunE: cmdutil.UpdateToken(ch.Config),
	}

	cmd.AddCommand(CreateCmd(ch))
	cmd.AddCommand(ListCmd(ch))
	cmd.AddCommand(GetCmd(ch))
	cmd.AddCommand(DeleteCmd(ch))
	return cmd
}
