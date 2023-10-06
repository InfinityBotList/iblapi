/*
Copyright © 2022 Infinity Bot List
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/InfinityBotList/iblapi/internal/ui"
	"github.com/InfinityBotList/iblapi/internal/views"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:     "login TYPE",
	Short:   "Login to the IBL API",
	Long:    `Login to the IBL API using a bot/user/server token.`,
	Aliases: []string{"auth", "a", "l"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := views.AccountSwitcher(args[0])

		if err != nil {
			fmt.Print(ui.RedText("Error logging in: " + err.Error()))
			os.Exit(1)
		}
	},
}

func init() {
	// login
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
