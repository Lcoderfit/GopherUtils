package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cloneCmd = &cobra.Command{
	Use:   "clone url [destination]",
	Short: "根据url指定的仓库地址拉取仓库代码到指定路径，如果不指定destination，则默认拉取到当前地址",
	Long: `根据url指定的仓库地址拉取仓库代码到指定路径，url是必传参数，destination是可选参数，如果不传入destination，
则默认会将代码保存到当前路径下
`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecuteCommand("git", "clone", args...)
		if err != nil {
			Error(cmd, args, err)
		}
		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
