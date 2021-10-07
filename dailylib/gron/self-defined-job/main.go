package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"sync"
	"time"
)

/*
1.c.Add(gron.Every(), j Job)
	1.1 Job是一个接口类型，只要实现了Run方法则实现了Job接口，c.Add第二个参数接收任意实现了Job接口的结构体对象
		type Job interface{
			Run()
		}

	1.2 实现Job接口示例：
		type Test struct{}
		func (t Test) Run() {xxx}

	1.3 调用Add方法
		t := Test{}
		c.Add(gron.Every(5*time.Second), t)
*/

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	g1 := GreetingJob{Name: "lh"}
	g2 := GreetingJob{Name: "ls"}

	c := gron.New()
	c.Add(gron.Every(5*time.Second), g1)
	c.Add(gron.Every(10*time.Second), g2)
	c.Start()
	wg.Wait()
}
