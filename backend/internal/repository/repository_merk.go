package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryMerk interface {
	GetAllMerk() ([]model.Merk, error)
	GetMerkById(id int) (model.Merk, error)
	ExistByNama(nama string) bool
	SearchMerk(nama string) ([]model.Merk, error)
	CreateMerk(merk model.Merk) (model.Merk, error)
}

type repositoryMerk struct {
	db *gorm.DB
}

func NewRepositoryMerk(db *gorm.DB) RepositoryMerk {
	return &repositoryMerk{db}
}

func (r *repositoryMerk) GetAllMerk() ([]model.Merk, error) {
	var merk []model.Merk
	err := r.db.Find(&merk).Error
	return merk, err
}

func (r *repositoryMerk) GetMerkById(id int) (model.Merk, error) {
	var merk model.Merk
	err := r.db.First(&merk, id).Error
	return merk, err
}

func (r *repositoryMerk) ExistByNama(nama string) bool {
	var merk model.Merk
	var count int64
	r.db.Where("nama = ?", nama).First(&merk).Count(&count)
	return count > 0 // bool
}

func (r *repositoryMerk) SearchMerk(nama string) ([]model.Merk, error) { // Lazy load
	var merk []model.Merk
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&merk).Error
	return merk, err
}

func (r *repositoryMerk) CreateMerk(merk model.Merk) (model.Merk, error) {
	err := r.db.Create(&merk).Error
	if err != nil {
		return model.Merk{}, err
	}

	return merk, nil
}
