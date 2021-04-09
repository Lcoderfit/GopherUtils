package math

import (
	"fmt"
	"math"
	"reflect"
)

func Example() {
	num := 5.3
	fmt.Println("原始数据: ", num, reflect.TypeOf(num))

	// 1.向上取整, 返回的结果为float64类型，ceil是天花板的意思
	c := math.Ceil(num)
	fmt.Println("向上取整: ", c, reflect.TypeOf(c))

	// 2.向下取整, 返回的结果为float64类型, floor是地板的意思
	f := math.Floor(num)
	fmt.Println("向下取整: ", f, reflect.TypeOf(f))

	// 3.四舍五入, 算法相当于math.Floor(num+0.5), 所以math.Round(-11.5)的值为-11
	r := math.Round(num)
	fmt.Println("四舍五入: ", r, reflect.TypeOf(r))
}
