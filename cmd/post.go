/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/kuchihira-bot/internal/core"
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "post kuchiwohiraku new item",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("post called")
		if err := core.Start(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}
