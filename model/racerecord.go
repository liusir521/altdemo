package model

// 赛程记录
type RaceRecord struct {
	Id         int64  `json:"id"`
	Ogmgid     int64  `gorm:"column:ogmgid",json:"ogmgid"`
	Name       string `gorm:"column:name",json:"name"`
	Ogid       int64  `gorm:"column:ogid",json:"ogid"`
	Placeid    int64  `gorm:"column:placeid",json:"placeid"`
	Stime      string `gorm:"column:stime",json:"stime"`
	Status     string `gorm:"column:status",json:"status"`
	Teamid1    int64  `gorm:"column:teamid1",json:"teamid1"`
	Teamscore1 int64  `gorm:"column:teamscore1",json:"teamscore1"`
	Teamid2    int64  `gorm:"column:teamid2",json:"teamid2"`
	Teamscore2 int64  `gorm:"column:teamscore2",json:"teamscore2"`
}

func (RaceRecord) TableName() string {
	return "racerecord"
}
