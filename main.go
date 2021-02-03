package main

import (
	"GopherUtils/timeutils"
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("times: ", i)
		timeutils.PrintSpendTime(0, func() {
			BitOpr1()
		})
		timeutils.PrintSpendTime(0, func() {
			BitOpr2()
		})
		timeutils.PrintSpendTime(0, func() {
			BitOpr3()
		})
		fmt.Println()
	}
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
