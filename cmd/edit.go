package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit [RESOURCE_TYPE] [RESOURCE_ID]",
	Aliases: []string{"desc"},
	GroupID: "basic",
	Short:   "Edit a resource",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires resource type and resource id")
		} else if len(args) < 2 {
			return errors.New("requires a resource_id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//resourceType := args[0]
		//id := args[1]
		//GetDescribeTask(resourceType, id)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
