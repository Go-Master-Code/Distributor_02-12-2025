package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerSales struct {
	service service.ServiceSales
}

func NewHandlerSales(service service.ServiceSales) *handlerSales {
	return &handlerSales{service}
}

func (h *handlerSales) GetAllSales(c *gin.Context) {
	sales, err := h.service.GetAllSales()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, sales)
}

func (h *handlerSales) GetSales(c *gin.Context) {
	// coba ambil isi query "id" dan "nama"
	id := c.Query("id")
	nama := c.Query("nama")

	// jika kedua query tidak ditemukan
	if id == "" && nama == "" {
		h.GetAllSales(c)
	}

	if id != "" { // jika id tidak kosong
		// convert id to int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			helper.ErrorParsingAtoi(c, err)
			return
		}

		sales, err := h.service.GetSalesById(idInt)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}

		helper.StatusSuksesGetData(c, sales)
	}

	if nama != "" { // jika nama tidak kosong, lakukan lazy search per 20 data
		sales, err := h.service.SearchSales(nama)
		if err != nil {
			helper.ErrorDataNotFound(c)
			return
		}
		helper.StatusSuksesGetData(c, sales)
	}
}

func (h *handlerSales) CreateSales(c *gin.Context) {
	// parsing body json
	var sales dto.CreateSalesRequest
	err := c.ShouldBindJSON(&sales)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// insert data ke db
	newSales, err := h.service.CreateSales(sales)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newSales)
}
