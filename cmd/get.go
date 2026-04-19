package cmd

import (
	"fmt"
	"os"

	"github.com/AmadlaOrg/doorman/plugin"
	"github.com/spf13/cobra"
)

var (
	getFrom string

	getPluginNew = plugin.New

	// GetCmd retrieves a secret via a doorman plugin.
	GetCmd = &cobra.Command{
		Use:   "get <key>",
		Short: "Retrieve a secret from a plugin",
		Long:  "Retrieves a secret by key using the specified doorman-* plugin (--from flag).",
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func init() {
	GetCmd.Flags().StringVar(&getFrom, "from", "", "Plugin name to retrieve from (e.g. vault, keepassxc, bitwarden)")
	_ = GetCmd.MarkFlagRequired("from")
}

func runGet(cmd *cobra.Command, args []string) error {
	key := args[0]
	pluginName := "doorman-" + getFrom

	svc := getPluginNew()
	err := svc.Get(pluginName, key, os.Stdout)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	return nil
}
