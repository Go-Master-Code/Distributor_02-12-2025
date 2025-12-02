package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceJenisBarang interface {
	GetAllJenisBarang() ([]dto.JenisBarangResponse, error)
	GetJenisBarangById(id int) (dto.JenisBarangResponse, error)
	SearchJenisBarang(nama string) ([]dto.JenisBarangResponse, error)
	CreateJenisBarang(jenisBarang dto.CreateJenisBarangRequest) (dto.JenisBarangResponse, error)
}

type serviceJenisBarang struct {
	repo repository.RepositoryJenisBarang
}

func NewServiceJenisBarang(repo repository.RepositoryJenisBarang) ServiceJenisBarang {
	return &serviceJenisBarang{repo}
}

func (s *serviceJenisBarang) GetAllJenisBarang() ([]dto.JenisBarangResponse, error) {
	jenisBarang, err := s.repo.GetAllJenisBarang()
	if err != nil {
		return []dto.JenisBarangResponse{}, err
	}

	// convert model ke dto
	jenisBarangDTO := helper.ConvertToDTOJenisBarangPlural(jenisBarang)
	return jenisBarangDTO, nil
}

func (s *serviceJenisBarang) GetJenisBarangById(id int) (dto.JenisBarangResponse, error) {
	jenis, err := s.repo.GetJenisBarangById(id)
	if err != nil {
		return dto.JenisBarangResponse{}, err
	}

	// convert model to dto
	jenisDTO := helper.ConvertToDTOJenisBarangSingle(jenis)
	return jenisDTO, nil
}

func (s *serviceJenisBarang) SearchJenisBarang(nama string) ([]dto.JenisBarangResponse, error) {
	jb, err := s.repo.SearchJenisBarang(nama)
	if err != nil {
		return []dto.JenisBarangResponse{}, err
	}

	// convert model to dto
	jbDTO := helper.ConvertToDTOJenisBarangPlural(jb)
	return jbDTO, nil
}

func (s *serviceJenisBarang) CreateJenisBarang(jenisBarang dto.CreateJenisBarangRequest) (dto.JenisBarangResponse, error) {
	exist := s.repo.ExistByNama(jenisBarang.Nama)
	if exist {
		return dto.JenisBarangResponse{}, errors.New("jenis barang sudah ada")
	}

	// buat model dari dto
	req := model.JenisBarang{
		Nama: jenisBarang.Nama,
	}

	newJenisBarang, err := s.repo.CreateJenisBarang(req)
	if err != nil {
		return dto.JenisBarangResponse{}, err
	}

	// convert model to dto
	jenisBarangDTO := helper.ConvertToDTOJenisBarangSingle(newJenisBarang)
	return jenisBarangDTO, nil
}
