package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示git的当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
