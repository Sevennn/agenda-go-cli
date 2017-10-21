// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_t, _ := cmd.Flags().GetString("title")
		tmp_p, _ := cmd.Flags().GetString("participator")
		tmp_s, _ := cmd.Flags().GetString("starttime")
		tmp_e, _ := cmd.Flags().GetString("endtime")
		fmt.Println("createmeeting args : ", tmp_t, tmp_p, tmp_s, tmp_e)
	},
}

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Here you will define your flags and configuration settings.
	createmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	createmeetingCmd.Flags().StringP("participator", "p", "", "the participator(s) of the meeting, split by comma")
	createmeetingCmd.Flags().StringP("starttime","s","","the startTime of the meeting")
	createmeetingCmd.Flags().StringP("endtime", "e", "", "the endTime of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
