package model

import "time"

// 赛程记录
type RaceRecord struct {
	Id         int64     `json:"id"`
	Ogid       int64     `gorm:"column:ogid",json:"ogid"`
	Placeid    int64     `gorm:"column:placeid",json:"placeid"`
	Ftime      time.Time `gorm:"column:ftime",json:"ftime"`
	teamid1    int64     `gorm:"column:teamid1",json:"teamid1"`
	teamscore1 int64     `gorm:"column:teamscore1",json:"teamscore1"`
	teamid2    int64     `gorm:"column:teamid2",json:"teamid2"`
	teamscore2 int64     `gorm:"column:teamscore2",json:"teamscore2"`
}

func (RaceRecord) TableName() string {
	return "racerecord"
}
