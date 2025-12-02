package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositorySupplier interface {
	GetAllSupplier() ([]model.Supplier, error)
	GetSupplierByID(id int) (model.Supplier, error)
	ExistByNama(nama string) bool
	CreateSupplier(supplier model.Supplier) (model.Supplier, error)
	UpdateSupplier(id int, updateMap map[string]any) (model.Supplier, error)
	DeleteSupplier(id int) (model.Supplier, error)
}

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) RepositorySupplier {
	return &repositorySupplier{db}
}

func (r *repositorySupplier) GetAllSupplier() ([]model.Supplier, error) {
	var supplier []model.Supplier
	err := r.db.Find(&supplier).Error
	return supplier, err
}

func (r *repositorySupplier) GetSupplierByID(id int) (model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.First(&supplier, id).Error
	return supplier, err
}

func (r *repositorySupplier) ExistByNama(nama string) bool {
	var supplier model.Supplier
	var count int64
	r.db.Where("nama = ?", nama).First(&supplier).Count(&count)
	return count > 0 // bool
}

func (r *repositorySupplier) CreateSupplier(supplier model.Supplier) (model.Supplier, error) {
	err := r.db.Create(&supplier).Error
	if err != nil {
		return model.Supplier{}, err
	}

	return supplier, nil
}

func (r *repositorySupplier) UpdateSupplier(id int, updateMap map[string]any) (model.Supplier, error) {
	// get data dulu
	var supplier model.Supplier

	err := r.db.First(&supplier, id).Error
	if err != nil {
		return model.Supplier{}, err
	}

	// jika ketemu, update
	err = r.db.Model(&supplier).Updates(updateMap).Error
	if err != nil {
		return model.Supplier{}, err
	}

	return supplier, nil
}

func (r *repositorySupplier) DeleteSupplier(id int) (model.Supplier, error) {
	// find data
	var supplier model.Supplier
	err := r.db.First(&supplier, id).Error
	if err != nil {
		return model.Supplier{}, err
	}

	// delete data
	err = r.db.Delete(&supplier).Error
	if err != nil {
		return model.Supplier{}, err
	}

	return supplier, nil
}
