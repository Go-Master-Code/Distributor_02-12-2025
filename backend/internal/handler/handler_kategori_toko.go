package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerKategoriToko struct {
	service service.ServiceKategoriToko
}

// constructor
func NewHandlerKategoriToko(service service.ServiceKategoriToko) *handlerKategoriToko {
	return &handlerKategoriToko{service}
}

func (h *handlerKategoriToko) GetAllKategoriToko(c *gin.Context) {
	kategoriToko, err := h.service.GetAllKategoriToko()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, kategoriToko)
}

func (h *handlerKategoriToko) GetKategoriToko(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllKategoriToko(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		kategori, err := h.service.GetKategoriTokoById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, kategori)
	}

	if nama != "" {
		kategori, err := h.service.SearchKategoriToko(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, kategori)
	}
}

func (h *handlerKategoriToko) CreateKategoriToko(c *gin.Context) {
	var kategoriToko dto.CreateKategoriTokoRequest
	// parsing json body
	err := c.ShouldBindJSON(&kategoriToko)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newKategoriToko, err := h.service.CreateKategoriToko(kategoriToko)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newKategoriToko)
}
