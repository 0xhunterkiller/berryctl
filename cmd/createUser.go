/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "user [Name of the User]",
	Args: cobra.ExactArgs(1),
	Short: "Create a new user",
	Run: func(cmd *cobra.Command, args []string) {
		// write this function to create a user
		name := args[0]
		desc, _ := cmd.Flags().GetString("desc")
		roles, _ := cmd.Flags().GetString("roles")
		rolesList := []string{}
		if roles != "" {
			rolesList = strings.Split(roles, ",")
		}
		// Output
		fmt.Printf("Name: %s\n", name)
		if desc != "" {
			fmt.Printf("Description: %s\n", desc)
		}
		if len(rolesList) > 0 {
			fmt.Println("Roles:")
			for _, role := range rolesList {
				fmt.Printf("- %s\n", role)
			}
		}
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("desc", "d", "", "Description of the User")
	createUserCmd.Flags().StringP("roles", "q", "", "Comma separated list of roles; role1,role2,role3")
}
