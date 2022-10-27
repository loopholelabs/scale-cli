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

package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/loopholelabs/scale-cli/cmd/auth"
	"github.com/loopholelabs/scale-cli/cmd/version"
	"github.com/loopholelabs/scale-cli/internal/cmdutil"
	"github.com/loopholelabs/scale-cli/internal/config"
	"github.com/loopholelabs/scale-cli/internal/printer"
	"github.com/loopholelabs/scale-cli/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	cfgFile  string
	replacer = strings.NewReplacer("-", "_", ".", "_")
)

var rootCmd = &cobra.Command{
	Use:              "scale",
	Short:            "A CLI for Scale Functions",
	Long:             `scale is a CLI for working with Scale Functions and communicating with the Scale API.`,
	TraverseChildren: true,
}

// Execute executes the command and returns the exit status of the finished
// command.
func Execute(ctx context.Context, ver, commit, buildDate string) int {
	var format printer.Format
	var debug bool

	if _, ok := os.LookupEnv("SCALE_DISABLE_DEV_WARNING"); !ok {
		if commit == "" || ver == "" || buildDate == "" {
			_, _ = fmt.Fprintf(os.Stderr, "!! WARNING: You are using a self-compiled binary which is not officially supported.\n!! To dismiss this warning, set SCALE_DISABLE_DEV_WARNING=true\n\n")
		}
	}

	err := runCmd(ctx, ver, commit, buildDate, &format, &debug)
	if err == nil {
		return 0
	}

	// print any user specific messages first
	switch format {
	case printer.JSON:
		_, _ = fmt.Fprintf(os.Stderr, `{"error": "%s"}`, err)
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	// check if a sub command wants to return a specific exit code
	var cmdErr *cmdutil.Error
	if errors.As(err, &cmdErr) {
		return cmdErr.ExitCode
	}

	return cmdutil.FatalErrExitCode
}

// runCmd adds all child commands to the root command, sets flags
// appropriately, and runs the root command.
func runCmd(ctx context.Context, ver, commit, buildDate string, format *printer.Format, debug *bool) error {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config",
		"", "Config file (default is $HOME/.config/scale/scale.yml)")
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	v := version.Format(ver, commit, buildDate)
	rootCmd.SetVersionTemplate(v)
	rootCmd.Version = v
	rootCmd.Flags().Bool("version", false, "Show scale cli version")

	cfg, err := config.New()
	if err != nil {
		return err
	}

	rootCmd.PersistentFlags().StringVar(&cfg.BaseURL, "api", "https://api.scale.sh", "The base URL for the Scale API.")

	rootCmd.PersistentFlags().VarP(printer.NewFormatValue(printer.Human, format), "format", "f",
		"Show output in a specific format. Possible values: [human, json, csv]")
	if err := viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format")); err != nil {
		return err
	}
	_ = rootCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"human", "json", "csv"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().BoolVar(debug, "debug", false, "Enable debug mode")
	if err := viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug")); err != nil {
		return err
	}

	ch := &cmdutil.Helper{
		Printer: printer.NewPrinter(format),
		Config:  cfg,
		Client: func() (*client.ScaleAPIV1, error) {
			userAgent := "scale-cli/" + ver
			headers := map[string]string{
				"scale-cli-version": ver,
			}
			return cfg.NewClientFromConfig(config.WithUserAgent(userAgent), config.WithRequestHeaders(headers))
		},
	}
	ch.SetDebug(debug)

	rootCmd.PersistentFlags().BoolVar(&color.NoColor, "no-color", false, "Disable color output")
	if err := viper.BindPFlag("no-color", rootCmd.PersistentFlags().Lookup("no-color")); err != nil {
		return err
	}

	loginCmd := auth.LoginCmd(ch)
	loginCmd.Hidden = true
	logoutCmd := auth.LogoutCmd(ch)
	logoutCmd.Hidden = true

	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(auth.Cmd(ch))
	rootCmd.AddCommand(version.Cmd(ch, ver, commit, buildDate))

	return rootCmd.ExecuteContext(ctx)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		defaultConfigDir, err := config.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(cmdutil.FatalErrExitCode)
		}

		// Order of preference for configuration files:
		// (1) $HOME/.config/scale
		viper.AddConfigPath(defaultConfigDir)
		viper.SetConfigName("scale")
		viper.SetConfigType("yml")
	}

	viper.SetEnvPrefix("SCALE")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Only handle errors when it's something unrelated to the config file not
			// existing.
			fmt.Println(err)
			os.Exit(cmdutil.FatalErrExitCode)
		}
	}

	// Check for a project-local configuration file to merge in if the user
	// has not specified a config file
	if rootDir, err := config.RootGitRepoDir(); err == nil && cfgFile == "" {
		viper.AddConfigPath(rootDir)
		viper.SetConfigName(config.ProjectConfigFile())
		_ = viper.MergeInConfig()
	}

	postInitCommands(rootCmd.Commands())
}

// Hacky fix for getting Cobra required flags and Viper playing well together.
// See: https://github.com/spf13/viper/issues/397
func postInitCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		presetRequiredFlags(cmd)
		if cmd.HasSubCommands() {
			postInitCommands(cmd.Commands())
		}
	}
}

func presetRequiredFlags(cmd *cobra.Command) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("error binding flags: %v", err)
	}

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			err = cmd.Flags().Set(f.Name, viper.GetString(f.Name))
			if err != nil {
				log.Fatalf("error setting flag %s: %v", f.Name, err)
			}
		}
	})
}
