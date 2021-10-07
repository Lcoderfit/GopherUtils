package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add用于计算所有传入参数的和",
	Run: func(cmd *cobra.Command, args []string) {
		values := convertArgsToFloat64Slice(args, errorHandlingType)
		result := calc(values, ADD)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "+"), result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
