package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	dividedByZeroHanding int // 除0如何处理
)

var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "divide计算所有参数的除",
	Run: func(cmd *cobra.Command, args []string) {
		values := convertArgsToFloat64Slice(args, errorHandlingType)
		result := calc(values, DIVIDE)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), result)
	},
}

func init() {
	divideCmd.Flags().IntVarP(
		&dividedByZeroHanding, "divide_by_zero", "d", PanicOnDivideByZero, "遇到除0情况时的处理方法",
	)
	rootCmd.AddCommand(divideCmd)
}
