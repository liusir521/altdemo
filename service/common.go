package service

import "github.com/gin-gonic/gin"

func IndexHtml(c *gin.Context) {
	c.HTML(200, "index-02.html", nil)
}
