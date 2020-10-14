/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/database"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new bot into the database.",
	Long:  `Adds a new bot into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.EnsureExists()
		if err != nil {
			log.Fatal(err)
		}

		err = database.AddBot(username, password, "twitter")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().StringVarP(&username, "username", "u", "", "Username to be used to log in the bot.")
	addCmd.MarkFlagRequired("username")

	addCmd.Flags().StringVarP(&password, "password", "p", "", "Password to be used to log in the bot.")
	addCmd.MarkFlagRequired("password")
}
