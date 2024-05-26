package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(c *gin.Context) (*model.Manager, string) {
	account := c.PostForm("username")
	password := c.PostForm("password")
	passwd := helper.Sha256cry(password)
	var manager model.Manager
	err := dao.DB.Model(model.Manager{}).Where("account=? and password=?", account, passwd).First(&manager).Error
	if err != nil {
		return nil, "账号或密码错误"
	}
	return &manager, "登录成功"
}

// 管理员创建赛事方
func ManagerCreateog(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pwd := helper.Sha256cry(password)
	nickname := c.PostForm("nickname")
	email := c.PostForm("email")
	image := c.PostForm("image")
	info := c.PostForm("info")
	pos := c.PostForm("pos")
	contact := c.PostForm("contact")
	err := dao.DB.Model(model.Organizer{}).Create(&model.Organizer{
		Username: username,
		Password: pwd,
		Nickname: nickname,
		Email:    email,
		Image:    image,
		Info:     info,
		Pos:      pos,
		Contact:  contact,
	}).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "创建失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
		})
	}
}

// 管理员删除赛事方账号
func ManagerDelog(c *gin.Context) {
	username := c.PostForm("username")
	err := dao.DB.Model(model.Organizer{}).Where("username=?", username).Delete(&model.Organizer{}).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "删除失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}
}
