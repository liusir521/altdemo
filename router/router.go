package router

import (
	"altdemo/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/templates", "./templates")

	// 登录页面
	r.GET("/", service.Login)

	// 管理员
	r.POST("/manager/login", service.ManagerLogin)

	return r
}
