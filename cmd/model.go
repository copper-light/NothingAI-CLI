package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/fetcher"
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {

		switch cmd.Parent().Use {
		case "get":
			get()
		case "create":
		case "delete":
		case "run":
		case "edit":
		default:
		}

	},
}

func get() {
	body, err := fetcher.Request("GET", "http://localhost:8000/api/v1/models")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}
}
