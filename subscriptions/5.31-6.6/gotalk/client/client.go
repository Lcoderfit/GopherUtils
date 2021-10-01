package main

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/rsms/gotalk"
	"log"
)

func main() {
	// 定义请求时发送的信息
	requestMsg := []string{
		"little ming1",
		"little ming2",
		"little ming3",
		"little ming4",
		"little ming5",
	}
	// 连接服务器，返回一个socket
	s, err := gotalk.Connect("tcp", ":8061")

	if err != nil {
		err = errors.Wrap(err, "client1")
		log.Printf("%+v\n", err)
	}
	var echo string
	for i := 0; i < len(requestMsg); i++ {
		// 第一个参数是消息名，对应服务器注册的消息名，如果请求不存在的消息会返回错误
		// 第二个参数是传给服务器处理器的参数（有且只能传一个）
		// 第三个参数用于接收服务器返回的响应信息
		if err := s.Request("echo", requestMsg[i], &echo); err != nil {
			err = errors.Wrap(err, "client2")
			log.Printf("%+v\n", err)
		}
		fmt.Println(echo)
	}

	s.Close()
}
