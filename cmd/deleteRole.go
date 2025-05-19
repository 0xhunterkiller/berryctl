/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteRoleCmd represents the deleteRole command
var deleteRoleCmd = &cobra.Command{
	Use:   "role [Name of the Role]",
	Args: cobra.ExactArgs(1),
	Short: "Delete a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the Role:", args[0])
		yes, _ := cmd.Flags().GetBool("yes")
		if yes {
			fmt.Println("Role deletion confirmed.")
		} else {
			fmt.Println("Role deletion not confirmed. Use --yes to skip confirmation.")
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteRoleCmd)
	
	deleteRoleCmd.Flags().BoolP("yes", "y", false, "Skip Confirmation by setting this flag")
}
