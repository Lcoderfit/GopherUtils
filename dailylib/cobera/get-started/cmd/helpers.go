package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "当前执行的命令为: %s, 参数: %v, err: %v\n", cmd.Name(), args, err)
}

func ExecuteCommand(name string, subName string, args ...string) (string, error) {
	args = append([]string{subName}, args...)
	// 使用exec.Command，将参数传递给系统执行
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}
