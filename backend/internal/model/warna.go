package model

import "time"

type Warna struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Barang    []Barang  // reverse relation: 1 warna ada di banyak barang
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Warna) TableName() string {
	return "warna"
}
