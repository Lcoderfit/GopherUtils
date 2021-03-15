1.cmd := exec.Command(`grep -q "d" /home/learnGoroutine/file/src.log && echo "contain"||echo "not contain"`)
e: fork/exec grep -q "d" /home/learnGoroutine/file/src.log: no such file or directory
s: exec.Command中的参数不能带空格，如果命令本身有空格，则需要根据空格分拆成多个参数传入
cmd := exec.Command("grep", "-q", "d", "/home/learnGoroutine/file/src.log&&", "echo", "contain", "||", "not contain")

2.长命令无效
s:使用 bash -c `command string`, command string 可以表示一整条命令
		cmd := exec.Command(
			"bash", "-c",
			fmt.Sprintf(`grep -q %s_ds %s && echo %d || echo %d`, key, path, GrepContain, GrepNotContain),
		)
3.
注：Go语言中字典是无序的
sort.Strings()可以对字符串列表进行排序：