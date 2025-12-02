package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerKota struct {
	service service.ServiceKota
}

// constructor
func NewHandlerKota(service service.ServiceKota) *handlerKota {
	return &handlerKota{service}
}

func (h *handlerKota) GetAllKota(c *gin.Context) {
	kota, err := h.service.GetAllKota()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, kota)
}

func (h *handlerKota) GetKota(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllKota(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetKotaById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchKota(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerKota) CreateKota(c *gin.Context) {
	var kota dto.CreateKotaRequest
	// parsing json body
	err := c.ShouldBindJSON(&kota)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newKota, err := h.service.CreateKota(kota)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newKota)
}
