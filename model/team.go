package model

// 篮球球队
type Team struct {
	Id       int64  `json:"id"`
	Name     string `gorm:"column:name",json:"name"`
	Win      int64  `gorm:"column:win",json:"win"`
	Lose     int64  `gorm:"column:lose",json:"lose"`
	Info     string `gorm:"column:info",json:"info"`
	Racetype string `gorm:"column:racetype",json:"racetype"`
	Image    string `gorm:"column:image",json:"image"`
	Ogid     int64  `gorm:"column:ogid",json:"ogid"`
}

func (Team) TableName() string {
	return "team"
}
