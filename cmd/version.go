package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version prints the information for simple-cli",
	Long: `Simple-cli version reports the simple-cli version used to build each of the named
	executable files.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
