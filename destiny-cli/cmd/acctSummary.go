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

// acctSummaryCmd represents the acctSummary command
var acctSummaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "get a summary of an account",
	Long:  "get a summary of an account",
	Run: func(cmd *cobra.Command, args []string) {
		for idx, member := range args {
			if idx != 0 {
				os.Stdout.Write([]byte("\n"))
			}
			membership, err := api.Memberships(member)
			if err != nil {
				os.Stdout.Write([]byte("{}"))
				log.Println(err.Error())
				continue
			}
			buf, err := membership.RawAccountSummary()
			if err != nil {
				log.Println(err.Error())
				os.Stdout.Write([]byte("{}"))
				continue
			}
			os.Stdout.Write(buf)
		}
	},
}

func init() {
	acctCmd.AddCommand(acctSummaryCmd)
}
