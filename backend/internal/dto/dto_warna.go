package dto

type CreateWarnaRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateWarnaRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type WarnaResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
