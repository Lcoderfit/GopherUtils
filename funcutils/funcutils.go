package funcutils

import (
	"fmt"
	"runtime"
)

func ParentFunc() {
	fmt.Printf("ParentFunc: %p\n", ParentFunc)
	ChildrenFunc()
}

func ChildrenFunc() {
	fmt.Printf("children: %p\n", ChildrenFunc)
	pc, _, _, _ := runtime.Caller(1)
	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Printf("%v\n", pc)
	f := runtime.FuncForPC(pc)
	fmt.Println()
}
