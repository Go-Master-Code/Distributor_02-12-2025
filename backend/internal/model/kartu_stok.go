package model

import (
	"time"

	"gorm.io/gorm"
)

type KartuStok struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	BarangID  int            `json:"barang_id"`
	Barang    Barang         `json:"barang"`
	Size      int            `json:"size"`
	Stok      int            `json:"stok"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"` //tipe datanya bukan time.Time tapi gorm.DeletedAt -> penanda soft delete
}

func (KartuStok) TableName() string {
	return "kartu_stok"
}
