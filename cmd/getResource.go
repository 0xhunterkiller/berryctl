/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getResourceCmd represents the getResource command
var getResourceCmd = &cobra.Command{
	Use:   "resource [Name of the Resource]",
	Args: cobra.ExactArgs(1),
	Short: "Get a resource",	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Name of the Resource:", args[0])
	},
}

func init() {
	getCmd.AddCommand(getResourceCmd)
}
