package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nothing",
	Short: "CLI to Nothing AI",
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println(fmt.Sprintf("%s version %s", constants.APP_NAME, constants.VERSION))
		} else {
			fmt.Println("Usage: nothing [command] [flags]")
		}

		fmt.Println()
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List resources",
}

var descCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Return information on resource",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	// Sub Commands
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(descCmd)

	// Resources
	getCmd.AddCommand(modelsCmd)
	getCmd.AddCommand(datasetsCmd)
	getCmd.AddCommand(experimentsCmd)
	getCmd.AddCommand(tasksCmd)

	descCmd.AddCommand(tasksCmd)
}
