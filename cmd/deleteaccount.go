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

// deleteaccountCmd represents the deleteaccount command
var deleteaccountCmd = &cobra.Command{
	Use:   "deleteaccount",
	Short: "to delete current user",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Delete account called")
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("You need login firstly!")
		} else {
			if dflag := service.DeleteUser(user.Name); dflag != true {
				fmt.Println("Error occurred when delete account.Please check error.log")
			} else {
				fmt.Println("Successfully Delete")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteaccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteaccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteaccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
