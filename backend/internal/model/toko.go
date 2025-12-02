package model

import (
	"time"
)

type Toko struct {
	ID             int          `json:"id" gorm:"primaryKey"`
	Kode           string       `json:"kode"`
	Nama           string       `json:"nama"`
	KategoriTokoID int          `json:"kategori_toko_id"`
	KategoriToko   KategoriToko `json:"kategori_toko"`
	KotaID         int          `json:"kota_id"`
	Kota           Kota         `json:"kota"`
	Alamat         string       `json:"alamat"`
	Disc1          float64      `json:"disc_1" gorm:"column:disc_1"` // tambahkan tag gorm, nama fieldnya harus sama dengan di mysql!
	Disc2          float64      `json:"disc_2" gorm:"column:disc_2"` // tambahkan tag gorm, nama fieldnya harus sama dengan di mysql!
	Disc3          float64      `json:"disc_3" gorm:"column:disc_3"` // tambahkan tag gorm, nama fieldnya harus sama dengan di mysql!
	EkspedisiID    int          `json:"ekspedisi_id"`
	Ekspedisi      Ekspedisi    `json:"ekspedisi"`
	OngkirID       int          `json:"ongkir_id"`
	Ongkir         Ongkir       `json:"ongkir"`
	CreatedAt      time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Penjualan      []Penjualan  // reverse relation 1 toko bisa ada di banyak penjualan
	// DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"` //tipe datanya bukan time.Time tapi gorm.DeletedAt -> penanda soft delete
}

func (Toko) TableName() string {
	return "toko"
}
