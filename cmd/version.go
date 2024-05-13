package cmd

import (
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/Constants"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s version %s", Constants.APP_NAME, Constants.VERSION))
	},
}
