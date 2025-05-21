package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getUserCmd represents the getUser command
var getUserCmd = &cobra.Command{
	Use:   "user [Name of the User]",
	Args:  cobra.ExactArgs(1),
	Short: "Get a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the User:", args[0])
	},
}

func init() {
	getCmd.AddCommand(getUserCmd)
}
