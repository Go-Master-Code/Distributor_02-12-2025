package helper

// import (
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// // ======================= STRUCT =======================
// type Produk struct {
// 	Nama     string
// 	Kategori string
// 	Size     int
// 	Stok     int
// 	Harga    int
// }

// type PivotRow struct {
// 	Nama     string
// 	Kategori string
// 	Harga    int
// 	QtyMap   map[string]map[int]int
// }

// type SizeMapping struct {
// 	KatAsal    string
// 	SizeAsal   int
// 	KatTujuan  string
// 	SizeTujuan int
// }

// // ======================= BUILD PIVOT =======================
// func BuildPivot(data []Produk, kategoriSizeRange map[string][]int) map[string]*PivotRow {
// 	pivot := map[string]*PivotRow{}

// 	for _, p := range data {
// 		if _, ok := pivot[p.Nama]; !ok {
// 			pivot[p.Nama] = &PivotRow{
// 				Nama:     p.Nama,
// 				Kategori: p.Kategori,
// 				Harga:    p.Harga,
// 				QtyMap:   map[string]map[int]int{},
// 			}
// 			for kat, sizes := range kategoriSizeRange {
// 				pivot[p.Nama].QtyMap[kat] = map[int]int{}
// 				for _, s := range sizes {
// 					pivot[p.Nama].QtyMap[kat][s] = 0
// 				}
// 			}
// 		}
// 		pivot[p.Nama].QtyMap[p.Kategori][p.Size] = p.Stok
// 	}

// 	// jalankan auto size mapping
// 	for _, row := range pivot {
// 		for kat, sizes := range row.QtyMap {
// 			for size, stok := range sizes {
// 				if stok == 0 {
// 					continue
// 				}
// 				mappings := AutoConvertSize(kat, size)
// 				for _, m := range mappings {
// 					row.QtyMap[m.KatTujuan][m.SizeTujuan] = stok
// 					row.QtyMap[m.KatAsal][m.SizeAsal] = 0
// 				}
// 			}
// 		}
// 	}

// 	return pivot
// }

// // ======================= AUTO SIZE =======================
// func AutoConvertSize(kat string, size int) []SizeMapping {
// 	var res []SizeMapping

// 	if kat == "D" {
// 		dest := size - 10
// 		if dest >= 26 && dest <= 31 {
// 			res = append(res, SizeMapping{"D", size, "K", dest})
// 		}
// 		if dest >= 32 && dest <= 37 {
// 			res = append(res, SizeMapping{"D", size, "A", dest})
// 		}
// 	}

// 	if kat == "A" {
// 		dest := size - 6
// 		if dest >= 26 && dest <= 31 {
// 			res = append(res, SizeMapping{"A", size, "K", dest})
// 		}
// 	}

// 	return res
// }

// // ======================= PRINT HEADER =======================
// func PrintHeader(kategoriSizeRange map[string][]int, kategoriLabel map[string]string, totalSizes int) string {
// 	var b strings.Builder
// 	b.WriteString(strings.Repeat("─", 82) + "\n")
// 	fmt.Fprintf(&b, "%-10s", "Nama")
// 	for i, kat := range []string{"K", "A", "D"} {
// 		if i > 0 {
// 			if i == 2 {
// 				fmt.Fprintf(&b, "%10s%-4s", "", kategoriLabel[kat])
// 				for _, s := range kategoriSizeRange[kat] {
// 					fmt.Fprintf(&b, "%-4d", s)
// 				}
// 				fmt.Fprintf(&b, "%3s%8s%9s\n", "Qty", "Harga", "Total")
// 			} else {
// 				fmt.Fprintf(&b, "%10s%-4s", "", kategoriLabel[kat])
// 				for _, s := range kategoriSizeRange[kat] {
// 					fmt.Fprintf(&b, "%-4d", s)
// 				}
// 				b.WriteString("\n")
// 			}
// 		} else {
// 			fmt.Fprintf(&b, "%-4s", kategoriLabel[kat])
// 			for _, s := range kategoriSizeRange[kat] {
// 				fmt.Fprintf(&b, "%-4d", s)
// 			}
// 			pad := totalSizes - len(kategoriSizeRange[kat])
// 			for i := 0; i < pad; i++ {
// 				fmt.Fprintf(&b, "%-6s", "")
// 			}
// 			b.WriteString("\n")
// 		}
// 	}
// 	b.WriteString(strings.Repeat("─", 82) + "\n")
// 	return b.String()
// }

// // ======================= PRINT DATA =======================
// func PrintData(pivot map[string]*PivotRow, kategoriSizeRange map[string][]int) string {
// 	var b strings.Builder
// 	names := []string{}
// 	for nama := range pivot {
// 		names = append(names, nama)
// 	}
// 	sort.Strings(names)

// 	for _, nama := range names {
// 		p := pivot[nama]
// 		fmt.Fprintf(&b, "%-10s %-3s", p.Nama, p.Kategori)

// 		row := []int{}
// 		for _, kat := range []string{"K", "A", "D"} {
// 			for _, s := range kategoriSizeRange[kat] {
// 				row = append(row, p.QtyMap[kat][s])
// 			}
// 		}

// 		if len(row) > 12 {
// 			row = row[:len(row)-12]
// 		}

// 		totalQty := 0
// 		for _, v := range row {
// 			if v == 0 {
// 				fmt.Fprintf(&b, "%-4s", "")
// 			} else {
// 				fmt.Fprintf(&b, "%-4d", v)
// 				totalQty += v
// 			}
// 		}

// 		totalUang := totalQty * p.Harga
// 		fmt.Fprintf(&b, "%3d %7d %8d\n", totalQty, p.Harga, totalUang)
// 		b.WriteString(strings.Repeat("─", 82) + "\n")
// 	}

// 	return b.String()
// }

// // ======================= WRITE FILE =======================
// func WriteToFile(filename string, content string) error {
// 	return os.WriteFile(filename, []byte(content), 0644)
// }

// // ======================= FORMAT RIBUAN =======================
// func FormatRibuan(n int) string {
// 	s := fmt.Sprintf("%d", n)
// 	l := len(s)
// 	if l <= 3 {
// 		return s
// 	}

// 	var result strings.Builder
// 	pre := l % 3
// 	if pre > 0 {
// 		result.WriteString(s[:pre])
// 		if l > pre {
// 			result.WriteString(".")
// 		}
// 	}

// 	for i := pre; i < l; i += 3 {
// 		result.WriteString(s[i : i+3])
// 		if i+3 < l {
// 			result.WriteString(".")
// 		}
// 	}

// 	return result.String()
// }
