package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryBarang interface {
	GetAllBarang() ([]model.Barang, error)
	GetBarangById(id int) (model.Barang, error)
	SearchBarang(nama string) ([]model.Barang, error)
	CreateBarang(brg model.Barang) (model.Barang, error)
	UpdateSupplier(id int, updateMap map[string]any) (model.Barang, error)
	UpdateStok(tx *gorm.DB, barang model.Barang) error
}

type repositoryBarang struct {
	db *gorm.DB
}

func NewRepositoryBarang(db *gorm.DB) RepositoryBarang {
	return &repositoryBarang{db}
}

func (r *repositoryBarang) GetAllBarang() ([]model.Barang, error) {
	var barang []model.Barang
	err := r.db.Preload("KategoriBarang").Preload("Merk").Preload("Artikel").Preload("Warna").Preload("JenisBarang").Preload("Ukuran").Find(&barang).Error

	return barang, err
}

func (r *repositoryBarang) GetBarangById(id int) (model.Barang, error) {
	var barang model.Barang
	err := r.db.First(&barang, id).Error
	if err != nil {
		return model.Barang{}, err
	}

	err = r.db.Preload("Merk").Preload("Artikel").Preload("Warna").Preload("KategoriBarang").Preload("JenisBarang").Preload("Ukuran").First(&barang).Error
	if err != nil {
		return model.Barang{}, err
	}

	return barang, nil
}

func (r *repositoryBarang) SearchBarang(nama string) ([]model.Barang, error) { // Lazy load
	var barang []model.Barang
	err := r.db.Preload("Merk").Preload("Ukuran").Preload("Artikel").Preload("Warna").Joins("LEFT JOIN merk on merk.id = barang.merk_id").Where("kode LIKE ? OR merk.nama LIKE ?", "%"+nama+"%", "%"+nama+"%").Limit(20).Find(&barang).Error
	return barang, err
}

func (r *repositoryBarang) CreateBarang(brg model.Barang) (model.Barang, error) {
	err := r.db.Create(&brg).Error
	if err != nil {
		return model.Barang{}, err
	}

	// get data berikut semua preload table
	err = r.db.Preload("Merk").Preload("Artikel").Preload("Warna").Preload("KategoriBarang").Preload("JenisBarang").Preload("Ukuran").First(&brg).Error
	if err != nil {
		return model.Barang{}, err
	}

	return brg, nil
}

func (r *repositoryBarang) UpdateSupplier(id int, updateMap map[string]any) (model.Barang, error) {
	var barang model.Barang
	// get data dulu
	err := r.db.First(&barang, id).Error
	if err != nil {
		return model.Barang{}, err
	}

	// untuk menghindari relasi tabel, kosongkan data struct Jenjang yang terhubung dengan struct Barang ini
	barang.Merk = model.Merk{}
	barang.Artikel = model.Artikel{}
	barang.Warna = model.Warna{}
	barang.KategoriBarang = model.KategoriBarang{}
	barang.JenisBarang = model.JenisBarang{}
	barang.Ukuran = model.Ukuran{}

	// lakukan update
	err = r.db.Model(&barang).Updates(updateMap).Error
	if err != nil {
		return model.Barang{}, err
	}

	// reload data + preload struct terkait untuk lihat perubahan
	err = r.db.Preload("Merk").Preload("Artikel").Preload("Warna").Preload("KategoriBarang").Preload("JenisBarang").Preload("Ukuran").First(&barang).Error
	if err != nil {
		return model.Barang{}, err
	}

	// preload semua dependency
	return barang, nil
}

func (r *repositoryBarang) UpdateStok(tx *gorm.DB, barang model.Barang) error {
	return tx.Save(&barang).Error
}
