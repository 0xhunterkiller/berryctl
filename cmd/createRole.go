package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var createRoleCmd = &cobra.Command{
	Use:   "role [Name of the Role]",
	Short: "Create a new role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		desc, _ := cmd.Flags().GetString("desc")
		accessList, _ := cmd.Flags().GetStringArray("access")

		accessMap := parseAccess(accessList)

		// Output
		fmt.Printf("Name: %s\n", name)
		if desc != "" {
			fmt.Printf("Description: %s\n", desc)
		}
		if len(accessMap) > 0 {
			fmt.Println("Access Rules:")
			for resource, verbs := range accessMap {
				fmt.Printf("  %s: %s\n", resource, strings.Join(verbs, ", "))
			}
		}
	},
}

func init() {
	createCmd.AddCommand(createRoleCmd)

	createRoleCmd.Flags().StringP("desc", "d", "", "Description of the Role")
	createRoleCmd.Flags().StringArray("access", []string{}, "Access rule in format resource=verb1,verb2 (repeatable)")
}

func parseAccess(entries []string) map[string][]string {
	accessMap := make(map[string][]string)

	for _, entry := range entries {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Skipping invalid access entry: %s (expected format: resource=verb1,verb2)\n", entry)
			continue
		}
		resource := strings.TrimSpace(parts[0])
		verbs := strings.Split(parts[1], ",")
		for i := range verbs {
			verbs[i] = strings.TrimSpace(verbs[i])
		}
		accessMap[resource] = verbs
	}

	return accessMap
}
