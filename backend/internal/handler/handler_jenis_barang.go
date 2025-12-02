package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerJenisBarang struct {
	service service.ServiceJenisBarang
}

// constructor
func NewHandlerJenisBarang(service service.ServiceJenisBarang) *handlerJenisBarang {
	return &handlerJenisBarang{service}
}

func (h *handlerJenisBarang) GetAllJenisBarang(c *gin.Context) {
	JenisBarang, err := h.service.GetAllJenisBarang()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, JenisBarang)
}

func (h *handlerJenisBarang) GetJenisBarang(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllJenisBarang(c)
	}

	if id != "" {
		// cari data by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		merk, err := h.service.GetJenisBarangById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}

	if nama != "" {
		merk, err := h.service.SearchJenisBarang(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, merk)
	}
}

func (h *handlerJenisBarang) CreateJenisBarang(c *gin.Context) {
	var JenisBarang dto.CreateJenisBarangRequest
	// parsing json body
	err := c.ShouldBindJSON(&JenisBarang)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newJenisBarang, err := h.service.CreateJenisBarang(JenisBarang)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newJenisBarang)
}

func (h *handlerJenisBarang) SearchJenisBarang(c *gin.Context) {
	query := c.Param("nama")
	jb, err := h.service.SearchJenisBarang(query)
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, jb)
}
