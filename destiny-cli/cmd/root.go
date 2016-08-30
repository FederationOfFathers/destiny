// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FederationOfFathers/destiny"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var apiKey = os.Getenv("DESTINY_API_KEY")
var platform = os.Getenv("DESTINY_API_PLATFORM")

var api *destiny.Platform

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "destiny-cli",
	Short: "Access the destiny API from the command line",
	Long: `See the following for more details:

	https://www.bungie.net/en/User/API

	https://www.bungie.net/en/Clan/Forum/39966

	https://www.bungie.net/platform/destiny/help/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(platform[:1]) {
		case "x":
			api = destiny.New(apiKey, nil).XBL()
		case "p":
			api = destiny.New(apiKey, nil).PSN()
		default:
			log.Fatal("Valid platforms are x[bl] or p[sn]")
		}
	},
}

// Execute ...
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	if platform == "" {
		platform = "xbl"
	}
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.destiny-cli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", apiKey, "your API key (see: https://www.bungie.net/en/User/API)")
	RootCmd.PersistentFlags().StringVarP(&platform, "platform", "p", platform, "platform to access the api for")
	RootCmd.PersistentFlags().BoolVarP(&destiny.Debug, "debug", "d", destiny.Debug, "debug http requests")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".destiny-cli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")        // adding home directory as first search path
	viper.AutomaticEnv()                // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
