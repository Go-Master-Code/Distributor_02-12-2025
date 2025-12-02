package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceArtikel interface {
	GetAllArtikel() ([]dto.ArtikelResponse, error)
	GetArtikelByID(id int) (dto.ArtikelResponse, error)
	SearchArtikel(nama string) ([]dto.ArtikelResponse, error)
	CreateArtikel(artikel dto.CreateArtikelRequest) (dto.ArtikelResponse, error)
	UpdateArtikel(id int, req dto.UpdateArtikelRequest) (dto.ArtikelResponse, error)
}

type serviceArtikel struct {
	repo repository.RepositoryArtikel
}

func NewServiceArtikel(repo repository.RepositoryArtikel) ServiceArtikel {
	return &serviceArtikel{repo}
}

func (s *serviceArtikel) GetAllArtikel() ([]dto.ArtikelResponse, error) {
	artikel, err := s.repo.GetAllArtikel()
	if err != nil {
		return []dto.ArtikelResponse{}, err
	}

	// convert model ke dto
	artikelDTO := helper.ConvertToDTOArtikelPlural(artikel)
	return artikelDTO, nil
}

func (s *serviceArtikel) GetArtikelByID(id int) (dto.ArtikelResponse, error) {
	artikel, err := s.repo.GetArtikelByID(id)
	if err != nil {
		return dto.ArtikelResponse{}, err
	}

	// convert model to dto
	artikelDTO := helper.ConvertToDTOArtikelSingle(artikel)
	return artikelDTO, nil
}

func (s *serviceArtikel) SearchArtikel(nama string) ([]dto.ArtikelResponse, error) {
	artikel, err := s.repo.SearchArtikel(nama)
	if err != nil {
		return []dto.ArtikelResponse{}, err
	}

	// convert model to dto
	artikelDTO := helper.ConvertToDTOArtikelPlural(artikel)
	return artikelDTO, nil
}

func (s *serviceArtikel) CreateArtikel(artikel dto.CreateArtikelRequest) (dto.ArtikelResponse, error) {
	exist := s.repo.ExistByNama(artikel.Nama)
	if exist {
		return dto.ArtikelResponse{}, errors.New("artikel sudah ada")
	}

	// buat model dari dto
	req := model.Artikel{
		Nama: artikel.Nama,
	}

	newArtikel, err := s.repo.CreateArtikel(req)
	if err != nil {
		return dto.ArtikelResponse{}, err
	}

	// convert model to dto
	artikelDTO := helper.ConvertToDTOArtikelSingle(newArtikel)
	return artikelDTO, nil
}

func (s *serviceArtikel) UpdateArtikel(id int, req dto.UpdateArtikelRequest) (dto.ArtikelResponse, error) {
	var updateMap = map[string]any{}
	if req.Nama != nil {
		updateMap["nama"] = *req.Nama // dereference
	}

	updatedArtikel, err := s.repo.UpdateArtikel(id, updateMap)
	if err != nil {
		return dto.ArtikelResponse{}, err
	}

	// convert to dto
	artikelDTO := helper.ConvertToDTOArtikelSingle(updatedArtikel)

	return artikelDTO, nil
}
