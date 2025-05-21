package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// canCmd represents the can command
var canCmd = &cobra.Command{
	Use:   "can [Name of the User]",
	Args:  cobra.ExactArgs(1),
	Short: "check the access level of a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("can called")
	},
}

func init() {
	rootCmd.AddCommand(canCmd)
	canCmd.Flags().String("do", "", "resource|verb - in this format")
}
