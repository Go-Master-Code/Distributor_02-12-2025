package dto

type CreateAreaRequest struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateAreaRequest struct {
	Nama string `json:"nama" binding:"omitempty"`
}

type AreaResponse struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}
