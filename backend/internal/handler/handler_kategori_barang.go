package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerkategoriBarang struct {
	service service.ServiceKategoriBarang
}

// constructor
func NewHandlerkategoriBarang(service service.ServiceKategoriBarang) *handlerkategoriBarang {
	return &handlerkategoriBarang{service}
}

func (h *handlerkategoriBarang) GetAllkategoriBarang(c *gin.Context) {
	kategori, err := h.service.GetAllKategoriBarang()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, kategori)
}

func (h *handlerkategoriBarang) GetKategoriBarang(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllkategoriBarang(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetKategoriBarangById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchKategoriBarang(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerkategoriBarang) CreatekategoriBarang(c *gin.Context) {
	var kategori dto.CreateKategoriBarangRequest
	// parsing json body
	err := c.ShouldBindJSON(&kategori)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newKategori, err := h.service.CreateKategoriBarang(kategori)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newKategori)
}
