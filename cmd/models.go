package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"git.datacentric.kr/handh/NothingAI-CLI/fetch"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:     "models",
	Aliases: []string{"mod"},
	Short:   "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {

		switch cmd.Parent().Use {
		case "get":
			getModel()
		case "create":
		case "delete":
		case "run":
		case "edit":
		default:
		}

	},
}

func getModel() {
	data, err := fetch.GetResources(constants.MODELS_GET_URL, "model")
	if err != nil {
		fmt.Println(err)
	} else if data == nil || len(data) == 0 {
		fmt.Println("No models found")
	} else {
		output.PrintTable(data, nil, true)
	}
}
