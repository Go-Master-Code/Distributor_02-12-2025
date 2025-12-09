package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/repository"
)

type ServiceArea interface {
	GetAllArea() ([]dto.AreaResponse, error)
	SearchArea(nama string) ([]dto.AreaResponse, error)
}

type serviceArea struct {
	repo repository.RepositoryArea
}

// constructor
func NewServiceArea(repo repository.RepositoryArea) ServiceArea {
	return &serviceArea{repo}
}

func (s *serviceArea) GetAllArea() ([]dto.AreaResponse, error) {
	area, err := s.repo.GetAllArea()
	if err != nil {
		return []dto.AreaResponse{}, err
	}

	// convert model to dto
	areaDTO := helper.ConvertToDTOAreaPlural(area)

	return areaDTO, nil
}

func (s *serviceArea) SearchArea(nama string) ([]dto.AreaResponse, error) {
	area, err := s.repo.SearchArea(nama)
	if err != nil {
		return []dto.AreaResponse{}, err
	}

	// convert model to dto
	areaDTO := helper.ConvertToDTOAreaPlural(area)

	return areaDTO, nil
}
