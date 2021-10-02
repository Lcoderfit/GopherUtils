package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
)

/*
1.自定义选项
	1.1 必须实现flag.Value接口，即String() string 和 Set(string) error 方法
	// src/flag/flag.go
	type Value interface {
	  String() string
	  Set(string) error
	}
	注意：String() string 方法用于格式化输出该类型的值，Set(string) error 方法用于对传入的选项值字符串进行解析

2.flag.Var()用法
	2.1 flag.Var(&variable, "name", "usage") 第一个参数传入自定义变量类型，第二个是选项名，第三个是选项的含义

3.flag.Var/flag.Type/flag.TypeVar的区别
	3.1 flag.Var用于自定义选项的解析
	3.2 flag.Type用于解析Type类型的变量，返回Type类型变量的指针，传入的第一个参数是选项名，第二个是默认值，第三个是含义
	3.3 flag.TypeVar() 传入Type类型变量指针，将解析后的选项值赋予该指针对应的变量，
					传入第一个参数为Type类型的变量指针,第二个参数是选项名，第三个参数是默认值，第四个参数是含义
*/

type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var (
	intervalFlag interval
)

func init() {
	flag.Var(&intervalFlag, "i", "custom interval flag")
}

func main() {
	flag.Parse()
	fmt.Println(intervalFlag)
}
