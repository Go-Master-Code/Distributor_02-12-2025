package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerUkuran struct {
	service service.ServiceUkuran
}

// constructor
func NewHandlerUkuran(service service.ServiceUkuran) *handlerUkuran {
	return &handlerUkuran{service}
}

func (h *handlerUkuran) GetAllUkuran(c *gin.Context) {
	ukuran, err := h.service.GetAllUkuran()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, ukuran)
}

func (h *handlerUkuran) GetUkuran(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllUkuran(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetUkuranById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchUkuran(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerUkuran) CreateUkuran(c *gin.Context) {
	var ukuran dto.CreateUkuranRequest
	// parsing json body
	err := c.ShouldBindJSON(&ukuran)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newUkuran, err := h.service.CreateUkuran(ukuran)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newUkuran)
}

func (h *handlerUkuran) SearchUkuran(c *gin.Context) {
	query := c.Param("nama")
	ukuran, err := h.service.SearchUkuran(query)
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, ukuran)
}
