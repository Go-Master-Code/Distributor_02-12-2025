package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerArtikel struct {
	service service.ServiceArtikel
}

// constructor
func NewHandlerArtikel(service service.ServiceArtikel) *handlerArtikel {
	return &handlerArtikel{service}
}

func (h *handlerArtikel) GetAllArtikel(c *gin.Context) {
	artikel, err := h.service.GetAllArtikel()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, artikel)
}

func (h *handlerArtikel) GetArtikelByID(c *gin.Context) {
	// ambil param id
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}
	artikel, err := h.service.GetArtikelByID(idInt)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	helper.StatusSuksesGetData(c, artikel)
}

func (h *handlerArtikel) GetArtikel(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllArtikel(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetArtikelByID(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchArtikel(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerArtikel) CreateArtikel(c *gin.Context) {
	var artikel dto.CreateArtikelRequest
	// parsing json body
	err := c.ShouldBindJSON(&artikel)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newArtikel, err := h.service.CreateArtikel(artikel)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newArtikel)
}

func (h *handlerArtikel) UpdateArtikel(c *gin.Context) {
	var artikel dto.UpdateArtikelRequest
	err := c.ShouldBindJSON(&artikel)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// get param
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	artikelDTO, err := h.service.UpdateArtikel(idInt, artikel)
	if err != nil {
		helper.ErrorUpdateData(c, err)
		return
	}

	helper.StatusSuksesUpdateData(c, artikelDTO)
}
