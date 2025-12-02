package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceKategoriToko interface {
	GetAllKategoriToko() ([]dto.KategoriTokoResponse, error)
	GetKategoriTokoById(id int) (dto.KategoriTokoResponse, error)
	SearchKategoriToko(nama string) ([]dto.KategoriTokoResponse, error)
	CreateKategoriToko(kategori dto.CreateKategoriTokoRequest) (dto.KategoriTokoResponse, error)
}

type serviceKategoriToko struct {
	repo repository.RepositoryKategoriToko
}

func NewServiceKategoriToko(repo repository.RepositoryKategoriToko) ServiceKategoriToko {
	return &serviceKategoriToko{repo}
}

func (s *serviceKategoriToko) GetAllKategoriToko() ([]dto.KategoriTokoResponse, error) {
	kategori, err := s.repo.GetAllKategoriToko()
	if err != nil {
		return []dto.KategoriTokoResponse{}, err
	}

	// convert model ke dto
	kategoriDTO := helper.ConvertToDTOKategoriTokoPlural(kategori)
	return kategoriDTO, nil
}

func (s *serviceKategoriToko) GetKategoriTokoById(id int) (dto.KategoriTokoResponse, error) {
	kategori, err := s.repo.GetKategoriTokoById(id)
	if err != nil {
		return dto.KategoriTokoResponse{}, err
	}

	// convert model to dto
	kategoriDTO := helper.ConvertToDTOKategoriTokoSingle(kategori)
	return kategoriDTO, nil
}

func (s *serviceKategoriToko) SearchKategoriToko(nama string) ([]dto.KategoriTokoResponse, error) {
	kategori, err := s.repo.SearchKategoriToko(nama)
	if err != nil {
		return []dto.KategoriTokoResponse{}, err
	}

	// convert model to dto
	kategoriDTO := helper.ConvertToDTOKategoriTokoPlural(kategori)

	return kategoriDTO, nil
}

func (s *serviceKategoriToko) CreateKategoriToko(kategori dto.CreateKategoriTokoRequest) (dto.KategoriTokoResponse, error) {
	exist := s.repo.ExistByNama(kategori.Nama)
	if exist {
		return dto.KategoriTokoResponse{}, errors.New("kategori toko sudah ada")
	}

	// buat model dari dto
	req := model.KategoriToko{
		Nama: kategori.Nama,
	}

	newKategoriToko, err := s.repo.CreateKategoriToko(req)
	if err != nil {
		return dto.KategoriTokoResponse{}, err
	}

	// convert model to dto
	kategoriDTO := helper.ConvertToDTOKategoriTokoSingle(newKategoriToko)
	return kategoriDTO, nil
}
