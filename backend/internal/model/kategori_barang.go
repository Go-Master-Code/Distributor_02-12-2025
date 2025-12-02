package model

import "time"

type KategoriBarang struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Huruf     string    `json:"huruf"`
	Nama      string    `json:"nama"`
	Barang    []Barang  // reverse relation: 1 kategori barang ada di banyak barang
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (KategoriBarang) TableName() string {
	return "kategori_barang"
}
