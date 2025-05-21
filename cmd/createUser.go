package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xhunterkiller/berry/pkgs/models"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "user [Name of the User]",
	Args:  cobra.ExactArgs(1),
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
		res, err := models.NewUser("v1", name, desc, rolesList)
		if err != nil {
			panic(err.Error())
		}

		jsonData, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("desc", "d", "", "Description of the User")
	createUserCmd.Flags().StringP("roles", "q", "", "Comma separated list of roles; role1,role2,role3")
}
