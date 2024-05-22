package router

import (
	"altdemo/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/templates", "./templates")
	r.Static("/images", "/templates/assets/images")
	r.Static("/css", "/templates/assets/css")
	r.Static("/scss", "/templates/assets/scss")
	r.Static("/vendor", "/templates/assets/vendor")
	r.Static("/js", "/templates/assets/js")

	// 默认页面
	r.GET("/", service.IndexHtml)

	r.POST("/login", service.Login)

	// 管理员

	return r
}
