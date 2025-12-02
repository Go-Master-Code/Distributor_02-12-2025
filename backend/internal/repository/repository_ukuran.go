package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryUkuran interface {
	GetAllUkuran() ([]model.Ukuran, error)
	GetUkuranById(id int) (model.Ukuran, error)
	ExistByNama(nama string) bool
	SearchUkuran(nama string) ([]model.Ukuran, error)
	CreateUkuran(ukuran model.Ukuran) (model.Ukuran, error)
}

type repositoryUkuran struct {
	db *gorm.DB
}

func NewRepositoryUkuran(db *gorm.DB) RepositoryUkuran {
	return &repositoryUkuran{db}
}

func (r *repositoryUkuran) GetAllUkuran() ([]model.Ukuran, error) {
	var ukuran []model.Ukuran
	err := r.db.Find(&ukuran).Error
	return ukuran, err
}

func (r *repositoryUkuran) GetUkuranById(id int) (model.Ukuran, error) {
	var ukuran model.Ukuran
	err := r.db.First(&ukuran, id).Error
	return ukuran, err
}

func (r *repositoryUkuran) ExistByNama(nama string) bool {
	var ukuran model.Ukuran
	var count int64
	r.db.Where("nama = ?", nama).First(&ukuran).Count(&count)
	return count > 0 // bool
}

func (r *repositoryUkuran) CreateUkuran(ukuran model.Ukuran) (model.Ukuran, error) {
	err := r.db.Create(&ukuran).Error
	if err != nil {
		return model.Ukuran{}, err
	}

	return ukuran, nil
}

func (r *repositoryUkuran) SearchUkuran(nama string) ([]model.Ukuran, error) {
	var ukuran []model.Ukuran
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&ukuran).Error
	return ukuran, err
}
