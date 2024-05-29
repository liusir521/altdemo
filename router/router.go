package router

import (
	"altdemo/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", service.Login)
	r.POST("/register", service.Register)

	// 管理员
	r.POST("/manager/createog", service.ManagerCreateog)
	r.POST("/manager/delog", service.ManagerDelog)

	// 赛事方
	r.POST("/organizer/createogmanager", service.Createogmanager)
	r.POST("/organizer/createteam", service.CreateTeam)
	r.POST("/organizer/createplayer", service.CreatePlayer)
	r.POST("/organizer/pubtickets", service.Pubtickets)
	r.POST("/organizer/pubgoods", service.PubGoods)

	// 赛事管理员
	r.POST("/ogmanager/createrace", service.OgmanagerCreateRace)
	r.POST("/ogmanager/startrace", service.StartRace)
	r.POST("/ogmanager/changescore", service.ChangeScore)
	r.POST("/ogmanager/endrace", service.EndRace)
	r.POST("/ogmanager/addplace", service.AddPlace)
	r.POST("/ogmanager/pubplaces", service.PubPlace)

	// 用户
	r.POST("/user/buytickets", service.BuyTickets)
	r.POST("/user/buygoods", service.BuyGoods)

	// get资源信息
	r.GET("/team", service.GetTeam)
	r.GET("/raceunstart", service.GetRaceUnstart)
	r.GET("/teaminfo", service.GetTeamInfo)
	r.GET("/tickets", service.GetTickets)
	r.GET("/goods", service.GetGoods)
	r.GET("/userticketshis", service.GetUserTicketsHis)
	r.GET("/usergoodshis", service.GetUserGoodsHis)
	r.GET("/usableplaces", service.UsablePlaces)
	return r
}
