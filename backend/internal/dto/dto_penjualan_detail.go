package dto

type CreatePenjualanDetailRequest struct {
	// PenjualanID int `json:"penjualan_id" binding:"required"`
	BarangID int `json:"barang_id" binding:"required"`
	Size     int `json:"size" binding:"required"`
	Qty      int `json:"qty" binding:"required"`
	Harga    int `json:"harga" binding:"required"`
	Subtotal int `json:"subtotal"`
}

type PenjualanDetailResponse struct {
	// ID          int             `json:"id"`
	// PenjualanID int             `json:"penjualan_id"`
	// kedua atribut di atas dimatikan karena ID digenerate otomatis dan PenjualanID baru ada setelah data master penjualan dibuat
	BarangID       int             `json:"barang_id"`
	BarangKode     string          `json:"barang_kode"`
	BarangArtikel  string          `json:"barang_artikel"`
	BarangWarna    string          `json:"barang_warna"`
	BarangUkuran   string          `json:"barang_ukuran"` // ukuran yang ada di tabel barang (38-43) misalnya
	BarangKategori string          `json:"barang_kategori"`
	Size           int             `json:"size"`
	Qty            int             `json:"qty"`
	Harga          int             `json:"harga"`
	Subtotal       int             `json:"subtotal"`
	Barang         *BarangResponse `json:"barang,omitempty"`
	// BarangNama  string          `json:"barang_nama"`
}
