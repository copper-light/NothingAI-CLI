package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
)

func deleteResource(resourceType string, id string) {
	resourceType = common.GetResourceType(resourceType)
	ok, err := common.DeleteResource(resourceType, id)
	if err != nil {
		fmt.Println(err)
	} else if ok {
		fmt.Printf("%v \"%v\" is deleted\n", resourceType, id)
	}
}

var deleteCmd = &cobra.Command{
	Use:     "delete [RESOURCE_TYPE] [RESOURCE_ID]",
	Aliases: []string{"del"},
	GroupID: "basic",
	Short:   "Delete the resource",
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
		deleteResource(resourceType, id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
