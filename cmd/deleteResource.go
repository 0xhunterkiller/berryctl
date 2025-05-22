package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// deleteResourceCmd represents the deleteResource command
var deleteResourceCmd = &cobra.Command{
	Use:   "resource [name]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a resource",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		yes, _ := cmd.Flags().GetBool("yes")
		if !yes {
			var input string
			fmt.Printf("Are you sure you want to delete resource '%s'? [yes/No]: ", name)
			fmt.Scanln(&input)

			if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
				fmt.Println("Aborted: Resource deletion not confirmed.")
				return
			}
			fmt.Println("Resource deletion confirmed.")
		} else {
			fmt.Printf("Resource: %s deletion confirmed via --yes flag.\n", name)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteResourceCmd)

	deleteResourceCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
