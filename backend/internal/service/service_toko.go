package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
)

type ServiceToko interface {
	GetAllToko() ([]dto.TokoResponse, error)
	GetTokoById(id int) (dto.TokoResponse, error)
	SearchToko(nama string) ([]dto.TokoResponse, error)
	CreateToko(brg dto.CreateTokoRequest) (dto.TokoResponse, error)
}

type serviceToko struct {
	repo repository.RepositoryToko
}

func NewServiceToko(repo repository.RepositoryToko) ServiceToko {
	return &serviceToko{repo}
}

func (s *serviceToko) GetAllToko() ([]dto.TokoResponse, error) {
	toko, err := s.repo.GetAllToko()
	if err != nil {
		return []dto.TokoResponse{}, err
	}

	// convert to dto
	tokoDTO := helper.ConvertToDTOTokoPlural(toko)
	return tokoDTO, nil
}

func (s *serviceToko) GetTokoById(id int) (dto.TokoResponse, error) {
	toko, err := s.repo.GetTokoById(id)
	if err != nil {
		return dto.TokoResponse{}, err
	}

	// convert to dto
	tokoDTO := helper.ConvertToDTOTokoSingle(toko)
	return tokoDTO, nil
}

func (s *serviceToko) SearchToko(nama string) ([]dto.TokoResponse, error) {
	toko, err := s.repo.SearchToko(nama)
	if err != nil {
		return []dto.TokoResponse{}, err
	}

	// convert model to dto
	tokoDTO := helper.ConvertToDTOTokoPlural(toko)

	return tokoDTO, nil
}

func (s *serviceToko) CreateToko(toko dto.CreateTokoRequest) (dto.TokoResponse, error) {
	// buat model toko
	req := model.Toko{
		Kode:           toko.Kode,
		Nama:           toko.Nama,
		KategoriTokoID: toko.KategoriTokoID,
		KotaID:         toko.KotaID,
		Alamat:         toko.Alamat,
		Disc1:          toko.Disc1,
		Disc2:          toko.Disc2,
		Disc3:          toko.Disc3,
		EkspedisiID:    toko.EkspedisiID,
		OngkirID:       toko.OngkirID,
	}

	newToko, err := s.repo.CreateToko(req)
	if err != nil {
		return dto.TokoResponse{}, err
	}

	// convert to dto
	tokoDTO := helper.ConvertToDTOTokoSingle(newToko)
	return tokoDTO, nil
}
