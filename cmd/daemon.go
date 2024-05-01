/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/kuchihira-bot/internal/core"
	"github.com/spf13/cobra"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "post kuchiwohiraku new item in daemon mode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("daemon called")
		if err := core.DaemonStart(debug); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().BoolVar(&debug, "debug", false, "debug(default: false)")
}
