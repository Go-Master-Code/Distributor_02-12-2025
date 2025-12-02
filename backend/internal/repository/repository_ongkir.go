package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryOngkir interface {
	GetAllOngkir() ([]model.Ongkir, error)
	ExistByNama(nama string) bool
	GetOngkirById(id int) (model.Ongkir, error)
	SearchOngkir(nama string) ([]model.Ongkir, error)
	CreateOngkir(ongkir model.Ongkir) (model.Ongkir, error)
}

type repositoryOngkir struct {
	db *gorm.DB
}

func NewRepositoryOngkir(db *gorm.DB) RepositoryOngkir {
	return &repositoryOngkir{db}
}

func (r *repositoryOngkir) GetAllOngkir() ([]model.Ongkir, error) {
	var ongkir []model.Ongkir
	err := r.db.Find(&ongkir).Error
	return ongkir, err
}

func (r *repositoryOngkir) ExistByNama(nama string) bool {
	var ongkir model.Ongkir
	var count int64
	r.db.Where("nama = ?", nama).First(&ongkir).Count(&count)
	return count > 0 // bool
}

func (r *repositoryOngkir) GetOngkirById(id int) (model.Ongkir, error) {
	var ongkir model.Ongkir
	err := r.db.First(&ongkir, id).Error
	return ongkir, err
}

func (r *repositoryOngkir) SearchOngkir(nama string) ([]model.Ongkir, error) { // Lazy load
	var ongkir []model.Ongkir
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&ongkir).Error
	return ongkir, err
}

func (r *repositoryOngkir) CreateOngkir(ongkir model.Ongkir) (model.Ongkir, error) {
	err := r.db.Create(&ongkir).Error
	if err != nil {
		return model.Ongkir{}, err
	}

	return ongkir, nil
}
