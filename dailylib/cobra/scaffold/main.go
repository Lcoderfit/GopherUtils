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
	// scaffold指定应用程序名称，--pkg-name指定包路径（如果使用go mod则直接指定项目相对路径）
	cobra init scaffold --pkg-name GopherUtils/dailylib/cobera/scaffold/
*/
package main

import "GopherUtils/dailylib/cobera/scaffold/cmd"

func main() {
	cmd.Execute()
}
