package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var minusCmd = &cobra.Command{
	Use:   "minus",
	Short: "minus计算所有参数的差",
	Run: func(cmd *cobra.Command, args []string) {
		values := convertArgsToFloat64Slice(args, errorHandlingType)
		result := calc(values, MINUS)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "-"), result)
	},
}

func init() {
	rootCmd.AddCommand(minusCmd)
}
