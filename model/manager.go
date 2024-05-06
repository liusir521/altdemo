package model

// 系统管理员
type Manager struct {
	Id       int64  `json:"id"`
	Account  string `gorm:"column:account",json:"account"`
	Password string `gorm:"column:password",json:"password"`
	Nickname string `gorm:"column:nickname",json:"nickname"`
}

func (Manager) TableName() string {
	return "manager"
}
