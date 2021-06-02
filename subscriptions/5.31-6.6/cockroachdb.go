package main

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"sync"
)

/*
cockroach: [ˈkɑːkroʊtʃ]
蟑螂
*/

func main() {
	// 1.打印错误时附带堆栈信息
	//err := errors.New("custom error")
	//fmt.Println("first: ", err)
	//fmt.Printf("second: %v\n", err)
	//fmt.Printf("third: %+v\n", err)
	//
	//e := errors.Newf("%s\n", "aksdjf")
	//fmt.Printf("1: %v\n", e)
	//fmt.Printf("2: %+v\n", e)

	// 2.添加错误前缀
	//n := rand.Int()
	//if n%2 == 1 {
	//	err := FirstPreError()
	//	fmt.Printf("%+v\n", err)
	//} else {
	//	err := SecondPreError()
	//	fmt.Printf("%+v\n", err)
	//}

	// 3.打印两个错误, 注意活学活用，WithSecondaryError可以将两个错误合并成一个错误，假设叫e1, 那么可以将e1再与其他错误继续嵌套，最终可以嵌套n个错误
	// CombineErrors也是同理，可以将多个err进行嵌套
	err := errors.New("first-error")
	// 第一个参数为主要的err，如果第一个参数为nil，则该api返回nil, 如果第一个参数不为nil，则第二个参数会嵌入第一个错误的堆栈树中
	// 还有一个CombineErrors(err1, err2), 如果err1和err2均不为nil，则将err2作为err1的second error， 如果err1为nil则返回err2，如果err2为nil则返回err1
	err = errors.WithSecondaryError(err, errors.New("second-error"))
	fmt.Printf("%+v\n", err)

	// 判断是否是该类型的错误
	fmt.Println(errors.Is(err, errors.New("first-error")))
}


/*****************************返回多个错误*******************************************/
type Mu struct {
	sync.Mutex
	err error
}

type Group struct {
	// 同步锁
	errOnce sync.Once
	mu      Mu
}

func (g *Group) Wait() error {
	return g.mu.err
}

func (g *Group) GO(f func() error) {
	go func() {
		if err := f(); !errors.Is(err, context.Canceled) {
			g.mu.Lock()
			defer g.mu.Unlock()
			// 可以将多个协成报错信息合并成一个错误, 以第一报错的错误为主要错误，其他为辅助错误
			// 所有协程共享一个g.mu.err??????
			g.mu.err = errors.CombineErrors(g.mu.err, err)
		}
	}()
}


/*****************************添加错误前缀*******************************************/
func BaseError() error {
	return errors.New("this is base error")
}

// 使用errors.Wrap添加错误前缀，可以直接return errors.Wrap(....), 这样的话如果err为nil则Wrap返回的也是nil
func FirstPreError() error {
	if err := BaseError(); err != nil {
		return errors.Wrap(err, "FirstPreError")
	}
	return nil
}

func SecondPreError() error {
	if err := BaseError(); err != nil {
		return errors.Wrap(err, "SecondPreError")
	}
	return nil
}
