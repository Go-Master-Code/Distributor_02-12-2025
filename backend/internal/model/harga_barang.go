package model

import (
	"time"
)

type HargaBarang struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	BarangID     int       `json:"barang_id"`
	Barang       Barang    `json:"barang"`
	Harga        int       `json:"harga"`
	MulaiBerlaku time.Time `json:"mulai_berlaku"`
}

func (HargaBarang) TableName() string {
	return "harga_barang"
}
