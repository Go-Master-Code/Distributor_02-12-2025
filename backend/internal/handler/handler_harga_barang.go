package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerHargaBarang struct {
	service service.ServiceHargaBarang
}

// constructor
func NewHandlerHargaBarang(service service.ServiceHargaBarang) *handlerHargaBarang {
	return &handlerHargaBarang{service}
}

// func (h *handlerHargaBarang) GetAllHargaBarang(c *gin.Context) {
// 	harga, err := h.service.GetAllHargaBarang()
// 	if err != nil {
// 		helper.ErrorFetchDataFromDB(c, err)
// 		return
// 	}

// 	helper.StatusSuksesGetData(c, harga)
// }

func (h *handlerHargaBarang) GetHargaBarang(c *gin.Context) {
	idBarang := c.Query("id")

	if idBarang != "" {
		// cari data harga by ID barang
		idBarangInt, err := strconv.Atoi(idBarang)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		hb, err := h.service.GetHargaBarangById(idBarangInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, hb)
	}
}

// func (h *handlerHargaBarang) CreateHargaBarang(c *gin.Context) {
// 	var harga dto.CreateHargaBarangRequest
// 	// parsing json body
// 	err := c.ShouldBindJSON(&harga)
// 	if err != nil {
// 		helper.ErrorParsingRequestBody(c, err)
// 		return
// 	}

// 	newHargaBarang, err := h.service.CreateHargaBarang(harga)
// 	if err != nil {
// 		helper.ErrorCreateData(c, err)
// 		return
// 	}

// 	helper.StatusSuksesCreateData(c, newHargaBarang)
// }
