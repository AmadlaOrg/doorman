package main

import (
	"github.com/AmadlaOrg/LibraryFramework/cli"
	"github.com/AmadlaOrg/doorman/internal/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cli.New(
		"doorman",
		"Doorman",
		"1.0.0",
		func(rootCmd *cobra.Command) {
			rootCmd.AddCommand(cmd.SettingsCmd)
			//rootCmd.AddCommand(cmd.CollectionCmd)
			//rootCmd.AddCommand(cmd.ComposeCmd)
		})
}
