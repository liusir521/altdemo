package model

// 用户
type User struct {
	Id       int64   `json:"id"`
	Account  string  `gorm:"column:account",json:"account"`
	Password string  `gorm:"column:password",json:"password"`
	Nickname string  `gorm:"column:nickname",json:"nickname"`
	Email    string  `gorm:"column:email",json:"email"`
	Image    string  `gorm:"column:image",json:"image"`
	Money    float64 `gorm:"column:money",json:"money"`
}

func (User) TableName() string {
	return "user"
}
