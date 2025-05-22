package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xhunterkiller/berry/pkgs/models"
	"github.com/spf13/cobra"
)

// createResourceCmd represents the createResource command
var createResourceCmd = &cobra.Command{
	Use:   "resource [Name of the Resource]",
	Args:  cobra.ExactArgs(1),
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

		res, err := models.NewResource("v1", name, desc, verbs)
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
	createCmd.AddCommand(createResourceCmd)

	createResourceCmd.Flags().StringP("desc", "d", "", "Description of the resource")
	createResourceCmd.Flags().StringP("verbs", "e", "", "Comma separated list of verbs; verb1,verb2,verb3")
}
