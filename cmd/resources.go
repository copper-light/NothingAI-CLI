package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var resourcesCmd = &cobra.Command{
	Use:     "resources",
	Short:   "List types of a resource",
	GroupID: "other",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RESOURCE-TYPE   DESCRIPTION")
		fmt.Println("models          Manage models")
		fmt.Println("datasets        Manage datasets")
		fmt.Println("experiments     Manage experiments")
		fmt.Println("tasks           Manage tasks")
	},
}

func init() {
	rootCmd.AddCommand(resourcesCmd)
}
