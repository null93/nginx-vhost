package internal

import (
	"fmt"
	"os"

	"github.com/jetrails/proposal-nginx/sdk/vhost"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete DOMAIN",
	Short:   "delete a vhost given the domain name",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domainName := args[0]
		if err := vhost.Delete(domainName); err != nil {
			fmt.Printf("\nError: %s\n\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}