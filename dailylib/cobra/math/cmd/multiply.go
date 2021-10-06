package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "multiply计算所有参数的乘积",
	Run: func(cmd *cobra.Command, args []string) {
		values := convertArgsToFloat64Slice(args, errorHandlingType)
		result := calc(values, MULTIPLY)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "*"), result)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}
