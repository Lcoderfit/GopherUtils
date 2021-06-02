package main

import "fmt"

// 定义一个名为Animal的接口，该接口声明了一个Speak方法
type Animal interface {
	Speak()
}

type Cat struct {
}

// Cat结构体实现了Animal接口中的Speak方法
func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct {
}

// Dog类实现了Animal接口的Speak方法
func (d Dog) Speak() {
	fmt.Println("Bark")
}

func main() {
	// 必须实现了Animal接口所有方法的变量，才能赋给Animal接口类型的变量
	// 接口变量包括两部分，类型和值，类型是赋值给接口变量的的值的类型，值是赋值给接口变量的值
	var a Animal
	a = Cat{}
	a.Speak()

	a = Dog{}
	a.Speak()
}
