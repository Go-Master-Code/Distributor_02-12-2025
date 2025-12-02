package dto

type CreateJenisBarangRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateJenisBarangRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type JenisBarangResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
