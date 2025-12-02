package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/repository"
)

type ServicePenjualanDetail interface {
	GetPenjualanDetailById(id int) ([]dto.PenjualanDetailResponse, error)
	// CreatePenjualan(penjualan dto.CreatePenjualanRequest) (dto.PenjualanResponse, error)
}

type servicePenjualanDetail struct {
	repo repository.RepositoryPenjualanDetail
}

func NewServicePenjualanDetail(repo repository.RepositoryPenjualanDetail) ServicePenjualanDetail {
	return &servicePenjualanDetail{repo}
}

func (s *servicePenjualanDetail) GetPenjualanDetailById(id int) ([]dto.PenjualanDetailResponse, error) {
	dj, err := s.repo.GetPenjualanDetailById(id)
	if err != nil {
		return []dto.PenjualanDetailResponse{}, err
	}

	// convert model to dto
	djDTO := helper.ConvertToDTOPenjualanDetailPlural(dj)

	return djDTO, nil
}
