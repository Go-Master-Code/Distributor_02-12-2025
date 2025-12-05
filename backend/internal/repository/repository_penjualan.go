package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryPenjualan interface {
	GetAllPenjualan() ([]model.Penjualan, error)
	GetPenjualanPerPeriode(awal string, akhir string) ([]model.Penjualan, error)
	GetPenjualanById(id int) (model.Penjualan, error)
	CreatePenjualan(tx *gorm.DB, penjualan *model.Penjualan) error
	CreatePenjualanDetail(tx *gorm.DB, detail *model.PenjualanDetail) error
}

type repositoryPenjualan struct {
	db *gorm.DB
}

func NewRepositoryPenjualan(db *gorm.DB) RepositoryPenjualan {
	return &repositoryPenjualan{db}
}

// Ambil semua data master detail penjualan
func (r *repositoryPenjualan) GetAllPenjualan() ([]model.Penjualan, error) {
	var penjualan []model.Penjualan
	err := r.db.Preload("Sales").Preload("Toko").Preload("Items").Preload("Items.Barang").Find(&penjualan).Error
	return penjualan, err
}

func (r *repositoryPenjualan) GetPenjualanPerPeriode(awal string, akhir string) ([]model.Penjualan, error) {
	var penjualan []model.Penjualan
	err := r.db.Preload("Sales").Preload("Toko").Preload("Toko.Kota").Preload("Items").Preload("Items.Barang").Where("tgl_penjualan BETWEEN ? AND ?", awal, akhir).Order("tgl_penjualan ASC, no_faktur ASC").Find(&penjualan).Error
	return penjualan, err
}

func (r *repositoryPenjualan) GetPenjualanById(id int) (model.Penjualan, error) {
	var penjualan model.Penjualan
	err := r.db.Preload("Sales").Preload("Toko").Preload("Toko.Kota").Preload("Items").Preload("Items.Barang").First(&penjualan, id).Error
	return penjualan, err
}

// CreatePenjualan insert master penjualan
func (r *repositoryPenjualan) CreatePenjualan(tx *gorm.DB, penjualan *model.Penjualan) error {
	return tx.Create(penjualan).Error
}

// CreatePenjualanDetail insert detail penjualan
func (r *repositoryPenjualan) CreatePenjualanDetail(tx *gorm.DB, detail *model.PenjualanDetail) error {
	return tx.Create(detail).Error
}

/* 	METHOD CreatePenjualan normal untuk CEK API tanpa transaction
func (r *repositoryPenjualan) CreatePenjualan(penjualan model.Penjualan) (model.Penjualan, error) {
	err := r.db.Create(&penjualan).Error
	if err != nil {
		return model.Penjualan{}, err
	}

	err = r.db.Preload("Toko").First(&penjualan).Error
	if err != nil {
		return model.Penjualan{}, err
	}

	return penjualan, nil
}
*/
