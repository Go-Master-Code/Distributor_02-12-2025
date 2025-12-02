package dto

type CreateKategoriTokoRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateKategoriTokoRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type KategoriTokoResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
