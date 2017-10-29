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
	"agenda-go-cli/entity"
	"github.com/spf13/cobra"
)

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "query meetings in a time interval",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Query Meeting called")
		tmp_s, _ := cmd.Flags().GetString("starttime")
		tmp_e, _ := cmd.Flags().GetString("endtime")
		if tmp_s == "" || tmp_e == "" {
			fmt.Println("Please input start time and end time both")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Please log in firstly")
		} else {
			if ml,flag :=service.QueryMeeting(user.Name, tmp_s,tmp_e); flag != true {
				fmt.Println("Wrong Date!please input the date as yyyy-mm-dd/hh:mm and make sure that starttiem <= endtime")
			} else {
				for _, m := range ml {
					fmt.Println("----------------")
					fmt.Println("Title: ", m.Title)
					ts,_ := entity.DateToString(m.StartDate)
					fmt.Println("Start Time", ts)
					te,_ := entity.DateToString(m.EndDate)
					fmt.Println("End Time", te)
					fmt.Printf("Participator(s): ")
					for _, p := range m.Participators {
						fmt.Printf(p," ")
					}
					fmt.Printf("\n")
					fmt.Println("----------------")
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(querymeetingCmd)

	// Here you will define your flags and configuration settings.
	querymeetingCmd.Flags().StringP("starttime", "s", "", "the start time of the meeting")
	querymeetingCmd.Flags().StringP("endtime", "e", "", "the end time of the meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
