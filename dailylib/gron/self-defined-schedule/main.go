package main

import (
	"fmt"
	"github.com/roylee0704/gron"
	"math"
	"sync"
	"time"
)

// 指数退避(1, 2, 4, 8, ...2^n)
type ExponentialBackOffSchedule struct {
	last int
}

func (e *ExponentialBackOffSchedule) Next(t time.Time) time.Time {
	// 计算下一次触发的时间间隔
	interval := time.Duration(math.Pow(2.0, float64(e.last))) * time.Second
	e.last += 1
	return t.Truncate(time.Second).Add(interval)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	c := gron.New()
	c.AddFunc(&ExponentialBackOffSchedule{}, func() {
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"), "hello")
	})
	c.Start()

	wg.Wait()
}
