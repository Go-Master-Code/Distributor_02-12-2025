package dto

type CreateSalesRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateSalesRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type SalesResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
