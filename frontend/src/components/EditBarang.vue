<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
        
            <div class="row mb-3 align-items-center">
                <label for="id" class="col-sm-2 col-form-label">ID:</label>
                <div class="col-sm-10">
                    <input id="id" v-model="form.id" type="text" class="form-control" disabled />
                </div>
            </div>

            <div class="row mb-3 align-items-center">
                <label for="kode" class="col-sm-2 col-form-label">Kode Barang:</label>
                <div class="col-sm-10">
                    <input id="kode" v-model="form.kode" type="text" class="form-control" autocomplete="off" autofocus required />
                </div>
            </div>

            <!-- Multiselect merk -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="merk" class="col-sm-2 col-form-label">Merk:</label> 
                <div class="col-sm-10"> 
                    <Multiselect
                        ref="merkRef"
                        id="merk"
                        v-model="selectedMerk"
                        mode="single"
                        :options="merkList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari merk..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchMerk"
                        @select="onSelectMerk"
                    />
                </div>   
            </div>

            <!-- Multiselect artikel -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="artikel" class="col-sm-2 col-form-label">Artikel:</label> 
                <div class="col-sm-10"> 
                    <Multiselect
                        ref="artikelRef"
                        id="artikel"
                        v-model="selectedArtikel"
                        mode="single"
                        :options="artikelList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari artikel..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchArtikel"
                        @select="onSelectArtikel"
                    />
                </div>   
            </div>

            <!-- Multiselect warna -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="warna" class="col-sm-2 col-form-label">Warna:</label> 
                <div class="col-sm-10"> 
                    <Multiselect
                        ref="warnaRef"
                        id="warna"
                        v-model="selectedWarna"
                        mode="single"
                        :options="warnaList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari warna..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchWarna"
                        @select="onSelectWarna"
                    />
                </div>   
            </div>

            <!-- Multiselect kategori barang: dewasa, anak, kids -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="kategori" class="col-sm-2 col-form-label">Kategori:</label> 
                <div class="col-sm-10"> 
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
                        placeholder="Ketik untuk mencari kategori barang..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchKategori"
                        @select="onSelectKategori"
                    />
                </div>   
            </div>

            <!-- Multiselect jenis barang: sepatu, sandal, kaos kaki, dll -->
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="jenis" class="col-sm-2 col-form-label">Jenis Barang:</label> 
                <div class="col-sm-10"> 
                    <Multiselect
                        ref="jenisRef"
                        id="jenis"
                        v-model="selectedJenis"
                        mode="single"
                        :options="jenisList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari jenis barang..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchJenis"
                        @select="onSelectJenis"
                    />
                </div>   
            </div>

            <!-- Multiselect ukuran -->
            <div class="row mb-3 align-items-center">
                <label for="ukuran" class="col-sm-2 col-form-label">Ukuran:</label> 
                <div class="col-sm-10"> 
                    <Multiselect
                        ref="ukuranRef"
                        id="ukuran"
                        v-model="selectedUkuran"
                        mode="single"
                        :options="ukuranList"
                        :searchable="true"
                        :filter-results="false"
                        :clear-on-select="true"
                        :close-on-select="true"
                        :internal-search="false"
                        :loading="loading"
                        autocomplete="off"
                        placeholder="Ketik untuk mencari ukuran..."
                        label="nama"
                        value-prop="id"
                        track-by="id"
                        @search-change="onSearchUkuran"
                        @select="onSelectUkuran"
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
        id: null,
        kode: '',
        merk_id: null,
        artikel_id: null,
        warna_id: null,
        kategori_barang_id: null,
        jenis_barang_id: null,
        ukuran_id: null,
        // kontak: '',
    })

    // State untuk multiselect
    const selectedMerk = ref(null)
    const selectedArtikel = ref(null)
    const selectedWarna = ref(null)
    const selectedKategori = ref(null)
    const selectedJenis = ref(null)
    const selectedUkuran = ref(null)

    // ambil dulu semua merkList dari API
    const merkList = ref([])
    const artikelList = ref([])
    const warnaList = ref([])
    const kategoriList = ref([])
    const jenisList = ref([])
    const ukuranList = ref([])

    // func onSelect masing-masing artibut
    const onSelectMerk = (option) => {
        selectedMerk.value = option
        console.log('Merk dipilih:', option)
    }

    const onSelectArtikel = (option) => {
        selectedArtikel.value = option
        console.log('Artikel dipilih:', option)
    }

    const onSelectWarna = (option) => {
        selectedWarna.value = option
        console.log('Warna dipilih:', option)
    }

    const onSelectKategori = (option) => {
        selectedKategori.value = option
        console.log('Kategori dipilih:', option)
    }

    const onSelectJenis = (option) => {
        selectedJenis.value = option
        console.log('Jenis dipilih:', option)
    }

    const onSelectUkuran = (option) => {
        selectedUkuran.value = option
        console.log('Ukuran dipilih:', option)
    }

    // --- HELPER FUNCTION ---
    // Taruh setelah merkList dan selectedMerk dideklarasikan
    function setSelectedMerkById(merkId) {
        if (!merkId) {
            selectedMerk.value = null
            return
        }
        const merk = merkList.value.find(m => m.id === merkId) || null
        console.log('Merk:', merk.id)
        selectedMerk.value = merk.id
        console.log('Set selected merk:', merk.id)
    }

    function setSelectedArtikelById(artikelId) {
        if (!artikelId) {
            selectedArtikel.value = null
            return
        }
        const artikel = artikelList.value.find(a => a.id === artikelId) || null
        console.log('Artikel:', artikel.id)
        selectedArtikel.value = artikel.id
        console.log('Set selected artikel:', artikel.id)
    }

    function setSelectedWarnaById(warnaId) {
        if (!warnaId) {
            selectedWarna.value = null
            return
        }
        const warna = warnaList.value.find(w => w.id === warnaId) || null
        console.log('Warna:', warna.id)
        selectedWarna.value = warna.id
        console.log('Set selected warna:', warna.id)
    }

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

    function setSelectedJenisById(jenisId) {
        if (!jenisId) {
            selectedJenis.value = null
            return
        }
        const jenis = jenisList.value.find(j => j.id === jenisId) || null
        console.log('Jenis:', jenis.id)
        selectedJenis.value = jenis.id
        console.log('Set selected jenis:', jenis.id)
    }

    function setSelectedUkuranById(ukuranId) {
        if (!ukuranId) {
            selectedUkuran.value = null
            return
        }
        const ukuran = ukuranList.value.find(u => u.id === ukuranId) || null
        console.log('Ukuran:', ukuran.id)
        selectedUkuran.value = ukuran.id
        console.log('Set selected ukuran:', ukuran.id)
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
    const onSearchMerk = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                merkList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/merk?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    merkList.value = data
                    console.log('Merk hasil pencarian:', data)
                } else {
                    merkList.value = []
                }
            } catch (err) {
                console.error('Gagal load merk:', err)
                merkList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI ARTIKEL (DEBOUNCE) =====
    const onSearchArtikel = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                artikelList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/artikel?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    artikelList.value = data
                    console.log('Artikel hasil pencarian:', data)
                } else {
                    artikelList.value = []
                }
            } catch (err) {
                console.error('Gagal load artikel:', err)
                artikelList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI WARNA (DEBOUNCE) =====
    const onSearchWarna = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                warnaList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/warna?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    warnaList.value = data
                    console.log('Warna hasil pencarian:', data)
                } else {
                    warnaList.value = []
                }
            } catch (err) {
                console.error('Gagal load warna:', err)
                warnaList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

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
                const res = await axios.get(`/api/kategori_barang?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    kategoriList.value = data
                    console.log('Kategori hasil pencarian:', data)
                } else {
                    kategoriList.value = []
                }
            } catch (err) {
                console.error('Gagal load kategori barang:', err)
                kategoriList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI JENIS (DEBOUNCE) =====
    const onSearchJenis = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                jenisList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/jenis_barang?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    jenisList.value = data
                    console.log('Jenis barang hasil pencarian:', data)
                } else {
                    jenisList.value = []
                }
            } catch (err) {
                console.error('Gagal load jenis barang:', err)
                jenisList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // ===== FUNGSI CARI UKURAN (DEBOUNCE) =====
    const onSearchUkuran = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
            // Kalau query kosong, biarkan list terakhir atau kosongkan (opsional)
            if (!query || query.trim().length < 1) {
                ukuranList.value = []
                return
            }

            loading.value = true
            try {
                const res = await axios.get(`/api/ukuran?nama=${encodeURIComponent(query)}`)
                const data = res.data.data

                // Pastikan hasil dari API adalah array
                if (Array.isArray(data)) {
                    ukuranList.value = data
                    console.log('Ukuran hasil pencarian:', data)
                } else {
                    ukuranList.value = []
                }
            } catch (err) {
                console.error('Gagal load ukuran:', err)
                ukuranList.value = []
            } finally {
                loading.value = false
            }
        }, 400) // delay agar tidak spam API
    }

    // fungsi fetch list masing-masing atribut (get all data)
    async function fetchMerkList() {
        try {
            const res = await axios.get('/api/merk')
            merkList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Merk list:',merkList.value)
        } catch (err) {
            console.error('Gagal load merk:', err)
        }
    }

    async function fetchArtikelList() {
        try {
            const res = await axios.get('/api/artikel')
            artikelList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Artikel list:',artikelList.value)
        } catch (err) {
            console.error('Gagal load artikel:', err)
        }
    }

    async function fetchWarnaList() {
        try {
            const res = await axios.get('/api/warna')
            warnaList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Warna list:',warnaList.value)
        } catch (err) {
            console.error('Gagal load warna:', err)
        }
    }

    async function fetchKategoriList() {
        try {
            const res = await axios.get('/api/kategori_barang')
            kategoriList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Kategori list:',kategoriList.value)
        } catch (err) {
            console.error('Gagal load kategori:', err)
        }
    }

    async function fetchJenisList() {
        try {
            const res = await axios.get('/api/jenis_barang')
            jenisList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Jenis list:',jenisList.value)
        } catch (err) {
            console.error('Gagal load jenis:', err)
        }
    }

    async function fetchUkuranList() {
        try {
            const res = await axios.get('/api/ukuran')
            ukuranList.value = Array.isArray(res.data.data) ? res.data.data : []
            console.log('Ukuran list:',ukuranList.value)
        } catch (err) {
            console.error('Gagal load ukuran:', err)
        }
    }

    

    async function fetchBarang() {
        loading.value = true
        error.value = ''
        success.value = false

        try {
            const barangId = route.params.id
            const response = await axios.get(`/api/barang?id=${barangId}`)
            const data = response.data // axios otomatis parsing json

            if (data.code === 200) {
                // masukan response api (data) ke dalam const
                const barang = data.data
                form.id = barang.id
                form.kode = barang.kode
                form.merk_id = barang.merk_id
                form.artikel_id = barang.artikel_id
                form.warna_id = barang.warna_id
                form.warna_id = barang.warna_id
                form.kategori_barang_id = barang.kategori_barang_id
                form.jenis_barang_id = barang.jenis_barang_id
                form.ukuran_id = barang.ukuran_id

                 // tunggu next tick supaya Multiselect render options dulu
                await nextTick()

                // Gunakan helper function
                setSelectedMerkById(barang.merk_id)
                setSelectedArtikelById(barang.artikel_id)
                setSelectedWarnaById(barang.warna_id)
                setSelectedKategoriById(barang.kategori_barang_id)
                setSelectedJenisById(barang.jenis_barang_id)
                setSelectedUkuranById(barang.ukuran_id)
            } else {
                throw new Error(data.message || 'Gagal mengambil data barang')
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
            await fetchMerkList()  // ambil data master merk dulu
            await fetchArtikelList() // ambil data master artikel dulu
            await fetchWarnaList() // ambil data master warna dulu
            await fetchKategoriList() // ambil data master kategori barang dulu
            await fetchJenisList() // ambil data master jenis barang dulu
            await fetchUkuranList() // ambil data master ukuran dulu
            await fetchBarang()    // baru set selectedMerk
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
        form.merk_id = selectedMerk.value
        form.artikel_id = selectedArtikel.value
        form.warna_id = selectedWarna.value
        form.kategori_barang_id = selectedKategori.value
        form.jenis_barang_id = selectedJenis.value
        form.ukuran_id = selectedUkuran.value || null
        
        console.log(form)

        submitting.value = true
        successMessage.value = ''
        errorMessage.value = ''

        // validasi data mandatory
        if (!selectedMerk.value) {
            errorMessage.value = 'Merk wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            merkRef.value?.focus()
            return
        }

        if (!selectedArtikel.value) {
            errorMessage.value = 'Artikel wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            artikelRef.value?.focus()
            return
        }

        if (!selectedWarna.value) {
            errorMessage.value = 'Warna wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            warnaRef.value?.focus()
            return
        }

        if (!selectedKategori.value) {
            errorMessage.value = 'Kategori wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            kategoriRef.value?.focus()
            return
        }

        if (!selectedJenis.value) {
            errorMessage.value = 'Jenis wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            jenisRef.value?.focus()
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
            await axios.put(`/api/barang/${form.id}`, form) // parameter id
            successMessage.value = 'Barang berhasil diupdate!'
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
            errorMessage.value = 'Gagal update barang.'
        } finally {
            submitting.value = false
        }
    }
</script>
