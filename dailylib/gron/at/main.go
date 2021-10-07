package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
	"sync"
)

/*
1.xtime(相比于time包，多了Day和Week的时间间隔)
	1.1 内置常量
	xtime.Second
	xtime.Minute
	xtime.Hour
	xtime.Day
	xtime.Week

	1.2 第一次执行任务的时间大于今天的日期
		1.2.1 如果需每间隔一段时间执行，用gron.Every(xx * xtime.xxx)，后面不需要跟At()方法，例如：
			gron.Every(1 * xtime.Day)，假设当前时间为14:00，则第一次执行定时任务的时间为第二天的14:00

		1.2.2 如果需要指定一个明确的时间执行任务，可以使用
			gron.Every(1 * xtime.Day).At("14:30") 则第一次执行任务的时间是第二天的14:30，且之后每天的14:30均会执行一次
			注意：如果使用了At，则传给Every的参数，即时间间隔必须大于1天

*/

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 创建一个定时任务管理器
	c := gron.New()
	c.AddFunc(gron.Every(1*xtime.Day).At("15:07"), func() {
		fmt.Println("runs at 15:07 every day")
	})
	c.Start()

	wg.Wait()
}
