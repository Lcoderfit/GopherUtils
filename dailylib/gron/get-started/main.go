package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"sync"
	"time"
)

/*
1.gron.New()
	返回一个*Cron类型的对象，创建一个新的定时任务管理器

2.c.AddFunc(gron.Every(time.Duration), job func() {})   c.Start()
	*Cron对象的内置方法(或者使用Add)，通过AddFunc()方法添加一个定时任务，gron.Every(5*time.Second)表示每5s执行一次任务
	然后通过c.Start()启动一个协程运行定时任务(go c.Run())

3.gron.Every() 传入一个time.Duration类型的参数，指定任务执行的周期；

4.为什么使用WaitGroup？？
	因为c.Start()会启动一个协程运行定时任务，如果主协程比该协程先结束，则整个程序就停止了；
	为了确保运行定时任务的协程比主协程后退出，所以这里使用等待组阻塞主协程运行

5.定时任务最小周期
	5.1 最小周期为1s，也就是最短只能设置每间隔1s执行一次，间隔不能小于1s
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(gron.Every(5*time.Second), func() {
		fmt.Println("runs every 5 seconds")
	})
	c.Start()

	wg.Wait()
}
