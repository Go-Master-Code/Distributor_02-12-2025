package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryWarna interface {
	GetAllWarna() ([]model.Warna, error)
	ExistByNama(nama string) bool
	GetWarnaById(id int) (model.Warna, error)
	SearchWarna(nama string) ([]model.Warna, error)
	CreateWarna(warna model.Warna) (model.Warna, error)
}

type repositoryWarna struct {
	db *gorm.DB
}

func NewRepositoryWarna(db *gorm.DB) RepositoryWarna {
	return &repositoryWarna{db}
}

func (r *repositoryWarna) GetAllWarna() ([]model.Warna, error) {
	var warna []model.Warna
	err := r.db.Find(&warna).Error
	return warna, err
}

func (r *repositoryWarna) ExistByNama(nama string) bool {
	var warna model.Warna
	var count int64
	r.db.Where("nama = ?", nama).First(&warna).Count(&count)
	return count > 0 // bool
}

func (r *repositoryWarna) GetWarnaById(id int) (model.Warna, error) {
	var warna model.Warna
	err := r.db.First(&warna, id).Error
	return warna, err
}

func (r *repositoryWarna) SearchWarna(nama string) ([]model.Warna, error) { // Lazy load
	var warna []model.Warna
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&warna).Error
	return warna, err
}

func (r *repositoryWarna) CreateWarna(warna model.Warna) (model.Warna, error) {
	err := r.db.Create(&warna).Error
	if err != nil {
		return model.Warna{}, err
	}

	return warna, nil
}
