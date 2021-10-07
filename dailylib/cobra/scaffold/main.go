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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
*/
package main

import "GopherUtils/dailylib/cobra/scaffold/cmd"

func main() {
	cmd.Execute()
}
