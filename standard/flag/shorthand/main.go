package main

import (
	"flag"
	"fmt"
)

/*
1.shorthand: 简记，速记

2.定义短选项
	2.1 flag并不支持短选项，但是可以通过给同一变量设置多个不同选项来实现短选项功能
	2.2 由于为一个相同变量设置多个选项，变量的默认值是由最后一个定义的选项决定的，
		为了避免出现行为不一致，所有对同一变量设置的选项均需要设置相同的默认值
*/

var (
	logLevel string
)

func init() {
	const (
		defaultLogLevel = "DEBUG"
		usage           = "set log level value"
	)

	flag.StringVar(&logLevel, "log-level", defaultLogLevel, usage)
	flag.StringVar(&logLevel, "l", defaultLogLevel, usage+"shorthand")
}

func main() {
	flag.Parse()
	fmt.Println("log level: ", logLevel)
}
