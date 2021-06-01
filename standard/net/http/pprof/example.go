package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		content := []byte("hello world")
		_, err := w.Write(content)
		if err != nil {
			log.Println("write body error!")
		}
	}
}

func Init() {
	// 另起一个协程用来运行pprof，
	// "_"导包方式仅调用包中的init函数 "."导包方式将包中的所有可导部分导入，调用时候可以不带包名
	go func() {
		// cup load avg： 在一段时间内，CPU正在处理和CPU处理的进程数量的统计信息
		// CPU 使用率，在一段时间内，CPU被占用的时间与这段时间的比值
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe(":8060", nil)
	if err != nil {
		log.Fatal("server error")
	}
}

func main() {
	Init()
}
