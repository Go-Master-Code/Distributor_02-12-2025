package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceWarna interface {
	GetAllWarna() ([]dto.WarnaResponse, error)
	GetWarnaById(id int) (dto.WarnaResponse, error)
	SearchWarna(nama string) ([]dto.WarnaResponse, error)
	CreateWarna(warna dto.CreateWarnaRequest) (dto.WarnaResponse, error)
}

type serviceWarna struct {
	repo repository.RepositoryWarna
}

func NewServiceWarna(repo repository.RepositoryWarna) ServiceWarna {
	return &serviceWarna{repo}
}

func (s *serviceWarna) GetAllWarna() ([]dto.WarnaResponse, error) {
	warna, err := s.repo.GetAllWarna()
	if err != nil {
		return []dto.WarnaResponse{}, err
	}

	// convert model ke dto
	warnaDTO := helper.ConvertToDTOWarnaPlural(warna)
	return warnaDTO, nil
}

func (s *serviceWarna) GetWarnaById(id int) (dto.WarnaResponse, error) {
	warna, err := s.repo.GetWarnaById(id)
	if err != nil {
		return dto.WarnaResponse{}, err
	}

	// convert model to dto
	warnaDTO := helper.ConvertToDTOWarnaSingle(warna)
	return warnaDTO, nil
}

func (s *serviceWarna) SearchWarna(nama string) ([]dto.WarnaResponse, error) {
	warna, err := s.repo.SearchWarna(nama)
	if err != nil {
		return []dto.WarnaResponse{}, err
	}

	// convert model to dto
	warnaDTO := helper.ConvertToDTOWarnaPlural(warna)

	return warnaDTO, nil
}

func (s *serviceWarna) CreateWarna(warna dto.CreateWarnaRequest) (dto.WarnaResponse, error) {
	exist := s.repo.ExistByNama(warna.Nama)
	if exist {
		return dto.WarnaResponse{}, errors.New("warna sudah ada")
	}

	// buat model dari dto
	req := model.Warna{
		Nama: warna.Nama,
	}

	newWarna, err := s.repo.CreateWarna(req)
	if err != nil {
		return dto.WarnaResponse{}, err
	}

	// convert model to dto
	warnaDTO := helper.ConvertToDTOWarnaSingle(newWarna)
	return warnaDTO, nil
}
