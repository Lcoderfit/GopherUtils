package qrcode

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	_ "image/gif"
	_ "image/jpeg"
)

// CreateStaticQrCode 创建普通静态二维码
// barcode: 条形码 	QR code：二维码
// 下载qrcode:go get -u -v github.com/skip2/go-qrcode
// 如果显示443： Timed out, 则需要设置GOPROXY（有的时候在Goland中设置不会及时生效，不知道为什么）
// Terminal手动设置：set GOPROXY=https://goproxy.cn,direct
func CreateStaticQrCode(content string, quality qrcode.RecoveryLevel, size int, sourceFile string) (err error) {
	err = qrcode.WriteFile(content, quality, size, sourceFile)
	return
}

// CreateDynamicQrCode 创建动态二维码
func CreateDynamicQrCode() {
	// 打开二维码图片
	qrFile, err := os.Open("a.png")
	if err != nil {
		log.Fatalf("打开二维码失败 %s", err.Error())
		return
	}
	defer qrFile.Close()

	// 将二维码文件接码成图片
	img, _, err := image.Decode(qrFile)
	if err != nil {
		log.Fatalf("image decode error %s", err.Error())
		return
	}
	// 获取二维码的宽高
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	// 打开要填充的图片
	bgFile, err := os.Open("b.jpg")
	if err != nil {
		log.Fatalf("打开填充图失败 %s", err.Error())
		return
	}
	defer bgFile.Close()

	// 将填充图解码成png图片
	bgImg, _, err := image.Decode(bgFile)
	if err != nil {
		log.Fatalf("填充图解码失败  %s", err.Error())
		return
	}

	// 获取填充图的宽高
	bgWidth, bgHeight := bgImg.Bounds().Max.X, bgImg.Bounds().Max.Y
	// 检测二维码和填充图宽高是否一致
	if width != bgWidth || height != bgHeight {
		// 如果不一致将填充图剪裁
		bgImg = ImageResize(bgImg, width, height)
	}

	// 开始填充二维码
	for y := 0; y < img.Bounds().Max.X; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			qrImgColor := img.At(x, y)
			// 检测图片颜色 如果rgb值是 255 255 255 255 则像素点为白色 跳过
			// 如果rgba值是 0 0 0 0 则为透明色 跳过
			switch img.(type) {
			case *image.NRGBA:
				c := qrImgColor.(color.NRGBA)
				if (c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0) || (c.R == 255 && c.G == 255 && c.B == 255 && c.A == 255) {
					continue
				}

			case *image.RGBA:
				c := qrImgColor.(color.RGBA)
				if (c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0) || (c.R == 255 && c.G == 255 && c.B == 255 && c.A == 255) {
					continue
				}
			}
			// 获取要填充的图片的颜色
			bgImgColor := bgImg.At(x, y)
			// 填充颜色
			switch bgImg.(type) {
			case *image.RGBA64:
				c := bgImgColor.(color.RGBA64)
				img.(draw.Image).Set(x, y, color.RGBA64{R: c.R, G: c.G, B: c.B, A: c.A})

			case *image.NRGBA:
				c := bgImgColor.(color.NRGBA)
				img.(draw.Image).Set(x, y, color.NRGBA{R: c.R, G: c.G, B: c.B, A: c.A})

			case *image.RGBA:
				c := bgImgColor.(color.RGBA)
				img.(draw.Image).Set(x, y, color.RGBA{R: c.R, G: c.G, B: c.B, A: c.A})

			case *image.YCbCr:
				c := bgImgColor.(color.YCbCr)
				img.(draw.Image).Set(x, y, color.YCbCr{Y: c.Y, Cb: c.Cb, Cr: c.Cr})
			default:
				fmt.Println("error")
			}

		}
	}

	filename := fmt.Sprintf("%s.png", time.Now().Format("20060102150405"))
	// 写入文件
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	_ = png.Encode(outFile, img)
}

func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
