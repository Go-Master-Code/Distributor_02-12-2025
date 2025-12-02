package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryJenisBarang interface {
	GetAllJenisBarang() ([]model.JenisBarang, error)
	GetJenisBarangById(id int) (model.JenisBarang, error)
	ExistByNama(nama string) bool
	SearchJenisBarang(nama string) ([]model.JenisBarang, error)
	CreateJenisBarang(jenisBarang model.JenisBarang) (model.JenisBarang, error)
}

type repositoryJenisBarang struct {
	db *gorm.DB
}

func NewRepositoryJenisBarang(db *gorm.DB) RepositoryJenisBarang {
	return &repositoryJenisBarang{db}
}

func (r *repositoryJenisBarang) GetAllJenisBarang() ([]model.JenisBarang, error) {
	var jenisBarang []model.JenisBarang
	err := r.db.Find(&jenisBarang).Error
	return jenisBarang, err
}

func (r *repositoryJenisBarang) GetJenisBarangById(id int) (model.JenisBarang, error) {
	var jenisBarang model.JenisBarang
	err := r.db.First(&jenisBarang, id).Error
	return jenisBarang, err
}

func (r *repositoryJenisBarang) ExistByNama(nama string) bool {
	var jenisBarang model.JenisBarang
	var count int64
	r.db.Where("nama = ?", nama).First(&jenisBarang).Count(&count)
	return count > 0 // bool
}

func (r *repositoryJenisBarang) SearchJenisBarang(nama string) ([]model.JenisBarang, error) { // Lazy load
	var jb []model.JenisBarang
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&jb).Error
	return jb, err
}

func (r *repositoryJenisBarang) CreateJenisBarang(jenisBarang model.JenisBarang) (model.JenisBarang, error) {
	err := r.db.Create(&jenisBarang).Error
	if err != nil {
		return model.JenisBarang{}, err
	}

	return jenisBarang, nil
}
