/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/kuchihira-bot/internal/bsky"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "setup login info to bluesky",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		bsky.DoLogin(host, handle, password)
	},
}

var (
	host     string
	handle   string
	password string
)

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&host, "host", "s", "https://bsky.social", "Using host")
	loginCmd.Flags().StringVarP(&handle, "handle", "u", "", "handle (Username)")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "password")
}
