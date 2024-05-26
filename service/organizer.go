package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func OrganizerLogin(c *gin.Context) (*model.Organizer, string) {
	account := c.PostForm("username")
	password := c.PostForm("password")
	passwd := helper.Sha256cry(password)
	var og model.Organizer
	err := dao.DB.Model(model.Organizer{}).Where("username=? and password=?", account, passwd).First(&og).Error
	if err != nil {
		return nil, "账号或密码错误"
	}
	return &og, "登录成功"
}

func Createogmanager(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	pwd := helper.Sha256cry(password)
	nickname := c.PostForm("nickname")
	image := c.PostForm("image")
	ogid := c.PostForm("ogid")
	atoi, _ := strconv.Atoi(ogid)
	racetype := c.PostForm("racetype")
	err := dao.DB.Create(&model.OgManager{
		Username: username,
		Password: pwd,
		Nickname: nickname,
		Image:    image,
		Ogid:     int64(atoi),
		Racetype: racetype,
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

// 赛事方创建球队
func CreateTeam(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	racetype := c.PostForm("racetype")
	image := c.PostForm("image")
	ogid, _ := strconv.Atoi(c.PostForm("ogid"))
	err := dao.DB.Create(&model.Team{
		Name:     name,
		Info:     info,
		Racetype: racetype,
		Image:    image,
		Ogid:     int64(ogid),
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

// 赛事方创建球队队员
func CreatePlayer(c *gin.Context) {
	name := c.PostForm("name")
	height := c.PostForm("height")
	heightnum, _ := strconv.ParseFloat(height, 32)
	weight := c.PostForm("weight")
	weightnum, _ := strconv.ParseFloat(weight, 32)
	teamid := c.PostForm("teamid")
	teamidnum, _ := strconv.Atoi(teamid)
	sex := c.PostForm("sex")
	image := c.PostForm("image")
	err := dao.DB.Create(&model.Players{
		Name:   name,
		Height: float32(heightnum),
		Weight: float32(weightnum),
		Teamid: int64(teamidnum),
		Sex:    sex,
		Image:  image,
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

// 赛事方发布门票
func Pubtickets(c *gin.Context) {
	raceid := c.PostForm("raceid")
	raceidnum, _ := strconv.Atoi(raceid)
	price := c.PostForm("price")
	pricenum, _ := strconv.ParseFloat(price, 64)
	count := c.PostForm("count")
	countnum, _ := strconv.Atoi(count)
	err := dao.DB.Create(&model.Tickets{
		Raceid: int64(raceidnum),
		Price:  float64(pricenum),
		Count:  int64(countnum),
	}).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "发布失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "发布成功",
		})
	}
}

// 赛事方发布周边商品
func PubGoods(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	image := c.PostForm("image")
	price := c.PostForm("price")
	pricenum, _ := strconv.ParseFloat(price, 64)
	count := c.PostForm("count")
	countnum, _ := strconv.Atoi(count)
	err := dao.DB.Create(&model.Goods{
		Name:  name,
		Info:  info,
		Image: image,
		Price: float64(pricenum),
		Count: int64(countnum),
	}).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "发布失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "发布成功",
		})
	}
}
