package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceUkuran interface {
	GetAllUkuran() ([]dto.UkuranResponse, error)
	GetUkuranById(id int) (dto.UkuranResponse, error)
	SearchUkuran(nama string) ([]dto.UkuranResponse, error)
	CreateUkuran(ukuran dto.CreateUkuranRequest) (dto.UkuranResponse, error)
}

type serviceUkuran struct {
	repo repository.RepositoryUkuran
}

func NewServiceUkuran(repo repository.RepositoryUkuran) ServiceUkuran {
	return &serviceUkuran{repo}
}

func (s *serviceUkuran) GetAllUkuran() ([]dto.UkuranResponse, error) {
	ukuran, err := s.repo.GetAllUkuran()
	if err != nil {
		return []dto.UkuranResponse{}, err
	}

	// convert model ke dto
	ukuranDTO := helper.ConvertToDTOUkuranPlural(ukuran)
	return ukuranDTO, nil
}

func (s *serviceUkuran) GetUkuranById(id int) (dto.UkuranResponse, error) {
	ukuran, err := s.repo.GetUkuranById(id)
	if err != nil {
		return dto.UkuranResponse{}, err
	}

	// convert model to dto
	ukuranDTO := helper.ConvertToDTOUkuranSingle(ukuran)
	return ukuranDTO, nil
}

func (s *serviceUkuran) CreateUkuran(ukuran dto.CreateUkuranRequest) (dto.UkuranResponse, error) {
	exist := s.repo.ExistByNama(ukuran.Nama)
	if exist {
		return dto.UkuranResponse{}, errors.New("ukuran sudah ada")
	}

	// buat model dari dto
	req := model.Ukuran{
		Nama: ukuran.Nama,
	}

	newUkuran, err := s.repo.CreateUkuran(req)
	if err != nil {
		return dto.UkuranResponse{}, err
	}

	// convert model to dto
	ukuranDTO := helper.ConvertToDTOUkuranSingle(newUkuran)
	return ukuranDTO, nil
}

func (s *serviceUkuran) SearchUkuran(nama string) ([]dto.UkuranResponse, error) {
	ukuran, err := s.repo.SearchUkuran(nama)
	if err != nil {
		return []dto.UkuranResponse{}, err
	}

	// convert model to dto
	ukuranDTO := helper.ConvertToDTOUkuranPlural(ukuran)
	return ukuranDTO, nil
}
