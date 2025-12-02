package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerMerk struct {
	service service.ServiceMerk
}

// constructor
func NewHandlerMerk(service service.ServiceMerk) *handlerMerk {
	return &handlerMerk{service}
}

func (h *handlerMerk) GetAllMerk(c *gin.Context) {
	merk, err := h.service.GetAllMerk()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, merk)
}

func (h *handlerMerk) GetMerk(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllMerk(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetMerkById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchMerk(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerMerk) CreateMerk(c *gin.Context) {
	var merk dto.CreateMerkRequest
	// parsing json body
	err := c.ShouldBindJSON(&merk)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newMerk, err := h.service.CreateMerk(merk)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newMerk)
}
