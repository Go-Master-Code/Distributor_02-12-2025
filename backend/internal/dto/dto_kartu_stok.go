package dto

type KartuStokResponse struct {
	ID          int    `json:"id"`
	BarangID    int    `json:"barang_id"`
	BarangKode  string `json:"barang_kode"`
	BarangMerk  string `json:"barang_merk"`
	BarangWarna string `json:"barang_warna"`
	Size        int    `json:"size"`
	Stok        int    `json:"stok"`
}
