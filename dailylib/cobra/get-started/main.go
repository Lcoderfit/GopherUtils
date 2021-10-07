package main

import "GopherUtils/dailylib/cobra/get-started/cmd"

/*
1.创建cobra命令系统的流程
	1.1 在项目根目录下创建cmd目录和main.go文件
	1.2 在cmd目录下创建root.go文件，然后有如下定义:
		var rootCmd = &cobra.Command{
			Use: "git",
			Short: "short description",
			Long: "Long description",
			Run: func(cmd *cobra.Command, args []string) {
				// 输入该命令后需要执行的处理逻辑
			}
		}
	1.3 在cmd/root.go中编写Execute方法供main.go调用
		func Execute() {
			// 任何*cobra.Command类型均可以调用Execute()方法执行命令，调用Execute方法即会运行*cobra.Command中定义的Run方法
			rootCmd.Execute()
		}
	1.4 在main.go中
		func main() {
			// 调用命令系统的入口
			cmd.Execute()
		}

	1.5 root.go中创建的根命令，每添加一个子命令，例如命令A，即在cmd目录中添加一个A.go文件，然后按照如下格式定义命令;
		注意：root命令和子命令有一个不同，root命令中需要定义一个Execute()方法，方法中式调用rootCmd.Execute();
		而子命令中是直接再init中调用 rootCmd.AddCommand(ACmd)将子命令添加到根命令中
		var ACmd = &cobra.Command{
			Use: "A",
			Short: "这是命令A，用来xxx",
			Run: func(cmd *cobra.Command, args []string) {
				// 输入该命令时需要执行的逻辑
			}
		}

		func init() {
			rooCmd.AddCommand(ACmd)
		}
	1.6 如果需要什么辅助函数，可以再helpers中定义,例如：
		func ExecuteCommand(name string, subName string, args ...string) (string, error) {
			args = append([]string{subName}, args...)
			// 使用exec.Command，将参数传递给系统执行
			cmd := exec.Command(name, args...)
			bytes, err := cmd.CombinedOutput()
			return string(bytes), err
		}

2.Use/Short/Long/Run的含义
	2.1 Use表示命令的名称，例如Use: "clone url [destination]", 则表示子命令为clone，url是必传参数，destination是可选参数
	2.2 Short 是在用help命令时的输出，例如m.exe -h 或者 m.exe help 或者 m.exe --help
	2.3 Long 是在用help <command>时的输出，例如m.exe clone -h 就会输出Long的内容
	2.4 Run 当传入该命令时候所执行的实际操作

3.exec.Command
	3.1 基本使用: 可以传入命令和参数，然后交给系统执行
		// 会返回一个*Cmd类型，然后调用*Cmd.CombinedOutput() 会返回[]byte, error
		// []byte即为系统执行命令后的输出内容
		cmd := exec.Command("git", "branch", "-a")
		// 执行命令并将标准输出和标准错误输出合并再一起返回
		bytes, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bytes))

		注意：exec.Command(name string, args ...string) 命令、选项、参数必须按照顺序传入，第一个参数必须是命令，
		例如git branch -a应该这样调用 exec.Command("git", "branch", "-a")

4.fmt.Fprint(w io.Writer, a ...interface{})
	将a写入到io.Writer，常用的是将字符串内容输出到标准输出
	fmt.Fprint(os.Stdout, "xxxx")

5.helpers
	一般不定义子命令，而用来编写一些辅助函数，供其他命令调用

6.go build -o m.exe main.go
	6.1 首先m.exe是会在使用命令的目录下生成
	6.2 生成了m.exe之后，m.exe就相当于是root命令，如果需要调用子命令，例如clone子命令，直接 m.exe clone xxx即可

*/

func main() {
	cmd.Execute()
}
