<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <label for="kode" class="col-sm-2 col-form-label">Kode Toko:</label>
                <div class="col-sm-2">
                    <input id="kode" v-model="form.kode" type="text" class="form-control" autocomplete="off" autofocus required />
                </div>
                <label for="nama" class="col-sm-2 col-form-label">Nama Toko:</label>
                <div class="col-sm-6">
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

            <!-- Tombol Simpan -->
            <button
                type="submit"
                :disabled="submitting"
                class="btn btn-primary mt-2"
            >
                {{ submitting ? 'Menyimpan...' : 'Tambah Toko' }}
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
    const kategoriRef = ref(null)
    const kotaRef = ref(null)
    const ekspedisiRef = ref(null)
    const ongkirRef = ref(null)

    // Form utama
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

    // State untuk multiselect kategori toko
    const selectedKategori = ref(null)
    const kategoriList = ref([])
    const loading = ref(false) // deklarasi sekali saja

    const onSelectKategori = (option) => {
        selectedKategori.value = option
        console.log('Kategori dipilih:', option)
    }

    // State untuk multiselect kota
    const selectedKota = ref(null)
    const kotaList = ref([])
    // const loading = ref(false)

    const onSelectKota = (option) => {
        selectedKota.value = option
        console.log('Kota dipilih:', option)
    }

    // State untuk multiselect ekspedisi
    const selectedEkspedisi = ref(null)
    const ekspedisiList = ref([])
    // const loading = ref(false)

    const onSelectEkspedisi = (option) => {
        selectedEkspedisi.value = option
        console.log('Ekspedisi dipilih:', option)
    }

    // State untuk multiselect ongkir
    const selectedOngkir = ref(null)
    const ongkirList = ref([])
    // const loading = ref(false)

    const onSelectOngkir = (option) => {
        selectedOngkir.value = option
        console.log('Ongkir dipilih:', option)
    }

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    // ===== FUNGSI CARI WARNA (DEBOUNCE) =====
    let searchTimeout = null // deklarasi sekali saja
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
                    console.log('Kategori toko hasil pencarian:', data)
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

    // ===== FUNGSI CARI ARTIKEL (DEBOUNCE) =====
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

    // ===== FUNGSI SUBMIT =====
    const handleSubmit = async () => {
        form.kategori_toko_id = selectedKategori.value
        form.kota_id = selectedKota.value
        form.ekspedisi_id = selectedEkspedisi.value
        form.ongkir_id = selectedOngkir.value
        form.disc_1 = form.disc_1/100
        form.disc_2 = form.disc_2/100
        form.disc_3 = form.disc_3/100
        
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

        try {
            await axios.post('/api/toko', form)
            successMessage.value = 'Toko berhasil disimpan!'
            // reset seluruh komponen form
            form.kode = ''
            form.nama = ''
            form.alamat = ''
            form.disc_1 = ''
            form.disc_2 = ''
            form.disc_3 = ''
            selectedKategori.value = null
            selectedKota.value = null
            selectedEkspedisi.value = null
            selectedOngkir.value = null
        } catch (err) {
            console.error('Gagal simpan:', err)
            errorMessage.value = 'Gagal menyimpan toko.'
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