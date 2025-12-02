package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceEkspedisi interface {
	GetAllEkspedisi() ([]dto.EkspedisiResponse, error)
	GetEkspedisiById(id int) (dto.EkspedisiResponse, error)
	SearchEkspedisi(nama string) ([]dto.EkspedisiResponse, error)
	CreateEkspedisi(ekspedisi dto.CreateEkspedisiRequest) (dto.EkspedisiResponse, error)
}

type serviceEkspedisi struct {
	repo repository.RepositoryEkspedisi
}

func NewServiceEkspedisi(repo repository.RepositoryEkspedisi) ServiceEkspedisi {
	return &serviceEkspedisi{repo}
}

func (s *serviceEkspedisi) GetAllEkspedisi() ([]dto.EkspedisiResponse, error) {
	ekspedisi, err := s.repo.GetAllEkspedisi()
	if err != nil {
		return []dto.EkspedisiResponse{}, err
	}

	// convert model ke dto
	ekspedisiDTO := helper.ConvertToDTOEkspedisiPlural(ekspedisi)
	return ekspedisiDTO, nil
}

func (s *serviceEkspedisi) GetEkspedisiById(id int) (dto.EkspedisiResponse, error) {
	ekspedisi, err := s.repo.GetEkspedisiById(id)
	if err != nil {
		return dto.EkspedisiResponse{}, err
	}

	// convert model to dto
	ekspedisiDTO := helper.ConvertToDTOEkspedisiSingle(ekspedisi)
	return ekspedisiDTO, nil
}

func (s *serviceEkspedisi) SearchEkspedisi(nama string) ([]dto.EkspedisiResponse, error) {
	ekspedisi, err := s.repo.SearchEkspedisi(nama)
	if err != nil {
		return []dto.EkspedisiResponse{}, err
	}

	// convert model to dto
	ekspedisiDTO := helper.ConvertToDTOEkspedisiPlural(ekspedisi)

	return ekspedisiDTO, nil
}

func (s *serviceEkspedisi) CreateEkspedisi(ekspedisi dto.CreateEkspedisiRequest) (dto.EkspedisiResponse, error) {
	exist := s.repo.ExistByNama(ekspedisi.Nama)
	if exist {
		return dto.EkspedisiResponse{}, errors.New("ekspedisi sudah ada")
	}

	// buat model dari dto
	req := model.Ekspedisi{
		Nama: ekspedisi.Nama,
	}

	newEkspedisi, err := s.repo.CreateEkspedisi(req)
	if err != nil {
		return dto.EkspedisiResponse{}, err
	}

	// convert model to dto
	ekspedisiDTO := helper.ConvertToDTOEkspedisiSingle(newEkspedisi)
	return ekspedisiDTO, nil
}
