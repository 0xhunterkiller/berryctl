/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createResourceCmd represents the createResource command
var createResourceCmd = &cobra.Command{
	Use:   "resource [Name of the Resource]",
	Args: cobra.ExactArgs(1),
	Short: "Create a new resource",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		desc := ""
		verbs := []string{}

		if cmd.Flags().Lookup("desc").Changed {
			desc, _ = cmd.Flags().GetString("desc")
		}

		if cmd.Flags().Lookup("verbs").Changed {
			verbsStr, _ := cmd.Flags().GetString("verbs")
			verbs = strings.Split(verbsStr, ",")
		}

		fmt.Printf("Name: %s\n", name)
		if desc != "" {
			fmt.Printf("Description: %s\n", desc)
		}
		if len(verbs) > 0 {
			fmt.Printf("Verbs: %s\n", strings.Join(verbs, ", "))
		}
	},
}

func init() {
	createCmd.AddCommand(createResourceCmd)

	createResourceCmd.Flags().StringP("desc", "d", "", "Description of the resource")
	createResourceCmd.Flags().StringP("verbs", "e", "", "Comma separated list of verbs; verb1,verb2,verb3")
}
