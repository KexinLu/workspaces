package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	userLicense string
	verbose bool

	rootCmd = &cobra.Command{
		Use: "works",
		Short: "works is a workspace management tool",
		Long: `Set up your workspace with one click on a brand new work station, or navigate between your projects with ease`,
	}
)

func Execute() {
	if verbose {
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func preRun() {


}
