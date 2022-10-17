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
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "scale",
	Short: "The Scale CLI is a command line interface for working with Scale Functions",
	Long: `The Scale CLI is a command line interface for working with Scale Functions, the 
Scale Registry, and the Scale Build Service.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.config/scale/scale.json)")
	rootCmd.PersistentFlags().Bool("debug", false, "debug output")
	rootCmd.PersistentFlags().String("app", "https://app.scale.sh", "Scale API URL")
	rootCmd.PersistentFlags().String("builder", "build.scale.sh:8192", "Scale Build Service URL")

	err := viper.BindPFlag("app", rootCmd.PersistentFlags().Lookup("app"))
	if err != nil {
		panic(err)
	}

	err = viper.BindPFlag("builder", rootCmd.PersistentFlags().Lookup("builder"))
	if err != nil {
		panic(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigFile(path.Join(homeDir, ".config", "scale", "config.json"))
}
