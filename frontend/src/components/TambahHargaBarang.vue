<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <!-- nama label harus sesuai dengan response api json (nama) -->
                <label for="barang" class="col-sm-1 col-form-label">Barang:</label> 
                <div class="col-sm-5"> 
                    <Multiselect
                        ref="barangRef"
                        id="barang"
                        v-model="selectedBarang"
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
                        @select="onSelectBarang"
                    />
                </div> 
            </div>

            <div class="row mb-3 align-items-center">
                <label for="harga" class="col-sm-1 col-form-label">Harga</label>
                <div class="col-sm-2">
                    <input id="harga" v-model.number="form.harga" type="number" min="1" class="form-control" autocomplete="off" required />
                </div>

                <label for="mulai_berlaku" class="col-sm-1 col-form-label">Berlaku:</label>
                <div class="col-sm-2">
                    <input
                        id="mulai_berlaku"
                        v-model="form.mulai_berlaku"
                        type="date"
                        class="form-control"
                        autocomplete="off"
                        required
                    />
                </div>
            </div>

            <!-- Tombol Simpan -->
            <button
                type="submit"
                :disabled="submitting"
                class="btn btn-primary mt-2"
            >
                {{ submitting ? 'Menyimpan...' : 'Tambah Harga Barang' }}
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
    const barangRef = ref(null)

    // Form utama (reactive form)
    const form = reactive({
        mulai_berlaku: new Date().toISOString().substr(0, 10), // default hari ini
        harga: '',
        barang_id: null,
    })

    // State untuk multiselect merk
    const selectedBarang = ref(null)
    const barangList = ref([])
    const loading = ref(false) // deklarasi sekali saja

    const onSelectBarang = (option) => {
        selectedBarang.value = option
        console.log('Barang dipilih:', option)
    }

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    // ===== FUNGSI CARI WARNA (DEBOUNCE) =====
    let searchTimeout = null // deklarasi sekali saja
        const onSearchBarang = (query) => {
        clearTimeout(searchTimeout)
        searchTimeout = setTimeout(async () => {
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
                    label: `${b.kode} : ${b.merk_nama} | ${b.artikel_nama} | ${b.warna_nama} | ${b.ukuran_nama}`
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

    // ===== FUNGSI SUBMIT =====
    const handleSubmit = async () => {
        form.barang_id = selectedBarang.value
        
        console.log(form)

        successMessage.value = ''
        errorMessage.value = ''

        // validasi data mandatory
        if (!selectedBarang.value) {
            errorMessage.value = 'Barang wajib dipilih.'
            // 🧭 fokuskan kembali ke Multiselect
            barangRef.value?.focus()
            return
        }

        // proses submit dimulai setelah validasi selesai
        submitting.value = true

        try {
            await axios.post('/api/harga_barang', form)
            successMessage.value = 'Harga barang berhasil disimpan!'
            // reset seluruh komponen form
            selectedBarang.value = null
            form.harga = ''
            form.mulai_berlaku = new Date().toISOString().substr(0, 10)
        } catch (err) {
            console.error('Gagal simpan:', err)
            errorMessage.value = 'Gagal menyimpan harga barang.'
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