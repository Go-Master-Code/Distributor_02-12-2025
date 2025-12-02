package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct implementasi
type handlerSupplier struct {
	service service.ServiceSupplier
}

// constructor
func NewHandlerSupplier(service service.ServiceSupplier) *handlerSupplier {
	return &handlerSupplier{service}
}

func (h *handlerSupplier) GetAllSupplier(c *gin.Context) {
	supplier, err := h.service.GetAllSupplier()
	if err != nil {
		helper.ErrorFetchDataFromDB(c, err)
		return
	}

	helper.StatusSuksesGetData(c, supplier)
}

func (h *handlerSupplier) CreateSupplier(c *gin.Context) {
	var supplier dto.CreateSupplierRequest
	// parsing json body
	err := c.ShouldBindJSON(&supplier)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newSupplier, err := h.service.CreateSupplier(supplier)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	helper.StatusSuksesCreateData(c, newSupplier)
}

func (h *handlerSupplier) UpdateSupplier(c *gin.Context) {
	var req dto.UpdateSupplierRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	// ambil param id
	id := c.Param("id")
	// convert id ke int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	supplierDTO, err := h.service.UpdateSupplier(idInt, req)
	if err != nil {
		helper.ErrorUpdateData(c, err)
		return
	}

	helper.StatusSuksesUpdateData(c, supplierDTO)
}

func (h *handlerSupplier) DeleteSupplier(c *gin.Context) {
	// ambil param id
	id := c.Param("id")
	// convert id ke int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	supplier, err := h.service.DeleteSupplier(idInt)
	if err != nil {
		helper.ErrorDeleteData(c)
		return
	}

	helper.StatusSuksesDeleteData(c, supplier)
}

func (h *handlerSupplier) GetSupplierByID(c *gin.Context) {
	// get param id
	// ambil param id
	id := c.Param("id")
	// convert id ke int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorParsingAtoi(c, err)
		return
	}

	supplier, err := h.service.GetSupplierByID(idInt)

	if err != nil {
		helper.ErrorDeleteData(c)
		return
	}

	helper.StatusSuksesGetData(c, supplier)
}
