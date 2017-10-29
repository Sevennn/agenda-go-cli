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
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// quitmeetingCmd represents the quitmeeting command
var quitmeetingCmd = &cobra.Command{
	Use:   "quitmeeting",
	Short: "quit a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Quit Meeting called")
		tmp_t, _ := cmd.Flags().GetString("title")
		if tmp_t == "" {
			fmt.Println("Please input meeting title")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: Please login firstly!")
		} else {
			if flag := service.QuitMeeting(user.Name, tmp_t); flag == false {
				fmt.Println("Error: Meeting not exist or you're not a participator of it")
			} else {
				fmt.Println("Quit Successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(quitmeetingCmd)

	// Here you will define your flags and configuration settings.
	quitmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
