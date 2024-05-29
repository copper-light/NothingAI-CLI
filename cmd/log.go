package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [TASK_ID]",
	Short: "Show log of task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("TASK_ID is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		logs, err := common.LogTask(id)
		if err != nil {
			fmt.Println(err)
		}
		for _, log := range logs {
			fmt.Println(log)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
