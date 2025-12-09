package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryArea interface {
	GetAllArea() ([]model.Area, error)
	SearchArea(nama string) ([]model.Area, error)
}

// struct implementasi
type repositoryArea struct {
	db *gorm.DB
}

// constructor
func NewRepositoryArea(db *gorm.DB) RepositoryArea {
	return &repositoryArea{db}
}

func (r *repositoryArea) GetAllArea() ([]model.Area, error) {
	var area []model.Area
	err := r.db.Find(&area).Error
	return area, err
}

func (r *repositoryArea) SearchArea(nama string) ([]model.Area, error) {
	var area []model.Area
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&area).Error // cari 20 area termirip dengan input
	return area, err
}
