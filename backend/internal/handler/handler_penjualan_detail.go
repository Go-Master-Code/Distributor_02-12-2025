package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerPenjualanDetail struct {
	service service.ServicePenjualanDetail
}

func NewHandlerPenjualanDetail(service service.ServicePenjualanDetail) *handlerPenjualanDetail {
	return &handlerPenjualanDetail{service}
}

func (h *handlerPenjualanDetail) GetAllPenjualanDetail(c *gin.Context) {
	// ambil param
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	dj, err := h.service.GetPenjualanDetailById(idInt)
	if err != nil {
		helper.ErrorDataNotFound(c)
		return
	}

	helper.StatusSuksesGetData(c, dj)
}

// func (h *handlerPenjualan) CreatePenjualan(c *gin.Context) {
// 	// parsing json body
// 	var penjualan dto.CreatePenjualanRequest
// 	err := c.ShouldBindJSON(&penjualan)
// 	if err != nil {
// 		helper.ErrorParsingRequestBody(c, err)
// 		return
// 	}

// 	newPenjualan, err := h.service.CreatePenjualan(penjualan)
// 	if err != nil {
// 		helper.ErrorCreateData(c, err)
// 		return
// 	}

// 	helper.StatusSuksesCreateData(c, newPenjualan)
// }
