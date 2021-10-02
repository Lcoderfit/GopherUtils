package main

import "GopherUtils/dailylib/cobera/math/cmd"

/*
1.result := make([]float64, 0, len(args))
2.m.exe divide 10 5 // args就是10 5
3.
rootCmd.PersistentFlags().IntVarP(
		&errorHandlingType, "err_handling_type", "p", ContinueOnParseError, "当解析错误时运行",
	)

4.
	divideCmd.Flags().IntVarP(
		&dividedByZeroHanding, "divide_by_zero", "d", PanicOnDivideByZero, "遇到除0情况时的处理方法",
	)

5.末尾带有P和不带有P的参数的区别
	带有
*/

func main() {
	cmd.Execute()
}
