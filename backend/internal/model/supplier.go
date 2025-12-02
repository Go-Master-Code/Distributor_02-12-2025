package model

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Nama      string         `json:"nama"`
	Alamat    string         `json:"alamat"`
	Kontak    string         `json:"kotak"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"` //tipe datanya bukan time.Time tapi gorm.DeletedAt -> penanda soft delete
}

func (Supplier) TableName() string {
	return "supplier"
}
