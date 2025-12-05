package service

import (
	"api-distributor/helper"
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
	"api-distributor/internal/repository"
	"api-distributor/internal/utils/report/penjualan"
	"errors"
	"fmt"
	"sort"
	"time"

	"gorm.io/gorm"
)

type ServicePenjualan interface {
	GetAllPenjualan() ([]dto.PenjualanResponse, error)
	GeneratePenjualanPerPeriode(awal string, akhir string) ([]byte, error)
	CreatePenjualan(penjualan dto.CreatePenjualanRequest) (dto.PenjualanResponse, string, string, error)
}

type servicePenjualan struct {
	repo          repository.RepositoryPenjualan
	repoBarang    repository.RepositoryBarang
	repoKartuStok repository.RepositoryKartuStok
	repoDetilJual repository.RepositoryPenjualanDetail
	db            *gorm.DB
}

func NewServicePenjualan(db *gorm.DB, repo repository.RepositoryPenjualan, repoKartuStok repository.RepositoryKartuStok, repoBarang repository.RepositoryBarang, repoDetilJual repository.RepositoryPenjualanDetail) ServicePenjualan {
	return &servicePenjualan{
		repo:          repo,
		repoBarang:    repoBarang,
		repoKartuStok: repoKartuStok,
		repoDetilJual: repoDetilJual,
		db:            db,
	}
}

func (s *servicePenjualan) GetAllPenjualan() ([]dto.PenjualanResponse, error) {
	penjualan, err := s.repo.GetAllPenjualan()
	if err != nil {
		return []dto.PenjualanResponse{}, err
	}

	// convert model to dto
	jualDTO := helper.ConvertToDTOPenjualanPlural(penjualan)

	return jualDTO, nil
}

func (s *servicePenjualan) GeneratePenjualanPerPeriode(awal string, akhir string) ([]byte, error) {
	// get master penjualan
	jual, err := s.repo.GetPenjualanPerPeriode(awal, akhir)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	jualDTO := helper.ConvertToDTOPenjualanPlural(jual)

	// ambil semua id penjualan
	ids := make([]uint, len(jualDTO))
	for i, j := range jualDTO {
		ids[i] = uint(j.ID)
	}

	// get detil penjualan
	detilJual, err := s.repoDetilJual.GetPenjualanDetailByFakturIds(ids)

	if err != nil {
		return nil, err
	}

	detilJualDTO := helper.ConvertToDTOPenjualanDetailPlural(detilJual)

	// penjualan -> dari folder utils
	return penjualan.GenerateReportPresensiAllPerPeriode(awal, akhir, jualDTO, detilJualDTO)
}

func (s *servicePenjualan) CreatePenjualan(req dto.CreatePenjualanRequest) (dto.PenjualanResponse, string, string, error) {
	var response dto.PenjualanResponse // untuk tampung response akhir

	// insert data master dan detil penjualan menggunakan transaction
	tx := s.db.Begin() // mulai transaction

	if tx.Error != nil {
		return response, "", "", tx.Error // jika terjadi error
	}

	// parse tanggal dari string (input) ke time.Time
	tglJual, err := time.Parse("2006-01-02", req.TglPenjualan)
	if err != nil {
		return dto.PenjualanResponse{}, "", "", err
	}

	tglJatuhTempo, err := time.Parse("2006-01-02", req.TglJatuhTempo)
	if err != nil {
		return dto.PenjualanResponse{}, "", "", err
	}

	// mapping dto ke model penjualan
	penjualan := model.Penjualan{
		NoFaktur:      req.NoFaktur,
		TglPenjualan:  tglJual,
		TglJatuhTempo: tglJatuhTempo,
		TokoID:        req.TokoID,
		SalesID:       req.SalesID,
		Total:         req.Total,
		TotalNetto:    req.TotalNetto,
		Keterangan:    req.Keterangan,
	}

	// insert master penjualan
	if err := s.repo.CreatePenjualan(tx, &penjualan); err != nil {
		tx.Rollback() // jika terjadi error, rollback
		return response, "", "", err
	}

	// insert detail penjualan
	for _, item := range req.Items {
		detilJual := model.PenjualanDetail{
			PenjualanID: penjualan.ID, // auto return id dari hasil insert data master penjualan
			BarangID:    item.BarangID,
			Size:        item.Size,
			Qty:         item.Qty,
			Harga:       item.Harga,
			Subtotal:    item.Qty * item.Harga,
		}

		// insert data detil penjualan
		if err := s.repo.CreatePenjualanDetail(tx, &detilJual); err != nil {
			tx.Rollback()
			return response, "", "", err
		}

		// -- SECTION update stok barang
		// ****TIDAK DIPAKAI LAGI KARENA UPDATE STOK DARI KARTU STOK ****//
		/*
			// validasi apakah barang ada di db
			barang, err := s.repoBarang.GetBarangById(item.BarangID)
			if err != nil {
				tx.Rollback()
				return dto.PenjualanResponse{}, err
			}

			// validasi jumlah stok
			if barang.Stok < item.Qty {
				tx.Rollback()
				return dto.PenjualanResponse{}, errors.New("Stok barang tidak cukup (ID:" + barang.Artikel.Nama + ")")
			}

			// jika stok cukup, kurangi stok barang aktif
			barang.Stok -= item.Qty
			if err := s.repoBarang.UpdateStok(tx, barang); err != nil {
				return dto.PenjualanResponse{}, err
			}
		*/

		// validasi apakah barang ada di db
		barang, err := s.repoBarang.GetBarangById(item.BarangID)
		if err != nil {
			tx.Rollback()
			return dto.PenjualanResponse{}, "", "", err
		}

		// -- SECTION update stok barang di kartu stok
		// validasi apakah barang ada di tabel kartu stok
		kartuStok, err := s.repoKartuStok.GetStokByBarangIdAndSize(item.BarangID, item.Size)
		if err != nil {
			tx.Rollback()
			return dto.PenjualanResponse{}, "", "", err
		}

		// validasi stok barang di kartu stok
		if kartuStok.Stok < item.Qty {
			tx.Rollback()
			return dto.PenjualanResponse{}, "", "", errors.New("Stok barang tidak cukup (ID:" + barang.Artikel.Nama + ")")
		}

		// jika stok cukup, kurangi stok barang aktif
		kartuStok.Stok -= item.Qty
		if err := s.repoKartuStok.UpdateStokDiKartuStok(tx, kartuStok); err != nil {
			return dto.PenjualanResponse{}, "", "", err
		}

	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return response, "", "", err
	}

	// convert dulu req detil penjualan ke dalam response detil penjualan
	itemsDTO := helper.ConvertItemRequestToItemResponse(req.Items)

	// get penjualan by id untuk memperoleh relasi dengan master toko untuk tampilkan nama toko, alamat, dan kotanya berikut diskon1,diskon2,diskon3
	newPenjualan, err := s.repo.GetPenjualanById(penjualan.ID)
	if err != nil {
		return dto.PenjualanResponse{}, "", "", err
	}
	fmt.Println("New penjualan: ", newPenjualan)

	// mapping model ke response dto
	response = dto.PenjualanResponse{
		ID:            penjualan.ID,
		NoFaktur:      penjualan.NoFaktur,
		TglPenjualan:  penjualan.TglPenjualan.Format("2006-01-02"),
		TglJatuhTempo: penjualan.TglJatuhTempo.Format("2006-01-02"),
		TokoID:        penjualan.TokoID,
		TokoNama:      newPenjualan.Toko.Nama,
		TokoAlamat:    newPenjualan.Toko.Alamat,
		TokoKota:      newPenjualan.Toko.Kota.Nama,
		Disc1:         newPenjualan.Toko.Disc1,
		Disc2:         newPenjualan.Toko.Disc2,
		Disc3:         newPenjualan.Toko.Disc3,
		SalesID:       penjualan.SalesID,
		SalesNama:     penjualan.Sales.Nama,
		Total:         penjualan.Total,
		TotalNetto:    penjualan.TotalNetto,
		Keterangan:    penjualan.Keterangan,
		Items:         itemsDTO,
	}

	// ==============================
	// Tambahkan bagian cetak faktur
	// ==============================
	// invoiceText := s.generateInvoiceText(newPenjualan)

	// // Simpan ke file (opsional)
	// fileName := fmt.Sprintf("faktur_%s.txt", penjualan.NoFaktur)
	// if err := os.WriteFile(fileName, []byte(invoiceText), 0644); err != nil {
	// 	fmt.Println("Gagal menulis faktur:", err)
	// }

	// percobaan pivot table
	fmt.Println("Test Pivot Faktur")
	// ambil data detil penjualan berdasarkan nomor faktur
	newDetilJual, err := s.repoDetilJual.GetPenjualanDetailById(penjualan.ID)

	if err != nil {
		panic(err)
	}

	// convert ke dto detil jual respons
	detilJualDTO := helper.ConvertToDTOPenjualanDetailPlural(newDetilJual)

	// debug (optional)
	// for _, detil := range detilJualDTO {
	// 	fmt.Println(detil)
	// }

	// 1. Ambil semua size unik
	sizeSet := map[int]bool{}
	for _, d := range detilJualDTO {
		sizeSet[d.Size] = true
	}

	sizes := make([]int, 0, len(sizeSet))
	for s := range sizeSet {
		sizes = append(sizes, s)
	}
	sort.Ints(sizes)

	// 2. Buat pivot: "Kode + warna" -> size -> qty dan map harga
	hargaMap := map[string]int{}
	pivot := map[string]map[int]int{}

	for _, d := range detilJualDTO {
		// key pivot: kode + warna
		key := fmt.Sprintf("%-7s %-13s %-10s", d.BarangKode, d.BarangArtikel, d.BarangWarna)

		if _, ok := pivot[key]; !ok {
			pivot[key] = map[int]int{}
		}
		// key = barang kode + warna
		pivot[key][d.Size] += d.Qty
		// simpan harga (ambil dari dto)
		hargaMap[key] = d.Harga
	}

	//=======OUTPUT CONSOLE=======
	// fmt.Println("===============================================================================")
	// // 3. Print pivot table
	// fmt.Printf("%-7s %-13s %-10s", "Kode", "Artikel", "Warna")
	// for _, s := range sizes {
	// 	fmt.Printf("%5d", s)
	// }

	// // header kolom total
	// fmt.Printf("%5s", "Qty")

	// // header kolom harga
	// fmt.Printf("%8s", "Harga")

	// // header kolom total
	// fmt.Printf("%9s", "Total")

	// fmt.Println("\n===============================================================================")

	// for key, row := range pivot {
	// 	fmt.Printf("%-20s", key)
	// 	// initial value total per row
	// 	total := 0

	// 	for _, s := range sizes {
	// 		val := row[s]
	// 		if val == 0 {
	// 			fmt.Printf("%5s", " ") // kosong jika 0
	// 		} else {
	// 			fmt.Printf("%5d", val) // qty sebenarnya
	// 			total += val
	// 		}
	// 	}
	// 	// tampilkan total di kolom terakhir
	// 	fmt.Printf("%5d", total)
	// 	fmt.Printf("%8d", hargaMap[key])       // harga dari dto
	// 	fmt.Printf("%9d", total*hargaMap[key]) // hitung total harga * qty

	// 	fmt.Println()
	// }
	// fmt.Println("===============================================================================")

	// OUTPUT FAKTUR FILE .TXT
	kategoriSizeRange := map[string][]int{
		"K": {26, 27, 28, 29, 30, 31},
		"A": {32, 33, 34, 35, 36, 37},
		"D": {36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47},
	}

	kategoriLabel := map[string]string{"K": "|K|", "A": "|A|", "D": "|D|"}

	totalSizes := len(kategoriSizeRange["K"]) + len(kategoriSizeRange["A"]) + len(kategoriSizeRange["D"])

	pivotMultiHeader := helper.BuildPivot(detilJualDTO, kategoriSizeRange)

	// ambil data brand dari slice detil jual pertama
	merk := detilJualDTO[0].BarangMerk

	totalQty := 0
	for _, d := range detilJualDTO {
		totalQty += d.Qty
	}

	// ==== Generate text faktur ====
	content := helper.PrintHeader(merk, response, kategoriSizeRange, kategoriLabel, totalSizes)
	content += helper.PrintData(pivotMultiHeader, kategoriSizeRange)
	content += helper.PrintSumarry(response, totalQty)

	// =======BUAT FILE FAKTUR.TXT=======
	// err = helper.WriteToFile("fakturMultiHeader.txt", output)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Berhasil membuat file fakturMultiHeader.txt")
	// }
	// end of faktur multi header experimental

	// =======LANGSUNG DOWNLOAD FAKTUR=======
	// Nama file
	fileName := fmt.Sprintf("Faktur-%s.txt", response.NoFaktur)

	return response, content, fileName, nil
}

// helper function untuk cetak faktur
// func (s *servicePenjualan) generateInvoiceText(penjualan model.Penjualan) string {
// 	var sb strings.Builder
// 	total := 0

// 	sb.WriteString(strings.Repeat("=", 80) + "\n")
// 	sb.WriteString(fmt.Sprintf("MITRA SUKSES BERSAMA%10sNama Toko: %s\n", "", penjualan.Toko.Nama))
// 	sb.WriteString(fmt.Sprintf("FAKTUR PENJUALAN%13s %s (%s)\n", "", penjualan.Toko.Alamat, penjualan.Toko.Kota.Nama))
// 	sb.WriteString(fmt.Sprintf("No.Faktur: %-10s         Tanggal Faktur: %s\n",
// 		penjualan.NoFaktur,
// 		penjualan.TglPenjualan.Format("02-01-2006")))
// 	sb.WriteString(fmt.Sprintf("%30sJatuh Tempo: %s\n", "", penjualan.TglJatuhTempo.Format("02-01-2006")))
// 	sb.WriteString(strings.Repeat("=", 80) + "\n")
// 	sb.WriteString("| ID / Artikel              | Warna        | Size | Harga    | Qty | Subtotal  |\n")
// 	sb.WriteString(strings.Repeat("-", 80) + "\n")

// 	//ambil data detil penjualan berdasarkan nomor faktur
// 	newDetilJual, err := s.repoDetilJual.GetPenjualanDetailById(penjualan.ID)
// 	// fmt.Println(newDetilJual)
// 	if err != nil {
// 		panic(err)
// 	}

// 	detilJualDTO := helper.ConvertToDTOPenjualanDetailPlural(newDetilJual)
// 	// fmt.Println(detilJualDTO)

// 	for _, item := range detilJualDTO {
// 		sub := item.Qty * item.Harga
// 		total += sub
// 		sb.WriteString(fmt.Sprintf("| %-5s/%-18s | %-12s | %4d | %8d | %3d | %9d |\n",
// 			item.BarangKode, item.BarangArtikel, item.BarangWarna, item.Size, item.Harga, item.Qty, item.Subtotal))
// 	}

// 	// perhitungan diskon 1,2, dan 3
// 	var grandTotal float64
// 	grandTotal = float64(total)
// 	fmt.Println("Total bruto:", grandTotal)
// 	diskon1 := penjualan.Toko.Disc1 * grandTotal
// 	fmt.Println("Diskon 1:", diskon1)
// 	grandTotal -= diskon1
// 	fmt.Println("Grand total after diskon 1:", grandTotal)
// 	diskon2 := penjualan.Toko.Disc2 * grandTotal
// 	fmt.Println("Diskon 2:", diskon2)
// 	grandTotal -= diskon2
// 	fmt.Println("Grand total after diskon 2:", grandTotal)
// 	diskon3 := penjualan.Toko.Disc3 * grandTotal
// 	fmt.Println("Diskon 3:", diskon3)
// 	grandTotal -= diskon3
// 	fmt.Println("Grand total after diskon 3:", grandTotal)
// 	//grandTotal = grandTotal - diskon3
// 	sb.WriteString(strings.Repeat("-", 80) + "\n")
// 	sb.WriteString(fmt.Sprintf("| %64s | %9d |\n", "Total", total))

// 	sb.WriteString(fmt.Sprintf("| %64s | %9.0f |\n", "Disc 1", diskon1)) // %9.0F ARTINYA pembualan ke bilangan bulat tanpa desimal
// 	sb.WriteString(fmt.Sprintf("| %64s | %9.0f |\n", "Disc 2", diskon2))
// 	sb.WriteString(fmt.Sprintf("| %64s | %9.0f |\n", "Disc 3", diskon3))

// 	sb.WriteString(fmt.Sprintf("| %64s | %9d |\n", "Grand Total", penjualan.TotalNetto))
// 	sb.WriteString(strings.Repeat("=", 80) + "\n")
// 	sb.WriteString(centerText("TERIMA KASIH ATAS PEMBELIAN ANDA", 80) + "\n")
// 	sb.WriteString(strings.Repeat("=", 80) + "\n")

// 	return sb.String()
// }

// func centerText(text string, width int) string {
// 	padding := (width - len(text)) / 2
// 	if padding < 0 {
// 		padding = 0
// 	}
// 	return strings.Repeat(" ", padding) + text
// }
