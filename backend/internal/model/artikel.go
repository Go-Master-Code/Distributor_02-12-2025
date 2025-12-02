package model

import "time"

type Artikel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Barang    []Barang  // reverse relation: 1 artikel ada di banyak barang
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Artikel) TableName() string {
	return "artikel"
}
