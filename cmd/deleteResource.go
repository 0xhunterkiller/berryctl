package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteResourceCmd represents the deleteResource command
var deleteResourceCmd = &cobra.Command{
	Use:   "resource [name]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the Resource:", args[0])
		yes, _ := cmd.Flags().GetBool("yes")
		if yes {
			fmt.Println("Resource deletion confirmed.")
		} else {
			fmt.Println("Resource deletion not confirmed. Use --yes to skip confirmation.")
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteResourceCmd)

	deleteResourceCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
