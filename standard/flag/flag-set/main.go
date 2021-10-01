package main

import (
	"flag"
	"fmt"
	"time"
)

/*
1.Duration类型需要传入：1s 1m 1h -1s 100ms 这种格式，不能单纯传入数字
2.flag.TypeVar和flag.Type的区别
	2.1 flag.TypeVar
		flag.IntVar(&a, "i", 10, "int flag -i")
		如果定义的是flag.TypeVar类型选项，则第一个参数必须传入一个Type类型的指针，
		第二个参数是选项名，第三个参数是默认值，第四个参数是选项含义

	2.2 flag.Type
		intFlag := flag.Int("i", 10, "int flag -i")
		第一个参数传入选项名，第二个参数传入默认值，第三个参数传入选项含义
		flag.Type会返回一个Type类型的值,而flag.TypeVar则不会,TypeVar会将传入的选项值赋给第一个参数对应的变量

3.flag.NewFlagSet("MyFlagSet", flag.xxxx)
	3.1 创建一个*FlagSet实例，如果传入的选项列表中包含没有定义的选项，则会打印如下信息
		flag provided but not defined: -stringflag
		Usage of MyFlagSet:
		  -bool-flag
				bool flag value
		  -int-flag int
				int flag value
		  -string-flag string
				string flag value (default "default")
	3.2 flag.NewFlagSet("f1", flag.xxx)
		3.2.1 flag.ContinueError 如果有错误，则会直接打印出来
		flag provided but not defined: -stringflag

		3.2.2 flag.ExitOnError 如果有错误，直接调用os.Exit(2)退出，例如：
		flag provided but not defined: -stringflag
		Usage of MyFlagSet:
		  -bool-flag
				bool flag value
		  -int-flag int
				int flag value
		  -string-flag string
				string flag value (default "default")
		exit status 2

		3.2.3 flag.PanicOnError 以panic的形式打印错误信息

	3.3 *FlagSet.Parse(arguments []string)
		传入一个包含选项名和参数值的列表，批量设置选项值

*/

var (
	period time.Duration
)

func init() {
	flag.DurationVar(&period, "period", 1*time.Second, "sleep period")
}

func main() {
	// 参数解析
	flag.Parse()
	fmt.Printf("Sleeping for %v...", period)
	time.Sleep(period)
	fmt.Println()

	args := []string{"-int-flag", "12", "-stringflag", "test"}

	var intFlag int
	var boolFlag bool
	var stringFlag string

	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)
	fs.IntVar(&intFlag, "int-flag", 0, "int flag value")
	fs.BoolVar(&boolFlag, "bool-flag", false, "bool flag value")
	fs.StringVar(&stringFlag, "string-flag", "default", "string flag value")

	// 如果传入的参数列表中包含未定义的参数，则会报错：flag provided but not defined: -stringflag
	fs.Parse(args)

	fmt.Println("int flag: ", intFlag)
	fmt.Println("bool flag: ", boolFlag)
	fmt.Println("string flag: ", stringFlag)
}
