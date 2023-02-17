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

package utils

import (
	"errors"
	"fmt"
	runtimeClient "github.com/go-openapi/runtime/client"
	"github.com/loopholelabs/auth"
	"github.com/loopholelabs/auth/pkg/client/session"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/log"
	"github.com/spf13/cobra"
)

var (
	ErrNotAuthenticated = errors.New("You must be authenticated to use this command. Please run 'scale auth login' to authenticate.")
)

func PreRunAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		log.Init(ch.Config.GetLogFile())
		err := ch.Config.GlobalRequiredFlags(cmd)
		if err != nil {
			return err
		}
		err = ch.Config.Validate()
		if err != nil {
			return err
		}

		if !ch.Config.IsAuthenticated() {
			return ErrNotAuthenticated
		}

		c, err := ch.Config.NewAuthenticatedAPIClient()
		if err != nil {
			return err
		}

		ch.Config.SetAPIClient(c)

		return nil
	}
}

func PostRunAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := ch.Config.APIClient()
		cookies := c.Transport.(*runtimeClient.Runtime).Jar.Cookies(config.DefaultCookieURL)
		if len(cookies) == 0 {
			return nil
		}
		ch.Config.Session = session.New(auth.KindSession, cookies[0].Value, cookies[0].Expires)

		err := ch.Config.WriteSession()
		if err != nil {
			return fmt.Errorf("error updating session: %w", err)
		}

		return nil
	}
}
