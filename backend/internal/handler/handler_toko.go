package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerToko struct {
	service service.ServiceToko
}

func NewHandlerToko(service service.ServiceToko) *handlerToko {
	return &handlerToko{service}
}

func (h *handlerToko) GetAllToko(c *gin.Context) {
	toko, err := h.service.GetAllToko()
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, toko)
}

func (h *handlerToko) GetTokoById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	toko, err := h.service.GetTokoById(idInt)
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, toko)
}

func (h *handlerToko) GetToko(c *gin.Context) {
	id := c.Query("id")
	nama := c.Query("nama")

	if id == "" && nama == "" {
		h.GetAllToko(c)
	}

	if id != "" {
		// cari data by ID, hanya tampilkan data toko by ID
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		toko, err := h.service.GetTokoById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, toko)
	}

	if nama != "" {
		toko, err := h.service.SearchToko(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, toko)
	}

}

func (h *handlerToko) CreateToko(c *gin.Context) {
	// parsing request body
	var toko dto.CreateTokoRequest
	err := c.ShouldBindJSON(&toko)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newToko, err := h.service.CreateToko(toko)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newToko)
}
