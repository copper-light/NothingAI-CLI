package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"git.datacentric.kr/handh/NothingAI-CLI/fetch"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

var experimentsCmd = &cobra.Command{
	Use:     "experiments",
	Aliases: []string{"exp"},
	Short:   "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {

		switch cmd.Parent().Use {
		case "get":
			getExperiments()
		case "create":
		case "delete":
		case "run":
		case "edit":
		default:
		}

	},
}

func getExperiments() {
	data, err := fetch.GetResources(constants.EXPERIMENTS_GET_URL, "experiment")
	if err != nil {
		fmt.Println(err)
	} else if data == nil || len(data) == 0 {
		fmt.Println("No experiments found")
	} else {
		output.PrintTable(data, nil, true)
	}
}
