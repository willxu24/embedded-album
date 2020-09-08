package api

import (
	"embedded-album/config"
	"embedded-album/models"
	"embedded-album/utils"
	"embedded-album/views"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

func PictureGetHandler(c *gin.Context) {
	// 查找图片
	res, err := models.PictureFindAll()
	if err != nil {
		log.Println(err)
		c.JSON(200, views.ErrorResponse(views.Error))
		return
	}

	// 返回结果
	c.JSON(200, views.Response(res))
	return
}

func PictureUploadHandler(c *gin.Context) {
	// 接收参数
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, views.ErrorResponse(views.ErrorInput))
		return
	}

	// 生成哈希值
	fileContent, _ := file.Open()
	fileHash := utils.File2md5(fileContent)

	// 生成文件名
	fileSuffix := path.Ext(filepath.Base(file.Filename))
	uploadName := fileHash + fileSuffix

	// 上传图片
	err = c.SaveUploadedFile(file, config.StaticFilePath+uploadName)
	if err != nil {
		log.Println(err)
		c.JSON(200, views.ErrorResponse(views.Error))
		return
	}

	// 插入数据
	if err := models.PictureInsertOne(models.Picture{
		//ID:         0,
		CreatedAt:  time.Now().Unix(),
		OriginName: file.Filename,
		UploadName: uploadName,
	}); err != nil {
		log.Println(err)
		c.JSON(200, views.ErrorResponse(views.Error))
		return
	}

	// 返回结果
	c.JSON(200, views.Response(nil))
}

func PictureDeleteHandler(c *gin.Context) {
	// 接收参数
	uploadName := c.Param("upload_name")

	// 删除图片
	os.Remove(config.StaticFilePath + uploadName)

	// 删除数据
	err := models.PictureDeleteOne(models.Picture{
		UploadName: uploadName,
	})
	if err != nil {
		log.Println(err)
		c.JSON(200, views.ErrorResponse(views.Error))
		return
	}

	// 返回结果
	c.JSON(200, views.Response(nil))
}
