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
	// "strings"
	"agenda-go-cli/service"
)

// addparticipatorCmd represents the addparticipator command
var addparticipatorCmd = &cobra.Command{
	Use:   "addparticipator",
	Short: "add participators",
	Long: `This is a command to add participator(s) to a a meeting specified by title`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Add participator called")
		tmp_p, _ := cmd.Flags().GetStringSlice("participator")
		tmp_t, _ := cmd.Flags().GetString("title")
		if len(tmp_p) == 0 || tmp_t == "" {
			fmt.Println("Please input title and participator(s)(input like \"name1, name2\")")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please login firstly")
		} else {
			// participators := strings.Split(tmp_p,",")
			flag := service.AddMeetingParticipator(user.Name, tmp_t, tmp_p)
			if flag != true {
				fmt.Println("Unexpected error. Check error.log for detail")
			} else {
				fmt.Println("Successfully add")
			}
		}
	},
}


func init() {
	RootCmd.AddCommand(addparticipatorCmd)

	// Here you will define your flags and configuration settings.
	addparticipatorCmd.Flags().StringSliceP("participator", "p", nil, "participator(s) you want to add, input like \"name1, name2\"")
	addparticipatorCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addparticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addparticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
