import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SupplierList from '../views/SupplierList.vue'
import SupplierView from '../views/SupplierView.vue'
import ArtikelList from '../views/ArtikelList.vue'
import ArtikelView from '../views/ArtikelView.vue'
import BarangList from '../views/BarangList.vue'
import BarangView from '../views/BarangView.vue'
import TokoList from '../views/TokoList.vue'
import TokoView from '../views/TokoView.vue'
import KartuStokList from '../views/KartuStokList.vue'
import KartuStokView from '../views/KartuStokView.vue'
import PenjualanView from '../views/PenjualanView.vue'
import ViewEditSupplier from '../views/ViewEditSupplier.vue'
import ViewEditArtikel from '../views/ViewEditArtikel.vue'
import ViewEditBarang from '../views/ViewEditBarang.vue'
import ViewEditToko from '../views/ViewEditToko.vue'

const routes = [
    { path: '/', name: 'Home', component: HomeView},
    { path: '/supplier', name: 'Suppliers', component: SupplierList}, // path yang akan diakses ketika running web
    { path: '/supplier/tambah', name: 'TambahSupplier', component: SupplierView}, // tambah supplier
    { path: '/supplier/edit/:id', name: 'EditSupplier', component: ViewEditSupplier}, // edit supplier

    { path: '/artikel', name: 'Artikel', component: ArtikelList}, // path yang akan diakses ketika running web
    { path: '/artikel/tambah', name: 'TambahArtikel', component: ArtikelView}, // tambah supplier
    { path: '/artikel/edit/:id', name: 'EditArtikel', component: ViewEditArtikel}, // edit supplier

    { path: '/barang', name: 'Barang', component: BarangList}, // path yang akan diakses ketika running web
    { path: '/barang/tambah', name: 'TambahBarang', component: BarangView}, // tambah supplier
    { path: '/barang/edit/:id', name: 'EditBarang', component: ViewEditBarang}, // edit supplier

    { path: '/toko', name: 'Toko', component: TokoList}, // path yang akan diakses ketika running web
    { path: '/toko/tambah', name: 'TambahToko', component: TokoView}, // tambah toko
    { path: '/toko/edit/:id', name: 'EditToko', component: ViewEditToko}, // edit toko

    { path: '/penjualan', name: 'Penjualan', component: PenjualanView}, // tambah penjualan

    { path: '/kartu_stok', name: 'KartuStok', component: KartuStokList}, // path yang akan diakses ketika running web
    { path: '/kartu_stok/tambah', name: 'TambahKartuStok', component: KartuStokView}, // tambah stok spesifik
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router