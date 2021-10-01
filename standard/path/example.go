package path

import (
	"fmt"
	"path"
)

func Example() {
	// 1.传入一个表示路径的字符串参数，返回该路径的上一级路径，例如传入"/a/b/c"则返回/a/b; 传入/a/b/c/则返回/a/b/c
	a := "/a/b/c"
	b := "/a/b/c/"

	ans1 := path.Dir(a)
	ans2 := path.Dir(b)
	fmt.Printf("before: %s, path.Dir: %s\n", a, ans1)
	fmt.Printf("before: %s, path.Dir: %s\n", b, ans2)

	// 2.拼接路径/path.Join()
	a = "/a/b"
	b = "c/d"
	join := path.Join(a, b)
	fmt.Printf("before: %s, %s; path.Join: %s", a, b, join)
}
