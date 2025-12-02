package dto

type CreateEkspedisiRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateEkspedisiRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type EkspedisiResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
