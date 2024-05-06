package model

// 篮球球队
type BasketballTeam struct {
	Id   int64  `json:"id"`
	Name string `gorm:"column:name",json:"name"`
	Win  int64  `gorm:"column:win",json:"win"`
	Lose int64  `gorm:"column:lose",json:"lose"`
	Info string `gorm:"column:info",json:"info"`
}

func (BasketballTeam) TableName() string {
	return "basketballteam"
}
