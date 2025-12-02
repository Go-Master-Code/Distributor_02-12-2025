package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryEkspedisi interface {
	GetAllEkspedisi() ([]model.Ekspedisi, error)
	ExistByNama(nama string) bool
	GetEkspedisiById(id int) (model.Ekspedisi, error)
	SearchEkspedisi(nama string) ([]model.Ekspedisi, error)
	CreateEkspedisi(ekspedisi model.Ekspedisi) (model.Ekspedisi, error)
}

type repositoryEkspedisi struct {
	db *gorm.DB
}

func NewRepositoryEkspedisi(db *gorm.DB) RepositoryEkspedisi {
	return &repositoryEkspedisi{db}
}

func (r *repositoryEkspedisi) GetAllEkspedisi() ([]model.Ekspedisi, error) {
	var ekspedisi []model.Ekspedisi
	err := r.db.Find(&ekspedisi).Error
	return ekspedisi, err
}

func (r *repositoryEkspedisi) ExistByNama(nama string) bool {
	var ekspedisi model.Ekspedisi
	var count int64
	r.db.Where("nama = ?", nama).First(&ekspedisi).Count(&count)
	return count > 0 // bool
}

func (r *repositoryEkspedisi) GetEkspedisiById(id int) (model.Ekspedisi, error) {
	var ekspedisi model.Ekspedisi
	err := r.db.First(&ekspedisi, id).Error
	return ekspedisi, err
}

func (r *repositoryEkspedisi) SearchEkspedisi(nama string) ([]model.Ekspedisi, error) { // Lazy load
	var ekspedisi []model.Ekspedisi
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&ekspedisi).Error
	return ekspedisi, err
}

func (r *repositoryEkspedisi) CreateEkspedisi(ekspedisi model.Ekspedisi) (model.Ekspedisi, error) {
	err := r.db.Create(&ekspedisi).Error
	if err != nil {
		return model.Ekspedisi{}, err
	}

	return ekspedisi, nil
}
