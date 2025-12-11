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
	pdf.SetAutoPageBreak(true, 15) // spasi tambahan jika perlu

	// // Setbackground color dan font color header
	// pdf.SetFillColor(255, 140, 60)  // Latar belakang orange (misalnya)
	// pdf.SetTextColor(255, 255, 255) // Teks putih
	// pdf.SetDrawColor(0, 0, 0)       // Border hitam

	// Header Tabel
	pdf.SetFont("Arial", "B", 12)
	headers := []string{"Faktur", "Area", "Sales", "Brand", "Tanggal", "Toko", "Kota", "Qty", "Bruto", "Disc", "Netto"} // judul header
	widths := []float64{20, 18, 20, 26, 25, 52, 32, 13, 25, 23, 25}                                                     // width masing-masing kolom
	aligns := []string{"C", "C", "C", "C", "C", "C", "C", "R", "R", "R", "R"}                                           // text-alignment masing-masing kolom

	// for i, str := range headers {
	// 	pdf.CellFormat(widths[i], 10, str, "1", 0, "C", true, 0, "")
	// }
	// pdf.Ln(-1)

	// ===================================================
	// ===============   HEADER FUNCTION   ===============
	// ===================================================
	pdf.SetHeaderFunc(func() {

		// Logo
		pdf.Image("internal/utils/report/penjualan/logo_msb.png", 62, 8, 40, 0, false, "", 0, "")

		// Judul
		pdf.SetFont("Arial", "B", 20)
		pdf.Ln(1)
		pdf.CellFormat(0, 10, "LAPORAN PENJUALAN", "", 1, "C", false, 0, "")
		pdf.Ln(1)

		// Nama Perusahaan
		pdf.CellFormat(0, 10, "MITRA SUKSES BERSAMA", "", 1, "C", false, 0, "")
		pdf.Ln(1)

		// Periode
		pdf.SetFont("Arial", "", 14)
		periode := "Dari " + awal + " Sampai " + akhir
		pdf.CellFormat(0, 6, periode, "", 1, "C", false, 0, "")
		pdf.Ln(4)

		// Header Table
		pdf.SetFont("Arial", "B", 12)
		pdf.SetFillColor(140, 180, 230) // #8CB4E6 biru muda soft
		pdf.SetTextColor(0, 0, 0)       // warna teks header hitam
		pdf.SetDrawColor(0, 0, 0)       // border hitam

		for i, title := range headers {
			pdf.CellFormat(widths[i], 10, title, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		// Balik warna isi tabel ke hitam
		pdf.SetTextColor(0, 0, 0)
	})

	// ===================================================
	// ===============         FOOTER       ===============
	// ===================================================
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 10)
		pdf.CellFormat(0, 10, "Halaman "+strconv.Itoa(pdf.PageNo()), "", 0, "C", false, 0, "")
	})

	// Halaman pertama
	pdf.AddPage()

	// ===================================================
	// ===============       TABLE BODY      ==============
	// ===================================================
	pdf.SetTextColor(0, 0, 0) // Supaya isi tabel kembali teks hitam
	pdf.SetFont("Arial", "", 10)

	// ambil nama brand
	brand := detil[0].BarangMerk

	// grand total bruto
	grandTotalBruto := 0

	// grand total netto
	grandTotalNetto := 0

	// grand total qty (summary)
	grandTotalQty := 0

	for _, penjualan := range master {
		// update += total bruto per master faktur
		grandTotalBruto += penjualan.Total

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

				// increment grand total qty per row data
				grandTotalQty += d.Qty
			}
		}

		row := []string{
			penjualan.NoFaktur,
			penjualan.TokoArea,
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

	// ===================================================
	// ===============       TOTAL AKHIR     ==============
	// ===================================================
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(widths[0]+widths[1]+widths[2]+widths[3]+widths[4]+widths[5]+widths[6], 10, "Total Qty, Bruto, Netto:", "1", 0, "R", false, 0, "")
	pdf.CellFormat(widths[7], 10, FormatAngka(grandTotalQty), "1", 0, "R", false, 0, "")
	pdf.CellFormat(widths[8], 10, FormatAngka(grandTotalBruto), "1", 0, "R", false, 0, "")
	pdf.CellFormat(widths[9], 10, "", "1", 0, "R", false, 0, "")
	pdf.CellFormat(widths[10], 10, FormatAngka(grandTotalNetto), "1", 0, "R", false, 0, "")
	//pdf.CellFormat(28, 10, , "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
