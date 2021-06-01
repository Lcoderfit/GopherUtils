package main

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 第二个参数接收一个接口类型，将第二个参数写入到io.Writer类型中
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		err = errors.Wrap(err, "main")
		log.Printf("%+v\n", err)
	}
}
