package main

import (
	"fmt"
	"reflect"
)

/*
一共 26 种，我们可以分类如下：

基础类型Bool、String以及各种数值类型（有符号整数Int/Int8/Int16/Int32/Int64，无符号整数Uint/Uint8/Uint16/Uint32/Uint64/Uintptr，浮点数Float32/Float64，复数Complex64/Complex128）
复合（聚合）类型Array和Struct
引用类型Chan、Func、Ptr、Slice和Map（值类型和引用类型区分不明显，这里不引战，大家理解意思就行）
接口类型Interface
非法类型Invalid，表示它还没有任何值（reflect.Value的零值就是Invalid类型）
Go 中所有的类型（包括自定义的类型），都是上面这些类型或它们的组合。

const (
  Invalid Kind = iota
  Bool
  Int
  Int8
  Int16
  Int32
  Int64
  Uint
  Uint8
  Uint16
  Uint32
  Uint64
  Uintptr
  Float32
  Float64
  Complex64
  Complex128
  Array
  Chan
  Func
  Interface
  Map
  Ptr
  Slice
  String
  Struct
  UnsafePointer
)
 */

type Cat struct {
	Name string
}

func main() {
	// Go 语言是静态类型的，每个变量在编译期有且只能有一个确定的、已知的类型，即变量的静态类型。静态类型在变量声明的时候就已经确定了，无法修改。
	// 一个接口变量，它的静态类型就是该接口类型。虽然在运行时可以将不同类型的值赋值给它，改变的也只是它内部的动态类型和动态值。它的静态类型始终没有改变。
	var f float64 = 3.5
	c := Cat{Name: "kitty"}

	// TypeOf和ValueOf接收一个interface{} （空接口）值作为参数，如果传入的不是一个interface类型，则会先转换为接口类型
	t1 := reflect.TypeOf(f)
	t2 := reflect.TypeOf(c)
	fmt.Println(t1.String())
	fmt.Println(t2.String())

	// fmt.Println会对reflect.Value做额外处理，获取其内部值
	// Go 语言中类型是无限的，而且可以通过type定义新的类型。但是类型的种类是有限的，reflect包中定义了所有种类的枚举：
	v1 := reflect.ValueOf(f)
	v2 := reflect.ValueOf(c)
	fmt.Println(v1)
	fmt.Println(v1.String())
	fmt.Println(v2)
	fmt.Println(v2.String())
}
