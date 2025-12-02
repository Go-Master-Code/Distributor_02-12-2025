package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerWarna struct {
	service service.ServiceWarna
}

// constructor
func NewHandlerWarna(service service.ServiceWarna) *handlerWarna {
	return &handlerWarna{service}
}

func (h *handlerWarna) GetAllWarna(c *gin.Context) {
	warna, err := h.service.GetAllWarna()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, warna)
}

func (h *handlerWarna) GetWarna(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllWarna(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetWarnaById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchWarna(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerWarna) CreateWarna(c *gin.Context) {
	var warna dto.CreateWarnaRequest
	// parsing json body
	err := c.ShouldBindJSON(&warna)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newWarna, err := h.service.CreateWarna(warna)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newWarna)
}
