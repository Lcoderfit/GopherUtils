package main

import (
	"errors"
	"fmt"
	"github.com/roylee0704/gron"
	"github.com/roylee0704/gron/xtime"
	"sync"
	"time"
)

/*
1.time.Second和time.Duration的关系
	1.1 time.Second/time.Minute/time.Nanosecond/time.Millisecond/time.Microsecond...均是time.Duration类型
	1.2 time.Duration的最小单位为1纳秒，即 Nanosecond Duration = 1 (type Duration int64)
	1.3 要想进行加减乘除运算，必须先转换成time.Duration类型才可，例如:
		a := 100 * time.Second //100会自动转换为time.Duration类型进行计算
		但是：b := 100 a := b * time.Second则会报错，因为int类型不能直接与int64类型进行算术运算
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)

2.指数退避
	每次间隔1 2 4 .... 2^n次方秒即为指数退避
	interval := time.Duration(math.Pow(2.0, float64(last))) * time.Second
	last += 1 // 每次指数+1

3.自定义Schedule
	3.1 无论c.Add() 还是 c.AddFunc()；第一个参数均为Schedule类型，这是一个接口类型：
		type Schedule struct {
			Next(t time.Time) time.Time
		}
	只要实现了Next(t time.Time) time.Time 方法，即可作为参数传入c.Add() 或c.AddFunc()

	3.2 如果不调用c.Start()启动定时任务，则WaitGroup会报死锁错误

4.如何设置当即运行一次/当天指定时间也会运行一次
	4.1 为FourModeSchedule添加一个Every()方法，类似于gron.Every()，不同的是多了一个参数runImmediately，
		若其值为true则表示程序启动后立即执行一次定时任务

	4.2 为FourModeSchedule添加一个At()方法，类似于gron.Every().At()，不同的是多了一个runToday参数，
		若其值为true表示如果当天还没到指定的时间，则当天到指定时间后也会执行一次定时任务，不用等到第二天

	4.3 运行结果，可以看到，在16:30的时候，运行了一次At()指定时间的定时任务
		runs every one day 2021-10-07 16:29:44
		runs every 5 seconds 2021-10-07 16:29:43
		runs every 5 seconds 2021-10-07 16:29:48
		runs every 5 seconds 2021-10-07 16:29:53
		runs every 5 seconds 2021-10-07 16:29:58
		runs every one day 2021-10-07 16:30:00
		runs every 5 seconds 2021-10-07 16:30:03
		runs every 5 seconds 2021-10-07 16:30:08
		runs every 5 seconds 2021-10-07 16:30:13

5.程序模板
	5.1 var wg sync.WaitGroup
		wg.Add(1)
		wg.Wait()

	5.2 在wg.Add(1)和wg.Wait()添加定时任务代码
		c := gron.New() // 创建一个新的定时任务
		c.AddFunc(gron.Every(xxx), func() {xxx})

	5.3 千万别忘记c.Start()，否则主goroutine会一直在wg.Wait()这一步阻塞，导致死锁错误

	5.4 如果需要自定义Schedule，则按照下面模板来编写代码
		type Test struct {xxx}
		func (t *Test) Next(t time.Time) time.Time {xxx}
		xxxxx
		t := Test{}
		c.AddFunc(t, func() {xxx})
*/

type FourModeSchedule struct {
	period         time.Duration
	runImmediately bool
	runToday       bool
	h              int
	m              int
}

func (f *FourModeSchedule) Next(t time.Time) time.Time {
	if f.runImmediately {
		f.runImmediately = false
		return t.Truncate(time.Second)
	}
	if f.runToday {
		f.runToday = false
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day(), f.h, f.m, 0, 0, now.Location())
		if t.After(next) {
			return next.Add(f.period)
		} else {
			return next
		}
	}
	return t.Truncate(time.Second).Add(f.period)
}

func (f *FourModeSchedule) Every(p time.Duration, runImmediately bool) *FourModeSchedule {
	if runImmediately {
		f.runImmediately = true
	}

	// 最小时间间隔为1s
	if p < time.Second {
		p = xtime.Second
	}
	// time.Second为time.Duration类型，且time.Nanosecond数值为1，time.Second是10^9的Nanosecond
	// 所以需要先将p转换为单位为Nanoseconds的数值表示，而由于time.xxxsecond为time.Duration类型，所以需要将
	// 数值转换为time.Duration类型进行除法运算，%time.Second表示求出p中小于1s的部分，然后用原来的p减去该部分即得到去除小于秒的部分的时间
	p = p - time.Duration(p.Nanoseconds())%time.Second
	f.period = p

	return f
}

func (f *FourModeSchedule) At(hm string, runToday bool) *FourModeSchedule {
	if f.period < xtime.Day {
		panic("period must be at least in days")
	}

	h, m, err := Parse(hm)
	if err != nil {
		panic(err.Error())
	}
	f.h = h
	f.m = m
	if runToday {
		f.runToday = runToday
	}
	return f
}

func Parse(hm string) (hh int, mm int, err error) {

	hh = int(hm[0]-'0')*10 + int(hm[1]-'0')
	mm = int(hm[3]-'0')*10 + int(hm[4]-'0')

	if hh < 0 || hh > 24 {
		hh, mm = 0, 0
		err = errors.New("invalid hh format")
	}
	if mm < 0 || mm > 59 {
		hh, mm = 0, 0
		err = errors.New("invalid mm format")
	}

	return
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	f1 := &FourModeSchedule{}
	f2 := &FourModeSchedule{}
	c := gron.New()
	c.AddFunc(f1.Every(5*time.Second, true), func() {
		fmt.Println("runs every 5 seconds", time.Now().Local().Format("2006-01-02 15:04:05"))
	})
	c.AddFunc(f2.Every(1*xtime.Day, true).At("16:30", true), func() {
		fmt.Println("runs every one day", time.Now().Local().Format("2006-01-02 15:04:05"))
	})
	c.Start()

	wg.Wait()
}
