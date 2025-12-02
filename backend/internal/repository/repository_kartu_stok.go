package repository

import (
	"api-distributor/internal/model"

	"gorm.io/gorm"
)

type RepositoryKartuStok interface {
	GetAllKartuStok() ([]model.KartuStok, error)
	GetKartuStokByIdBarang(idBarang int) ([]model.KartuStok, error)
	GetStokByBarangIdAndSize(idBarang int, size int) (model.KartuStok, error)
	UpdateStokDiKartuStok(tx *gorm.DB, kartuStok model.KartuStok) error
	// SearchBarang(nama string) ([]model.KartuStok, error)
	// CreateBarang(brg model.KartuStok) (model.KartuStok, error)
	// UpdateSupplier(id int, updateMap map[string]any) (model.KartuStok, error)
	// UpdateStok(tx *gorm.DB, KartuStok model.KartuStok) error
}

type repositoryKartuStok struct {
	db *gorm.DB
}

func NewRepositoryKartuStok(db *gorm.DB) RepositoryKartuStok {
	return &repositoryKartuStok{db}
}

func (r *repositoryKartuStok) GetAllKartuStok() ([]model.KartuStok, error) {
	var kartuStok []model.KartuStok
	err := r.db.Preload("Barang").Preload("Barang.Artikel").Preload("Barang.Merk").Preload("Barang.Warna").Preload("Barang.KategoriBarang").Preload("Barang.JenisBarang").Find(&kartuStok).Error
	return kartuStok, err
}

func (r *repositoryKartuStok) GetKartuStokByIdBarang(idBarang int) ([]model.KartuStok, error) {
	var kartuStok []model.KartuStok
	err := r.db.Preload("Barang").Preload("Barang.Artikel").Preload("Barang.Warna").Preload("Barang.Merk").Preload("Barang.KategoriBarang").Preload("Barang.JenisBarang").Where("barang_id=?", idBarang).Find(&kartuStok).Error
	return kartuStok, err
}

func (r *repositoryKartuStok) GetStokByBarangIdAndSize(idBarang int, size int) (model.KartuStok, error) {
	var kartuStok model.KartuStok
	err := r.db.Preload("Barang").Preload("Barang.Artikel").Preload("Barang.Merk").Preload("Barang.Warna").Where("barang_id=? AND size=?", idBarang, size).First(&kartuStok).Error
	return kartuStok, err
}

func (r *repositoryKartuStok) UpdateStokDiKartuStok(tx *gorm.DB, kartuStok model.KartuStok) error {
	return tx.Save(&kartuStok).Error
}

// func (r *repositoryBarang) GetAllBarang() ([]model.KartuStok, error) {
// 	var KartuStok []model.KartuStok
// 	err := r.db.Preload("KategoriBarang").Preload("Merk").Preload("Artikel").Preload("Warna").Preload("JenisBarang").Preload("Ukuran").Find(&KartuStok).Error

// 	return KartuStok, err
// }

// func (r *repositoryBarang) SearchBarang(nama string) ([]model.KartuStok, error) { // Lazy load
// 	var KartuStok []model.KartuStok
// 	err := r.db.Preload("Merk").Preload("Ukuran").Preload("Artikel").Preload("Warna").Joins("LEFT JOIN merk on merk.id = KartuStok.merk_id").Where("kode LIKE ? OR merk.nama LIKE ?", "%"+nama+"%", "%"+nama+"%").Limit(20).Find(&KartuStok).Error
// 	return KartuStok, err
// }

// func (r *repositoryBarang) CreateBarang(brg model.KartuStok) (model.KartuStok, error) {
// 	err := r.db.Create(&brg).Error
// 	if err != nil {
// 		return model.KartuStok{}, err
// 	}

// 	// get data berikut semua preload table
// 	err = r.db.Preload("Merk").Preload("Artikel").Preload("Warna").Preload("KategoriBarang").Preload("JenisBarang").Preload("Ukuran").First(&brg).Error
// 	if err != nil {
// 		return model.KartuStok{}, err
// 	}

// 	return brg, nil
// }

// func (r *repositoryBarang) UpdateSupplier(id int, updateMap map[string]any) (model.KartuStok, error) {
// 	var KartuStok model.KartuStok
// 	// get data dulu
// 	err := r.db.First(&KartuStok, id).Error
// 	if err != nil {
// 		return model.KartuStok{}, err
// 	}

// 	// untuk menghindari relasi tabel, kosongkan data struct Jenjang yang terhubung dengan struct KartuStok ini
// 	KartuStok.Merk = model.Merk{}
// 	KartuStok.Artikel = model.Artikel{}
// 	KartuStok.Warna = model.Warna{}
// 	KartuStok.KategoriBarang = model.KategoriBarang{}
// 	KartuStok.JenisBarang = model.JenisBarang{}
// 	KartuStok.Ukuran = model.Ukuran{}

// 	// lakukan update
// 	err = r.db.Model(&KartuStok).Updates(updateMap).Error
// 	if err != nil {
// 		return model.KartuStok{}, err
// 	}

// 	// reload data + preload struct terkait untuk lihat perubahan
// 	err = r.db.Preload("Merk").Preload("Artikel").Preload("Warna").Preload("KategoriBarang").Preload("JenisBarang").Preload("Ukuran").First(&KartuStok).Error
// 	if err != nil {
// 		return model.KartuStok{}, err
// 	}

// 	// preload semua dependency
// 	return KartuStok, nil
// }

// func (r *repositoryBarang) UpdateStok(tx *gorm.DB, KartuStok model.KartuStok) error {
// 	return tx.Save(&KartuStok).Error
// }
