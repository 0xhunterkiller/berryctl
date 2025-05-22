package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/0xhunterkiller/berry/pkgs/models"
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
		resourceItems := []models.ResourceItem{}

		if len(accessMap) > 0 {
			for resource, verbs := range accessMap {
				resItem, err := models.NewResourceItem(resource, verbs)
				if err != nil {
					panic(err)
				}
				resourceItems = append(resourceItems, *resItem)
			}
		}

		res, err := models.NewRole("v1", name, desc, resourceItems)
		if err != nil {
			panic(err)
		}

		jsonData, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))
	},
}

func init() {
	createCmd.AddCommand(createRoleCmd)

	createRoleCmd.Flags().StringP("desc", "d", "", "Description of the Role")
	createRoleCmd.Flags().StringArray("access", []string{}, "Access rule in format resource=verb1,verb2 (repeatable)")
}
