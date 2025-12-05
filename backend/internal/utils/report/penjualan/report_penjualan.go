package penjualan

import (
	"api-distributor/internal/dto"
	"bytes"
	"strconv"

	"github.com/go-pdf/fpdf"
)

func FormatAngka(n int) string {
	s := strconv.Itoa(n)
	l := len(s)
	if l <= 3 {
		return s
	}

	var result string
	for i, v := range s {
		if (l-i)%3 == 0 && i != 0 {
			result += "."
		}
		result += string(v)
	}
	return result
}

func GenerateReportPresensiAllPerPeriode(awal, akhir string, master []dto.PenjualanResponse, detil []dto.PenjualanDetailResponse) ([]byte, error) {
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.SetTitle("Laporan Penjualan", false)

	// Footer halaman otomatis
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Halaman "+strconv.Itoa(pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	pdf.AddPage()

	// add logo
	// pdf.Image("internal/utils/report/pf.jpg", 10, 10, 30, 0, false, "", 0, "")

	// Judul
	pdf.SetFont("Arial", "B", 20)

	// beri spasi agar judul bisa sejajar logo
	pdf.Ln(3)
	pdf.CellFormat(0, 10, "Laporan Penjualan", "", 1, "C", false, 0, "") // judul rata tengah
	pdf.Ln(1)                                                            // spasi tambahan jika perlu

	// Periode
	pdf.SetFont("Arial", "", 14)

	// nama bulan string
	periode := "Dari " + awal + " Sampai " + akhir

	// rata tengah
	pdf.CellFormat(0, 6, periode, "", 1, "C", false, 0, "")
	pdf.Ln(2)

	// Geser posisi Y ke bawah agar logo tidak tertimpa judul / tabel
	// pdf.SetY(42)

	// Setbackground color dan font color header
	pdf.SetFillColor(255, 140, 60)  // Latar belakang orange (misalnya)
	pdf.SetTextColor(255, 255, 255) // Teks putih
	pdf.SetDrawColor(0, 0, 0)       // Border hitam

	// Header Tabel
	pdf.SetFont("Arial", "B", 12)
	headers := []string{"No. Faktur", "Sales", "Brand", "Tanggal", "Toko", "Kota", "Qty", "Bruto", "Disc", "Netto"} // judul header
	widths := []float64{26, 20, 26, 25, 52, 32, 15, 28, 25, 28}                                                     // width masing-masing kolom
	aligns := []string{"C", "C", "C", "C", "C", "C", "R", "R", "R", "R"}                                            // text-alignment masing-masing kolom

	for i, str := range headers {
		pdf.CellFormat(widths[i], 10, str, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// Isi Tabel
	pdf.SetTextColor(0, 0, 0) // Supaya isi tabel kembali teks hitam
	pdf.SetFont("Arial", "", 10)

	// jumlah karyawan
	// jmlKaryawan := 0

	// dummy data
	// pdf.CellFormat(widths[10], 10, "str", "1", 0, aligns[5], false, 0, "") // C = Center

	// ambil nama brand
	brand := detil[0].BarangMerk

	// grand total netto
	grandTotalNetto := 0

	for _, penjualan := range master {
		// update += total netto per master faktur
		grandTotalNetto += penjualan.TotalNetto

		// Hitung Total Qty per Faktur
		totalQty := 0
		for _, d := range detil {
			//log.Println("ID detil: ", d.PenjualanID)
			//log.Println("ID master jual: ", penjualan.ID)
			if d.PenjualanID == penjualan.ID { // match faktur berdasarkan foreign key ID
				totalQty += d.Qty
				// log.Println("Cocok!", totalQty)
			}
		}

		row := []string{
			penjualan.NoFaktur,
			penjualan.SalesNama,
			brand,
			penjualan.TglPenjualan,
			penjualan.TokoNama,
			penjualan.TokoKota,
			FormatAngka(totalQty),
			FormatAngka(penjualan.Total),
			FormatAngka(penjualan.Total - penjualan.TotalNetto),
			FormatAngka(penjualan.TotalNetto),
		}

		for j, str := range row {
			pdf.CellFormat(widths[j], 10, str, "1", 0, aligns[j], false, 0, "") // C = Center
		}

		pdf.Ln(-1)
	}
	// =======ISI DATA=======
	// for _, penjualan := range detil {
	// 	row := []string{
	// 		// nomor row
	// 		presensi.KaryawanID,
	// 		presensi.Nama,
	// 		strconv.Itoa(presensi.Kehadiran),
	// 		strconv.Itoa(report.HariKerja - presensi.Kehadiran),
	// 	}

	// 	for j, str := range row {
	// 		pdf.CellFormat(widths[j], 10, str, "1", 0, aligns[j], false, 0, "") // C = Center
	// 	}

	// 	// tambah counter jmlKaryawan
	// 	jmlKaryawan += 1

	// 	pdf.Ln(-1)
	// }

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(widths[0]+widths[1]+widths[2]+widths[3]+widths[4]+widths[5]+widths[6]+widths[7]+widths[8]+widths[9], 10, "Total Penjualan Netto: "+FormatAngka(grandTotalNetto), "1", 0, "R", false, 0, "")
	//pdf.CellFormat(28, 10, , "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
