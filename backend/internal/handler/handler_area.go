package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/service"

	"github.com/gin-gonic/gin"
)

type handlerArea struct {
	service service.ServiceArea
}

// constructor
func NewHandlerArea(service service.ServiceArea) *handlerArea {
	return &handlerArea{service}
}

// func
func (h *handlerArea) GetArea(c *gin.Context) {
	// nantinya akan ada 3 fungsi, getAll, search by id, search by nama
	// ambil query nama
	nama := c.Query("nama")

	if nama != "" {
		areaDTO, err := h.service.SearchArea(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, areaDTO)
	} else {
		// jika query nama = nil, getAllArea
		h.GetAllArea(c)
	}

}

func (h *handlerArea) GetAllArea(c *gin.Context) {
	area, err := h.service.GetAllArea()
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, area)
}
