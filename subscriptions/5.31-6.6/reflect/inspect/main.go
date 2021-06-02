package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	u := User{
		Name:    "robert lu",
		Age:     24,
		Married: false,
	}
	inspectStruct(u)
	fmt.Println()

	// 透视不可导出字段
	inspectStruct(bytes.Buffer{})

	// 透视map
	m := map[int]string{
		1: "lu",
		2: "ji",
	}
	inspectMap(m)
	fmt.Println()

	// 透视slice
	inspectSliceArray([]int{1, 3, 6})
	fmt.Println()

	// 透视func
	inspectFunc("Add", Add)
	inspectFunc("Greeting", Greeting)

	// 透视结构体中的定义的func
	u1 := User{
		Name:    "jsy",
		Age:     24,
		Married: false,
	}
	inspectMethod(&u1)
}

// reflect.ValueOf()：获取反射值对象；
// reflect.Value.NumField()：从结构体的反射值对象中获取它的字段个数；
// reflect.Value.Field(i)：从结构体的反射值对象中获取第i个字段的反射值对象；
// reflect.Kind()：从反射值对象中获取种类；
// reflect.Int()/reflect.Uint()/reflect.String()/reflect.Bool()：这些方法从反射值对象获取出具体类型。
func inspectStruct(u interface{}) {
	// 获取反射值对象
	v := reflect.ValueOf(u)
	// NumField和Field函数只有对象是结构体时才能调用，否则会panic
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// 这里是对field的底层类型进(反射类型)行判断
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// field.Int()返回int64类型，field.Uint()返回Uint64类型（返回相应的最大范围的类型）
			fmt.Printf("field: %d, type: %s, value: %d\n", i, field.Type().Name(), field.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fmt.Printf("field: %d, type: %s, value: %d\n", i, field.Type().Name(), field.Uint())
		case reflect.Bool:
			// %t表示布尔类型的值，%q表示单引号围绕的字符串字面值
			fmt.Printf("field: %d, type: %s, value: %t\n", i, field.Type().Name(), field.Bool())
		case reflect.String:
			fmt.Printf("field: %d, type: %s, value: %q\n", i, field.Type().Name(), field.String())
		default:
			fmt.Printf("field: %d unhandled kind:%s\n", i, field.Kind())
		}
	}
}

// reflect.Value.MapKeys()：将每个键的reflect.Value对象组成一个切片返回；
// reflect.Value.MapIndex(k)：传入键的reflect.Value对象，返回值的reflect.Value；
// 然后可以对键和值的reflect.Value进行和上面一样的处理。
func inspectMap(m interface{}) {
	v := reflect.ValueOf(m)
	// 可以对k和field使用switch语句进行类型判断
	// MapKeys()和MapIndex(k)方法只能在原对象是map类型时才能调用，否则会panic
	for _, k := range v.MapKeys() {
		field := v.MapIndex(k)
		// Interface()方法以空接口的形式返回内部包含的值
		fmt.Printf("%v => %v, type: %s, value: %v\n", k.Interface(), field.Interface(), k.Kind(), field.Kind())
	}
}

// Len()和Index(i)方法只能在原对象是切片，数组或字符串时才能调用，其他类型会panic。
func inspectSliceArray(sa interface{}) {
	v := reflect.ValueOf(sa)
	fmt.Printf("%c", '[')
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		fmt.Printf("%v ", elem.Interface())
	}
	fmt.Printf("%c", ']')
	fmt.Printf(", type: %v\n", v.Kind())
}

func Add(a, b int) int {
	return a + b
}

func Greeting(name string) string {
	return "hello, " + name
}

// 只有在原对象是函数类型的时候才能调用NumIn()/In()/NumOut()/Out()这些方法，其他类型会panic。
// reflect.Type.NumIn()：获取函数参数个数；
// reflect.Type.In(i)：获取第i个参数的reflect.Type；
// reflect.Type.NumOut()：获取函数返回值个数；
// reflect.Type.Out(i)：获取第i个返回值的reflect.Type。
func inspectFunc(name string, f interface{}) {
	t := reflect.TypeOf(f)
	fmt.Printf("%v input: ", name)
	for i := 0; i < t.NumIn(); i++ {
		ti := t.In(i)
		fmt.Print(ti.Name())
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Printf("%v output: ", name)
	for i := 0; i < t.NumOut(); i++ {
		to := t.Out(i)
		fmt.Print(to.Name())
		fmt.Print(" ")
	}
	fmt.Println("\n===========")
}

type User struct {
	Name    string
	Age     int
	Married bool
}

// 设置名字
func (u *User) SetName(name string) {
	u.Name = name
}

// 设置年龄
func (u *User) SetAge(age int) {
	u.Age = age
}

// reflect.Value也定义了NumMethod()/Method(i)这些方法。区别在于：reflect.Type.Method(i)返回的是一个reflect.Method对象，
// 可以获取方法名、类型、是结构体中的第几个方法等信息。如果要通过这个reflect.Method调用方法，必须使用Func字段，而且要传入接收器的reflect.Value作为第一个参数
//
// reflect.Value.Method(i)返回一个reflect.Value对象，它总是以调用Method(i)方法的reflect.Value作为接收器对象，
// 不需要额外传入。而且直接使用Call()发起方法调用
func inspectMethod(o interface{}) {
	t := reflect.TypeOf(o)
	sl := []interface{}{1, "jsy"}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		// 如果要通过这个reflect.Method调用方法，必须使用Func字段，而且要传入接收器的reflect.Value作为第一个参数
		fmt.Println(m.Func.Call([]reflect.Value{reflect.ValueOf(o), reflect.ValueOf(sl[i])}))
	}
	fmt.Println(o)
}
