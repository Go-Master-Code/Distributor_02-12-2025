package test

import (
	"api-distributor/internal/dto"
	"api-distributor/internal/utils/report/penjualan"
	"os"
	"testing"
	"time"
)

func TestGenerateReportPenjualanPerPeriode(t *testing.T) {
	// dummy data master penjualan
	master := []dto.PenjualanResponse{
		{
			ID:           1,
			NoFaktur:     "CONTOH 1",
			TglPenjualan: time.Now().Format("2006-01-02"),
			TokoID:       1,
			TokoNama:     "Toko Jual",
			TokoKota:     "SURABAYA",
			SalesID:      1,
			SalesNama:    "Umam",
			Total:        5125000,
			TotalNetto:   4571000,
			Keterangan:   "master dummy 1",
			//Items:        itemsDTO,
		},
		{
			ID:           2,
			NoFaktur:     "CONTOH 2",
			TglPenjualan: time.Now().Format("2006-01-02"),
			TokoID:       2,
			TokoNama:     "Toko Dummy",
			TokoKota:     "BANDUNG",
			SalesID:      2,
			SalesNama:    "Baskoro",
			Total:        10174846,
			TotalNetto:   8513800,
			Keterangan:   "master dummy 2",
			//Items:        itemsDTO,
		},
	}

	detil := []dto.PenjualanDetailResponse{
		{
			BarangID:       1,
			BarangKode:     "VP001",
			BarangArtikel:  "BLAZE",
			BarangWarna:    "ALL BLACK",
			BarangUkuran:   "37-44",
			BarangKategori: "D",
			Size:           41,
			Qty:            5,
			Harga:          139800,
			Subtotal:       699000,
		},
		{
			BarangID:       1,
			BarangKode:     "VP001",
			BarangArtikel:  "BLAZE",
			BarangWarna:    "ALL BLACK",
			BarangUkuran:   "37-44",
			BarangKategori: "D",
			Size:           40,
			Qty:            4,
			Harga:          139800,
			Subtotal:       559200,
		},
		{
			BarangID:       258,
			BarangKode:     "CAV001",
			BarangArtikel:  "LEMBATA",
			BarangWarna:    "HTM/HTM",
			BarangUkuran:   "38-43",
			BarangKategori: "D",
			Size:           39,
			Qty:            3,
			Harga:          140000,
			Subtotal:       420000,
		},
		{
			BarangID:       258,
			BarangKode:     "CAV001",
			BarangArtikel:  "LEMBATA",
			BarangWarna:    "HTM/HTM",
			BarangUkuran:   "38-43",
			BarangKategori: "D",
			Size:           40,
			Qty:            1,
			Harga:          140000,
			Subtotal:       140000,
		},
		{
			BarangID:       258,
			BarangKode:     "CAV001",
			BarangArtikel:  "LEMBATA",
			BarangWarna:    "HTM/HTM",
			BarangUkuran:   "38-43",
			BarangKategori: "D",
			Size:           42,
			Qty:            6,
			Harga:          140000,
			Subtotal:       840000,
		},
		{
			BarangID:       258,
			BarangKode:     "CAV001",
			BarangArtikel:  "LEMBATA",
			BarangWarna:    "HTM/HTM",
			BarangUkuran:   "38-43",
			BarangKategori: "D",
			Size:           43,
			Qty:            8,
			Harga:          140000,
			Subtotal:       1120000,
		},
	}

	// build into bytes
	pdfBytes, err := penjualan.GenerateReportPresensiAllPerPeriode("2025-01-01", "2025-01-31", master, detil)
	if err != nil {
		t.Fatalf("error generating report penjualan: %v", err)
	}

	// simpan ke file
	err = os.WriteFile("test_penjualan.pdf", pdfBytes, 0644)
	if err != nil {
		t.Fatalf("error write file: %v", err)
	}

	t.Log("Laporan Penjualan pdf berhasil dibuat!")
}
