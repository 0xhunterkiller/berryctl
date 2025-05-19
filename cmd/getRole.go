/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getRoleCmd represents the getRole command
var getRoleCmd = &cobra.Command{
	Use:   "role [Name of the Role]",
	Args: cobra.ExactArgs(1),
	Short: "Get a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the Role:", args[0])
	},
}

func init() {
	getCmd.AddCommand(getRoleCmd)
}
