package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexHtml(c *gin.Context) {
	c.HTML(200, "index-02.html", nil)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Identity string `json:"identity"`
}

func Login(c *gin.Context) {
	// 从请求体中解析JSON数据到User结构体
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	fmt.Println(user)
}
