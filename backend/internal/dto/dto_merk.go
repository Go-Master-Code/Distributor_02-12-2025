package dto

type CreateMerkRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateMerkRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type MerkResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
