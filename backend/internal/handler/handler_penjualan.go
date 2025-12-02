package handler

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type handlerPenjualan struct {
	service service.ServicePenjualan
}

func NewHandlerPenjualan(service service.ServicePenjualan) *handlerPenjualan {
	return &handlerPenjualan{service}
}

// func (h *handlerPenjualan) GetAllPenjualan(c *gin.Context) {
// 	penjualan, err := h.service.GetAllPenjualan()
// 	if err != nil {
// 		helper.ErrorDataNotFound(c)
// 		return
// 	}

// 	helper.StatusSuksesGetData(c, penjualan)
// }

func (h *handlerPenjualan) CreatePenjualan(c *gin.Context) {
	// parsing json body
	var penjualan dto.CreatePenjualanRequest
	err := c.ShouldBindJSON(&penjualan)

	if err != nil {
		helper.ErrorParsingRequestBody(c, err)
		return
	}

	newPenjualan, fakturContent, fileName, err := h.service.CreatePenjualan(penjualan)
	if err != nil {
		helper.ErrorCreateData(c, err)
		return
	}

	// ---- Nama file memakai nomor faktur ----
	fmt.Println("Nomor faktur:", newPenjualan.NoFaktur)
	fileName = fmt.Sprintf("Faktur_%s.txt", newPenjualan.NoFaktur)
	fmt.Println("Nama file:", fileName)

	// Header download
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Header("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	c.Data(200, "text/plain; charset=utf-8", []byte(fakturContent))

	// === KIRIM FILE ===
	// c.String(200, fakturContent) jangan panggil c.String lagi karena sudah pakai c.Data, nanti konten file notepad nya tercetak 2x

	// response API
	// helper.StatusSuksesCreateData(c, newPenjualan)
}
