package helper

import (
	"api-distributor/internal/dto"
	"api-distributor/internal/model"
)

func ConvertToDTOKotaPlural(kota []model.Kota) []dto.KotaResponse {
	var kotaDTO []dto.KotaResponse
	for _, k := range kota {
		kotaDTO = append(kotaDTO, dto.KotaResponse{
			ID:   k.ID,
			Nama: k.Nama,
		})
	}
	return kotaDTO
}

func ConvertToDTOKotaSingle(kota model.Kota) dto.KotaResponse {
	var kotaDTO dto.KotaResponse
	kotaDTO.ID = kota.ID
	kotaDTO.Nama = kota.Nama
	return kotaDTO
}

func ConvertToDTOAreaPlural(area []model.Area) []dto.AreaResponse {
	var areaDTO []dto.AreaResponse
	for _, a := range area {
		areaDTO = append(areaDTO, dto.AreaResponse{
			ID:   a.ID,
			Nama: a.Nama,
		})
	}
	return areaDTO
}

func ConvertToDTOAreaSingle(area model.Area) dto.AreaResponse {
	var areaDTO dto.AreaResponse
	areaDTO.ID = area.ID
	areaDTO.Nama = area.Nama
	return areaDTO
}

func ConvertToDTOArtikelPlural(artikel []model.Artikel) []dto.ArtikelResponse {
	var artikelDTO []dto.ArtikelResponse
	for _, a := range artikel {
		artikelDTO = append(artikelDTO, dto.ArtikelResponse{
			ID:   a.ID,
			Nama: a.Nama,
		})
	}
	return artikelDTO
}

func ConvertToDTOArtikelSingle(artikel model.Artikel) dto.ArtikelResponse {
	var artikelDTO dto.ArtikelResponse
	artikelDTO.ID = artikel.ID
	artikelDTO.Nama = artikel.Nama
	return artikelDTO
}

func ConvertToDTOJenisBarangPlural(jenisBarang []model.JenisBarang) []dto.JenisBarangResponse {
	var jenisBarangDTO []dto.JenisBarangResponse
	for _, jb := range jenisBarang {
		jenisBarangDTO = append(jenisBarangDTO, dto.JenisBarangResponse{
			ID:   jb.ID,
			Nama: jb.Nama,
		})
	}
	return jenisBarangDTO
}

func ConvertToDTOJenisBarangSingle(jenisBarang model.JenisBarang) dto.JenisBarangResponse {
	var jenisBarangDTO dto.JenisBarangResponse
	jenisBarangDTO.ID = jenisBarang.ID
	jenisBarangDTO.Nama = jenisBarang.Nama
	return jenisBarangDTO
}

func ConvertToDTOKategoriBarangPlural(kategori []model.KategoriBarang) []dto.KategoriBarangResponse {
	var kategoriDTO []dto.KategoriBarangResponse
	for _, k := range kategori {
		kategoriDTO = append(kategoriDTO, dto.KategoriBarangResponse{
			ID:    k.ID,
			Huruf: k.Huruf,
			Nama:  k.Nama,
		})
	}
	return kategoriDTO
}

func ConvertToDTOKategoriBarangSingle(kategori model.KategoriBarang) dto.KategoriBarangResponse {
	var kategoriDTO dto.KategoriBarangResponse
	kategoriDTO.ID = kategori.ID
	kategoriDTO.Huruf = kategori.Huruf
	kategoriDTO.Nama = kategori.Nama
	return kategoriDTO
}

func ConvertToDTOKategoriTokoPlural(kategoriToko []model.KategoriToko) []dto.KategoriTokoResponse {
	var kategoriTokoDTO []dto.KategoriTokoResponse
	for _, k := range kategoriToko {
		kategoriTokoDTO = append(kategoriTokoDTO, dto.KategoriTokoResponse{
			ID:   k.ID,
			Nama: k.Nama,
		})
	}
	return kategoriTokoDTO
}

func ConvertToDTOKategoriTokoSingle(kategoriToko model.KategoriToko) dto.KategoriTokoResponse {
	var kategoriTokoDTO dto.KategoriTokoResponse
	kategoriTokoDTO.ID = kategoriToko.ID
	kategoriTokoDTO.Nama = kategoriToko.Nama
	return kategoriTokoDTO
}

func ConvertToDTOMerkPlural(merk []model.Merk) []dto.MerkResponse {
	var merkDTO []dto.MerkResponse
	for _, m := range merk {
		merkDTO = append(merkDTO, dto.MerkResponse{
			ID:   m.ID,
			Nama: m.Nama,
		})
	}
	return merkDTO
}

func ConvertToDTOMerkSingle(merk model.Merk) dto.MerkResponse {
	var merkDTO dto.MerkResponse
	merkDTO.ID = merk.ID
	merkDTO.Nama = merk.Nama
	return merkDTO
}

func ConvertToDTOSupplierPlural(supplier []model.Supplier) []dto.SupplierResponse {
	var supplierDTO []dto.SupplierResponse
	for _, s := range supplier {
		supplierDTO = append(supplierDTO, dto.SupplierResponse{
			ID:     s.ID,
			Nama:   s.Nama,
			Alamat: s.Alamat,
			Kontak: s.Kontak,
		})
	}
	return supplierDTO
}

func ConvertToDTOSupplierSingle(supplier model.Supplier) dto.SupplierResponse {
	var supplierDTO dto.SupplierResponse
	supplierDTO.ID = supplier.ID
	supplierDTO.Nama = supplier.Nama
	supplierDTO.Alamat = supplier.Alamat
	supplierDTO.Kontak = supplier.Kontak
	return supplierDTO
}

func ConvertToDTOWarnaPlural(warna []model.Warna) []dto.WarnaResponse {
	var warnaDTO []dto.WarnaResponse
	for _, w := range warna {
		warnaDTO = append(warnaDTO, dto.WarnaResponse{
			ID:   w.ID,
			Nama: w.Nama,
		})
	}
	return warnaDTO
}

func ConvertToDTOWarnaSingle(warna model.Warna) dto.WarnaResponse {
	var warnaDTO dto.WarnaResponse
	warnaDTO.ID = warna.ID
	warnaDTO.Nama = warna.Nama
	return warnaDTO
}

func ConvertToDTOUkuranPlural(ukuran []model.Ukuran) []dto.UkuranResponse {
	var ukuranDTO []dto.UkuranResponse
	for _, w := range ukuran {
		ukuranDTO = append(ukuranDTO, dto.UkuranResponse{
			ID:   w.ID,
			Nama: w.Nama,
		})
	}
	return ukuranDTO
}

func ConvertToDTOUkuranSingle(ukuran model.Ukuran) dto.UkuranResponse {
	var ukuranDTO dto.UkuranResponse
	ukuranDTO.ID = ukuran.ID
	ukuranDTO.Nama = ukuran.Nama
	return ukuranDTO
}

func ConvertToDTOEkspedisiPlural(ekspedisi []model.Ekspedisi) []dto.EkspedisiResponse {
	var ekspedisiDTO []dto.EkspedisiResponse
	for _, e := range ekspedisi {
		ekspedisiDTO = append(ekspedisiDTO, dto.EkspedisiResponse{
			ID:   e.ID,
			Nama: e.Nama,
		})
	}
	return ekspedisiDTO
}

func ConvertToDTOEkspedisiSingle(ekspedisi model.Ekspedisi) dto.EkspedisiResponse {
	var ekspedisiDTO dto.EkspedisiResponse
	ekspedisiDTO.ID = ekspedisi.ID
	ekspedisiDTO.Nama = ekspedisi.Nama
	return ekspedisiDTO
}

func ConvertToDTOOngkirPlural(ongkir []model.Ongkir) []dto.OngkirResponse {
	var ongkirDTO []dto.OngkirResponse
	for _, o := range ongkir {
		ongkirDTO = append(ongkirDTO, dto.OngkirResponse{
			ID:   o.ID,
			Nama: o.Nama,
		})
	}
	return ongkirDTO
}

func ConvertToDTOOngkirSingle(ongkir model.Ongkir) dto.OngkirResponse {
	var ongkirDTO dto.OngkirResponse
	ongkirDTO.ID = ongkir.ID
	ongkirDTO.Nama = ongkir.Nama
	return ongkirDTO
}

func ConvertToDTOBarangPlural(barang []model.Barang) []dto.BarangResponse {
	var barangDTO []dto.BarangResponse
	for _, b := range barang {
		barangDTO = append(barangDTO, dto.BarangResponse{
			ID:                 b.ID,
			Kode:               b.Kode,
			MerkID:             b.MerkID,
			MerkNama:           b.Merk.Nama,
			ArtikelID:          b.ArtikelID,
			ArtikelNama:        b.Artikel.Nama,
			WarnaID:            b.WarnaID,
			WarnaNama:          b.Warna.Nama,
			KategoriBarangID:   b.KategoriBarangID,
			KategoriBarangNama: b.KategoriBarang.Nama,
			JenisBarangID:      b.JenisBarangID,
			JenisBarangNama:    b.JenisBarang.Nama,
			UkuranID:           b.UkuranID,
			UkuranNama:         b.Ukuran.Nama,
		})
	}
	return barangDTO
}

func ConvertToDTOBarangSingle(barang model.Barang) dto.BarangResponse {
	var barangDTO dto.BarangResponse
	barangDTO.ID = barang.ID
	barangDTO.Kode = barang.Kode
	barangDTO.ArtikelID = barang.ArtikelID
	barangDTO.ArtikelNama = barang.Artikel.Nama
	barangDTO.MerkID = barang.MerkID
	barangDTO.MerkNama = barang.Merk.Nama
	barangDTO.WarnaID = barang.WarnaID
	barangDTO.WarnaNama = barang.Warna.Nama
	barangDTO.KategoriBarangID = barang.KategoriBarangID
	barangDTO.KategoriBarangNama = barang.KategoriBarang.Nama
	barangDTO.JenisBarangID = barang.JenisBarangID
	barangDTO.JenisBarangNama = barang.JenisBarang.Nama
	barangDTO.UkuranID = barang.UkuranID
	barangDTO.UkuranNama = barang.Ukuran.Nama

	return barangDTO
}

func ConvertToDTOTokoPlural(toko []model.Toko) []dto.TokoResponse {
	var tokoDTO []dto.TokoResponse
	for _, t := range toko {
		tokoDTO = append(tokoDTO, dto.TokoResponse{
			ID:               t.ID,
			Kode:             t.Kode,
			Nama:             t.Nama,
			KategoriTokoID:   t.KategoriTokoID,
			KategoriTokoNama: t.KategoriToko.Nama,
			KotaID:           t.KotaID,
			KotaNama:         t.Kota.Nama,
			AreaID:           t.AreaID,
			AreaNama:         t.Area.Nama,
			Alamat:           t.Alamat,
			Disc1:            t.Disc1,
			Disc2:            t.Disc2,
			Disc3:            t.Disc3,
			EkspedisiID:      t.EkspedisiID,
			EkspedisiNama:    t.Ekspedisi.Nama,
			OngkirID:         t.OngkirID,
			OngkirNama:       t.Ongkir.Nama,
		})
	}
	return tokoDTO
}

func ConvertToDTOTokoSingle(toko model.Toko) dto.TokoResponse {
	var tokoDTO dto.TokoResponse
	tokoDTO.ID = toko.ID
	tokoDTO.Kode = toko.Kode
	tokoDTO.Nama = toko.Nama
	tokoDTO.KategoriTokoID = toko.KategoriTokoID
	tokoDTO.KategoriTokoNama = toko.KategoriToko.Nama
	tokoDTO.KotaID = toko.KotaID
	tokoDTO.KotaNama = toko.Kota.Nama
	tokoDTO.AreaID = toko.AreaID
	tokoDTO.AreaNama = toko.Area.Nama
	tokoDTO.Alamat = toko.Alamat
	tokoDTO.Disc1 = toko.Disc1
	tokoDTO.Disc2 = toko.Disc2
	tokoDTO.Disc3 = toko.Disc3
	tokoDTO.EkspedisiID = toko.EkspedisiID
	tokoDTO.EkspedisiNama = toko.Ekspedisi.Nama
	tokoDTO.OngkirID = toko.OngkirID
	tokoDTO.OngkirNama = toko.Ongkir.Nama

	return tokoDTO
}

func ConvertToDTOPenjualanSingle(penjualan model.Penjualan) dto.PenjualanResponse {
	var jualDTO dto.PenjualanResponse
	jualDTO.ID = penjualan.ID
	jualDTO.NoFaktur = penjualan.NoFaktur
	jualDTO.TglPenjualan = penjualan.TglPenjualan.Format("2006-01-02")
	jualDTO.TglJatuhTempo = penjualan.TglJatuhTempo.Format("2006-01-02")
	jualDTO.TokoID = penjualan.TokoID
	jualDTO.TokoNama = penjualan.Toko.Nama
	jualDTO.TokoAlamat = penjualan.Toko.Alamat
	jualDTO.TokoKota = penjualan.Toko.Kota.Nama
	jualDTO.Disc1 = penjualan.Toko.Disc1
	jualDTO.Disc2 = penjualan.Toko.Disc2
	jualDTO.Disc3 = penjualan.Toko.Disc3
	jualDTO.SalesID = penjualan.SalesID
	jualDTO.SalesNama = penjualan.Sales.Nama
	jualDTO.Total = penjualan.Total
	jualDTO.TotalNetto = penjualan.TotalNetto
	jualDTO.Keterangan = penjualan.Keterangan
	return jualDTO
}

func ConvertToDTOPenjualanPlural(penjualan []model.Penjualan) []dto.PenjualanResponse {
	var jualDTO []dto.PenjualanResponse
	for _, p := range penjualan {
		jualDTO = append(jualDTO, dto.PenjualanResponse{
			ID:            p.ID,
			NoFaktur:      p.NoFaktur,
			TglPenjualan:  p.TglPenjualan.Format("2006-01-02"),
			TglJatuhTempo: p.TglJatuhTempo.Format("2006-01-02"),
			TokoID:        p.TokoID,
			TokoNama:      p.Toko.Nama,
			TokoAlamat:    p.Toko.Alamat,
			TokoKota:      p.Toko.Kota.Nama,
			Disc1:         p.Toko.Disc1,
			Disc2:         p.Toko.Disc2,
			Disc3:         p.Toko.Disc3,
			SalesID:       p.SalesID,
			SalesNama:     p.Sales.Nama,
			Total:         p.Total,
			TotalNetto:    p.TotalNetto,
			Keterangan:    p.Keterangan,
		})
	}
	return jualDTO
}

func ConvertToDTOPenjualanDetailSingle(dj model.PenjualanDetail) dto.PenjualanDetailResponse {
	var djDTO dto.PenjualanDetailResponse
	djDTO.ID = dj.ID
	// djDTO.PenjualanID = dj.PenjualanID
	djDTO.BarangID = dj.BarangID
	djDTO.BarangArtikel = dj.Barang.Artikel.Nama
	djDTO.BarangWarna = dj.Barang.Warna.Nama
	djDTO.BarangMerk = dj.Barang.Merk.Nama
	djDTO.BarangUkuran = dj.Barang.Ukuran.Nama
	djDTO.BarangKategori = dj.Barang.KategoriBarang.Huruf
	djDTO.Size = dj.Size
	djDTO.Qty = dj.Qty
	djDTO.Harga = dj.Harga
	djDTO.Subtotal = dj.Subtotal
	return djDTO
}

func ConvertToDTOPenjualanDetailPlural(dj []model.PenjualanDetail) []dto.PenjualanDetailResponse {
	var djDTO []dto.PenjualanDetailResponse
	for _, data := range dj {
		djDTO = append(djDTO, dto.PenjualanDetailResponse{
			ID:             data.ID,
			PenjualanID:    data.PenjualanID,
			BarangID:       data.BarangID,
			BarangKode:     data.Barang.Kode,
			BarangArtikel:  data.Barang.Artikel.Nama,
			BarangWarna:    data.Barang.Warna.Nama,
			BarangMerk:     data.Barang.Merk.Nama,
			BarangUkuran:   data.Barang.Ukuran.Nama,
			BarangKategori: data.Barang.KategoriBarang.Huruf,
			Size:           data.Size,
			Qty:            data.Qty,
			Harga:          data.Harga,
			Subtotal:       data.Subtotal,
		})
	}
	return djDTO
}

func ConvertToDTOHargaBarangPlural(hb []model.HargaBarang) []dto.HargaBarangResponse {
	var hbDTO []dto.HargaBarangResponse
	for _, h := range hb {
		hbDTO = append(hbDTO, dto.HargaBarangResponse{
			ID:           h.ID,
			BarangID:     h.BarangID,
			BarangKode:   h.Barang.Kode,
			Harga:        h.Harga,
			MulaiBerlaku: h.MulaiBerlaku.Format("2006-01-02"),
		})
	}
	return hbDTO
}

// helper convert item dari create request detil penjualan menjadi response detil penjualan
func ConvertItemRequestToItemResponse(req []dto.CreatePenjualanDetailRequest) []dto.PenjualanDetailResponse {
	var detilJualDTO []dto.PenjualanDetailResponse
	for _, r := range req {
		detilJualDTO = append(detilJualDTO, dto.PenjualanDetailResponse{
			//PenjualanID: r.PenjualanID,
			BarangID: r.BarangID,
			Qty:      r.Qty,
			Harga:    r.Harga,
			Subtotal: r.Qty * r.Harga,
		})
	}
	return detilJualDTO
}

func ConvertToDTOKartuStokSingle(ks model.KartuStok) dto.KartuStokResponse {
	var ksDTO dto.KartuStokResponse
	ksDTO.ID = ks.ID
	ksDTO.BarangID = ks.BarangID
	ksDTO.BarangKode = ks.Barang.Kode
	ksDTO.BarangMerk = ks.Barang.Merk.Nama
	ksDTO.BarangWarna = ks.Barang.Warna.Nama
	ksDTO.Size = ks.Size
	ksDTO.Stok = ks.Stok
	return ksDTO
}

func ConvertToDTOKartuStokPlural(ks []model.KartuStok) []dto.KartuStokResponse {
	var ksDTO []dto.KartuStokResponse
	for _, row := range ks {
		ksDTO = append(ksDTO, dto.KartuStokResponse{
			ID:          row.ID,
			BarangID:    row.BarangID,
			BarangKode:  row.Barang.Kode,
			BarangMerk:  row.Barang.Merk.Nama,
			BarangWarna: row.Barang.Warna.Nama,
			Size:        row.Size,
			Stok:        row.Stok,
		})
	}
	return ksDTO
}

func ConvertToDTOSalesPlural(sales []model.Sales) []dto.SalesResponse {
	var salesDTO []dto.SalesResponse
	for _, s := range sales {
		salesDTO = append(salesDTO, dto.SalesResponse{
			ID:   s.ID,
			Nama: s.Nama,
		})
	}
	return salesDTO
}

func ConvertToDTOSalesSingle(sales model.Sales) dto.SalesResponse {
	var salesDTO dto.SalesResponse
	salesDTO.ID = sales.ID
	salesDTO.Nama = sales.Nama
	return salesDTO
}
