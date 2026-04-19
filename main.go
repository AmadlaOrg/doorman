package main

import (
	"fmt"
	"os"

	"github.com/AmadlaOrg/doorman/cmd"
	"github.com/spf13/cobra"
)

const (
	appName = "doorman"
	version = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:     appName,
	Short:   "Secrets management CLI with doorman-* plugins",
	Version: version,
}

func init() {
	rootCmd.AddCommand(cmd.GetCmd)
	rootCmd.AddCommand(cmd.PluginsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
