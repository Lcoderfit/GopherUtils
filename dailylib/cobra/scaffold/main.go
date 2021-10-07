/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
1.利用cobra自带命令初始化项目
	1.1 cobra init scaffold --pkg-name GopherUtils/dailylib/cobera/scaffold/
		首先会在使用该命令的路径下(例如在GopherUtils目录下使用该命令，则会在GopherUtils目录下生成scaffold目录)
		创建一个scaffold目录，目录结构如下：
		./scaffold
			main.go
			./cmd
				root.go
		scaffold指定生成的根命令名和存放该命令相关程序的目录名，--pkg-name指定包路径，
		例如使用cobra init test --pkg-name GopherUtils/dailylib/cobera/g 则生成的test目录下的main.go
		文件中，导入的cmd包路径为 import "GopherUtils/dailylib/cobera/g/cmd"

2.cobra add date
	2.1 在scaffold目录下使用该命令，会在/scaffold/cmd/目录下生成date.go文件，其中包含自动生成的date命名作为scaffold的子命令;
	自动生成的代码如下，可以分别定义永久选项和本地选项

	func init() {
		rootCmd.AddCommand(dateCmd)

		// Here you will define your flags and configuration settings.

		// Cobra supports Persistent Flags which will work for this command
		// and all subcommands, e.g.:
		dateCmd.PersistentFlags().IntVarP(
			&year, "year", "y", time.Now().Year(), "year to show(should in [1000, 9999])",
		)
		dateCmd.PersistentFlags().IntVarP(
			&month, "month", "m", int(time.Now().Month()), "month to show(should in [1, 12])",
		)

		// Cobra supports local flags which will only run when this command
		// is called directly, e.g.:
		// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}


3.time用法
	3.1 time.Now() 返回当前时间，Time类型

	3.2 time.Date
	// 参数顺序分别为年份，月份（注意月份需要转换为time.Month类型），天，时，分，秒，纳秒，时区
	// 获取当前时区时，可以先通过time.Now()获取当前时间now，然后通过now.Location()获取时区
	time.Date(2021, time.Month(10), 1, 0, 0, 0, 0, now.Location())

	3.3
	time.Now().Year()返回一个表示当前年份的整数，time.Now.Month()返回一个表示当前月份的time.Month类型(type Month int)，
	如果需要将其转换为整数： int(time.Now().Month())

	3.4 Time类型的内置方法：WeekDay() Month() Day()
		showTime.WeekDay() 表示当前日期是星期几
		showTime.WeekDay() 表示当前日期是在哪一个月
		showTime.Day() 表示当前日期是在本月的哪一天

	3.5 计算时间差
		showTime.Add(time.Hour * 24) 计算一天后的日期，返回Time类型（因为没有表示一天的内置变量，所以用time.Hour*24）
		time.Saturday 是一个time.WeekDay类型（type WeekDay int），从time.Sunday -> time.Saturday的值 0 -> 6

4.strings.Repeat(s string, count int)
	返回count个s拼接后的结果，如果count为0，则返回空字符串


5.cobra还包含钩子函数（在Run之前或之后执行某些操作），
	5.1 钩子函数：https://github.com/spf13/cobra/blob/master/user_guide.md#prerun-and-postrun-hooks
	5.2 生成文档：https://github.com/spf13/cobra/blob/master/doc/yaml_docs.md

6.viper/cobra
	cobra是眼镜蛇
	viper是蝰蛇
*/
package main

import (
	"GopherUtils/dailylib/cobra/scaffold/cmd"
)

func main() {
	cmd.Execute()
}
