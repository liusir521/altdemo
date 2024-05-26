package model

// 赛事方
type Organizer struct {
	Id       int64  `json:"id"`
	Username string `gorm:"column:username",json:"username"`
	Password string `gorm:"column:password",json:"password"`
	Nickname string `gorm:"column:nickname",json:"nickname"`
	Info     string `gorm:"column:info",json:"info"`
	Pos      string `gorm:"column:pos",json:"pos"`
	Contact  string `gorm:"column:contact",json:"contact"`
	Email    string `gorm:"column:email",json:"email"`
	Image    string `gorm:"column:image",json:"image"`
}

func (Organizer) TableName() string {
	return "organizer"
}
