package utils

import (
	"github.com/gonutz/framebuffer"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func Image2screen(picPath string) {
	fb, err := framebuffer.Open("/dev/fb0")
	if err != nil {
		panic(err)
	}

	magenta := image.NewUniform(color.RGBA{0, 0, 0, 255})
	draw.Draw(fb, fb.Bounds(), magenta, image.ZP, draw.Src)

	defer fb.Close()
	// 加载图像文件
	file, err := os.Open(picPath)
	if err != nil {
		log.Println("加载图像文件失败", err)
		return
	}
	defer file.Close()

	// 读取图像信息
	pic, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("读取图像信息失败", err)
	}

	draw.Draw(fb, fb.Bounds(), pic, image.ZP, draw.Src)
}
