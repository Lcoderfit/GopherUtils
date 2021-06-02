package main

import "fmt"

// 声明一个animal接口
type Animal interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func main() {
	var a Animal
	a = Cat{Name: "coder"}
	a.Speak()

	// 如果类型断言与变量的实际类型不符，则会panic
	// 可以使用t, ok := a.(cat)， 如果类型不符不会返回panic，ok的值会为false
	// t := a.(Cat)
	t, ok := a.(Cat)
	if !ok {
		fmt.Println("类型断言不符")
	} else {
		fmt.Println(t.Name)
	}
}
