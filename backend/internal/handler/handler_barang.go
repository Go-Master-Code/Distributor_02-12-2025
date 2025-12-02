package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerBarang struct {
	service service.ServiceBarang
}

// constructor
func NewHandlerBarang(service service.ServiceBarang) *handlerBarang {
	return &handlerBarang{service}
}

func (h *handlerBarang) GetAllBarang(c *gin.Context) {
	barang, err := h.service.GetAllBarang()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, barang)
}

func (h *handlerBarang) CreateBarang(c *gin.Context) {
	// parsing request body
	var barang dto.CreateBarangRequest
	err := c.ShouldBindJSON(&barang)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newBarang, err := h.service.CreateBarang(barang)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newBarang)
}

func (h *handlerBarang) GetBarangById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	barang, err := h.service.GetBarangById(idInt)
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, barang)
}

func (h *handlerBarang) GetBarang(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllBarang(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetBarangById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchBarang(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerBarang) UpdateBarang(c *gin.Context) {
	// parsing json
	var req dto.UpdateBarangRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	barangDTO, err := h.service.UpdateBarang(idInt, req)

	if err != nil {
		helper.ErrorUpdateData(c, err)
		return
	}

	helper.StatusSuksesUpdateData(c, barangDTO)
}
