package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"git.datacentric.kr/handh/NothingAI-CLI/fetch"
	"git.datacentric.kr/handh/NothingAI-CLI/output"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:     "tasks",
	Aliases: []string{"tas"},
	Short:   "Print the version number",
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	fmt.Println("hi", cmd.Parent().Use)
	//	switch cmd.Parent().Use {
	//	case "describe":
	//		fallthrough
	//	case "delete":
	//		if len(args) < 0 {
	//			print("error")
	//			return errors.New("error: You must specify the id of resource to describe")
	//		}
	//	}
	//	return nil
	//},
	Run: func(cmd *cobra.Command, args []string) {

		switch cmd.Parent().Use {
		case "get":
			getTasks()
		case "describe":
			id := args[0]
			describeTask(id)
		case "create":
		case "delete":
		case "run":
		case "edit":
		default:
		}

	},
}

func getTasks() {
	data, err := fetch.GetResources(constants.TASKS_GET_URL, "task")
	if err != nil {
		fmt.Println(err)
	} else if data == nil || len(data) == 0 {
		fmt.Println("No tasks found")
	} else {
		output.PrintTable(data, nil, true)
	}
}

func describeTask(id string) {
	data, err := fetch.DescribeResource(constants.TASK_DESC_URL, id, "task")
	if err != nil {
		fmt.Println(err)
	} else if data == nil {
		fmt.Printf("No task %s found\n", id)
	} else {
		output.PrintKeyValue(data)
	}
}
