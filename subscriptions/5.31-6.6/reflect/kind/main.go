package main

import (
	"fmt"
	"reflect"
)

type MyInt int

/*
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

func main() {
	var i int
	var j MyInt

	// 两个类型之间的转换需要进行强制类型转换
	i = int(j)
	ti := reflect.TypeOf(i)
	fmt.Println("type of i: ", ti.String())

	tj := reflect.TypeOf(j)
	fmt.Println("type of j: ", tj.String())

	// 打印底层数据类型
	// 两个类型的底层数据类型是一样的
	// 从反射值对象中获取种类，总共26种，定义在reflect中
	fmt.Println("kind of i: ", ti.Kind())
	fmt.Println("kind of j: ", tj.Kind())
}
