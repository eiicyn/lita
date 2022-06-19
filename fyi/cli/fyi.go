package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// root command
var fyiCmd = &cobra.Command{
	Use:   "fyi",
	Short: "for your information",
	Long:  `For your information!`,
}

func Execute() {
	if err := fyiCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	fyiCmd.CompletionOptions.DisableDefaultCmd = true
	fyiCmd.CompletionOptions.DisableNoDescFlag = true

	fyiCmd.AddCommand(eatCmd)
	fyiCmd.AddCommand(drinkCmd)
	fyiCmd.AddCommand(chatCmd)
	fyiCmd.AddCommand(hangoutCmd)
	fyiCmd.AddCommand(callbackCmd)
	fyiCmd.AddCommand(sayCmd)
}
