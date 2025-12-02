package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerOngkir struct {
	service service.ServiceOngkir
}

// constructor
func NewHandlerOngkir(service service.ServiceOngkir) *handlerOngkir {
	return &handlerOngkir{service}
}

func (h *handlerOngkir) GetAllOngkir(c *gin.Context) {
	ongkir, err := h.service.GetAllOngkir()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, ongkir)
}

func (h *handlerOngkir) GetOngkir(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllOngkir(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetOngkirById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchOngkir(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerOngkir) CreateOngkir(c *gin.Context) {
	var ongkir dto.CreateOngkirRequest
	// parsing json body
	err := c.ShouldBindJSON(&ongkir)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newOngkir, err := h.service.CreateOngkir(ongkir)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newOngkir)
}
