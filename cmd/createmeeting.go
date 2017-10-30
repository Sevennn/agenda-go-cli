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
	// "strings"
	"github.com/spf13/cobra"
	"agenda-go-cli/service"
)

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "create meeting command",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Create Meeting Called")
		tmp_t, _ := cmd.Flags().GetString("title")
		tmp_p, _ := cmd.Flags().GetStringSlice("participator")
		tmp_s, _ := cmd.Flags().GetString("starttime")
		tmp_e, _ := cmd.Flags().GetString("endtime")
		if tmp_t == "" || len(tmp_p) == 0 || tmp_s == "" || tmp_e == "" {
			fmt.Println("Please input title, starttime[yyyy-mm-dd/hh:mm],endtime,participator(input like \"name1, name2\")")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login firstly")
			return
		} else {
			// participators := strings.Split(tmp_p,",")
			if flag := service.CreateMeeting(user.Name, tmp_t,tmp_s,tmp_e,tmp_p); flag != true {
				fmt.Println("Error: create Failed. Please check error.log for more detail")
				return
			} else {
				fmt.Println("Create meeting successfully!")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Here you will define your flags and configuration settings.
	createmeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")
	createmeetingCmd.Flags().StringSliceP("participator", "p", nil, "the participator(s) of the meeting, split by comma,input like \"name1, name2\"")
	createmeetingCmd.Flags().StringP("starttime","s","","the startTime of the meeting")
	createmeetingCmd.Flags().StringP("endtime", "e", "", "the endTime of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
