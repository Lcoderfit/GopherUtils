package timeutils

import (
	"fmt"
	"time"
)

const (
	LoopTimes = 1000000
)

// PrintSpendTime 用于计算程序运行的时间，需要计算运行耗时的程序应放在f中，
// loopTimes用于设置循环次数,例如：
//
// PrintSpendTime(100, f func(){
//		a = math.Pow(2.0, 10.0)
// })
//
// 将会打印a = math.Pow(2.0, 10.0)运行100次所消耗的时间，单位为纳秒；
// 如果loopTimes <= 0,则循环次数默认为1000000次
func PrintSpendTime(loopTimes int, f func()) {
	if loopTimes <= 0 {
		loopTimes = LoopTimes
	}
	startTime := time.Now()
	for i := 0; i < loopTimes; i++ {
		f()
	}
	fmt.Println(time.Since(startTime).Nanoseconds())
}
