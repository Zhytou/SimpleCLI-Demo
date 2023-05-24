package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "simple-cli",
    Short: "A simple command line interface demo",
    Long: `Simple-cli is a CLI tool based on Cobra.
This application is a demo to show how to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}


// Execute executes the root command.
func Execute() error {
    return rootCmd.Execute()
}

func init(){
    rootCmd.AddCommand(versionCmd)
    rootCmd.AddCommand(echoCmd)
}