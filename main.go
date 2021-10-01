/*
1.获取不同分支的包：go get -u -v github.com/rif/cache2go@muesli
 */
package main

import (
	"fmt"
	"math"
)

type Type int

const (
	// Null is a null json value
	Null Type = iota
	// False is a json false boolean
	False
	// Number is json number
	Number
	// String is a json string
	String
	// True is a json true boolean
	True
	// JSON is a raw block of JSON
	JSON
)

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println("times: ", i)
	//	timeutils.PrintSpendTime(0, func() {
	//		BitOpr1()
	//	})
	//	timeutils.PrintSpendTime(0, func() {
	//		BitOpr2()
	//	})
	//	timeutils.PrintSpendTime(0, func() {
	//		BitOpr3()
	//	})
	//	fmt.Println()
	//}
	// https://mp.weixin.qq.com/s/dgVv6p8HQtc-krPGEdU2cQ
	// https://mp.weixin.qq.com/s/ysOlv8xsc7wF7cPcV5pu7g
	// https://mp.weixin.qq.com/s/qFAtgpbAsHSPVLuo3PYIhg

	//filePath := "log/log"
	//src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	//if err != nil {
	//	// TODO:创建失败了怎么办？？
	//	fmt.Println(err)
	//}
	//fmt.Println(src)

	// 判断系统内存排序是大端序还是小端序
	// go:generate fmt.Println("asdlfkj")

	//binaryutils.SystemEndian()
	//fmt.Println()
	//binaryutils.TestBigEndian()

	a := []int{1, 2, 3}
	fmt.Printf("%v, %p\n", a, a)
	fmt.Printf("%v, %p, %v", a[:0], a[:0], cap(a[:0]))
}

func BitOpr1() {
	_ = -1 ^ (-1 << 5)
}

func BitOpr2() {
	_ = 1<<5 - 1
}

func BitOpr3() {
	_ = math.Pow(2.0, 5.0) - 1
}
