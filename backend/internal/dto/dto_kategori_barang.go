package dto

type CreateKategoriBarangRequest struct {
	Huruf string `json:"huruf" binding:"required"`
	Nama  string `json:"nama" binding:"required"`
}

type UpdateKategoriBarangRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type KategoriBarangResponse struct {
	ID    int    `json:"id"`
	Huruf string `json:"huruf"`
	Nama  string `json:"nama"`
}
