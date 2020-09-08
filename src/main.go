package main

import (
	"embedded-album/config"
	api "embedded-album/controllers"
	"embedded-album/middlewares"
	"embedded-album/models"
	"embedded-album/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// 相片展示
	//go slidePicture()

	// 创建路由
	router := gin.Default()

	// 跨域
	router.Use(middlewares.Cors())

	// 静态资源
	router.Static("/index", "./assets")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.Static("/images", config.StaticFilePath)

	// 登录路由
	router.POST("/login", api.LoginHandler)

	// 中间件
	router.Use(middlewares.Jwt())

	// 图片管理
	router.GET("/pictures", api.PictureGetHandler)
	router.POST("/pictures", api.PictureUploadHandler)
	router.DELETE("/pictures/:upload_name", api.PictureDeleteHandler)

	// 监听端口
	router.Run()
}

func slidePicture() {
	for {
		// 查找图片
		pics, err := models.PictureFindAll()
		if err != nil {
			log.Fatal(err)
		}

		// 轮播图片
		picCount := len(pics)
		if picCount == 0 {
			utils.Image2screen("./assets/index.png")
			time.Sleep(10 * time.Second)
			continue
		}
		for i := 0; i < picCount; i++ {
			utils.Image2screen(config.StaticFilePath + pics[i].UploadName)
			time.Sleep(3 * time.Second)
		}
	}
}
