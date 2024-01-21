package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	loadConfig("$HOME", ".")
}

var rootCmd = &cobra.Command{
	Use:   "datw",
	Short: "DAtW is a website for tracking your Epcot drink around the world challenge",
	Long: `A simple website / web app for tracking a users drink around the world challenge

	- website: https://drinkaroundthe.world`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
