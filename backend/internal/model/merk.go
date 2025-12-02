package model

import "time"

type Merk struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Barang    []Barang  // reverse relation: 1 warna ada di banyak barang
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Merk) TableName() string {
	return "merk"
}
