package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryHargaBarang interface {
	GetHargaBarangById(id int) ([]model.HargaBarang, error)
}

type repositoryHargaBarang struct {
	db *gorm.DB
}

func NewRepositoryHargaBarang(db *gorm.DB) RepositoryHargaBarang {
	return &repositoryHargaBarang{db}
}

func (r *repositoryHargaBarang) GetHargaBarangById(id int) ([]model.HargaBarang, error) {
	var hb []model.HargaBarang
	err := r.db.Preload("Barang").Where("barang_id=?", id).Order("harga DESC").Find(&hb).Error
	return hb, err
}
