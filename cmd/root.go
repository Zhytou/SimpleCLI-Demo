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

func Execute() error {
    return rootCmd.Execute()
}

func init(){
    // add flag
    rootCmd.Flags().BoolP("version", "v", false, "print simple-cli version")

    // add subcommand flag
    echoCmd.Flags().BoolP("flip", "f", false, "print in reverse order")
    
    // add subcommand
    rootCmd.AddCommand(versionCmd)
    rootCmd.AddCommand(echoCmd)
}