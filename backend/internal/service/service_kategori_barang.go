package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceKategoriBarang interface {
	GetAllKategoriBarang() ([]dto.KategoriBarangResponse, error)
	GetKategoriBarangById(id int) (dto.KategoriBarangResponse, error)
	SearchKategoriBarang(nama string) ([]dto.KategoriBarangResponse, error)
	CreateKategoriBarang(kategori dto.CreateKategoriBarangRequest) (dto.KategoriBarangResponse, error)
}

type serviceKategoriBarang struct {
	repo repository.RepositoryKategoriBarang
}

func NewServiceKategoriBarang(repo repository.RepositoryKategoriBarang) ServiceKategoriBarang {
	return &serviceKategoriBarang{repo}
}

func (s *serviceKategoriBarang) GetAllKategoriBarang() ([]dto.KategoriBarangResponse, error) {
	kategori, err := s.repo.GetAllKategoriBarang()
	if err != nil {
		return []dto.KategoriBarangResponse{}, err
	}

	// convert model ke dto
	kategoriDTO := helper.ConvertToDTOKategoriBarangPlural(kategori)
	return kategoriDTO, nil
}

func (s *serviceKategoriBarang) GetKategoriBarangById(id int) (dto.KategoriBarangResponse, error) {
	kategori, err := s.repo.GetKategoriBarangById(id)
	if err != nil {
		return dto.KategoriBarangResponse{}, err
	}

	// convert model to dto
	kategoriDTO := helper.ConvertToDTOKategoriBarangSingle(kategori)
	return kategoriDTO, nil
}

func (s *serviceKategoriBarang) SearchKategoriBarang(nama string) ([]dto.KategoriBarangResponse, error) {
	kb, err := s.repo.SearchKategoriBarang(nama)
	if err != nil {
		return []dto.KategoriBarangResponse{}, err
	}

	// convert model to dto
	kbDTO := helper.ConvertToDTOKategoriBarangPlural(kb)
	return kbDTO, nil
}

func (s *serviceKategoriBarang) CreateKategoriBarang(kategori dto.CreateKategoriBarangRequest) (dto.KategoriBarangResponse, error) {
	exist := s.repo.ExistByNama(kategori.Nama)
	if exist {
		return dto.KategoriBarangResponse{}, errors.New("kategori barang sudah ada")
	}

	// buat model dari dto
	req := model.KategoriBarang{
		Nama:  kategori.Nama,
		Huruf: kategori.Huruf,
	}

	newKategoriBarang, err := s.repo.CreateKategoriBarang(req)
	if err != nil {
		return dto.KategoriBarangResponse{}, err
	}

	// convert model to dto
	kategoriDTO := helper.ConvertToDTOKategoriBarangSingle(newKategoriBarang)
	return kategoriDTO, nil
}
