package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceKota interface {
	GetAllKota() ([]dto.KotaResponse, error)
	GetKotaById(id int) (dto.KotaResponse, error)
	SearchKota(nama string) ([]dto.KotaResponse, error)
	CreateKota(kota dto.CreateKotaRequest) (dto.KotaResponse, error)
}

type serviceKota struct {
	repo repository.RepositoryKota
}

func NewServiceKota(repo repository.RepositoryKota) ServiceKota {
	return &serviceKota{repo}
}

func (s *serviceKota) GetAllKota() ([]dto.KotaResponse, error) {
	kota, err := s.repo.GetAllKota()
	if err != nil {
		return []dto.KotaResponse{}, err
	}

	// convert model ke dto
	kotaDTO := helper.ConvertToDTOKotaPlural(kota)
	return kotaDTO, nil
}

func (s *serviceKota) GetKotaById(id int) (dto.KotaResponse, error) {
	kota, err := s.repo.GetKotaById(id)
	if err != nil {
		return dto.KotaResponse{}, err
	}

	// convert model to dto
	kotaDTO := helper.ConvertToDTOKotaSingle(kota)
	return kotaDTO, nil
}

func (s *serviceKota) SearchKota(nama string) ([]dto.KotaResponse, error) {
	kota, err := s.repo.SearchKota(nama)
	if err != nil {
		return []dto.KotaResponse{}, err
	}

	// convert model to dto
	kotaDTO := helper.ConvertToDTOKotaPlural(kota)

	return kotaDTO, nil
}

func (s *serviceKota) CreateKota(kota dto.CreateKotaRequest) (dto.KotaResponse, error) {
	exist := s.repo.ExistByNama(kota.Nama)
	if exist {
		return dto.KotaResponse{}, errors.New("kota sudah ada")
	}

	// buat model dari dto
	req := model.Kota{
		Nama: kota.Nama,
	}

	newKota, err := s.repo.CreateKota(req)
	if err != nil {
		return dto.KotaResponse{}, err
	}

	// convert model to dto
	kotaDTO := helper.ConvertToDTOKotaSingle(newKota)
	return kotaDTO, nil
}
