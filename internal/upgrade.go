package internal

import (
	"fmt"
	"os"

	"github.com/jetrails/proposal-nginx/sdk/vhost"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade DOMAIN KEY=VALUE...",
	Short:   "upgrade a vhost from a template",
	Args:    cobra.MinimumNArgs(1),
	PreRunE: ValidateKeyValueArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domainName := args[0]
		parameters := ParseKeyValueArgs(args[1:])
		if err := vhost.Upgrade(domainName, parameters); err != nil {
			fmt.Printf("\nError: %s\n\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(upgradeCmd)
}
