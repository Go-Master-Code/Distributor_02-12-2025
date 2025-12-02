package helper

import (
	"api-distributor/internal/dto"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// ======================= STRUCT =======================
type Produk struct {
	Kode     string
	Artikel  string
	Warna    string
	Kategori string
	Size     int
	Stok     int
	Harga    int
}

type PivotRow struct {
	Kode     string
	Artikel  string
	Warna    string
	Kategori string
	Harga    int
	QtyMap   map[string]map[int]int
}

type SizeMapping struct {
	KatAsal    string
	SizeAsal   int
	KatTujuan  string
	SizeTujuan int
}

// ======================= BUILD PIVOT =======================
func BuildPivot(data []dto.PenjualanDetailResponse, kategoriSizeRange map[string][]int) map[string]*PivotRow {
	pivot := make(map[string]*PivotRow)

	for _, d := range data {
		// gunakan Kode Barang sebagai key
		key := d.BarangKode

		if _, ok := pivot[key]; !ok {
			pivot[key] = &PivotRow{
				Kode:     d.BarangKode,
				Artikel:  d.BarangArtikel,
				Warna:    d.BarangWarna,
				Kategori: d.BarangKategori,
				Harga:    d.Harga,
				QtyMap:   make(map[string]map[int]int),
			}

			// siapkan semua kategori
			for k := range kategoriSizeRange {
				pivot[key].QtyMap[k] = make(map[int]int)
			}
		}

		kat := d.BarangKategori

		// pastikan hanya mengisi size yang valid di kategoriSizeRange
		for _, allowedSize := range kategoriSizeRange[kat] {
			if allowedSize == d.Size {
				// isi qty asli dulu
				pivot[key].QtyMap[kat][allowedSize] += d.Qty

				// === APPLY AUTO-CONVERT SIZE ===
				conversions := AutoConvertSize(kat, d.Size)
				for _, conv := range conversions {
					pivot[key].QtyMap[conv.KatTujuan][conv.SizeTujuan] += d.Qty
				}

				// === HAPUS QTY ASLI KECUALI K ===
				if kat == "A" || kat == "D" {
					pivot[key].QtyMap[kat][allowedSize] = 0
				}

				break
			}
		}
	}

	return pivot
}

// ======================= AUTO SIZE =======================
func AutoConvertSize(kat string, size int) []SizeMapping {
	var res []SizeMapping

	if kat == "D" {
		// convert size D 36 s.d 41
		dest := size - 10
		if dest >= 26 && dest <= 31 {
			res = append(res, SizeMapping{"D", size, "K", dest})
		}
		if dest >= 32 && dest <= 37 {
			res = append(res, SizeMapping{"D", size, "A", dest})
		}
	}

	if kat == "A" {
		dest := size - 6
		if dest >= 26 && dest <= 31 {
			res = append(res, SizeMapping{"A", size, "K", dest})
		}
	}

	return res
}

// ======================= PRINT HEADER =======================
func PrintHeader(jual dto.PenjualanResponse, kategoriSizeRange map[string][]int, kategoriLabel map[string]string, totalSizes int) string {
	var b strings.Builder

	// parsing tanggal jual dan jatuh tempo dari sting ke date
	tglPenjualan, err := time.Parse("2006-01-02", jual.TglPenjualan)
	if err != nil {
		return "Gagal parsing tanggal penjualan"
	}

	tglJatuhTempo, err := time.Parse("2006-01-02", jual.TglJatuhTempo)
	if err != nil {
		return "Gagal parsing tanggal jatuh tempo"
	}

	b.WriteString(strings.Repeat("─", 90) + "\n")
	b.WriteString(fmt.Sprintf("MITRA SUKSES BERSAMA%10sNama Toko: %s\n", "", jual.TokoNama))
	b.WriteString(fmt.Sprintf("FAKTUR PENJUALAN%13s %s (%s)\n", "", jual.TokoAlamat, jual.TokoKota))
	b.WriteString(fmt.Sprintf("No.Faktur: %-10s         Tanggal Faktur: %s\n",
		jual.NoFaktur,
		tglPenjualan.Format("2006-01-02")))
	b.WriteString(fmt.Sprintf("Sales ID: %1d %18sJatuh Tempo: %s\n", jual.SalesID, " ", tglJatuhTempo.Format("2006-01-02")))

	b.WriteString(strings.Repeat("─", 90) + "\n")
	fmt.Fprintf(&b, "%-8s", "Kode")
	fmt.Fprintf(&b, "%-10s", "Artikel")
	fmt.Fprintf(&b, "%-12s", "Warna")
	for i, kat := range []string{"K", "A", "D"} {
		if i > 0 {
			if i == 2 {
				fmt.Fprintf(&b, "%30s%-4s", "", kategoriLabel[kat])
				for _, s := range kategoriSizeRange[kat] {
					fmt.Fprintf(&b, "%-3d", s)
				}
				fmt.Fprintf(&b, "%3s%8s%9s\n", "Qty", "Harga", "Total")
			} else {
				fmt.Fprintf(&b, "%30s%-4s", "", kategoriLabel[kat])
				for _, s := range kategoriSizeRange[kat] {
					fmt.Fprintf(&b, "%-3d", s)
				}
				b.WriteString("\n")
			}
		} else {
			fmt.Fprintf(&b, "%-4s", kategoriLabel[kat])
			for _, s := range kategoriSizeRange[kat] {
				fmt.Fprintf(&b, "%-3d", s)
			}
			pad := totalSizes - len(kategoriSizeRange[kat])
			for i := 0; i < pad; i++ {
				fmt.Fprintf(&b, "%-6s", "")
			}
			b.WriteString("\n")
		}
	}
	b.WriteString(strings.Repeat("─", 90) + "\n")
	return b.String()
}

// ======================= PRINT DATA =======================
func PrintData(pivot map[string]*PivotRow, kategoriSizeRange map[string][]int) string {
	var b strings.Builder

	keys := []string{}
	for k := range pivot {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		p := pivot[key]

		// tampilkan artikel + warna
		fmt.Fprintf(&b, "%-7s %-9s %-12s %-3s", p.Kode, p.Artikel, p.Warna, p.Kategori)

		row := []int{}
		for _, kat := range []string{"K", "A", "D"} {
			for _, s := range kategoriSizeRange[kat] {
				row = append(row, p.QtyMap[kat][s])
			}
		}

		// pemotongan 12 kolom terakhir (size dewasa D)
		if len(row) > 12 {
			row = row[:len(row)-12]
		}

		totalQty := 0
		for _, v := range row {
			if v == 0 {
				fmt.Fprintf(&b, "%-3s", "")
			} else {
				fmt.Fprintf(&b, "%-3d", v)
				totalQty += v
			}
		}

		totalUang := totalQty * p.Harga
		fmt.Fprintf(&b, "%3d %7d %8d\n", totalQty, p.Harga, totalUang)
		b.WriteString(strings.Repeat("─", 90) + "\n")
	}

	return b.String()
}

// ======================= PRINT SUMMARY =======================
func PrintSumarry(response dto.PenjualanResponse) string {
	var b strings.Builder

	// dummy data penjualan
	// penjualan := model.Penjualan{
	// 	ID:            111,
	// 	NoFaktur:      "JT-001",
	// 	TglPenjualan:  time.Now(),
	// 	TglJatuhTempo: time.Now(),
	// 	TokoID:        1,
	// 	Toko: model.Toko{
	// 		Nama:   "AGRIS SPORT",
	// 		Alamat: "JL. KH Zainal Alim No. 31. Kemayoran.",
	// 		KotaID: 66,
	// 		Kota: model.Kota{
	// 			Nama: "BANGKALAN",
	// 		},
	// 		Disc1: 0.3,
	// 		Disc2: 0,
	// 		Disc3: 0,
	// 	},
	// 	Total:      2500000,
	// 	TotalNetto: 2644740,
	// 	CreatedAt:  time.Now(),
	// }

	// total := 3778200

	// perhitungan diskon 1,2, dan 3
	var grandTotal float64
	grandTotal = float64(response.Total)

	// fmt.Println("Total bruto:", grandTotal)
	diskon1 := response.Disc1 * grandTotal
	// fmt.Println("Diskon 1:", diskon1)
	grandTotal -= diskon1
	// fmt.Println("Grand total after diskon 1:", grandTotal)
	diskon2 := response.Disc2 * grandTotal
	// fmt.Println("Diskon 2:", diskon2)
	grandTotal -= diskon2
	// fmt.Println("Grand total after diskon 2:", grandTotal)
	diskon3 := response.Disc3 * grandTotal
	// fmt.Println("Diskon 3:", diskon3)
	grandTotal -= diskon3
	// fmt.Println("Grand total after diskon 3:", grandTotal)

	//grandTotal = grandTotal - diskon3
	b.WriteString(fmt.Sprintf("%81s %8d\n", "Total", response.Total))

	b.WriteString(fmt.Sprintf("%73s (%4.0f%%) %8.0f\n", "Disc 1", response.Disc1*100, diskon1)) // %9.0F ARTINYA pembualan ke bilangan bulat tanpa desimal
	b.WriteString(fmt.Sprintf("%73s (%4.0f%%) %8.0f\n", "Disc 2", response.Disc2*100, diskon2))
	b.WriteString(fmt.Sprintf("%73s (%4.0f%%) %8.0f\n", "Disc 3", response.Disc3*100, diskon3))

	b.WriteString(fmt.Sprintf("%81s %8d\n", "Grand Total", response.TotalNetto))
	b.WriteString(strings.Repeat("─", 90) + "\n")
	b.WriteString(centerText("TERIMA KASIH ATAS PEMBELIAN ANDA", 88) + "\n")
	b.WriteString(strings.Repeat("─", 90) + "\n")

	return b.String()
}

// agar text terima kasih berada di tengah
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text
}

// ======================= WRITE FILE =======================
func WriteToFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// ======================= FORMAT RIBUAN =======================
func FormatRibuan(n int) string {
	s := fmt.Sprintf("%d", n)
	l := len(s)
	if l <= 3 {
		return s
	}

	var result strings.Builder
	pre := l % 3
	if pre > 0 {
		result.WriteString(s[:pre])
		if l > pre {
			result.WriteString(".")
		}
	}

	for i := pre; i < l; i += 3 {
		result.WriteString(s[i : i+3])
		if i+3 < l {
			result.WriteString(".")
		}
	}

	return result.String()
}
