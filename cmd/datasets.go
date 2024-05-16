package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"git.datacentric.kr/handh/NothingAI-CLI/fetch"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

var datasetsCmd = &cobra.Command{
	Use:     "datasets",
	Aliases: []string{"dat"},
	Short:   "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {

		switch cmd.Parent().Use {
		case "get":
			getDatasets()
		case "create":
		case "delete":
		case "run":
		case "edit":
		default:
		}

	},
}

func getDatasets() {
	data, err := fetch.GetResources(constants.DATASETS_GET_URL, "dataset")
	if err != nil {
		fmt.Println(err)
	} else if data == nil || len(data) == 0 {
		fmt.Println("No datasets found")
	} else {
		output.PrintTable(data, nil, true)
	}
}
