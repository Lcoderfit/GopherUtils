package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
1.为什么线程和进程会有一个ID?
ID相当于一个标识符,通过ID可以获取线程中的数据,并且跨线程进行操作

2.为什么Go中没有暴露出方法来显示获取goroutine ID?
 ”thread-local storage 的成本远远超过了它们的收益。它们只是不适合 Go 语言。”
当goroutine消失时,Goroutine本地存储不会被GC(你可以得到 goid 的当前的 Goroutine，但你不能得到所有运行的 Goroutine 的列表))
当goroutine自己创建了新的Goroutine时,新的Goroutine失去了已有的Goroutine的本地存储
 */

// sync.Pool用于创建一个对象池，存放已分配但是暂时不用的对象，如果需要可以直接从pool中取出，减少Go的GC次数，从而提高性能
var littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}

func main() {
	go func() {
		fmt.Println("当前GoroutineID：", GetGoroutineId())
	}()

	time.Sleep(time.Second)
}

// 进程ID为uint64类型
func GetGoroutineId() uint64 {
	// 从进程池中获取对象, 类型断言为*[]byte
	bp := littleBuf.Get().(*[]byte)
	// 用完后放回池中
	defer littleBuf.Put(bp)
	// 取出来的是一个引用，我们需要用指针取出原来的变量
	b := *bp
	// 获取运行时的堆栈信息,如下：
	// goroutine 6 [running]:
	// main.GetGoroutineId(0x0)
	//     D:/PrivateProje
	b = b[:runtime.Stack(b, false)]
	fmt.Printf("%v\n", string(b))
	res := strings.Split(string(b), " ")[1]
	ans, err := strconv.Atoi(res)
	if err != nil {
		log.Fatal("error")
	}
	return uint64(ans)
}
