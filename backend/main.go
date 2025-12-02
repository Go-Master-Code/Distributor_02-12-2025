package main

import (
	"api-distributor/internal/database"
	"api-distributor/internal/handler"
	"api-distributor/internal/repository"
	"api-distributor/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB() // initDB + cek koneksi

	r := gin.Default()

	// Tambahkan ini
	r.Use(cors.Default()) // ⬅️ ini wajib agar Vue bisa akses API Go dari domain berbeda

	// dependency injection kota
	repoKota := repository.NewRepositoryKota(database.DB)
	serviceKota := service.NewServiceKota(repoKota)
	handlerKota := handler.NewHandlerKota(serviceKota)

	// list handler kota
	r.GET("/api/kota", handlerKota.GetKota) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/kota", handlerKota.CreateKota)

	// dependency injection artikel
	repoArtikel := repository.NewRepositoryArtikel(database.DB)
	serviceArtikel := service.NewServiceArtikel(repoArtikel)
	handlerArtikel := handler.NewHandlerArtikel(serviceArtikel)

	// list handler artikel
	r.GET("/api/artikel", handlerArtikel.GetArtikel) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/artikel", handlerArtikel.CreateArtikel)
	r.PUT("/api/artikel/:id", handlerArtikel.UpdateArtikel)

	// dependency injection jenis barang
	repoJenisBarang := repository.NewRepositoryJenisBarang(database.DB)
	serviceJenisBarang := service.NewServiceJenisBarang(repoJenisBarang)
	handlerJenisBarang := handler.NewHandlerJenisBarang(serviceJenisBarang)

	// list handler jenis barang
	r.GET("/api/jenis_barang", handlerJenisBarang.GetJenisBarang) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/jenis_barang", handlerJenisBarang.CreateJenisBarang)

	// dependency injection kategori barang
	repoKategoriBarang := repository.NewRepositoryKategoriBarang(database.DB)
	serviceKategoriBarang := service.NewServiceKategoriBarang(repoKategoriBarang)
	handlerKategoriBarang := handler.NewHandlerkategoriBarang(serviceKategoriBarang)

	// list handler kategori barang
	r.GET("/api/kategori_barang", handlerKategoriBarang.GetKategoriBarang) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/kategori_barang", handlerKategoriBarang.CreatekategoriBarang)

	// dependency injection merk
	repoMerk := repository.NewRepositoryMerk(database.DB)
	serviceMerk := service.NewServiceMerk(repoMerk)
	handlerMerk := handler.NewHandlerMerk(serviceMerk)

	// list handler merk
	r.GET("/api/merk", handlerMerk.GetMerk)
	r.POST("/api/merk", handlerMerk.CreateMerk)

	// dependency injection warna
	repoWarna := repository.NewRepositoryWarna(database.DB)
	serviceWarna := service.NewServiceWarna(repoWarna)
	handlerWarna := handler.NewHandlerWarna(serviceWarna)

	// list handler warna
	r.GET("/api/warna", handlerWarna.GetWarna) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/warna", handlerWarna.CreateWarna)

	// dependency injection ukuran
	repoUkuran := repository.NewRepositoryUkuran(database.DB)
	serviceUkuran := service.NewServiceUkuran(repoUkuran)
	handlerUkuran := handler.NewHandlerUkuran(serviceUkuran)

	// list handler merk
	r.GET("/api/ukuran", handlerUkuran.GetUkuran) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/ukuran", handlerUkuran.CreateUkuran)

	// dependency injection supplier
	repoSupplier := repository.NewRepositorySupplier(database.DB)
	serviceSupplier := service.NewServiceSupplier(repoSupplier)
	handlerSupplier := handler.NewHandlerSupplier(serviceSupplier)

	// list handler supplier
	r.GET("/api/supplier", handlerSupplier.GetAllSupplier)
	r.GET("/api/supplier/:id", handlerSupplier.GetSupplierByID)
	r.POST("/api/supplier", handlerSupplier.CreateSupplier)
	r.PUT("/api/supplier/:id", handlerSupplier.UpdateSupplier)
	r.DELETE("/api/supplier/:id", handlerSupplier.DeleteSupplier)

	// dependency injection kategori toko
	repoKategoriToko := repository.NewRepositoryKategoriToko(database.DB)
	serviceKategoriToko := service.NewServiceKategoriToko(repoKategoriToko)
	handlerKategoriToko := handler.NewHandlerKategoriToko(serviceKategoriToko)

	// list handler kategori barang
	r.GET("/api/kategori_toko", handlerKategoriToko.GetKategoriToko) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/kategori_toko", handlerKategoriToko.CreateKategoriToko)

	// dependency injection kategori toko
	repoEkspedisi := repository.NewRepositoryEkspedisi(database.DB)
	serviceEkspedisi := service.NewServiceEkspedisi(repoEkspedisi)
	handlerEkspedisi := handler.NewHandlerEkspedisi(serviceEkspedisi)

	// list handler kategori barang
	r.GET("/api/ekspedisi", handlerEkspedisi.GetEkspedisi) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/ekspedisi", handlerEkspedisi.CreateEkspedisi)

	// dependency injection kategori toko
	repoOngkir := repository.NewRepositoryOngkir(database.DB)
	serviceOngkir := service.NewServiceOngkir(repoOngkir)
	handlerOngkir := handler.NewHandlerOngkir(serviceOngkir)

	// list handler kategori barang
	r.GET("/api/ongkir", handlerOngkir.GetOngkir) // 3 case dlm 1 api pakai query parameter: get all, by id, atau lazy search by nama
	r.POST("/api/ongkir", handlerOngkir.CreateOngkir)

	// dependency injection barang
	repoBarang := repository.NewRepositoryBarang(database.DB)
	serviceBarang := service.NewServiceBarang(repoBarang)
	handlerBarang := handler.NewHandlerBarang(serviceBarang)

	// list handler barang
	r.GET("/api/barang", handlerBarang.GetBarang)
	r.POST("/api/barang", handlerBarang.CreateBarang)
	r.PUT("/api/barang/:id", handlerBarang.UpdateBarang)

	// dependency injection toko
	repoToko := repository.NewRepositoryToko(database.DB)
	serviceToko := service.NewServiceToko(repoToko)
	handlerToko := handler.NewHandlerToko(serviceToko)

	r.GET("/api/toko", handlerToko.GetToko) // 3 case dlm 1 api pakai query parameter: get all, by id, lazy search by nama
	r.POST("/api/toko", handlerToko.CreateToko)

	repoSales := repository.NewRepositorySales(database.DB)
	serviceSales := service.NewServiceSales(repoSales)
	handlerSales := handler.NewHandlerSales(serviceSales)

	r.GET("/api/sales", handlerSales.GetSales)
	r.POST("/api/sales", handlerSales.CreateSales)

	// dependency injection harga barang
	repoHarga := repository.NewRepositoryHargaBarang(database.DB)
	serviceHarga := service.NewServiceHargaBarang(repoHarga)
	handlerHarga := handler.NewHandlerHargaBarang(serviceHarga)

	r.GET("/api/harga", handlerHarga.GetHargaBarang)

	// dependency injection penjualan detail
	repoDetilJual := repository.NewRepositoryPenjualanDetail(database.DB)
	serviceDetilJual := service.NewServicePenjualanDetail(repoDetilJual)
	handlerDetilJual := handler.NewHandlerPenjualanDetail(serviceDetilJual)

	r.GET("/api/penjualan_detail/:id", handlerDetilJual.GetAllPenjualanDetail)

	// dependency injection kartu stok
	repoKartuStok := repository.NewRepositoryKartuStok(database.DB)
	serviceKartuStok := service.NewServiceKartuStok(repoKartuStok)
	handlerKartuStok := handler.NewHandlerKartuStok(serviceKartuStok)

	r.GET("/api/kartu_stok", handlerKartuStok.GetKartuStok) // 3 in one endpoint
	r.GET("/api/kartu_stok/:id", handlerKartuStok.GetKartuStokByIdBarang)

	// dependency injection penjualan
	repoPenjualan := repository.NewRepositoryPenjualan(database.DB)
	servicePenjualan := service.NewServicePenjualan(database.DB, repoPenjualan, repoKartuStok, repoBarang, repoDetilJual) // SEKALIAN UPDATE BARANG
	handlerPenjualan := handler.NewHandlerPenjualan(servicePenjualan)

	//r.GET("/api/penjualan", handlerPenjualan.GetAllPenjualan)
	r.POST("/api/penjualan", handlerPenjualan.CreatePenjualan)

	// handler untuk frontend
	r.Static("/css", "./frontend/css")
	r.Static("/js", "./frontend/js")
	r.Static("/assets", "./frontend/assets")
	r.Static("/pages", "./frontend/pages")

	// Semua route frontend diarahkan ke index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	// Route untuk ke sistem absensi
	r.GET("/main", func(c *gin.Context) {
		c.File("./frontend/main.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	// run server
	r.Run("localhost:3000")
}
