package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IndexHtml(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	select_role := c.PostForm("select_role")
	fmt.Println("username:", username)
	fmt.Println("password:", password)
	fmt.Println("select_role:", select_role)
}
