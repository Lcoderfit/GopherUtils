package main

import "fmt"

// 定义一个名为Animal的接口，该接口定义了一个Speak方法
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
	// 1. 必须实现了Animal接口所有方法的变量，才能赋给Animal接口类型的变量
	// 接口变量包括两部分，类型和值，类型是赋值给接口变量的的值的类型，值是赋值给接口变量的值
	// 2. 赋给接口变量的结构体所实现的方法必须包含接口中定义的所有方法
	// 3. bytes.Buffer同时实现了io.Reader和io.Writer两个接口，所以bytes.Buffer对象可以赋值给io.Reader变量r，
	// 同时r可以断言为io.Writer， 因为r中存储的值也实现了io.Writer接口
	// 4. 所有类型的值均可以赋给 没有约定任何方法的接口即空接口，
	// 5.接口变量之间类型断言也好，直接赋值也好，其内部存储的(type, value)类型-值对是没有变化的。只是通过不同的接口能调用的方法有所不同而已。
	// 也是由于这个原因，接口变量中存储的值一定不是接口类型。
	var a Animal
	a = Cat{}
	a.Speak()

	a = Dog{}
	a.Speak()
}
