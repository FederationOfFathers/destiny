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
	"os"

	"github.com/spf13/cobra"
)

// charStatsCmd represents the charStats command
var charStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get stats for a character id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("You must supply exactly one display name and 1-3 character IDs")
		}
		if len(args) > 4 {
			log.Fatal("You must supply exactly one display name and 1-3 character IDs")
		}
		member, err := api.Memberships(args[0])
		if err != nil {
			log.Fatal(err)
		}
		for idx, char := range args[1:] {
			if idx != 0 {
				os.Stdout.Write([]byte("\n"))
			}
			raw, err := member.RawCharacterStats(char)
			if err != nil {
				os.Stdout.Write([]byte("{}"))
				log.Println(err)
				continue
			}
			os.Stdout.Write(raw)
		}
	},
}

func init() {
	charCmd.AddCommand(charStatsCmd)
}
