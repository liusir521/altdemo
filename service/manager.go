package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(c *gin.Context) (*model.Manager,string) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	crypaswd := helper.Sha256cry(password)
	var manager model.Manager
	err := dao.DB.Model(model.Manager{}).Where("account=? and password=?", account, crypaswd).First(&manager).Error
	if err != nil {
		return nil, "账号或密码错误"
	}
	return &manager, "登录成功"
}
