package model

// 赛事方
type Organizer struct {
	Id           int64  `json:"id"`
	Account      string `gorm:"column:account",json:"account"`
	Password     string `gorm:"column:password",json:"password"`
	Nickname     string `gorm:"column:nickname",json:"nickname"`
	Introduction string `gorm:"column:introduction",json:"introduction"`
	Position     string `gorm:"column:position",json:"position"`
	Contact      string `gorm:"column:contact",json:"contact"`
	Email        string `gorm:"column:email",json:"email"`
	Image        string `gorm:"column:image",json:"image"`
}

func (Organizer) TableName() string {
	return "organizer"
}
