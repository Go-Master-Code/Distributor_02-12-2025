package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"time"
)

type ServiceHargaBarang interface {
	GetHargaBarangById(id int) ([]dto.HargaBarangResponse, error)
	CreateHargaBarang(hb dto.CreateHargaBarangRequest) (dto.HargaBarangResponse, error)
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

func (s *serviceHargaBarang) CreateHargaBarang(hb dto.CreateHargaBarangRequest) (dto.HargaBarangResponse, error) {
	mulaiBerlaku, err := time.Parse("2006-01-02", hb.MulaiBerlaku)
	if err != nil {
		return dto.HargaBarangResponse{}, err
	}

	req := model.HargaBarang{
		BarangID:     hb.BarangID,
		Harga:        hb.Harga,
		MulaiBerlaku: mulaiBerlaku,
	}

	hargaBrg, err := s.repo.CreateHargaBarang(req)

	if err != nil {
		return dto.HargaBarangResponse{}, err
	}

	// convert model to dto
	hbDTO := helper.ConvertToHargaBarangSingle(hargaBrg)

	return hbDTO, nil
}
