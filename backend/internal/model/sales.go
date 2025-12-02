package model

import "time"

type Sales struct {
	ID        int         `json:"id" gorm:"primaryKey"`
	Nama      string      `json:"nama"`
	CreatedAt time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Penjualan []Penjualan // reverse relation 1 sales bisa ada di banyak penjualan
}

func (Sales) TableName() string {
	return "sales"
}
