package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Display a line of text",
	Long: `Print the STRING(s) to standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flip, err := cmd.Flags().GetBool("flip"); err == nil {
			if flip {
				for i := 0; i < len(args)/2; i++ {
					j := len(args) - i - 1
					args[i], args[j] = args[j], args[i]
				}	
			} 
			fmt.Println(strings.Join(args," "))
		}
	},
}
