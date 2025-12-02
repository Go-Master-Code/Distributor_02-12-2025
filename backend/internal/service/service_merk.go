package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceMerk interface {
	GetAllMerk() ([]dto.MerkResponse, error)
	GetMerkById(id int) (dto.MerkResponse, error)
	SearchMerk(nama string) ([]dto.MerkResponse, error)
	CreateMerk(merk dto.CreateMerkRequest) (dto.MerkResponse, error)
}

type serviceMerk struct {
	repo repository.RepositoryMerk
}

func NewServiceMerk(repo repository.RepositoryMerk) ServiceMerk {
	return &serviceMerk{repo}
}

func (s *serviceMerk) GetAllMerk() ([]dto.MerkResponse, error) {
	merk, err := s.repo.GetAllMerk()
	if err != nil {
		return []dto.MerkResponse{}, err
	}

	// convert model ke dto
	merkDTO := helper.ConvertToDTOMerkPlural(merk)
	return merkDTO, nil
}

func (s *serviceMerk) GetMerkById(id int) (dto.MerkResponse, error) {
	merk, err := s.repo.GetMerkById(id)
	if err != nil {
		return dto.MerkResponse{}, err
	}

	// convert model to dto
	merkDTO := helper.ConvertToDTOMerkSingle(merk)
	return merkDTO, nil
}

func (s *serviceMerk) CreateMerk(merk dto.CreateMerkRequest) (dto.MerkResponse, error) {
	exist := s.repo.ExistByNama(merk.Nama)
	if exist {
		return dto.MerkResponse{}, errors.New("merk sudah ada")
	}

	// buat model dari dto
	req := model.Merk{
		Nama: merk.Nama,
	}

	newMerk, err := s.repo.CreateMerk(req)
	if err != nil {
		return dto.MerkResponse{}, err
	}

	// convert model to dto
	merkDTO := helper.ConvertToDTOMerkSingle(newMerk)
	return merkDTO, nil
}

func (s *serviceMerk) SearchMerk(nama string) ([]dto.MerkResponse, error) { // Lazy load
	merk, err := s.repo.SearchMerk(nama)
	if err != nil {
		return []dto.MerkResponse{}, err
	}

	// convert model to dto
	merkDTO := helper.ConvertToDTOMerkPlural(merk)
	return merkDTO, nil
}
