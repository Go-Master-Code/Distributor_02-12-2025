package model

import "time"

type Penjualan struct {
	ID            int               `gorm:"primaryKey" json:"id"`
	NoFaktur      string            `gorm:"unique;not null" json:"no_faktur"`
	TglPenjualan  time.Time         `gorm:"not null" json:"tgl_penjualan"`
	TglJatuhTempo time.Time         `gorm:"not null" json:"tgl_jatuh_tempo"`
	TokoID        int               `json:"toko_id"`
	Toko          Toko              `json:"toko"`
	SalesID       int               `json:"sales_id"`
	Sales         Sales             `json:"sales"`
	Total         int               `json:"total"`
	TotalNetto    int               `json:"total_netto"`
	Keterangan    string            `json:"keterangan"`
	CreatedAt     time.Time         `gorm:"column:created_at;autoCreateTime"`
	Items         []PenjualanDetail `gorm:"foreignKey:PenjualanID" json:"items"`
}

func (Penjualan) TableName() string {
	return "penjualan"
}
