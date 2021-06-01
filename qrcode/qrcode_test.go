package qrcode

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"testing"
)

func TestCreateStaticQrCode(t *testing.T) {
	// 扫描二维码之后显示的内容或跳转的链接
	content := "http://blessing.lcoderfit.com"
	// 二维码的像素，200表示像素为200*200
	size := 200
	// 生成的二维码的文件名
	dest := "a.png"
	// 二维码的容错率，原理是二维码在编码过程中进行了冗余操作（例如123被编码成123123），所以只需要扫描到一部分就能识别
	// 二维码的容错率就是指：二维码被遮挡多少后仍可以扫描出信息的能力，容错率越高则可被遮挡的部分越多
	// 容错级别：L:7% (low)  M:15%(medium)  Q:25%(high)	H:30%(highest)  qrcode默认是M级别
	quality := qrcode.Medium
	err := CreateStaticQrCode(content, quality, size, dest)
	if err != nil {
		fmt.Println(err)
	}
}

// TestCreateDynamicQrCode 测试创建动态二维码的函数
func TestCreateDynamicQrCode(t *testing.T) {
	CreateDynamicQrCode()
}
