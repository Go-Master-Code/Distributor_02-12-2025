package dto

import "time"

type CreatePenjualanRequest struct {
	NoFaktur      string                         `json:"no_faktur" binding:"required"`
	TglPenjualan  string                         `json:"tgl_penjualan" binding:"required"`
	TglJatuhTempo string                         `json:"tgl_jatuh_tempo" binding:"required"`
	TokoID        int                            `json:"toko_id" binding:"required"`
	SalesID       int                            `json:"sales_id" binding:"required"`
	Total         int                            `json:"total" binding:"required"`
	TotalNetto    int                            `json:"total_netto" binding:"required"`
	Keterangan    string                         `json:"keterangan"`
	Items         []CreatePenjualanDetailRequest `json:"items" binding:"required,dive"`
}

type PenjualanResponse struct {
	ID            int                       `json:"id"`
	NoFaktur      string                    `json:"no_faktur"`
	TglPenjualan  string                    `json:"tgl_penjualan"`
	TglJatuhTempo string                    `json:"tgl_jatuh_tempo"`
	TokoID        int                       `json:"toko_id"`
	TokoNama      string                    `json:"toko_nama"`
	TokoAlamat    string                    `json:"toko_alamat"`
	TokoKota      string                    `json:"toko_kota"`
	TokoArea      string                    `json:"toko_area"`
	Disc1         float64                   `json:"disc_1"`
	Disc2         float64                   `json:"disc_2"`
	Disc3         float64                   `json:"disc_3"`
	SalesID       int                       `json:"sales_id"`
	SalesNama     string                    `json:"sales_nama"`
	Total         int                       `json:"total"`
	TotalNetto    int                       `json:"total_netto"`
	Keterangan    string                    `json:"keterangan"`
	CreatedAt     time.Time                 `json:"created_at"`
	Items         []PenjualanDetailResponse `json:"items"`
	// TokoNama     string                    `json:"toko_nama"`
}
