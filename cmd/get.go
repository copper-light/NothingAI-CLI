package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

func getResource(resourceType string) {
	resourceType = common.GetResourceType(resourceType)
	results, err := common.GetResources(resourceType)
	if err != nil {
		fmt.Println(err)
	} else if results == nil || len(results) == 0 {
		fmt.Printf("No %s found\n", resourceType)
	} else {
		output.PrintTable(results, nil, true)
	}
}

var getCmd = &cobra.Command{
	Use:     "get [RESOURCE_TYPE]",
	Short:   "List resources",
	GroupID: "basic",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a resource argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		getResource(resourceType)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
