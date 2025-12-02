package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryKategoriBarang interface {
	GetAllKategoriBarang() ([]model.KategoriBarang, error)
	GetKategoriBarangById(id int) (model.KategoriBarang, error)
	ExistByNama(nama string) bool
	SearchKategoriBarang(nama string) ([]model.KategoriBarang, error)
	CreateKategoriBarang(kategori model.KategoriBarang) (model.KategoriBarang, error)
}

type repositoryKategoriBarang struct {
	db *gorm.DB
}

func NewRepositoryKategoriBarang(db *gorm.DB) RepositoryKategoriBarang {
	return &repositoryKategoriBarang{db}
}

func (r *repositoryKategoriBarang) GetAllKategoriBarang() ([]model.KategoriBarang, error) {
	var kategori []model.KategoriBarang
	err := r.db.Find(&kategori).Error
	return kategori, err
}

func (r *repositoryKategoriBarang) GetKategoriBarangById(id int) (model.KategoriBarang, error) {
	var kategori model.KategoriBarang
	err := r.db.First(&kategori, id).Error
	return kategori, err
}

func (r *repositoryKategoriBarang) ExistByNama(nama string) bool {
	var kategori model.KategoriBarang
	var count int64
	r.db.Where("nama = ?", nama).First(&kategori).Count(&count)
	return count > 0 // bool
}

func (r *repositoryKategoriBarang) SearchKategoriBarang(nama string) ([]model.KategoriBarang, error) { // Lazy load
	var kb []model.KategoriBarang
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&kb).Error
	return kb, err
}

func (r *repositoryKategoriBarang) CreateKategoriBarang(kategori model.KategoriBarang) (model.KategoriBarang, error) {
	err := r.db.Create(&kategori).Error
	if err != nil {
		return model.KategoriBarang{}, err
	}

	return kategori, nil
}
