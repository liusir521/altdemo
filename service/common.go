package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type_role := c.PostForm("type")
	if type_role == "1" {
		manager, str := ManagerLogin(c)
		fmt.Println(manager, str)
		if manager != nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": manager,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  str,
			})
		}
	} else if type_role == "2" {
		login, s := OrganizerLogin(c)
		if login != nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": login,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  s,
			})
		}
	} else if type_role == "3" {
		login, s := OgManagerLogin(c)
		if login != nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": login,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  s,
			})
		}
	} else if type_role == "4" {
		login, s := UserLogin(c)
		if login != nil {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"data": login,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  s,
			})
		}
	}
}

// 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pwd := helper.Sha256cry(password)
	nickname := c.PostForm("nickname")
	email := c.PostForm("email")
	image := c.PostForm("image")
	user := model.User{
		Username: username,
		Password: pwd,
		Nickname: nickname,
		Email:    email,
		Image:    image,
	}
	err := dao.DB.Create(&user).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "注册失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	}
}

func GetTeam() {
	
}
