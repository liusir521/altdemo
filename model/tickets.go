package model

// 门票
type Tickets struct {
	Id     int64   `json:"id"`
	Raceid int64   `gorm:"column:raceid",json:"raceid"`
	Price  float64 `gorm:"column:price",json:"price"`
	Count  int64   `gorm:"column:count",json:"count"`
}

func (Tickets) TableName() string {
	return "tickets"
}
