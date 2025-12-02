<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
        
            <div class="row mb-3 align-items-center">
                <label for="id" class="col-sm-2 col-form-label">ID Toko:</label>
                <div class="col-sm-2">
                    <input id="id" v-model="form.id" type="text" class="form-control" autocomplete="off" disabled required />
                </div>
                
            </div>

            <div class="row mb-3 align-items-center">
                <label for="kode" class="col-sm-2 col-form-label">Kode Toko:</label>
                <div class="col-sm-4">
                    <input id="kode" v-model="form.kode" type="text" class="form-control" autocomplete="off" autofocus required />
                </div>
                <label for="nama" class="col-sm-2 col-form-label">Nama Toko:</label>
                <div class="col-sm-4">
                    <input id="nama" v-model="form.nama" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <!-- Multiselect kategori toko -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="kota" class="col-sm-2 col-form-label">Kota:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="kotaRef"
                        id="kota"
                        v-model="selectedKota"
                        mode="single"
                        :options="kotaList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari kota..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchKota"
                        @select="onSelectKota"
                    />
                </div> 

                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="kategori" class="col-sm-2 col-form-label">Kategori:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="kategoriRef"
                        id="kategori"
                        v-model="selectedKategori"
                        mode="single"
                        :options="kategoriList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari kategori..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchKategori"
                        @select="onSelectKategori"
                    />
                </div>   
            </div>

            <div class="row mb-3 align-items-center">
                <label for="nama" class="col-sm-2 col-form-label">Alamat:</label>
                <div class="col-sm-10">
                    <input id="alamat" v-model="form.alamat" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <div class="row mb-3 align-items-center">
                <label for="disc_1" class="col-sm-2 col-form-label">Disc 1 (%):</label>
                <div class="col-sm-2">
                    <input id="disc_1" v-model.number="form.disc_1" type="number" class="form-control" autocomplete="off" required />
                </div>
                <label for="disc_2" class="col-sm-2 col-form-label">Disc 2 (%):</label>
                <div class="col-sm-2">
                    <input id="disc_2" v-model.number="form.disc_2" type="number" class="form-control" autocomplete="off" required />
                </div>
                <label for="disc_3" class="col-sm-2 col-form-label">Disc 3 (%):</label>
                <div class="col-sm-2">
                    <input id="disc_3" v-model.number="form.disc_3" type="number" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <!-- Multiselect ekspedisi -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="ekspedisi" class="col-sm-2 col-form-label">Ekspedisi:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="ekspedisiRef"
                        id="ekspedisi"
                        v-model="selectedEkspedisi"
                        mode="single"
                        :options="ekspedisiList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari ekspedisi..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchEkspedisi"
                        @select="onSelectEkspedisi"
                    />
                </div>
                
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="ongkir" class="col-sm-2 col-form-label">Ongkir:</label> 
                <div class="col-sm-4"> 
                    <Multiselect
                        ref="ongkirRef"
                        id="ongkir"
                        v-model="selectedOngkir"
                        mode="single"
                        :options="ongkirList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari jenis ongkir..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchOngkir"
                        @select="onSelectOngkir"
                    />
                </div> 
            </div>

        <button type="submit" class="btn btn-primary" :disabled="loading">
            Simpan Perubahan
        </button>

        <!-- Pesan sukses / error -->
        <p v-if="errorMessage" class="text-danger mt-3">⚠️ {{ errorMessage }}</p>
        <p v-if="successMessage" class="text-success mt-3">✔️ {{ successMessage }}</p>
        </form>
    </div>
</template>

<script setup>
    import { reactive, ref, onMounted } from 'vue'
    import axios from '@/axios'
    import { useRoute } from 'vue-router'
    import Multiselect from '@vueform/multiselect'
    import '@vueform/multiselect/themes/default.css'
    //import { watch } from 'vue'
    import { nextTick } from 'vue'

    const route = useRoute()
    const loading = ref(false)
    const error = ref('')
    const success = ref(false)

    const form = reactive({
        kode: '',
        nama: '',
        kategori_toko_id: null,
        kota_id: null,
        alamat: '',
        disc_1: '',
        disc_2: '',
        disc_3: '',
        ekspedisi_id: null,
        ongkir_id: null,
    })

    // State untuk multiselect
    const selectedKategori = ref(null)
    const selectedKota = ref(null)
    const selectedEkspedisi = ref(null)
    const selectedOngkir = ref(null)

    // ambil dulu semua merkList dari API
    const kategoriList = ref([])
    const kotaList = ref([])
    const ekspedisiList = ref([])
    const ongkirList = ref([])

    // func onSelect masing-masing artibut
    const onSelectKategori = (option) => {
        selectedKategori.value = option
        console.log('Kategori dipilih:', option)
    }

    const onSelectKota = (option) => {
        selectedKota.value = option
        console.log('Kota dipilih:', option)
    }

    const onSelectEkspedisi = (option) => {
        selectedEkspedisi.value = option
        console.log('Ekspedisi dipilih:', option)
    }

    const onSelectOngkir = (option) => {
        selectedOngkir.value = option
        console.log('Ongkir dipilih:', option)
    }

    // --- HELPER FUNCTION ---
    // Taruh setelah merkList dan selectedMerk dideklarasikan
    function setSelectedKategoriById(kategoriId) {
        if (!kategoriId) {
            selectedKategori.value = null
            return
        }
        const kategori = kategoriList.value.find(k => k.id === kategoriId) || null
        console.log('Kategori:', kategori.id)
        selectedKategori.value = kategori.id
        console.log('Set selected kategori:', kategori.id)
    }

    function setSelectedKotaById(kotaId) {
        if (!kotaId) {
            selectedKota.value = null
            return
        }
        const kota = kotaList.value.find(k => k.id === kotaId) || null
        console.log('Kota:', kota.id)
        selectedKota.value = kota.id
        console.log('Set selected kota:', kota.id)
    }

    function setSelectedEkspedisiById(ekspedisiId) {
        if (!ekspedisiId) {
            selectedEkspedisi.value = null
            return
        }
        const ekspedisi = ekspedisiList.value.find(e => e.id === ekspedisiId) || null
        console.log('Ekspedisi:', ekspedisi.id)
        selectedEkspedisi.value = ekspedisi.id
        console.log('Set selected ekspedisi:', ekspedisi.id)
    }

    function setSelectedOngkirById(ongkirId) {
        if (!ongkirId) {
            selectedOngkir.value = null
            return
        }
        const ongkir = ongkirList.value.find(o => o.id === ongkirId) || null
        console.log('Ongkir:', ongkir.id)
        selectedOngkir.value = ongkir.id
        console.log('Set selected ongkir:', ongkir.id)
    }

    // --- WATCHER ---
    // Pastikan watcher ini setelah merkList dan setSelectedMerkById dideklarasikan
    // watch(merkList, (newList) => {
    //     if (form.merk_id) {
    //         setSelectedMerkById(form.merk_id)
    //     }
    // })

    // // Watcher lain untuk update form.merk_id otomatis
    // watch(selectedMerk, (val) => {
    //     form.merk_id = val ? val.id : null
    // })

    // ===== FUNGSI CARI MERK (DEBOUNCE) =====
    let searchTimeout = null

    // ===== FUNGSI CARI KATEGORI (DEBOUNCE) =====
    const onSearchKategori = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                kategoriList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/kategori_toko?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    kategoriList.value = data
                    console.log('Kategori hasil pencarian:', data)
                } else {
                    kategoriList.value = []
                }
            } catch (err) {
                console.error('Gagal load kategori toko:', err)
                kategoriList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI KOTA (DEBOUNCE) =====
    const onSearchKota = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                kotaList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/kota?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    kotaList.value = data
                    console.log('Kota hasil pencarian:', data)
                } else {
                    kotaList.value = []
                }
            } catch (err) {
                console.error('Gagal load kota:', err)
                kotaList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI EKSPEDISI (DEBOUNCE) =====
    const onSearchEkspedisi = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                ekspedisiList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/ekspedisi?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    ekspedisiList.value = data
                    console.log('Ekspedisi hasil pencarian:', data)
                } else {
                    ekspedisiList.value = []
                }
            } catch (err) {
                console.error('Gagal load ekspedisi:', err)
                ekspedisiList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI ONGKIR (DEBOUNCE) =====
    const onSearchOngkir = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                ongkirList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/ongkir?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    ongkirList.value = data
                    console.log('Ongkir hasil pencarian:', data)
                } else {
                    ongkirList.value = []
                }
            } catch (err) {
                console.error('Gagal load ongkir:', err)
                ongkirList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // fungsi fetch list masing-masing atribut (get all data)
    async function fetchKategoriList() {
        try {
            const res = await axios.get('/api/kategori_toko')
            kategoriList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Kategori list:',kategoriList.value)
        } catch (err) {
            console.error('Gagal load kategori:', err)
        }
    }

    async function fetchKotaList() {
        try {
            const res = await axios.get('/api/kota')
            kotaList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Kota list:',kotaList.value)
        } catch (err) {
            console.error('Gagal load kota:', err)
        }
    }

    async function fetchEkspedisiList() {
        try {
            const res = await axios.get('/api/ekspedisi')
            ekspedisiList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Ekspedisi list:',ekspedisiList.value)
        } catch (err) {
            console.error('Gagal load ekspedisi:', err)
        }
    }

    async function fetchOngkirList() {
        try {
            const res = await axios.get('/api/ongkir')
            ongkirList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Ongkir list:',ongkirList.value)
        } catch (err) {
            console.error('Gagal load ongkir:', err)
        }
    }

    

    async function fetchToko() {
        loading.value = true
        error.value = ''
        success.value = false

        try {
            const tokoId = route.params.id
            const response = await axios.get(`/api/toko?id=${tokoId}`)
            const data = response.data // axios otomatis parsing json

            if (data.code === 200) {
                // masukan response api (data) ke dalam const
                const toko = data.data
                // alert(JSON.stringify(toko))

                form.id = toko.id
                form.kode = toko.kode
                form.nama = toko.nama
                form.kategori_toko_id = toko.kategori_toko_id
                form.kota_id = toko.kota_id
                form.alamat = toko.alamat
                form.disc_1 = toko.disc_1*100
                form.disc_2 = toko.disc_2*100
                form.disc_3 = toko.disc_3*100
                form.ekspedisi_id = toko.ekspedisi_id
                form.ongkir_id = toko.ongkir_id

                 // tunggu next tick supaya Multiselect render options dulu
                await nextTick()

                // Gunakan helper function
                setSelectedKategoriById(toko.kategori_toko_id)
                setSelectedKotaById(toko.kota_id)
                setSelectedEkspedisiById(toko.ekspedisi_id)
                setSelectedOngkirById(toko.ongkir_id)
            } else {
                throw new Error(data.message || 'Gagal mengambil data toko')
            }
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    }

    onMounted(async () => {
        loading.value = true
        try {
            await fetchKategoriList() // ambil data master kategori barang dulu
            await fetchKotaList() // ambil data master kota dulu
            await fetchEkspedisiList() // ambil data master ekspedisi dulu
            await fetchOngkirList() // ambil data master ongkir dulu
            await fetchToko()    // baru set fetch Toko
        } finally {
            loading.value = false
        }
    })

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    // ===== FUNGSI SUBMIT =====
    const handleSubmit = async () => {
        form.kategori_toko_id = selectedKategori.value
        form.kota_id = selectedKota.value
        form.ekspedisi_id = selectedEkspedisi.value
        form.ongkir_id = selectedOngkir.value
        
        console.log(form)

        submitting.value = true
        successMessage.value = ''
        errorMessage.value = ''

        // validasi data mandatory
        if (!selectedKategori.value) {
            errorMessage.value = 'Kategori toko wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            kategoriRef.value?.focus()
            return
        }

        if (!selectedKota.value) {
            errorMessage.value = 'Kota wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            kotaRef.value?.focus()
            return
        }

        if (!selectedEkspedisi.value) {
            errorMessage.value = 'Ekspedisi wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            ekspedisiRef.value?.focus()
            return
        }

        if (!selectedOngkir.value) {
            errorMessage.value = 'Ongkir wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            ongkirRef.value?.focus()
            return
        }

        // const payload = {
        //     merk_id: selectedMerk.value.id,
        //     artikel_id: selectedArtikel.value.id,
        //     warna_id: selectedWarna.value.id,
        //     kategori_barang_id: selectedKategori.value.id,
        //     jenis_barang_id: selectedJenis.value.id,
        //     ukuran_barang_id: selectedUkuran.value?.id || null,
        //     Kode: form.kode
        // }

        try {
            await axios.put(`/api/toko/${form.id}`, form) // parameter id
            successMessage.value = 'Toko berhasil diupdate!'
            // reset seluruh komponen form
            // form.kode = ''
            // selectedMerk.value = null
            // selectedArtikel.value = null
            // selectedWarna.value = null
            // selectedKategori.value = null
            // selectedJenis.value = null
            // selectedUkuran.value = null
        } catch (err) {
            console.error('Gagal simpan:', err)
            errorMessage.value = 'Gagal update toko.'
        } finally {
            submitting.value = false
        }
    }
</script>
