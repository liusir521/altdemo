package model

// 赛程记录
type TicketsHis struct {
	Id      int64  `json:"id"`
	Uid     int64  `gorm:"column:uid",json:"uid"`
	Tid     int64  `gorm:"column:tid",json:"tid"`
	BuyTime string `gorm:"column:buytime",json:"buytime"`
	Count   int64  `gorm:"column:count",json:"count"`
}

func (TicketsHis) TableName() string {
	return "ticketshis"
}
