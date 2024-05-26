package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func OgManagerLogin(c *gin.Context) (*model.OgManager, string) {
	account := c.PostForm("username")
	password := c.PostForm("password")
	passwd := helper.Sha256cry(password)
	var ogmg model.OgManager
	err := dao.DB.Model(model.OgManager{}).Where("username=? and password=?", account, passwd).First(&ogmg).Error
	if err != nil {
		return nil, "账号或密码错误"
	}
	return &ogmg, "登录成功"
}

// 赛事管理员创建赛事
func OgmanagerCreateRace(c *gin.Context) {
	name := c.PostForm("name")
	stime := c.PostForm("stime")
	placeid := c.PostForm("placeid")
	placeidnum, _ := strconv.Atoi(placeid)
	team1id := c.PostForm("team1id")
	team1idnum, _ := strconv.Atoi(team1id)
	team2id := c.PostForm("team2id")
	team2idnum, _ := strconv.Atoi(team2id)
	ogid := c.PostForm("ogid")
	ogidnum, _ := strconv.Atoi(ogid)
	ogmgid := c.PostForm("ogmgid")
	ogmgidnum, _ := strconv.Atoi(ogmgid)
	err := dao.DB.Create(&model.RaceRecord{
		Name:    name,
		Stime:   stime,
		Placeid: int64(placeidnum),
		Teamid1: int64(team1idnum),
		Teamid2: int64(team2idnum),
		Ogid:    int64(ogidnum),
		Ogmgid:  int64(ogmgidnum),
		Status:  "未开始",
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

// 赛事管理员开始比赛
func StartRace(c *gin.Context) {
	raceid := c.PostForm("raceid")
	raceidnum, _ := strconv.Atoi(raceid)
	err := dao.DB.Model(&model.RaceRecord{}).Where("id = ?", raceidnum).Update("status", "进行中").Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}

func ChangeScore(c *gin.Context) {
	raceid := c.PostForm("raceid")
	raceidnum, _ := strconv.Atoi(raceid)
	team1score := c.PostForm("team1score")
	team1scorenum, _ := strconv.Atoi(team1score)
	team2score := c.PostForm("team2score")
	team2scorenum, _ := strconv.Atoi(team2score)
	err := dao.DB.Model(&model.RaceRecord{}).Where("id = ?", raceidnum).Update("teamscore1", team1scorenum).Update("teamscore2", team2scorenum).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}

// 赛事管理员结束比赛
func EndRace(c *gin.Context) {
	raceid := c.PostForm("raceid")
	raceidnum, _ := strconv.Atoi(raceid)
	err := dao.DB.Model(&model.RaceRecord{}).Where("id = ?", raceidnum).Update("status", "已结束").Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}

// 赛事管理员添加场地
func AddPlace(c *gin.Context) {
	ogid := c.PostForm("ogid")
	ogidnum, _ := strconv.Atoi(ogid)
	name := c.PostForm("name")
	racetype := c.PostForm("racetype")
	err := dao.DB.Create(&model.Place{
		Ogid:     int64(ogidnum),
		Name:     name,
		Status:   "不可出租",
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

// 赛事管理员出租场地
func PubPlace(c *gin.Context) {
	placeid := c.PostForm("placeid")
	placeidnum, _ := strconv.Atoi(placeid)
	err := dao.DB.Model(&model.Place{}).Where("id = ?", placeidnum).Update("status", "可出租").Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "修改失败:" + err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}
