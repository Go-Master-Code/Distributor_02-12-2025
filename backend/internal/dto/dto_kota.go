package dto

type CreateKotaRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateKotaRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type KotaResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
