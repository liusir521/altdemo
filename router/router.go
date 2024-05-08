package router

import (
	"altdemo/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/templates", "./templates")
	r.Static("/images", "./templates/images")
	r.Static("/css", "./templates/css")
	r.Static("/fonts", "./templates/fonts")
	r.Static("/js", "./templates/js")

	// 默认页面
	r.GET("/", service.IndexHtml)

	// 管理员
	r.POST("/manager/login", service.ManagerLogin)

	return r
}
