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
	"log"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an xbox live membership",
	Long: `Search for an xbox live membership

destiny-cli --platform [x|p] search <GamerTag>`,
	Run: func(cmd *cobra.Command, args []string) {
		membership, err := api.Memberships(args[0])
		if err != nil {
			log.Fatal(err)
		}
		jsonOut(membership)
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
}
