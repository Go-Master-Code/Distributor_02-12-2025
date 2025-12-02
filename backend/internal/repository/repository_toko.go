package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryToko interface {
	GetAllToko() ([]model.Toko, error)
	GetTokoById(id int) (model.Toko, error)
	SearchToko(nama string) ([]model.Toko, error)
	CreateToko(brg model.Toko) (model.Toko, error)
}

type repositoryToko struct {
	db *gorm.DB
}

func NewRepositoryToko(db *gorm.DB) RepositoryToko {
	return &repositoryToko{db}
}

func (r *repositoryToko) GetAllToko() ([]model.Toko, error) {
	var toko []model.Toko
	err := r.db.Preload("KategoriToko").Preload("Kota").Preload("Ekspedisi").Preload("Ongkir").Find(&toko).Error
	return toko, err
}

func (r *repositoryToko) SearchToko(nama string) ([]model.Toko, error) { // Lazy load
	var toko []model.Toko
	err := r.db.Where("nama LIKE ? OR kode LIKE ?", "%"+nama+"%", "%"+nama+"%").Limit(20).Find(&toko).Error
	return toko, err
}

func (r *repositoryToko) GetTokoById(id int) (model.Toko, error) {
	var toko model.Toko
	err := r.db.First(&toko, id).Error
	if err != nil {
		return model.Toko{}, err
	}

	err = r.db.Preload("KategoriToko").Preload("Kota").Preload("Ekspedisi").Preload("Ongkir").First(&toko).Error
	if err != nil {
		return model.Toko{}, err
	}

	return toko, nil
}

func (r *repositoryToko) CreateToko(toko model.Toko) (model.Toko, error) {
	err := r.db.Create(&toko).Error
	if err != nil {
		return model.Toko{}, err
	}

	// get data berikut semua preload table
	err = r.db.Preload("KategoriToko").Preload("Kota").Preload("Ekspedisi").Preload("Ongkir").First(&toko).Error
	if err != nil {
		return model.Toko{}, err
	}

	return toko, nil
}
