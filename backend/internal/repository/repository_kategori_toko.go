package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryKategoriToko interface {
	GetAllKategoriToko() ([]model.KategoriToko, error)
	ExistByNama(nama string) bool
	GetKategoriTokoById(id int) (model.KategoriToko, error)
	SearchKategoriToko(nama string) ([]model.KategoriToko, error)
	CreateKategoriToko(kategori model.KategoriToko) (model.KategoriToko, error)
}

type repositoryKategoriToko struct {
	db *gorm.DB
}

func NewRepositoryKategoriToko(db *gorm.DB) RepositoryKategoriToko {
	return &repositoryKategoriToko{db}
}

func (r *repositoryKategoriToko) GetAllKategoriToko() ([]model.KategoriToko, error) {
	var kategori []model.KategoriToko
	err := r.db.Find(&kategori).Error
	return kategori, err
}

func (r *repositoryKategoriToko) ExistByNama(nama string) bool {
	var kategori model.KategoriToko
	var count int64
	r.db.Where("nama = ?", nama).First(&kategori).Count(&count)
	return count > 0 // bool
}

func (r *repositoryKategoriToko) GetKategoriTokoById(id int) (model.KategoriToko, error) {
	var kategori model.KategoriToko
	err := r.db.First(&kategori, id).Error
	return kategori, err
}

func (r *repositoryKategoriToko) SearchKategoriToko(nama string) ([]model.KategoriToko, error) { // Lazy load
	var kategori []model.KategoriToko
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&kategori).Error
	return kategori, err
}

func (r *repositoryKategoriToko) CreateKategoriToko(kategori model.KategoriToko) (model.KategoriToko, error) {
	err := r.db.Create(&kategori).Error
	if err != nil {
		return model.KategoriToko{}, err
	}

	return kategori, nil
}
