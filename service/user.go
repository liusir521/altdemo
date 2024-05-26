package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func UserLogin(c *gin.Context) (*model.User, string) {
	account := c.PostForm("username")
	password := c.PostForm("password")
	passwd := helper.Sha256cry(password)
	var user model.User
	err := dao.DB.Model(model.User{}).Where("username=? and password=?", account, passwd).First(&user).Error
	if err != nil {
		return nil, "账号或密码错误"
	}
	return &user, "登录成功"
}

func BuyTickets(c *gin.Context) {
	userid := c.PostForm("userid")
	uid, _ := strconv.Atoi(userid)
	ticketsid := c.PostForm("ticketsid")
	tid, _ := strconv.Atoi(ticketsid)
	price := c.PostForm("price")
	countstr := c.PostForm("count")
	countnum, _ := strconv.Atoi(countstr)
	floatprice, _ := strconv.ParseFloat(price, 64)
	var money float64
	dao.DB.Table("user").Select("money").Where("id = ?", userid).Find(&money)
	if money < floatprice*float64(countnum) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "余额不足",
		})
		return
	}
	var count int64
	dao.DB.Table("tickets").Select("count").Where("id = ?", ticketsid).Find(&count)
	if count < int64(countnum) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "剩余票不足",
		})
		return
	} else {
		dao.DB.Table("user").Where("id = ?", userid).Update("money", money-floatprice*float64(countnum))
		dao.DB.Table("tickets").Where("id = ?", ticketsid).Update("count", gorm.Expr("count - ?", countnum))
		err := dao.DB.Table("ticketshis").Create(&model.TicketsHis{
			Uid:     int64(uid),
			Tid:     int64(tid),
			BuyTime: time.Now().Format("2006-01-02 15:04:05"),
			Count:   int64(countnum),
		}).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "购买失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "购买成功",
			})
		}
	}
}

// 用户购买周边
func BuyGoods(c *gin.Context) {
	userid := c.PostForm("userid")
	uid, _ := strconv.Atoi(userid)
	goodsid := c.PostForm("goodsid")
	gid, _ := strconv.Atoi(goodsid)
	price := c.PostForm("price")
	countstr := c.PostForm("count")
	countnum, _ := strconv.Atoi(countstr)
	floatprice, _ := strconv.ParseFloat(price, 64)
	var money float64
	dao.DB.Table("user").Select("money").Where("id = ?", userid).Find(&money)
	if money < floatprice*float64(countnum) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "余额不足",
		})
		return
	}
	var count int64
	dao.DB.Table("goods").Select("count").Where("id = ?", goodsid).Find(&count)
	if count < int64(countnum) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "库存不足",
		})
		return
	} else {
		dao.DB.Table("user").Where("id = ?", userid).Update("money", money-floatprice*float64(countnum))
		dao.DB.Table("goods").Where("id = ?", gid).Update("count", gorm.Expr("count - ?", countnum))
		err := dao.DB.Table("boughthis").Create(&model.BuyHis{
			Userid:     int64(uid),
			Gid:        int64(gid),
			Cost:       floatprice * float64(countnum),
			Count:      int64(countnum),
			Boughttime: time.Now().Format("2006-01-02 15:04:05"),
		}).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "购买失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "购买成功",
			})
		}
	}
}
