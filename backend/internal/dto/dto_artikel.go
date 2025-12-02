package dto

type CreateArtikelRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateArtikelRequest struct {
	Nama *string `json:"nama" binding:"omitempty"`
}

type ArtikelResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
