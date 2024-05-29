package cmd

import (
	"errors"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common"
	"github.com/spf13/cobra"
)

var execExperiment = &cobra.Command{
	Use:   "exec [EXPERIMENT_ID]",
	Short: "Execute a experiment",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require a experiment id")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		success, err := common.ExecExperiment(id)
		if success {
			fmt.Printf("Experiment \"%v\" is Executed \n", id)
		} else {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execExperiment)
}
