1.cmd := exec.Command(`grep -q "d" /home/learnGoroutine/file/src.log && echo "contain"||echo "not contain"`)
e: fork/exec grep -q "d" /home/learnGoroutine/file/src.log: no such file or directory
s: exec.Command中的参数不能带空格，如果命令本身有空格，则需要根据空格分拆成多个参数传入
cmd := exec.Command("grep", "-q", "d", "/home/learnGoroutine/file/src.log&&", "echo", "contain", "||", "not contain")