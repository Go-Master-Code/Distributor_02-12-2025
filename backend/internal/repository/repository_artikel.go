package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryArtikel interface {
	GetAllArtikel() ([]model.Artikel, error)
	GetArtikelByID(id int) (model.Artikel, error)
	ExistByNama(nama string) bool
	SearchArtikel(nama string) ([]model.Artikel, error)
	CreateArtikel(artikel model.Artikel) (model.Artikel, error)
	UpdateArtikel(id int, updateMap map[string]any) (model.Artikel, error)
}

type repositoryArtikel struct {
	db *gorm.DB
}

func NewRepositoryArtikel(db *gorm.DB) RepositoryArtikel {
	return &repositoryArtikel{db}
}

func (r *repositoryArtikel) GetAllArtikel() ([]model.Artikel, error) {
	var artikel []model.Artikel
	err := r.db.Find(&artikel).Error
	return artikel, err
}

func (r *repositoryArtikel) GetArtikelByID(id int) (model.Artikel, error) {
	var artikel model.Artikel
	err := r.db.First(&artikel, id).Error
	return artikel, err
}

func (r *repositoryArtikel) ExistByNama(nama string) bool {
	var artikel model.Artikel
	var count int64
	r.db.Where("nama = ?", nama).First(&artikel).Count(&count)
	return count > 0 // bool
}

func (r *repositoryArtikel) SearchArtikel(nama string) ([]model.Artikel, error) { // Lazy load
	var artikel []model.Artikel
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&artikel).Error
	return artikel, err
}

func (r *repositoryArtikel) CreateArtikel(artikel model.Artikel) (model.Artikel, error) {
	err := r.db.Create(&artikel).Error
	if err != nil {
		return model.Artikel{}, err
	}

	return artikel, nil
}

func (r *repositoryArtikel) UpdateArtikel(id int, updateMap map[string]any) (model.Artikel, error) {
	// get data dulu
	var artikel model.Artikel
	err := r.db.First(&artikel, id).Error
	if err != nil {
		return model.Artikel{}, err
	}

	err = r.db.Model(&artikel).Updates(updateMap).Error
	if err != nil {
		return model.Artikel{}, err
	}

	return artikel, nil
}
