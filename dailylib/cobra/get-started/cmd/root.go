package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "git命令是一个分布式的版本控制系统",
	Long: `git命令是一个分布式的版本控制系统，可以非常便利的对项目仓库进行管理，
本项目使用exe.Command将在Cmd中自定义的命令传递给系统执行
`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("无法识别的命令"))
	},
}

func Execute() {
	rootCmd.Execute()
}
