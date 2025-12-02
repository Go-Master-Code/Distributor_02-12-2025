package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositorySales interface {
	GetAllSales() ([]model.Sales, error)
	GetSalesById(id int) (model.Sales, error)
	SearchSales(nama string) ([]model.Sales, error)
	CreateSales(sales model.Sales) (model.Sales, error)
}

type repositorySales struct {
	db *gorm.DB
}

func NewRepositorySales(db *gorm.DB) RepositorySales {
	return &repositorySales{db}
}

func (r *repositorySales) GetAllSales() ([]model.Sales, error) {
	var sales []model.Sales
	err := r.db.Find(&sales).Error
	return sales, err
}

func (r *repositorySales) GetSalesById(id int) (model.Sales, error) {
	var sales model.Sales
	err := r.db.First(&sales, id).Error
	return sales, err
}

func (r *repositorySales) SearchSales(nama string) ([]model.Sales, error) {
	var sales []model.Sales
	err := r.db.Where("nama LIKE ?", "%"+nama+"%").Limit(20).Find(&sales).Error
	return sales, err
}

func (r *repositorySales) CreateSales(sales model.Sales) (model.Sales, error) {
	err := r.db.Create(&sales).Error
	if err != nil {
		return model.Sales{}, err
	}

	return sales, nil
}
