/*
Copyright © 2020 Andreas Gajdosik <andreas.gajdosik@gmail.com>

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

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/agajdosi/tauto/pkg/twitter"
	"github.com/spf13/cobra"
)

var likeCmd = &cobra.Command{
	Use:   "like",
	Short: "Likes tweet on URL with one or more bots.",
	Long:  `Likes tweet on URL with one or more bots.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		err := like()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(likeCmd)

	likeCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the bot who will like. When left empty it will use all bots available in the database.")

	likeCmd.Flags().StringSliceVarP(&urls, "url", "U", nil, "Where to like - an url of the tweet(s).")
	likeCmd.MarkFlagRequired("url")
}

func like() error {
	bots, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, bot := range bots {
		b, cancel, err := twitter.NewUser(bot.ID, bot.Username, bot.Password, 300)
		if err != nil {
			return err
		}

		for _, url := range urls {
			b.EnsureLiked(url)
		}
		cancel()
	}

	return nil
}
