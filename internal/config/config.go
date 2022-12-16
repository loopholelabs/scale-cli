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

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/loopholelabs/auth/pkg/client"
	"github.com/loopholelabs/auth/pkg/token/tokenKind"
	apiClient "github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/mitchellh/go-homedir"
	exec "golang.org/x/sys/execabs"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const (
	defaultConfigPath = "~/.config/scale"
	projectConfigName = ".scale.yml"
	configName        = "scale.yml"
	TokenFileMode     = 0600
)

type Token struct {
	AccessToken  string         `yaml:"access_token"`
	TokenType    string         `yaml:"token_type,omitempty"`
	RefreshToken string         `yaml:"refresh_token,omitempty"`
	Expiry       time.Time      `yaml:"expiry,omitempty"`
	Kind         tokenKind.Kind `yaml:"kind"`
	Endpoint     string         `yaml:"endpoint"`
	BasePath     string         `yaml:"base_path"`
	ClientID     string         `yaml:"client_id"`
}

func FromClientToken(t *client.Token, kind tokenKind.Kind, endpoint string, basePath string, clientID string) *Token {
	return &Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		Kind:         kind,
		Endpoint:     endpoint,
		BasePath:     basePath,
		ClientID:     clientID,
	}
}

// Config is dynamically sourced from various files and environment variables.
type Config struct {
	Endpoint string `yaml:"endpoint"`
	Build    string `yaml:"build"`
	Token    *Token `yaml:"token"`
}

func New() (*Config, error) {
	var token Token
	tokenPath, err := TokenPath()
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(tokenPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	} else {
		if stat.Mode()&^TokenFileMode != 0 {
			err = os.Chmod(tokenPath, TokenFileMode)
			if err != nil {
				log.Printf("Unable to change %v file mode to 0%o: %v", tokenPath, TokenFileMode, err)
			}
		}
		tokenData, err := os.ReadFile(tokenPath)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(tokenData, &token)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &Config{
		Endpoint: "https://api.scale.sh",
		Build:    "build.scale.sh:8192",
		Token:    &token,
	}, nil
}

func (c *Config) IsAuthenticated() error {
	if c.Token == nil || c.Token.AccessToken == "" {
		return errors.New("access token is empty")
	}

	return nil
}

// NewAuthenticatedClientFromConfig creates an Authenticated Scale API client from our configuration
func (c *Config) NewAuthenticatedClientFromConfig(endpoint string, t *Token) (*apiClient.ScaleAPIV1, error) {
	_, cl, err := client.AuthenticatedClient(endpoint, apiClient.DefaultBasePath, apiClient.DefaultSchemes, nil, path.Join(t.Endpoint, t.BasePath), t.ClientID, t.Kind, client.NewToken(t.AccessToken, t.TokenType, t.RefreshToken, t.Expiry))
	if err != nil {
		return nil, err
	}

	return apiClient.New(cl, nil), nil
}

// NewUnauthenticatedClientFromConfig creates an Unauthenticated Scale API client from our configuration
func (c *Config) NewUnauthenticatedClientFromConfig(endpoint string) (*apiClient.ScaleAPIV1, error) {
	cl, _ := client.UnauthenticatedClient(endpoint, apiClient.DefaultBasePath, apiClient.DefaultSchemes, nil)
	return apiClient.New(cl, nil), nil
}

// TokenPath is the path for the access token file
func TokenPath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, "token"), nil
}

// Dir is the directory for Scale config.
func Dir() (string, error) {
	dir, err := homedir.Expand(defaultConfigPath)
	if err != nil {
		return "", fmt.Errorf("can't expand path %q: %s", defaultConfigPath, err)
	}

	return dir, nil
}

func RootGitRepoDir() (string, error) {
	tl := []string{"rev-parse", "--show-toplevel"}
	out, err := exec.Command("git", tl...).CombinedOutput()
	if err != nil {
		return "", errors.New("unable to find git root directory")
	}

	return strings.TrimSuffix(string(out), "\n"), nil
}

func ProjectConfigFile() string {
	return projectConfigName
}

// DefaultConfigPath returns the default path for the config file.
func DefaultConfigPath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}

	return path.Join(dir, configName), nil
}
