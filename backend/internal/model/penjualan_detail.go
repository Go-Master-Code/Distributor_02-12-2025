package model

type PenjualanDetail struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	PenjualanID int    `gorm:"not null" json:"penjualan_id"`
	BarangID    int    `gorm:"not null" json:"barang_id"`
	Barang      Barang `gorm:"not null" json:"barang"`
	Size        int    `gorm:"not null" json:"size"`
	Qty         int    `gorm:"not null" json:"qty"`
	Harga       int    `gorm:"not null" json:"harga"`
	Subtotal    int    `gorm:"->;type:int;generated always as (qty * harga) stored" json:"subtotal"`
}

func (PenjualanDetail) TableName() string {
	return "penjualan_detail"
}
