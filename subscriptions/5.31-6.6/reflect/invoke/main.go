package main

import (
	"fmt"
	"reflect"
)

func main() {
	invoke(Add, 1, 2)
	m := Math{
		a: 10,
		b: 2,
	}
	invokeMethod(m, "Add")
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello" + name
}

// ValueOf返回的反射值对象调用函数时直接调用v.Call
// TypeOf返回的反射类型对象调用函数时需要调用 t.Func.Call
func invoke(f interface{}, args ...interface{}) {
	v := reflect.ValueOf(f)

	argsV := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argsV = append(argsV, reflect.ValueOf(arg))
	}
	retList := v.Call(argsV)
	for _, ret := range retList {
		fmt.Println(ret.Interface())
	}
}

type M struct {
	a, b int
	op   rune
}

func (m M) Op() int {
	switch m.op {
	case '+':
		return m.a + m.b
	case '-':
		return m.a - m.b
	case '*':
		return m.a * m.b
	case '/':
		return m.a / m.b
	default:
		panic("invalid op")
	}
}

type Math struct {
	a, b int
}

func (m Math) Add() int {
	return m.a + m.b
}

func (m Math) Sub() int {
	return m.a - m.b
}

func (m Math) Mul() int {
	return m.a * m.b
}

func (m Math) Div() int {
	return m.a / m.b
}

// 通过反射调用接收器的方法
func invokeMethod(obj interface{}, name string, args ...interface{}) {
	v := reflect.ValueOf(obj)
	m := v.MethodByName(name)
	argVge  := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argV = append(argV, reflect.ValueOf(arg))
	}
	retList := m.Call(argV)
	fmt.Println("ret: ")
	for _, ret := range retList {
		fmt.Println(ret.Interface())
	}
}
