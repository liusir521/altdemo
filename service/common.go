package service

import (
	"altdemo/dao"
	"altdemo/helper"
	"altdemo/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetTeam(c *gin.Context) {
	ogid := c.Query("ogid")
	atoi, err := strconv.Atoi(ogid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
	} else {
		var teams []model.Team
		err := dao.DB.Table("team").Where("ogid = ?", atoi).Find(&teams).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": teams,
			})
		}
	}
}

func GetRaceUnstart(c *gin.Context) {
	ogid := c.Query("ogid")
	atoi, err := strconv.Atoi(ogid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
	} else {
		var races []model.RaceRecord
		err := dao.DB.Table("racerecord").Where("ogid = ? AND status = ?", atoi, "未开始").Find(&races).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": races,
			})
		}
	}
}

func GetTeamInfo(c *gin.Context) {
	teamid := c.Query("teamid")
	atoi, err := strconv.Atoi(teamid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
	} else {
		data := make(map[string]interface{}, 3)
		var team map[string]interface{}
		err := dao.DB.Table("team").Where("id = ?", atoi).Find(&team).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		}
		data["team"] = team
		var players []model.Players
		err = dao.DB.Table("players").Where("teamid = ?", atoi).Find(&players).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		}
		data["players"] = players
		var racehis []model.RaceRecord
		err = dao.DB.Table("racerecord").Where("(teamid1 = ? or teamid2 = ?) AND status = ?", atoi, atoi, "已结束").Find(&racehis).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		} else {
			data["racehis"] = racehis
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": data,
			})
		}
	}
}

// 查询所有在售票务
func GetTickets(c *gin.Context) {
	var tickets []map[string]interface{}
	err := dao.DB.Table("tickets").Where("status = ?", "在售").Find(&tickets).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询失败:" + err.Error(),
		})
		return
	}
	for _, ticket := range tickets {
		raceid := ticket["raceid"]
		var race []model.RaceRecord
		err := dao.DB.Table("racerecord").Where("id = ?", raceid).Find(&race).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		} else {
			ticket["race"] = race
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": tickets,
	})
}

// 查询所有周边商品
func GetGoods(c *gin.Context) {
	var goods []model.Goods
	err := dao.DB.Table("goods").Find(&goods).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询失败:" + err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": goods,
		})
	}
}

// 获取用户购票记录
func GetUserTicketsHis(c *gin.Context) {
	userid := c.Query("userid")
	atoi, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
	} else {
		var tickets []map[string]interface{}
		err := dao.DB.Table("ticketshis").Where("uid = ?", atoi).Find(&tickets).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		} else {
			for _, ticket := range tickets {
				tid := ticket["tid"]
				var ticketinfo model.Tickets
				err := dao.DB.Table("tickets").Where("id = ?", tid).Find(&ticketinfo).Error
				if err != nil {
					c.JSON(200, gin.H{
						"code": 0,
						"msg":  "查询失败:" + err.Error(),
					})
					return
				} else {
					ticket["price"] = ticketinfo.Price
					var race model.RaceRecord
					dao.DB.Table("racerecord").Where("id = ?", ticketinfo.Raceid).Find(&race)
					ticket["name"] = race.Name
				}
			}
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": tickets,
			})
		}
	}
}

// 获取用户周边商品记录
func GetUserGoodsHis(c *gin.Context) {
	userid := c.Query("userid")
	atoi, err := strconv.Atoi(userid)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "参数错误",
		})
	} else {
		var goods []map[string]interface{}
		err := dao.DB.Table("boughthis").Where("uid = ?", atoi).Find(&goods).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "查询失败:" + err.Error(),
			})
			return
		} else {
			for _, good := range goods {
				gid := good["gid"]
				var goodinfo model.Goods
				err := dao.DB.Table("goods").Where("id = ?", gid).Find(&goodinfo).Error
				if err != nil {
					c.JSON(200, gin.H{
						"code": 0,
						"msg":  "查询失败:" + err.Error(),
					})
					return
				} else {
					good["name"] = goodinfo.Name
					good["image"] = goodinfo.Image
				}
			}
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "查询成功",
				"data": goods,
			})
		}
	}
}

// 获取可租用场地
func UsablePlaces(c *gin.Context) {
	var palces []map[string]interface{}
	err := dao.DB.Table("place").Where("status = ?", "可出租").Find(&palces).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询失败:" + err.Error(),
		})
		return
	} else {
		for _, palce := range palces {
			ogid := palce["ogid"]
			var og model.Organizer
			err := dao.DB.Table("organizer").Where("id = ?", ogid).Find(&og).Error
			if err != nil {
				c.JSON(200, gin.H{
					"code": 0,
					"msg":  "查询失败:" + err.Error(),
				})
				return
			} else {
				palce["contact"] = og.Contact
			}
		}
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": palces,
		})
	}
}

// 获取所有赛事方
func GetAllOgs(c *gin.Context) {
	var ogs []model.Organizer
	err := dao.DB.Table("organizer").Find(&ogs).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "查询失败:" + err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": ogs,
		})
	}
}
