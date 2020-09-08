package api

import (
	"embedded-album/config"
	"embedded-album/utils"
	"embedded-album/views"
	"github.com/gin-gonic/gin"
	"log"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	// 接收参数
	var req LoginReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, views.ErrorResponse(views.ErrorInput))
		return
	}

	// 验证身份
	if req.Username != config.AdminName || req.Password != config.AdminPassword {
		c.JSON(200, views.ErrorResponse(views.ErrorAuthMsg))
		return
	}
	token, err := utils.GenerateToken("admin", []byte("admin"))
	if err != nil {
		log.Println(err)
		c.JSON(200, views.ErrorResponse(views.Error))
		return
	}
	c.JSON(200, views.Response(token))
}
