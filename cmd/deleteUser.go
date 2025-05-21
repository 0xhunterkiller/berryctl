package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "user [Name of the User]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the User:", args[0])
		yes, _ := cmd.Flags().GetBool("yes")
		if yes {
			fmt.Println("User deletion confirmed.")
		} else {
			fmt.Println("User deletion not confirmed. Use --yes to skip confirmation.")
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteUserCmd)

	deleteUserCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
