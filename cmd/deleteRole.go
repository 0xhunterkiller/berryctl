package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// deleteRoleCmd represents the deleteRole command
var deleteRoleCmd = &cobra.Command{
	Use:   "role [Name of the Role]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a role",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		yes, _ := cmd.Flags().GetBool("yes")
		if !yes {
			var input string
			fmt.Printf("Are you sure you want to delete role '%s'? [yes/No]: ", name)
			fmt.Scanln(&input)

			if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
				fmt.Println("Aborted: Role deletion not confirmed.")
				return
			}
			fmt.Println("Role deletion confirmed.")
		} else {
			fmt.Printf("Role: %s deletion confirmed via --yes flag.\n", name)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteRoleCmd)

	deleteRoleCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
