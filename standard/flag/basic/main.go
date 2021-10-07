package main

import (
	"flag"
	"fmt"
)

/*
1.flag.Type() （下面以flag.Int("i", 0, "flag of int")为例）
	1.1 Type指代各种基本类型：int，bool，string...), 返回Type类型的指针(*Type)
	1.2 包含3个参数，第一个是标志名(-i), 第二个是默认值，第三个是标志含义(通过go run main.go -h 可以查看所有标志的含义)

2.flag.Parse() 解析参数
	2.1 从os.Args[:1]中解析选项，例如 go run main.go -i 10 main.go为第一个参数，是不会被解析的，只会解析后面的部分
	2.2 flag支持的三种命令行选项格式(注意：-和--是等效的，可以使用-flag，也可以使用--flag)
		-flag	// 该格式仅支持bool类型，不支持其他类型
		-flag x // 该格式不支持bool类型，支持其他所有类型
		-flag=x // 支持任意类型
		注意：-flag格式只支持bool类型的选项，-flag x格式不支持bool类型的选项，如果要显示设置bool类型选项的值，用-flag=x格式
	2.3 当遇到第一个非选项参数时，解析就会停止
		go run main.go fuck -i 10  // 由于fuck为非选项参数，所以会导致后面的-i选失效
	2.4 取值为true和false可以输入：
		2.4.1 取值为true的：1 t T true True TRUE
		2.4.2 取值为false的：0 f F false False FALSE
	2.5 整型选项值，只要传入的是strconv.ParseInt可以接收的值，flag就能解析（十进制，八进制0664， 十六进制0x123均可以）

3.flag.Args() flag.NArg() flag.NFlag() flag.Arg(i)
	3.1 flag.Args()返回所有非标志（non-flag）参数 例如： go run main.go -i 10 20 30 则非标志参数为[20 30]
		注意，bool类型的标志比较特殊，假设设置了一个bool类型的标志flag.Bool("b", false, "xxx"),
		如果命令行中不带有-b标志，则其值为设置的默认值，只要带有-b标志（-b后面不用跟其他值，例如：go run main.go -b），其值就为true
	3.2 flag.NFlag() 返回所有标志参数的个数
	3.3 flag.NArg() 返回所有非标志参数的个数
	3.4 flag.Arg(i) 返回传入的第i个非标志参数
*/

var (
	intFlag    *int
	boolFlag   *bool
	stringFlag *string
)

func init() {
	intFlag = flag.Int("int", 0, "int flag")
	boolFlag = flag.Bool("bool", true, "bool flag")
	stringFlag = flag.String("string", "", "string flag")
}

func main() {
	flag.Parse()

	fmt.Println("int flag: ", *intFlag)
	fmt.Println("bool flag: ", *boolFlag)
	fmt.Println("string flag: ", *stringFlag)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
	}
}
