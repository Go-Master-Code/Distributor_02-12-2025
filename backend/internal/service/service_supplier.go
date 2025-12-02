package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"errors"
)

type ServiceSupplier interface {
	GetAllSupplier() ([]dto.SupplierResponse, error)
	GetSupplierByID(id int) (dto.SupplierResponse, error)
	CreateSupplier(supplier dto.CreateSupplierRequest) (dto.SupplierResponse, error)
	UpdateSupplier(id int, req dto.UpdateSupplierRequest) (dto.SupplierResponse, error)
	DeleteSupplier(id int) (dto.SupplierResponse, error)
}

type serviceSupplier struct {
	repo repository.RepositorySupplier
}

func NewServiceSupplier(repo repository.RepositorySupplier) ServiceSupplier {
	return &serviceSupplier{repo}
}

func (s *serviceSupplier) GetAllSupplier() ([]dto.SupplierResponse, error) {
	supplier, err := s.repo.GetAllSupplier()
	if err != nil {
		return []dto.SupplierResponse{}, err
	}

	// convert model ke dto
	supplierDTO := helper.ConvertToDTOSupplierPlural(supplier)
	return supplierDTO, nil
}

func (s *serviceSupplier) GetSupplierByID(id int) (dto.SupplierResponse, error) {
	supplier, err := s.repo.GetSupplierByID(id)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert model to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(supplier)
	return supplierDTO, nil
}

func (s *serviceSupplier) CreateSupplier(supplier dto.CreateSupplierRequest) (dto.SupplierResponse, error) {
	exist := s.repo.ExistByNama(supplier.Nama)
	if exist {
		return dto.SupplierResponse{}, errors.New("supplier sudah ada")
	}

	// buat model dari dto
	req := model.Supplier{
		Nama:   supplier.Nama,
		Alamat: supplier.Alamat,
		Kontak: supplier.Kontak,
	}

	newSupplier, err := s.repo.CreateSupplier(req)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert model to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(newSupplier)
	return supplierDTO, nil
}

func (s *serviceSupplier) UpdateSupplier(id int, req dto.UpdateSupplierRequest) (dto.SupplierResponse, error) {
	var updateMap = map[string]any{}

	// pengecekan isi request
	if req.Nama != nil {
		updateMap["nama"] = *req.Nama // dereference
	}
	if req.Alamat != nil {
		updateMap["alamat"] = *req.Alamat // dereference
	}
	if req.Kontak != nil {
		updateMap["kontak"] = *req.Kontak // dereference
	}

	updatedSupplier, err := s.repo.UpdateSupplier(id, updateMap)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(updatedSupplier)

	return supplierDTO, nil
}

func (s *serviceSupplier) DeleteSupplier(id int) (dto.SupplierResponse, error) {
	supplier, err := s.repo.DeleteSupplier(id)
	if err != nil {
		return dto.SupplierResponse{}, err
	}

	// convert to dto
	supplierDTO := helper.ConvertToDTOSupplierSingle(supplier)
	return supplierDTO, nil
}
