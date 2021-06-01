/*
1.下载：
go get github.com/cheggaaa/pb/v3

2.package命名规范
包名用小写,使用短命名(尽量短，如果有多个单词，则用a-b的形式),尽量和标准库不要冲突。包名统一使用单数形式。

3.文件命名规范
由于文件跟包无任何关系， 而又避免windows大小写的问题，所以推荐的文件规范如下：
文件名应一律使用小写， 不同单词之间用下划线分割, 不用驼峰式，命名应尽可能地见名知意。
尽量见名思义，看见文件名就可以知道这个文件下的大概内容.其中测试文件以_test.go结尾，除测试文件外，命名不出现_。
*/
package cli

import (
	"fmt"
)

// Bar 进度条结构体
type Bar struct {
	percent int64  //当前进度百分比
	cur     int64  //当前进度
	total   int64  //总进度
	rate    string //表示进度条的字符串
	graph   string //显示符号
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		//unicode：9608
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph
	}
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if (bar.percent > last) && (bar.percent%2 == 0) {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
