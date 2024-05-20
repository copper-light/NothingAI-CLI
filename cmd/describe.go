package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

func describeResource(resourceType string, id string) {
	resourceType = common.GetResourceType(resourceType)
	results, err := common.DescribeResource(resourceType, id)
	if err != nil {
		fmt.Println(err)
	} else if results == nil {
		fmt.Printf("No %v found in %v\n", id, resourceType)
	} else {
		output.PrintKeyValue(results)
	}
}

var describeCmd = &cobra.Command{
	Use:     "describe [RESOURCE_TYPE] [RESOURCE_ID]",
	Aliases: []string{"desc"},
	GroupID: "basic",
	Short:   "Show details of a specific resource",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires resource type and resource id")
		} else if len(args) < 2 {
			return errors.New("requires a resource_id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resourceType := args[0]
		id := args[1]
		describeResource(resourceType, id)
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
