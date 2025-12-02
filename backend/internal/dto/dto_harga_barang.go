package dto

type CreateHargaBarangRequest struct {
	BarangID     int    `json:"barang_id" binding:"required"`
	Harga        int    `json:"harga" binding:"required"`
	MulaiBerlaku string `json:"mulai_berlaku" binding:"required"`
}

type HargaBarangResponse struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	BarangID     int    `json:"barang_id"`
	BarangKode   string `json:"barang_kode"`
	Harga        int    `json:"harga"`
	MulaiBerlaku string `json:"mulai_berlaku"`
}
