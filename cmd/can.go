package cmd

import (
	"github.com/spf13/cobra"
)

// canCmd represents the can command
var canCmd = &cobra.Command{
	Use:   "can [Name of the User]",
	Args:  cobra.ExactArgs(1),
	Short: "check the access level of a user",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(canCmd)
	canCmd.Flags().StringArray("access", []string{}, "Access rule in format resource=verb1,verb2 (repeatable)")
}
