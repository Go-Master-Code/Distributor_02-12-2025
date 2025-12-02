package dto

type CreateTokoRequest struct {
	Kode           string  `json:"kode" binding:"required"`
	Nama           string  `json:"nama" binding:"required"`
	KategoriTokoID int     `json:"kategori_toko_id" binding:"required"`
	KotaID         int     `json:"kota_id" binding:"required"`
	Alamat         string  `json:"alamat" binding:"required"`
	Disc1          float64 `json:"disc_1" binding:"required"`
	Disc2          float64 `json:"disc_2" binding:"required"`
	Disc3          float64 `json:"disc_3" binding:"required"`
	EkspedisiID    int     `json:"ekspedisi_id" binding:"required"`
	OngkirID       int     `json:"ongkir_id" binding:"required"`
}

type UpdateTokoRequest struct {
	Kode           *string  `json:"kode" binding:"omitempty"`
	Nama           *string  `json:"nama" binding:"omitempty"`
	KategoriTokoID *int     `json:"kategori_toko_id" binding:"omitempty"`
	KotaID         *int     `json:"kota_id" binding:"omitempty"`
	Alamat         *string  `json:"alamat" binding:"omitempty"`
	Disc1          *float64 `json:"disc_1" binding:"omitempty"`
	Disc2          *float64 `json:"disc_2" binding:"omitempty"`
	Disc3          *float64 `json:"disc_3" binding:"omitempty"`
	EkspedisiID    *int     `json:"ekspedisi_id" binding:"omitempty"`
	OngkirID       *int     `json:"ongkir_id" binding:"omitempty"`
}

type TokoResponse struct {
	ID               int     `json:"id"`
	Kode             string  `json:"kode"`
	Nama             string  `json:"nama"`
	KategoriTokoID   int     `json:"kategori_toko_id"`
	KategoriTokoNama string  `json:"kategori_toko_nama"`
	KotaID           int     `json:"kota_id"`
	KotaNama         string  `json:"kota_nama"`
	Alamat           string  `json:"alamat"`
	Disc1            float64 `json:"disc_1"`
	Disc2            float64 `json:"disc_2"`
	Disc3            float64 `json:"disc_3"`
	EkspedisiID      int     `json:"ekspedisi_id"`
	EkspedisiNama    string  `json:"ekspedisi_nama"`
	OngkirID         int     `json:"ongkir_id"`
	OngkirNama       string  `json:"ongkir_nama"`
}
