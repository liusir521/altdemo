package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	crypaswd := helper.Sha256cry(password)
	var count int64
	dao.DB.Model(model.Manager{}).Where("account=? and password=?", account, crypaswd).Count(&count)
	if count == 1 {
		c.HTML(200, "index.html", gin.H{
			"code":     200,
			"msg":      "登录成功",
			"account": account,
		})
	} else {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "账号或密码错误",
		})
	}
}
