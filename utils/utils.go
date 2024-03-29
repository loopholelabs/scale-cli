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
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	runtimeClient "github.com/go-openapi/runtime/client"
	"github.com/loopholelabs/auth/pkg/client/session"
	"github.com/loopholelabs/auth/pkg/kind"
	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/releaser/pkg/client"
	"github.com/loopholelabs/scale-cli/analytics"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/log"
	"github.com/loopholelabs/scale-cli/version"
	"github.com/loopholelabs/scale/scalefunc"
	"github.com/spf13/cobra"
)

var (
	ErrNotAuthenticated = errors.New("you must be authenticated to use this command. Please run 'scale login' to authenticate")
)

var _ runtime.NamedReadCloser = (*ScaleFunctionNamedReadCloser)(nil)

type ScaleFunctionNamedReadCloser struct {
	reader io.ReadCloser
	name   string
}

func NewScaleFunctionNamedReadCloser(sf *scalefunc.V1BetaSchema) *ScaleFunctionNamedReadCloser {
	return &ScaleFunctionNamedReadCloser{
		reader: io.NopCloser(bytes.NewReader(sf.Encode())),
		name:   sf.Name,
	}
}

func (s *ScaleFunctionNamedReadCloser) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

func (s *ScaleFunctionNamedReadCloser) Close() error {
	return s.reader.Close()
}

func (s *ScaleFunctionNamedReadCloser) Name() string {
	return s.name
}

func PreRunUpdateCheck(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		log.Init(ch.Config.GetLogFile(), ch.Debug())
		err := ch.Config.GlobalRequiredFlags(cmd)
		if err != nil {
			return err
		}

		err = ch.Config.Validate()
		if err != nil {
			return err
		}

		if !ch.Config.DisableAutoUpdate {
			updateClient := client.New(fmt.Sprintf("https://%s", ch.Config.UpdateEndpoint))
			latestReleaseName, err := updateClient.GetLatestReleaseName()
			if err == nil {
				if latestReleaseName != version.Version {
					ch.Printer.Printf("A new version of the Scale CLI is available: %s. Please run 'scale update' to update.\n\n", latestReleaseName)
				}
			}
		}

		return nil
	}
}

func PreRunAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		log.Init(ch.Config.GetLogFile(), ch.Debug())
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

		if !ch.Config.DisableAutoUpdate {
			updateClient := client.New(fmt.Sprintf("https://%s", ch.Config.UpdateEndpoint))
			latestReleaseName, err := updateClient.GetLatestReleaseName()
			if err == nil {
				if latestReleaseName != version.Version {
					ch.Printer.Printf("A new version of the Scale CLI is available: %s. Please run 'scale update' to update.\n\n", latestReleaseName)
				}
			}
		}

		return nil
	}
}

func PreRunOptionalAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		log.Init(ch.Config.GetLogFile(), ch.Debug())
		err := ch.Config.GlobalRequiredFlags(cmd)
		if err != nil {
			return err
		}

		err = ch.Config.Validate()
		if err != nil {
			return err
		}

		if ch.Config.IsAuthenticated() {
			c, err := ch.Config.NewAuthenticatedAPIClient()
			if err == nil {
				ch.Config.SetAPIClient(c)
			}
		} else {
			ch.Config.SetAPIClient(ch.Config.NewUnauthenticatedAPIClient())
		}

		if !ch.Config.DisableAutoUpdate {
			updateClient := client.New(fmt.Sprintf("https://%s", ch.Config.UpdateEndpoint))
			latestReleaseName, err := updateClient.GetLatestReleaseName()
			if err == nil {
				if latestReleaseName != version.Version {
					ch.Printer.Printf("A new version of the Scale CLI is available: %s. Please run 'scale update' to update.\n\n", latestReleaseName)
				}
			}
		}

		return nil
	}
}

func PostRunAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := ch.Config.APIClient()
		if c != nil && c.Transport != nil {
			cookies := c.Transport.(*runtimeClient.Runtime).Jar.Cookies(ch.Config.SessionCookieURL())
			if len(cookies) == 0 {
				return nil
			}
			ch.Config.Session = session.New(kind.Session, cookies[0].Value, cookies[0].Expires)

			err := ch.Config.WriteSession()
			if err != nil {
				return fmt.Errorf("error updating session: %w", err)
			}

		}
		analytics.Cleanup()
		return nil
	}
}

func PostRunOptionalAuthenticatedAPI(ch *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := ch.Config.APIClient()
		if c != nil && c.Transport != nil && ch.Config.IsAuthenticated() {
			cookies := c.Transport.(*runtimeClient.Runtime).Jar.Cookies(ch.Config.SessionCookieURL())
			if len(cookies) == 0 {
				return nil
			}
			ch.Config.Session = session.New(kind.Session, cookies[0].Value, cookies[0].Expires)

			err := ch.Config.WriteSession()
			if err != nil {
				return fmt.Errorf("error updating session: %w", err)
			}
		}
		analytics.Cleanup()
		return nil
	}
}

func PostRunAnalytics(_ *cmdutils.Helper[*config.Config]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		analytics.Cleanup()
		return nil
	}
}

func InvalidStringError(kind string, str string) error {
	return fmt.Errorf("invalid %s '%s', %ss can only include letters, numbers, periods (`.`), and dashes (`-`)", kind, str, kind)
}

type Parsed struct {
	Organization string
	Name         string
	Tag          string
}

// Parse parses a function or signature name of the form <org>/<name>:<tag> into its organization, name, and tag
func Parse(name string) *Parsed {
	orgSplit := strings.Split(name, "/")
	if len(orgSplit) == 1 {
		orgSplit = []string{"", name}
	}
	tagSplit := strings.Split(orgSplit[1], ":")
	if len(tagSplit) == 1 {
		tagSplit = []string{tagSplit[0], ""}
	}
	return &Parsed{
		Organization: orgSplit[0],
		Name:         tagSplit[0],
		Tag:          tagSplit[1],
	}
}
