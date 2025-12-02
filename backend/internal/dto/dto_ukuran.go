package dto

type CreateUkuranRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateUkuranRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type UkuranResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
