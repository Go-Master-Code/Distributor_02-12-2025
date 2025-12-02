package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerEkspedisi struct {
	service service.ServiceEkspedisi
}

// constructor
func NewHandlerEkspedisi(service service.ServiceEkspedisi) *handlerEkspedisi {
	return &handlerEkspedisi{service}
}

func (h *handlerEkspedisi) GetAllEkspedisi(c *gin.Context) {
	ekspedisi, err := h.service.GetAllEkspedisi()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, ekspedisi)
}

func (h *handlerEkspedisi) GetEkspedisi(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllEkspedisi(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetEkspedisiById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchEkspedisi(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerEkspedisi) CreateEkspedisi(c *gin.Context) {
	var ekspedisi dto.CreateEkspedisiRequest
	// parsing json body
	err := c.ShouldBindJSON(&ekspedisi)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newEkspedisi, err := h.service.CreateEkspedisi(ekspedisi)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newEkspedisi)
}
