package main

import "GopherUtils/dailylib/cobera/math/cmd"

/*
1.result := make([]float64, 0, len(args))
	预分配内存，注意第二个参数表示长度，第三个表示容量，预分配内存时长度参数为0；
	容易出错的地方：预分配内存时长度不为0，然后result = append(result, xxx)时导致结果不符合预期

2.m.exe divide 10 5
	传入到Run: func(cmd *cobra.Cmd, args []string)中的args就是[10, 5]，即Use之后的部分

3.PersistentFlags()返回一个*pflag.FlagSet类型变量，内置Type() TypeVar() Var()三种大类的方法
且每一类方法均包含一个末尾带有P的方法，例如TypeP() TypeVarP() VarP(), 表示第三个参数可以定义一个短选项;
第一个参数是接收传给该选项的参数值的变量指针，第二个是选项名称，第三个是短选项名，第四个是默认值，第五个是含义

	rootCmd.PersistentFlags().IntVarP(
			&errorHandlingType, "err_handling_type", "p", ContinueOnParseError, "当解析错误时运行",
		)

	divideCmd.Flags().IntVarP(
		&dividedByZeroHanding, "divide_by_zero", "d", PanicOnDivideByZero, "遇到除0情况时的处理方法",
	)

4.本地选项和永久选项
	4.1 本地选项，只能在定义它的命令中使用
	// 本地选项与永久选项参数一致，参数顺序为接收选项的变量，选项名，短选项名，默认值，帮助信息
	rootCmd.IntVarP(&a, "int", "i", 10, "本地选项")

	4.2 永久选项，可以在定义它的选项和子选项中使用
	rootCmd.PersistentFlags().IntVarP(&a, "int", "i", 10, "永久选项")

*/

func main() {
	cmd.Execute()
}
