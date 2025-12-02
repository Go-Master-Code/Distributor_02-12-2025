<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
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

            <!-- Tombol Simpan -->
            <button
                type="submit"
                :disabled="submitting"
                class="btn btn-primary mt-2"
            >
                {{ submitting ? 'Menyimpan...' : 'Tambah Barang' }}
            </button>

            <!-- Pesan sukses / error -->
            <p v-if="errorMessage" class="text-danger mt-3">⚠️ {{ errorMessage }}</p>
            <p v-if="successMessage" class="text-success mt-3">✔️ {{ successMessage }}</p>
        </form>
    </div>
</template>

<script setup>
    import { ref, reactive } from 'vue'
    import axios from '@/axios'
    import Multiselect from '@vueform/multiselect'
    import '@vueform/multiselect/themes/default.css'

    // deklarasi ref komponen multiselect untuk request focus
    const merkRef = ref(null)
    const artikelRef = ref(null)
    const warnaRef = ref(null)
    const kategoriRef = ref(null)
    const jenisRef = ref(null)
    const ukuranRef = ref(null)

    // Form utama
    const form = reactive({
        kode: '',
        merk_id: null,
        artikel_id: null,
        warna_id: null,
        kategori_barang_id: null,
        jenis_barang_id: null,
        ukuran_id: null,
    })

    // State untuk multiselect merk
    const selectedMerk = ref(null)
    const merkList = ref([])
    const loading = ref(false) // deklarasi sekali saja

    const onSelectMerk = (option) => {
        selectedMerk.value = option
        console.log('Merk dipilih:', option)
    }

    // State untuk multiselect artikel
    const selectedArtikel = ref(null)
    const artikelList = ref([])
    // const loading = ref(false)

    const onSelectArtikel = (option) => {
        selectedArtikel.value = option
        console.log('Artikel dipilih:', option)
    }

    // State untuk multiselect warna
    const selectedWarna = ref(null)
    const warnaList = ref([])
    // const loading = ref(false)

    const onSelectWarna = (option) => {
        selectedWarna.value = option
        console.log('Warna dipilih:', option)
    }

    // State untuk multiselect kategori
    const selectedKategori = ref(null)
    const kategoriList = ref([])
    // const loading = ref(false)

    const onSelectKategori = (option) => {
        selectedKategori.value = option
        console.log('Kategori dipilih:', option)
    }

    // State untuk multiselect jenis
    const selectedJenis = ref(null)
    const jenisList = ref([])
    // const loading = ref(false)

    const onSelectJenis = (option) => {
        selectedJenis.value = option
        console.log('Jenis barang dipilih:', option)
    }

    // State untuk multiselect ukuran
    const selectedUkuran = ref(null)
    const ukuranList = ref([])
    // const loading = ref(false)

    const onSelectUkuran = (option) => {
        selectedUkuran.value = option
        console.log('Ukuran dipilih:', option)
    }

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    // ===== FUNGSI CARI WARNA (DEBOUNCE) =====
    let searchTimeout = null // deklarasi sekali saja
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

    // ===== FUNGSI CARI MERK (DEBOUNCE) =====
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
            await axios.post('/api/barang', form)
            successMessage.value = 'Barang berhasil disimpan!'
            // reset seluruh komponen form
            form.kode = ''
            selectedMerk.value = null
            selectedArtikel.value = null
            selectedWarna.value = null
            selectedKategori.value = null
            selectedJenis.value = null
            selectedUkuran.value = null
        } catch (err) {
            console.error('Gagal simpan:', err)
            errorMessage.value = 'Gagal menyimpan barang.'
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