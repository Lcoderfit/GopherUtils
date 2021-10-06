package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// 自定义类型

const (
	ContinueOnParseError = iota + 1 // 解析错误尝试继续处理
	ExitOnParseError                // 解析错误退出
	PanicOnParseError               // 解析错误panic
	ReturnOnDivideByZero            // 除0 返回
	PanicOnDivideByZero             // 除0 panic
)

// 加减乘除
const (
	ADD = iota + 1
	MINUS
	MULTIPLY
	DIVIDE
)

var (
	errorHandlingType int
)

var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "math计算累计的结果",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("无法识别命令"))
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(
		&errorHandlingType, "err_handling_type", "p", ContinueOnParseError, "当解析错误时运行",
	)
}

func Execute() {
	rootCmd.Execute()
}
