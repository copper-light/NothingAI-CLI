package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/constants"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nothing [COMMAND]",
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

var basicCmdGroup = &cobra.Group{
	Title: "Basic Commands",
	ID:    "basic",
}

var otherCmdGroup = &cobra.Group{
	Title: "Other Commands",
	ID:    "other",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	// Basic Commands
	rootCmd.AddGroup(basicCmdGroup)
	rootCmd.AddGroup(otherCmdGroup)
}
