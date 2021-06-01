package main

import (
	"github.com/cockroachdb/errors"
	"github.com/rsms/gotalk"
	"log"
)

func main() {
	// 新建一个消息类型
	gotalk.Handle("echo", func(in string) (string, error) {
		return "hello " + in, nil
	})
	// 启动服务器, 三个参数分别为协议类型、端口、处理器对象（传入nil表示使用默认的处理器对象）
	// 服务器内部会一直循处理请求
	if err := gotalk.Serve("tcp", ":8061", nil); err != nil {
		log.Printf("%+v\n", errors.Wrap(err, "server"))
	}
}
