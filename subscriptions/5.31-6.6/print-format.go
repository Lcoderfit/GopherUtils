package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	// 只打印值
	s := Student{"bob", 12345678}
	fmt.Printf("%v\n", s)
	// 会先打印字段类型，然后打印字段值
	fmt.Printf("%+v\n", s)
	// 会先打印结构体名称，再输出结构体(字段类型+字段值)
	fmt.Printf("%#v\n", s)
}
