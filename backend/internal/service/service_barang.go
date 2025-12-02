package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
)

type ServiceBarang interface {
	GetAllBarang() ([]dto.BarangResponse, error)
	GetBarangById(id int) (dto.BarangResponse, error)
	SearchBarang(nama string) ([]dto.BarangResponse, error)
	CreateBarang(brg dto.CreateBarangRequest) (dto.BarangResponse, error)
	UpdateBarang(id int, req dto.UpdateBarangRequest) (dto.BarangResponse, error)
}

type serviceBarang struct {
	repo repository.RepositoryBarang
}

func NewServiceBarang(repo repository.RepositoryBarang) ServiceBarang {
	return &serviceBarang{repo}
}

func (s *serviceBarang) GetAllBarang() ([]dto.BarangResponse, error) {
	barang, err := s.repo.GetAllBarang()
	if err != nil {
		return []dto.BarangResponse{}, err
	}

	// convert model ke dto
	barangDTO := helper.ConvertToDTOBarangPlural(barang)
	return barangDTO, nil
}

func (s *serviceBarang) GetBarangById(id int) (dto.BarangResponse, error) {
	barang, err := s.repo.GetBarangById(id)
	if err != nil {
		return dto.BarangResponse{}, err
	}

	// convert model to dto
	barangDTO := helper.ConvertToDTOBarangSingle(barang)
	return barangDTO, nil
}

func (s *serviceBarang) SearchBarang(nama string) ([]dto.BarangResponse, error) {
	barang, err := s.repo.SearchBarang(nama)
	if err != nil {
		return []dto.BarangResponse{}, err
	}

	// convert model to dto
	barangDTO := helper.ConvertToDTOBarangPlural(barang)

	return barangDTO, nil
}

func (s *serviceBarang) CreateBarang(brg dto.CreateBarangRequest) (dto.BarangResponse, error) {
	// buat model barang
	req := model.Barang{
		Kode:             brg.Kode,
		MerkID:           brg.MerkID,
		ArtikelID:        brg.ArtikelID,
		WarnaID:          brg.WarnaID,
		KategoriBarangID: brg.KategoriBarangID,
		JenisBarangID:    brg.JenisBarangID,
		UkuranID:         brg.UkuranID,
	}

	newBarang, err := s.repo.CreateBarang(req)
	if err != nil {
		return dto.BarangResponse{}, err
	}

	// convert to dto
	barangDTO := helper.ConvertToDTOBarangSingle(newBarang)
	return barangDTO, nil
}

func (s *serviceBarang) UpdateBarang(id int, req dto.UpdateBarangRequest) (dto.BarangResponse, error) {
	// convert dto.Update ke map
	var updateMap = map[string]any{}

	if req.Kode != nil {
		updateMap["kode"] = *req.Kode // dereference
	}

	if req.MerkID != nil {
		updateMap["merk_id"] = *req.MerkID // dereference
	}

	if req.ArtikelID != nil {
		updateMap["artikel_id"] = *req.ArtikelID // dereference
	}

	if req.WarnaID != nil {
		updateMap["warna_id"] = *req.WarnaID // dereference
	}

	if req.KategoriBarangID != nil {
		updateMap["kategori_barang_id"] = *req.KategoriBarangID // dereference
	}

	if req.JenisBarangID != nil {
		updateMap["jenis_barang_id"] = *req.JenisBarangID // dereference
	}

	if req.UkuranID != nil {
		updateMap["ukuran_id"] = *req.UkuranID // dereference
	}

	updatedBarang, err := s.repo.UpdateSupplier(id, updateMap)
	if err != nil {
		return dto.BarangResponse{}, err
	}

	// convert to dto
	barangDTO := helper.ConvertToDTOBarangSingle(updatedBarang)
	return barangDTO, nil
}
