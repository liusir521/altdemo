package model

// 购买记录
type BuyHis struct {
	Id         int64   `json:"id"`
	Userid     int64   `gorm:"column:uid",json:"uid"`
	Gid        int64   `gorm:"column:gid",json:"gid"`
	Cost       float64 `gorm:"column:cost",json:"cost"`
	Count      int64   `gorm:"column:count",json:"count"`
	Boughttime string  `gorm:"column:boughttime",json:"boughttime"`
}

func (BuyHis) TableName() string {
	return "boughthis"
}
