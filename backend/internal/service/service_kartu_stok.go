package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/repository"
)

type ServiceKartuStok interface {
	// GetAllKartuStok() ([]dto.KartuStokResponse, error)
	GetAllKartuStok() ([]dto.KartuStokResponse, error)
	GetKartuStokByIdBarang(idBarang int) ([]dto.KartuStokResponse, error)
	GetStokByBarangIdAndSize(idBarang int, size int) (dto.KartuStokResponse, error)
	// SearchKartuStok(nama string) ([]dto.KartuStokResponse, error)
	// CreateKartuStok(brg dto.CreateKartuStokRequest) (dto.KartuStokResponse, error)
	// UpdateKartuStok(id int, req dto.UpdateKartuStokRequest) (dto.KartuStokResponse, error)
}

type serviceKartuStok struct {
	repo repository.RepositoryKartuStok
}

func NewServiceKartuStok(repo repository.RepositoryKartuStok) ServiceKartuStok {
	return &serviceKartuStok{repo}
}

func (s *serviceKartuStok) GetAllKartuStok() ([]dto.KartuStokResponse, error) {
	kartuStok, err := s.repo.GetAllKartuStok()
	if err != nil {
		return []dto.KartuStokResponse{}, err
	}

	// convert model to dto
	kartuStokDTO := helper.ConvertToDTOKartuStokPlural(kartuStok)

	return kartuStokDTO, nil
}

func (s *serviceKartuStok) GetKartuStokByIdBarang(id int) ([]dto.KartuStokResponse, error) {
	kartuStok, err := s.repo.GetKartuStokByIdBarang(id)
	if err != nil {
		return []dto.KartuStokResponse{}, err
	}

	// convert model to dto
	ksDTO := helper.ConvertToDTOKartuStokPlural(kartuStok)
	return ksDTO, nil
}

func (s *serviceKartuStok) GetStokByBarangIdAndSize(idBarang int, size int) (dto.KartuStokResponse, error) {
	kartuStok, err := s.repo.GetStokByBarangIdAndSize(idBarang, size)
	if err != nil {
		return dto.KartuStokResponse{}, err
	}

	// convert model to dto
	kartuStokDTO := helper.ConvertToDTOKartuStokSingle(kartuStok)

	return kartuStokDTO, nil
}

// func (s *serviceKartuStok) GetAllKartuStok() ([]dto.KartuStokResponse, error) {
// 	KartuStok, err := s.repo.GetAllKartuStok()
// 	if err != nil {
// 		return []dto.KartuStokResponse{}, err
// 	}

// 	// convert model ke dto
// 	barangDTO := helper.ConvertToDTOKartuStokPlural(KartuStok)
// 	return barangDTO, nil
// }

// func (s *serviceKartuStok) SearchKartuStok(nama string) ([]dto.KartuStokResponse, error) {
// 	KartuStok, err := s.repo.SearchKartuStok(nama)
// 	if err != nil {
// 		return []dto.KartuStokResponse{}, err
// 	}

// 	// convert model to dto
// 	barangDTO := helper.ConvertToDTOKartuStokPlural(KartuStok)

// 	return barangDTO, nil
// }

// func (s *serviceKartuStok) CreateKartuStok(brg dto.CreateKartuStokRequest) (dto.KartuStokResponse, error) {
// 	// buat model KartuStok
// 	req := model.KartuStok{
// 		Kode:                brg.Kode,
// 		MerkID:              brg.MerkID,
// 		ArtikelID:           brg.ArtikelID,
// 		WarnaID:             brg.WarnaID,
// 		KategoriKartuStokID: brg.KategoriKartuStokID,
// 		JenisKartuStokID:    brg.JenisKartuStokID,
// 		UkuranID:            brg.UkuranID,
// 		Stok:                brg.Stok,
// 	}

// 	newKartuStok, err := s.repo.CreateKartuStok(req)
// 	if err != nil {
// 		return dto.KartuStokResponse{}, err
// 	}

// 	// convert to dto
// 	barangDTO := helper.ConvertToDTOKartuStokSingle(newKartuStok)
// 	return barangDTO, nil
// }

// func (s *serviceKartuStok) UpdateKartuStok(id int, req dto.UpdateKartuStokRequest) (dto.KartuStokResponse, error) {
// 	// convert dto.Update ke map
// 	var updateMap = map[string]any{}

// 	if req.Kode != nil {
// 		updateMap["kode"] = *req.Kode // dereference
// 	}

// 	if req.MerkID != nil {
// 		updateMap["merk_id"] = *req.MerkID // dereference
// 	}

// 	if req.ArtikelID != nil {
// 		updateMap["artikel_id"] = *req.ArtikelID // dereference
// 	}

// 	if req.WarnaID != nil {
// 		updateMap["warna_id"] = *req.WarnaID // dereference
// 	}

// 	if req.KategoriKartuStokID != nil {
// 		updateMap["kategori_barang_id"] = *req.KategoriKartuStokID // dereference
// 	}

// 	if req.JenisKartuStokID != nil {
// 		updateMap["jenis_barang_id"] = *req.JenisKartuStokID // dereference
// 	}

// 	if req.UkuranID != nil {
// 		updateMap["ukuran_id"] = *req.UkuranID // dereference
// 	}

// 	updatedKartuStok, err := s.repo.UpdateSupplier(id, updateMap)
// 	if err != nil {
// 		return dto.KartuStokResponse{}, err
// 	}

// 	// convert to dto
// 	barangDTO := helper.ConvertToDTOKartuStokSingle(updatedKartuStok)
// 	return barangDTO, nil
// }
