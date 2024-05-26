package model

// 赛事场地
type Place struct {
	Id       int64  `json:"id"`
	Ogid     int64  `gorm:"column:ogid",json:"ogid"`
	Name     string `gorm:"column:name",json:"name"`
	Status   string `gorm:"column:status",json:"status"`
	Racetype string `gorm:"column:racetype",json:"racetype"`
}

func (Place) TableName() string {
	return "place"
}
