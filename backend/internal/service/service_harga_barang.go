package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/repository"
)

type ServiceHargaBarang interface {
	GetHargaBarangById(id int) ([]dto.HargaBarangResponse, error)
}

type serviceHargaBarang struct {
	repo repository.RepositoryHargaBarang
}

func NewServiceHargaBarang(repo repository.RepositoryHargaBarang) ServiceHargaBarang {
	return &serviceHargaBarang{repo}
}

func (s *serviceHargaBarang) GetHargaBarangById(id int) ([]dto.HargaBarangResponse, error) {
	harga, err := s.repo.GetHargaBarangById(id)
	if err != nil {
		return []dto.HargaBarangResponse{}, err
	}

	// convert model to dto
	hargaDTO := helper.ConvertToDTOHargaBarangPlural(harga)
	return hargaDTO, nil
}
