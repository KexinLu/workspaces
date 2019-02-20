package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "works",
	Short: "works is a workspace management tool",
	Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
