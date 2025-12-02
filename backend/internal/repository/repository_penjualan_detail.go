package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryPenjualanDetail interface {
	GetPenjualanDetailById(id int) ([]model.PenjualanDetail, error)
	// CreatePenjualan(penjualan model.Penjualan) (model.Penjualan, error)
}

type repositoryPenjualanDetail struct {
	db *gorm.DB
}

func NewRepositoryPenjualanDetail(db *gorm.DB) RepositoryPenjualanDetail {
	return &repositoryPenjualanDetail{db}
}

func (r *repositoryPenjualanDetail) GetPenjualanDetailById(id int) ([]model.PenjualanDetail, error) {
	var detilJual []model.PenjualanDetail
	err := r.db.Preload("Barang.KategoriBarang").Preload("Barang.Artikel").Preload("Barang.Warna").Preload("Barang.Ukuran").Where("penjualan_id = ?", id).Find(&detilJual).Error
	return detilJual, err
}
