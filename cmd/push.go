package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push [RESOURCE_TYPE] [RESOURCE_ID] [filename or dirname]",
	Short:   "Push files to a server",
	GroupID: "file",
}

var pushModelCmd = &cobra.Command{
	Use:     "models [RESOURCE_TYPE] [RESOURCE_ID] [filename or dirname]",
	Aliases: []string{"model"},
	Short:   "Push model files to a server",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require a resource type")
		} else if len(args) < 2 {
			return errors.New("requires a resource type and a id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		path := args[1]
		result, err := common.SendFiles("models", id, path)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Pushed %d files\n", result)
		}
	},
}

var pushDatasetCmd = &cobra.Command{
	Use:   "dataset [RESOURCE_TYPE] [RESOURCE_ID]",
	Short: "Push dataset files to a server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.AddCommand(pushModelCmd)
	pushCmd.AddCommand(pushDatasetCmd)
}
