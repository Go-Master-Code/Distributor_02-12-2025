<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <label for="no_faktur" class="col-sm-2 col-form-label">Nomor Faktur:</label>
                <div class="col-sm-2">
                    <input id="no_faktur" v-model="form.no_faktur" type="text" class="form-control" autocomplete="off" maxlength="10" autofocus required />
                </div>
                <label for="tgl_penjualan" class="col-sm-2 col-form-label">Tanggal Faktur:</label>
                <div class="col-sm-2">
                    <input id="tgl_penjualan" v-model="form.tgl_penjualan" type="date" class="form-control" autocomplete="off" required />
                </div>
                <label for="tgl_jatuh_tempo" class="col-sm-2 col-form-label">Jatuh Tempo:</label>
                <div class="col-sm-2">
                    <input id="tgl_jatuh_tempo" v-model="form.tgl_jatuh_tempo" type="date" class="form-control" autocomplete="off" required />
                </div>
            </div>
            <!-- Multiselect toko -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="toko" class="col-sm-2 col-form-label">Toko:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="tokoRef"
                        id="toko"
                        v-model="selectedToko"
                        mode="single"
                        :options="tokoList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari toko..."
                        label="label"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchToko"
                        @select="onSelectToko"
                    >
                    <!-- <template #clear="{ clear }">
                        <button
                            type="button"
                            class="multiselect-clear"
                            @click="clear"
                            aria-label="Hapus pilihan"
                        >❌</button></template> -->
                    </Multiselect>
                </div>  
                <label for="disc_1" class="col-sm-1 col-form-label">Disc 1:</label>
                <div class="col-sm-1">
                    <input id="disc_1" v-model="form.disc_1" type="text" class="form-control" disabled />
                </div> 
                <label for="disc_2" class="col-sm-1 col-form-label">Disc 2:</label>
                <div class="col-sm-1">
                    <input id="disc_2" v-model="form.disc_2" type="text" class="form-control" disabled />
                </div> 
                <label for="disc_3" class="col-sm-1 col-form-label">Disc 3:</label>
                <div class="col-sm-1">
                    <input id="disc_3" v-model="form.disc_3" type="text" class="form-control" disabled />
                </div> 
            </div>

            <div class="row mb-3 align-items-center">
                <label for="toko" class="col-sm-2 col-form-label">Sales:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="salesRef"
                        id="sales"
                        v-model="selectedSales"
                        mode="single"
                        :options="salesList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari sales..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchSales"
                        @select="onSelectSales"
                    >
                    </Multiselect>
                </div>
                <label for="keterangan" class="col-sm-2 col-form-label">Keterangan:</label>
                <div class="col-sm-4">
                    <input id="keterangan" v-model="form.keterangan" type="text" class="form-control" autocomplete="off" />
                </div> 
            </div>
            
            <table class="table table-bordered align-middle">
                <thead class="table-light">
                    <tr>
                    <th style="width: 46%">Barang</th>
                    <th style="width: 8%">Size</th>
                    <th style="width: 8%">Stok</th>
                    <th style="width: 7%">Qty</th>
                    <th style="width: 11%">Harga</th>
                    <th style="width: 11%">Total</th>
                    <th style="width: 9%">Aksi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(item, index) in form.items" :key="index">
                    <td>
                        <div class="col-sm-12"> 
                            <Multiselect
                                ref="barangRef"
                                id="'barang-' + index"
                                v-model="item.barang_id"
                                mode="single"
                                :options="barangList"
                                :searchable="true"
                                :filter-results="false"
                                :clear-on-select="true"
                                :close-on-select="true"
                                :internal-search="false"
                                :loading="loading"
                                autocomplete="off"
                                placeholder="Ketik untuk mencari barang..."
                                label="label"
                                value-prop="id"
                                track-by="id"
                                @search-change="onSearchBarang"
                                @select="barang => onSelectBarang(barang, item)"
                                @clear="resetItemFields(item)"
                            />
                        </div> 
                    </td>
                    <td>
                        <select
                            v-model.number="item.size"
                            class="form-select text-end"
                            :disabled="!item.sizeList?.length"
                        >
                            <option value="" disabled>Pilih Size</option>
                            <option v-for="s in item.sizeList" :key="s.id" :value="s.size">
                                {{ s.size }}
                            </option>
                        </select>
                        
                    </td>
                    <td>
                        <input
                            type="number"
                            class="form-control text-end"
                            v-model.number="item.stok"
                            disabled
                        />
                    </td>
                    <td>
                        <input
                            type="number"
                            min="1"
                            :max="item.stok"
                            class="form-control text-end"
                            v-model.number="item.qty"
                            @input="hitungTotal(item)"
                        />
                    </td>
                    <td>
                        <select
                            v-model.number="item.harga"
                            class="form-select text-end"
                            :disabled="!item.hargaList?.length"
                        >
                        <option value="" disabled>Pilih Harga</option>
                        <option v-for="h in item.hargaList" :key="h.id" :value="h.harga">
                            {{ h.harga.toLocaleString('id-ID') }}
                        </option>
                        </select>
                    </td>
                    <!-- <td>
                        <input
                            type="number"
                            min="0"
                            class="form-control text-end"
                            v-model.number="item.harga"
                            @input="hitungTotal(item)"
                        />
                    </td> -->
                    <td class="text-end">
                        {{ item.total.toLocaleString('id-ID') || 0 }}
                    </td>
                    <td class="text-center">
                        <button type="button" class="btn btn-sm btn-danger" @click="hapusBaris(index)"><font-awesome-icon icon="trash" />Hapus</button>
                    </td>
                    </tr>
                </tbody>
            </table>

            <div class="d-flex justify-content-between align-items-start mb-3">
                <!-- Tombol kiri -->
                <div class="d-flex gap-2 align-items-center">
                    <button type="button" class="btn btn-sm btn-warning" @click="tambahBaris">
                        <font-awesome-icon icon="plus" /> Tambah Baris
                    </button>

                    <button type="submit" :disabled="submitting" class="btn btn-sm btn-primary">
                        <font-awesome-icon icon="floppy-disk" />
                        {{ submitting ? 'Menyimpan...' : 'Simpan Transaksi' }}
                    </button>
                    
                    <!-- Pesan sukses / error -->
                    <p v-if="errorMessage" class="text-danger mt-3">⚠️ {{ errorMessage }}</p>
                    <p v-if="successMessage" class="text-success mt-3">✔️ {{ successMessage }}</p>
                </div>

                <!-- Resume kanan -->
                <div class="d-flex flex-column text-end">
                    <div>Bruto: <strong>{{ totalPenjualanFormatted }}</strong></div>
                    <div>Diskon 1: <strong>{{ totalDiskon1Formatted }}</strong></div>
                    <div>Diskon 2: <strong>{{ totalDiskon2Formatted }}</strong></div>
                    <div>Diskon 3: <strong>{{ totalDiskon3Formatted }}</strong></div>
                    <div>Grand Total: <strong>{{ totalNetoPenjualanFormatted }}</strong></div>
                </div>
        </div>

            
        </form>
    </div>
</template>

<script setup>
    import { ref, reactive, computed, watch } from 'vue'
    import axios from '@/axios'
    import Multiselect from '@vueform/multiselect'
    import '@vueform/multiselect/themes/default.css'

    // Deklarasi form sebelum watcher
    const form = reactive({
        no_faktur: '',
        tgl_penjualan: '',
        tgl_jatuh_tempo: '',
        total: 0,
        total_netto: 0,
        keterangan: '',
        toko_id: null, // multiselect
        sales_id: null, // multiselect
        disc_1 : 0, // deklarasi field disc_1, 2, dan 3
        disc_2 : 0,
        disc_3 : 0,
        items: [
            { 
                barang: null,
                barang_id: null,
                sizeList: [], // ✅ penting!
                hargaList: [], // ✅ penting!
                size: 0,
                harga: 0,
                qty: 1,
                stok: 0,
                total: 0
            }
        ]
    })

    // Otomatis menghitung total tanpa panggilan manual.
    // Lebih mudah kalau nanti kamu menambah/hapus baris.
    watch(
        () => form.items,
        (items) => {
            items.forEach(item => {
                item.total = (item.qty || 0) * (item.harga || 0)
            })
        },
        { deep: true}
    )

    // watcher untuk stok barang berdasarkan id dan size
    watch(
        () => form.items.map(i => i.size),
        (newSizes, oldSizes) => {
            form.items.forEach((item, index) => {
                if (newSizes[index] !== oldSizes[index]) {
                    getStokByBarangAndSize(item)
                }
            })
        }
    )

    // deklarasi ref komponen multiselect
    const tokoRef = ref(null)
    const salesRef = ref(null)
    const barangRef = ref(null)

    // fungsi jika tombol x di multiselect barang diklik
    const resetItemFields = (item) => {
        item.size = null
        item.qty = 1
        item.harga = null
        item.total = 0
        item.sizeList = []  // Reset sizeList
        item.hargaList = [] // Reset hargaList
    }

    // function tambah baris, hapus baris, hitung total
    function tambahBaris() {
        form.items.push({ barang: '', qty: 1, harga: 0, total: 0 })
    }

    function hapusBaris(index) {
        form.items.splice(index, 1)
    }

    function hitungTotal(item) {
        const qty = Number(item.qty) || 0
        const stok = Number(item.stok) || 0

        // validasi qty terhadap stok
        if (qty > stok) {
            item.qty = stok // kunci kembali ke stok max
            alert("Qty tidak boleh melebihi stok tersedia!")
        }

        const harga = Number(item.harga) || 0
        item.total = qty * harga
    }

    const totalPenjualan = computed(() =>
        form.items.reduce((sum, i) => sum + i.total, 0)
    )

    // total penjualan (formatted) -> pakai formatter rupiah reusable
    const totalPenjualanFormatted = computed(() =>
        formatRupiah(totalPenjualan.value)
    )

    // formatter rupiah reusable
    const formatRupiah = (value) => {
        return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
            minimumFractionDigits: 0
        }).format(value || 0)
    }

    // diskon penjualan 1
    const totalDiskon1 = computed(() => {
        const bruto = totalPenjualan.value // ambil total penjualan sebagai bruto
        const persen1 = Number(form.disc_1) || 0
        var diskon1 = persen1/100 * bruto

        return Math.round(diskon1)
    })

    // dibuat agar field totalDiskon1Formatted (di elemen html) tidak warning
    const totalDiskon1Formatted = computed(() => {
        return formatRupiah(totalDiskon1.value)
    })

    // diskon penjualan 2
    const totalDiskon2 = computed(() => {
        const bruto = totalPenjualan.value - totalDiskon1.value // bruto - diskon1
        const persen2 = Number(form.disc_2) || 0
        var diskon2 = persen2/100 * bruto

        return Math.round(diskon2)
    })

    // dibuat agar field totalDiskon2Formatted (di elemen html) tidak warning
    const totalDiskon2Formatted = computed(() => {
        return formatRupiah(totalDiskon2.value)
    })

    // diskon penjualan 3
    const totalDiskon3 = computed(() => {
        const bruto = totalPenjualan.value - totalDiskon1.value - totalDiskon2.value// bruto - diskon1 - diskon2
        const persen3 = Number(form.disc_3) || 0
        var diskon3 = persen3/100 * bruto

        return Math.round(diskon3)
    })

    // dibuat agar field totalDiskon2Formatted (di elemen html) tidak warning
    const totalDiskon3Formatted = computed(() => {
        return formatRupiah(totalDiskon3.value)
    })

    // ======TOTAL DISKON TIDAK DITAMPILKAN======
    // const totalDiskonFormatted = computed(() => {
    //     const totalDiskon = totalDiskon1.value + totalDiskon2.value + totalDiskon3.value // harus pakai .value karena totalDiskon1, 2, dan 3 adalah computed
    //     return formatRupiah(totalDiskon)
    // })

    // computed neto numerik bulat
    const totalNetoPenjualan = computed(() => {
        const bruto = totalPenjualan.value
        const neto =
                bruto -
                totalDiskon1.value -
                totalDiskon2.value -
                totalDiskon3.value

        return Math.round(neto) // bulatkan ke satuan
    })

    // Computed neto formatted
    const totalNetoPenjualanFormatted = computed(() => {
        return formatRupiah(totalNetoPenjualan.value)
    })

    // State untuk multiselect toko
    const selectedToko = ref(null)
    const tokoList = ref([])
    const loading = ref(false) // deklarasi sekali saja

    const onSelectToko = (id) => { // ketika toko di select hanya response id saja 
        if (!id) { // jika id kosong
            form.disc_1 = 0
            form.disc_2 = 0
            form.disc_3 = 0
            return
        }

        // watcher untuk handle reset multiselect
        watch(selectedToko, (newVal) => {
            if (!newVal) { // kalo di reset (multiselect kosong) => kosongkan semua diskon
                form.disc_1 = 0
                form.disc_2 = 0
                form.disc_3 = 0
            }
        })

        selectedToko.value = id
        console.log('Toko dipilih:', id)

        const toko = tokoList.value.find(t => t.id === id) // compare it terhadap data list toko (untuk dapat reponse toko lengkap) kemudian masukkan ke var toko
        if (!toko) return

        // ubah diskon 1,2, dan 3 sesuai data pada db toko
        form.disc_1 = Number(toko.disc_1)*100 || 0 // validasi angka
        form.disc_2 = Number(toko.disc_2)*100 || 0
        form.disc_3 = Number(toko.disc_3)*100 || 0
    }

    // State untuk multiselect barang di detil penjualan
    const barangList = ref([])

    const onSelectBarang = async (barang, item) => {
        console.log("Barang dipilih:", barang) // ini akan menampilkan angka ID

        const barangId = item.barang_id || barang

        if (!barangId) {
            console.warn('Barang ID tidak ditemukan:', barang)
            return
        }

        // Reset data item terlebih dahulu
        resetItemFields(item)

        try { // blok try untuk harga
            loading.value = true

            // ✅ sesuaikan dengan API kamu
            const res = await axios.get(`/api/harga?id=${barangId}`)
            
            // asumsi struktur { code: 200, data: [ { id, harga, mulai_berlaku } ] }
            item.hargaList = res.data.data || [] // isikan data ke dalam list

            if (item.hargaList.length === 1) {
                // otomatis pilih jika hanya satu harga
                item.harga = item.hargaList[0].harga
            } else {
                // reset harga kalau ada beberapa pilihan
                item.harga = ''
            }
        } catch (err) {
            console.error('Gagal ambil harga:', err)
            item.hargaList = []
        } finally {
            loading.value = false
        }

        try { // blok try untuk size
            // =====TAMPILKAN SIZE BARANG DARI KARTU STOK=====
            const res2 = await axios.get(`/api/kartu_stok/${barangId}`)
            // asumsi struktur { code: 200, data: [ { id, barang_id, barang_kode, barang_merk } ] }
            item.sizeList = res2.data.data || [] // isikan data dari endpoint ke dalam list

            // jika data hanya 1, auto select row pertama
            if (item.sizeList.length === 1) {
                // otomatis pilih jika hanya satu harga
                item.size = item.sizeList[0].size
            } else {
                // reset harga kalau ada beberapa pilihan
                item.size = ''
            }
        } catch (err) {
            console.error('Gagal ambil size:', err)
            item.sizeList = []
        } finally {
            loading.value = false
        }
    }

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    // ===== FUNGSI CARI TOKO (DEBOUNCE) =====
    let searchTokoTimeout = null // pisahkan timeout toko dan barang
        const onSearchToko = (query) => {
        clearTimeout(searchTokoTimeout)
        searchTokoTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                tokoList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/toko?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    // Tambahkan properti label gabungan (misalnya kode dan nama toko)
                    tokoList.value = data.map(t => ({
                    ...t,
                    label: `${t.kode} : ${t.nama}` // gabungkan kode & nama toko
                    }))

                    console.log('Toko hasil pencarian (dengan label):', tokoList.value)
                } else {
                    tokoList.value = []
                }
            } catch (err) {
                console.error('Gagal load toko:', err)
                tokoList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI BARANG (DEBOUNCE) =====
    //let searchTimeout = null // deklarasi sekali saja
    let searchBarangTimeout = null // pisahkan timeout toko dan barang
    const onSearchBarang = (query) => {
        clearTimeout(searchBarangTimeout)
        searchBarangTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                barangList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/barang?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    // Tambahkan properti label gabungan (misalnya kode, artikel, dan warna)
                    barangList.value = data.map(b => ({
                    ...b,
                    label: `${b.kode} : ${b.merk_nama} | ${b.artikel_nama} | ${b.warna_nama} | ${b.ukuran_nama}` // gabungkan kode & nama toko
                    }))

                    console.log('Barang hasil pencarian:', data)
                } else {
                    barangList.value = []
                }
            } catch (err) {
                console.error('Gagal load barang:', err)
                barangList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // State untuk multiselect sales
    const selectedSales = ref(null)
    const salesList = ref([])
    // const loading = ref(false)

    const onSelectSales = (id) => {
        selectedSales.value = id
        console.log('Sales dipilih:', id)
    }

    // fungsi untuk mengambil stok barang berdasarkan barang dan size yang sudah dipilih
    const getStokByBarangAndSize = async (item) => {
        if (!item.barang_id || !item.size) return

        try {
            const res = await axios.get(
                `/api/kartu_stok?id=${item.barang_id}&size=${item.size}`
            )
            item.stok = res.data.data?.stok || 0
        } catch (err) {
            console.error("Gagal mengambil stok:", err)
            item.stok = 0
        }
    }

    // ===== FUNGSI CARI SALES (DEBOUNCE) =====
    //let searchTimeout = null // deklarasi sekali saja
    let searchSalesTimeout = null
        const onSearchSales = (query) => {
        clearTimeout(searchSalesTimeout)
        searchSalesTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                salesList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/sales?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    salesList.value = data
                    console.log('Sales hasil pencarian:', data)
                } else {
                    salesList.value = []
                }
            } catch (err) {
                console.error('Gagal load sales:', err)
                salesList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // fungsi resetForm dipanggil setelah penjualan berhasil diinput
    function resetForm() {
        form.no_faktur = ''
        form.tgl_penjualan = null
        form.tgl_jatuh_tempo = null
        form.total = 0
        form.keterangan = ''
        selectedToko.value = null
        selectedSales.value = null

        form.items = [
            {
                barang: null,
                barang_id: null,
                hargaList: [],
                sizeList: [],
                harga: 0,
                qty: 1,
                total: 0
            }
        ]
    }

    // ===== FUNGSI SUBMIT =====
    const handleSubmit = async () => {

        form.toko_id = selectedToko.value
        form.sales_id = selectedSales.value
        form.total = totalPenjualan.value
        form.total_netto = totalNetoPenjualan.value

        successMessage.value = ''
        errorMessage.value = ''

        if (!selectedToko.value) {
            errorMessage.value = 'Toko wajib dipilih.'
            tokoRef.value?.focus()
            return
        }

        if (!selectedSales.value) {
            errorMessage.value = 'Sales wajib dipilih.'
            salesRef.value?.focus()
            return
        }

        if (!form.items.length) {
            errorMessage.value = 'Harus ada minimal satu barang.'
            return
        }

        submitting.value = true

        try {
            const res = await axios.post(
                "/api/penjualan",
                form,
                { responseType: "blob" }   // ⬅⬅ PENTING!
            )

            // ---- Download file ----
            const blob = new Blob([res.data], { type: "text/plain" })
            console.log("Blob size:", blob.size)

            const url = URL.createObjectURL(blob)
            console.log("Object URL:",url)

            const a = document.createElement("a")
            a.href = url

            // Ambil nama file dari header Content-Disposition
            const cd = res.headers["content-disposition"] || res.headers["Content-Disposition"]
            console.log("Content-Disposition header:",cd)
            
            let fileName = "faktur.txt"

            if (cd) {
                const match = cd.match(/filename="?(.*?)"?$/)
                if (match) fileName = match[1]
            }
            console.log("Nama file hasil parsing:",fileName)

            a.download = fileName
            console.log("Atribut download:", a.download)
            a.click()

            URL.revokeObjectURL(url)

            // reset form
            successMessage.value = "Penjualan berhasil disimpan!"
            resetForm()

        } catch (err) {
            console.error("Gagal simpan:", err)

            if (err.response) {
                // backend mengembalikan error JSON dalam blob
                const reader = new FileReader()
                reader.onload = () => {
                    try {
                        const parsed = JSON.parse(reader.result)
                        errorMessage.value = parsed.message || "Gagal menyimpan penjualan."
                    } catch {
                        errorMessage.value = "Gagal menyimpan penjualan."
                    }
                }
                reader.readAsText(err.response.data)
            } else {
                errorMessage.value = "Network error."
            }
        } finally {
            submitting.value = false
        }
    }

</script>

<style scoped>
    /* Opsional: sedikit animasi untuk loading */
    .multiselect.is-loading .multiselect-spinner {
        border-top-color: #3b82f6; /* warna biru Tailwind */
    }
</style>