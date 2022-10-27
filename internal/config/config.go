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
	TokenFileMode     = 0o600
)

type Token struct {
	AccessToken  string         `yaml:"access_token"`
	TokenType    string         `yaml:"token_type,omitempty"`
	RefreshToken string         `yaml:"refresh_token,omitempty"`
	Expiry       time.Time      `yaml:"expiry,omitempty"`
	Kind         tokenKind.Kind `yaml:"kind"`
}

func FromClientToken(t *client.Token, kind tokenKind.Kind) *Token {
	return &Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		Kind:         kind,
	}
}

// Config is dynamically sourced from various files and environment variables.
type Config struct {
	BaseURL string `yaml:"base_url"`
	Token   *Token `yaml:"token"`
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
		BaseURL: "https://api.scale.sh",
		Token:   &token,
	}, nil
}

func (c *Config) IsAuthenticated() error {
	if c.Token == nil || c.Token.AccessToken == "" {
		return errors.New("access token is empty")
	}

	return nil
}

type ClientOption func(c *apiClient.ScaleAPIV1) error

var (
	WithUserAgent = func(userAgent string) ClientOption {
		return func(c *apiClient.ScaleAPIV1) error {
			//c.UserAgent = userAgent
			return nil
		}
	}
	WithRequestHeaders = func(headers map[string]string) ClientOption {
		return func(c *apiClient.ScaleAPIV1) error {
			//c.RequestHeaders = headers
			return nil
		}
	}
)

// NewClientFromConfig creates a Scale API client from our configuration
func (c *Config) NewClientFromConfig(opts ...ClientOption) (*apiClient.ScaleAPIV1, error) {
	return apiClient.NewHTTPClientWithConfig(nil, &apiClient.TransportConfig{
		Host:     "",
		BasePath: "",
		Schemes:  nil,
	}), nil
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
