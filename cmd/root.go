package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/Constants"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nothing",
	Short: "CLI to Nothing AI",
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println(fmt.Sprintf("%s version %s", Constants.APP_NAME, Constants.VERSION))
		} else {
			fmt.Println("Usage: nothing [command] [flags]")
		}

		fmt.Println()
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get desc",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(modelCmd)
}
