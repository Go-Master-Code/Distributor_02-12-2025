package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryKota interface {
	GetAllKota() ([]model.Kota, error)
	ExistByNama(nama string) bool
	GetKotaById(id int) (model.Kota, error)
	SearchKota(nama string) ([]model.Kota, error)
	CreateKota(kota model.Kota) (model.Kota, error)
}

type repositoryKota struct {
	db *gorm.DB
}

func NewRepositoryKota(db *gorm.DB) RepositoryKota {
	return &repositoryKota{db}
}

func (r *repositoryKota) GetAllKota() ([]model.Kota, error) {
	var kota []model.Kota
	err := r.db.Find(&kota).Error
	return kota, err
}

func (r *repositoryKota) ExistByNama(nama string) bool {
	var kota model.Kota
	var count int64
	r.db.Where("nama = ?", nama).First(&kota).Count(&count)
	return count > 0 // bool
}

func (r *repositoryKota) GetKotaById(id int) (model.Kota, error) {
	var kota model.Kota
	err := r.db.First(&kota, id).Error
	return kota, err
}

func (r *repositoryKota) SearchKota(nama string) ([]model.Kota, error) { // Lazy load
	var kota []model.Kota
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&kota).Error
	return kota, err
}

func (r *repositoryKota) CreateKota(kota model.Kota) (model.Kota, error) {
	err := r.db.Create(&kota).Error
	if err != nil {
		return model.Kota{}, err
	}

	return kota, nil
}
