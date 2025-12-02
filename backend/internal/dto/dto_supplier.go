package dto

type CreateSupplierRequest struct {
	Nama   string `json:"nama" binding:"required"`
	Alamat string `json:"alamat" binding:"required"`
	Kontak string `json:"kontak" binding:"required"`
}

type UpdateSupplierRequest struct {
	Nama   *string `json:"nama" binding:"omitempty"`
	Alamat *string `json:"alamat" binding:"omitempty"`
	Kontak *string `json:"kontak" binding:"omitempty"`
}

type SupplierResponse struct {
	ID     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Kontak string `json:"kontak"`
}
