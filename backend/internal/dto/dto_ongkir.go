package dto

type CreateOngkirRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateOngkirRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type OngkirResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
