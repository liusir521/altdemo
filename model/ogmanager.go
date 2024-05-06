package model

// 体育赛事管理员
type OgManager struct {
	Id       int64  `json:"id"`
	Account  string `gorm:"column:account",json:"account"`
	Password string `gorm:"column:password",json:"password"`
	Nickname string `gorm:"column:nickname",json:"nickname"`
	Ogid     int64  `gorm:"column:ogid",json:"ogid"`
	Racetype string `gorm:"column:racetype",json:"racetype"`
	Image    string `gorm:"column:image",json:"image"`
}

func (OgManager) TableName() string {
	return "ogmanager"
}
