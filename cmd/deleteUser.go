package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "user [Name of the User]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		yes, _ := cmd.Flags().GetBool("yes")
		if !yes {
			var input string
			fmt.Printf("Are you sure you want to delete user '%s'? [yes/No]: ", name)
			fmt.Scanln(&input)

			if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
				fmt.Println("Aborted: User deletion not confirmed.")
				return
			}
			fmt.Println("User deletion confirmed.")
		} else {
			fmt.Printf("User: %s deletion confirmed via --yes flag.\n", name)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteUserCmd)

	deleteUserCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
