package dto

type CreateBarangRequest struct {
	Kode             string `json:"kode" binding:"required"`
	MerkID           int    `json:"merk_id" binding:"required"`
	ArtikelID        int    `json:"artikel_id" binding:"required"`
	WarnaID          int    `json:"warna_id" binding:"required"`
	KategoriBarangID int    `json:"kategori_barang_id" binding:"required"`
	JenisBarangID    int    `json:"jenis_barang_id" binding:"required"`
	UkuranID         int    `json:"ukuran_id"`
}

type UpdateBarangRequest struct {
	Kode             *string `json:"kode" binding:"omitempty"`
	MerkID           *int    `json:"merk_id" binding:"omitempty"`
	ArtikelID        *int    `json:"artikel_id" binding:"omitempty"`
	WarnaID          *int    `json:"warna_id" binding:"omitempty"`
	KategoriBarangID *int    `json:"kategori_barang_id" binding:"omitempty"`
	JenisBarangID    *int    `json:"jenis_barang_id" binding:"omitempty"`
	UkuranID         *int    `json:"ukuran_id" binding:"omitempty"`
}

type BarangResponse struct {
	ID                 int    `json:"id"`
	Kode               string `json:"kode"`
	MerkID             int    `json:"merk_id"`
	MerkNama           string `json:"merk_nama"`
	ArtikelID          int    `json:"artikel_id"`
	ArtikelNama        string `json:"artikel_nama"`
	WarnaID            int    `json:"warna_id"`
	WarnaNama          string `json:"warna_nama"`
	KategoriBarangID   int    `json:"kategori_barang_id"`
	KategoriBarangNama string `json:"kategori_barang_nama"`
	JenisBarangID      int    `json:"jenis_barang_id"`
	JenisBarangNama    string `json:"jenis_barang_nama"`
	UkuranID           int    `json:"ukuran_id"`
	UkuranNama         string `json:"ukuran_nama"`
	// CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
	// UpdatedAt       time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	// DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"` //tipe datanya bukan time.Time tapi gorm.DeletedAt -> penanda soft delete
}
