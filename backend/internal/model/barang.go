package model

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	Kode             string         `json:"kode"`
	MerkID           int            `json:"merk_id"`
	Merk             Merk           `json:"merk"`
	ArtikelID        int            `json:"artikel_id"`
	Artikel          Artikel        `json:"artikel"`
	WarnaID          int            `json:"warna_id"`
	Warna            Warna          `json:"warna"`
	KategoriBarangID int            `json:"kategori_barang_id"`
	KategoriBarang   KategoriBarang `json:"kategori_barang"`
	JenisBarangID    int            `json:"jenis_barang_id"`
	JenisBarang      JenisBarang    `json:"jenis_barang"`
	UkuranID         int            `json:"ukuran_id"`
	Ukuran           Ukuran         `json:"ukuran"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"` //tipe datanya bukan time.Time tapi gorm.DeletedAt -> penanda soft delete
}

func (Barang) TableName() string {
	return "barang"
}
