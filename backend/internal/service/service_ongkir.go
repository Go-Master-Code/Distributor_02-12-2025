package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceOngkir interface {
	GetAllOngkir() ([]dto.OngkirResponse, error)
	GetOngkirById(id int) (dto.OngkirResponse, error)
	SearchOngkir(nama string) ([]dto.OngkirResponse, error)
	CreateOngkir(ongkir dto.CreateOngkirRequest) (dto.OngkirResponse, error)
}

type serviceOngkir struct {
	repo repository.RepositoryOngkir
}

func NewServiceOngkir(repo repository.RepositoryOngkir) ServiceOngkir {
	return &serviceOngkir{repo}
}

func (s *serviceOngkir) GetAllOngkir() ([]dto.OngkirResponse, error) {
	ongkir, err := s.repo.GetAllOngkir()
	if err != nil {
		return []dto.OngkirResponse{}, err
	}

	// convert model ke dto
	ongkirDTO := helper.ConvertToDTOOngkirPlural(ongkir)
	return ongkirDTO, nil
}

func (s *serviceOngkir) GetOngkirById(id int) (dto.OngkirResponse, error) {
	ongkir, err := s.repo.GetOngkirById(id)
	if err != nil {
		return dto.OngkirResponse{}, err
	}

	// convert model to dto
	ongkirDTO := helper.ConvertToDTOOngkirSingle(ongkir)
	return ongkirDTO, nil
}

func (s *serviceOngkir) SearchOngkir(nama string) ([]dto.OngkirResponse, error) {
	ongkir, err := s.repo.SearchOngkir(nama)
	if err != nil {
		return []dto.OngkirResponse{}, err
	}

	// convert model to dto
	ongkirDTO := helper.ConvertToDTOOngkirPlural(ongkir)

	return ongkirDTO, nil
}

func (s *serviceOngkir) CreateOngkir(ongkir dto.CreateOngkirRequest) (dto.OngkirResponse, error) {
	exist := s.repo.ExistByNama(ongkir.Nama)
	if exist {
		return dto.OngkirResponse{}, errors.New("nama penanggung ongkir sudah ada")
	}

	// buat model dari dto
	req := model.Ongkir{
		Nama: ongkir.Nama,
	}

	newOngkir, err := s.repo.CreateOngkir(req)
	if err != nil {
		return dto.OngkirResponse{}, err
	}

	// convert model to dto
	ongkirDTO := helper.ConvertToDTOOngkirSingle(newOngkir)
	return ongkirDTO, nil
}
