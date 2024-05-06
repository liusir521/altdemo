package model

// 运动员
type Players struct {
	Id     int64   `json:"id"`
	Name   string  `gorm:"column:name",json:"name"`
	Age    int64   `gorm:"column:age",json:"age"`
	Sex    string  `gorm:"column:sex",json:"sex"`
	Height float32 `gorm:"column:height",json:"height"`
	Weight float32 `gorm:"column:weight",json:"weight"`
	Image  string  `gorm:"column:image",json:"image"`
	Teamid int64   `gorm:"column:teamid",json:"teamid"`
}

func (Players) TableName() string {
	return "players"
}
