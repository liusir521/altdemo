package model

// 商品表
type Goods struct {
	Id    int64   `json:"id"`
	Name  string  `gorm:"column:name",json:"name"`
	Price float64 `gorm:"column:price",json:"price"`
	Info  string  `gorm:"column:info",json:"info"`
	Count int64   `gorm:"column:count",json:"count"`
	Image string  `gorm:"column:image",json:"image"`
}

func (Goods) TableName() string {
	return "goods"
}
