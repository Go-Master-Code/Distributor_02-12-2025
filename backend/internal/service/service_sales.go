package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
)

type ServiceSales interface {
	GetAllSales() ([]dto.SalesResponse, error)
	GetSalesById(id int) (dto.SalesResponse, error)
	SearchSales(nama string) ([]dto.SalesResponse, error)
	CreateSales(sales dto.CreateSalesRequest) (dto.SalesResponse, error)
}

type serviceSales struct {
	repo repository.RepositorySales
}

func NewServiceSales(repo repository.RepositorySales) ServiceSales {
	return &serviceSales{repo}
}

func (s *serviceSales) GetAllSales() ([]dto.SalesResponse, error) {
	sales, err := s.repo.GetAllSales()
	if err != nil {
		return []dto.SalesResponse{}, err
	}

	// convert model to dto
	salesDTO := helper.ConvertToDTOSalesPlural(sales)

	return salesDTO, nil
}

func (s *serviceSales) GetSalesById(id int) (dto.SalesResponse, error) {
	sales, err := s.repo.GetSalesById(id)
	if err != nil {
		return dto.SalesResponse{}, err
	}

	// convert to model dto
	salesDTO := helper.ConvertToDTOSalesSingle(sales)
	return salesDTO, err
}

func (s *serviceSales) SearchSales(nama string) ([]dto.SalesResponse, error) {
	sales, err := s.repo.SearchSales(nama)
	if err != nil {
		return []dto.SalesResponse{}, err
	}

	// convert model to dto
	salesDTO := helper.ConvertToDTOSalesPlural(sales)
	return salesDTO, nil
}

func (s *serviceSales) CreateSales(sales dto.CreateSalesRequest) (dto.SalesResponse, error) {
	// buat var model dari dto create request
	req := model.Sales{
		Nama: sales.Nama,
	}

	newSales, err := s.repo.CreateSales(req)
	if err != nil {
		return dto.SalesResponse{}, err
	}

	// convert model to dto
	salesDTO := helper.ConvertToDTOSalesSingle(newSales)
	return salesDTO, nil
}
