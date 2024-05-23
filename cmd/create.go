package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
)

func createResource(resourceType string, name string, options map[string]string) {
	resourceType = common.GetResourceType(resourceType)

	body := map[string]any{}
	body["name"] = name

	id, err := common.CreateResource(resourceType, body)
	if err != nil {
		fmt.Println(err)
	} else if id == -1 {
		fmt.Printf("Failed to create %s\n", resourceType)
	} else {
		fmt.Printf("Created %v %v(%v)\n", resourceType, name, id)
	}
}g

var createCmd = &cobra.Command{
	Use:     "create [RESOURCE_TYPE] [NAME]",
	Short:   "Create a resource",
	GroupID: "basic",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a resource type and a name")
		} else if len(args) < 2 {
			return errors.New("requires a resource type")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		name := args[1]
		createResource(resourceType, name, nil)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("description", "d", "", "Model description")
	createCmd.Flags().StringP("model_type", "m", "", "Model type")
	createCmd.Flags().StringP("source_type", "s", "", "Type storage of source file")
	//createCmd.Flags().StringP("visibility", "v", "", "Print version information")
}
