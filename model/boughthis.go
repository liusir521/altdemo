package model

import "time"

// 购买记录
type BuyHis struct {
	Id         int64     `json:"id"`
	Userid     int64     `gorm:"column:uid",json:"uid"`
	Gid        int64     `gorm:"column:gid",json:"gid"`
	Cost       float64   `gorm:"column:cost",json:"cost"`
	Boughttime time.Time `gorm:"column:boughttime",json:"boughttime"`
}

func (BuyHis) TableName() string {
	return "buyhis"
}
