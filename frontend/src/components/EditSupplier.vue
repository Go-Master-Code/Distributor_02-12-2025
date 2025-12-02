<template>
    <div class="container">
        <form @submit.prevent="handleSubmit">
        
        <div class="mb-3 row align-items-center">
            <label for="id" class="col-sm-2 col-form-label">ID Supplier:</label>
            <div class="col-sm-10">
            <input id="id" v-model="form.id" type="text" class="form-control" disabled />
            </div>
        </div>

        <div class="mb-3 row align-items-center">
            <label for="nama" class="col-sm-2 col-form-label">Nama Supplier:</label>
            <div class="col-sm-10">
            <input id="nama" v-model="form.nama" type="text" class="form-control" autofocus required />
            </div>
        </div>

        <div class="mb-3 row align-items-center">
            <label for="alamat" class="col-sm-2 col-form-label">Alamat Supplier:</label>
            <div class="col-sm-10">
            <input id="alamat" v-model="form.alamat" type="text" class="form-control" required />
            </div>
        </div>

        <div class="mb-3 row align-items-center">
            <label for="kontak" class="col-sm-2 col-form-label">Kontak Supplier:</label>
            <div class="col-sm-10">
            <input id="kontak" v-model="form.kontak" type="text" class="form-control" required />
            </div>
        </div>

        <button type="submit" class="btn btn-primary" :disabled="loading">
            Simpan Perubahan
        </button>

        <p v-if="error" class="text-danger mt-3">⚠️ {{ error }}</p>
        <p v-if="success" class="text-success mt-3">✔️ Data berhasil diupdate!</p>
        </form>
    </div>
</template>

<script setup lang="ts">
    import { reactive, ref, onMounted } from 'vue'
    import { useRoute } from 'vue-router'

    const route = useRoute()

    const form = reactive({
        id: null as number | null,
        nama: '',
        alamat: '',
        kontak: '',
    })

    const loading = ref(false)
    const error = ref('')
    const success = ref(false)

    async function fetchSupplier() {
        loading.value = true
        error.value = ''
        success.value = false

        try {
            const supplierId = route.params.id
            const response = await fetch(`http://localhost:3000/api/supplier/${supplierId}`)
            const data = await response.json()

            if (response.ok && data.code === 200) {
                // Response sesuai format API yang kamu kasih
                // Isi seluruh komponen input dengan data yang didapat
                form.id = data.data.id
                form.nama = data.data.nama
                form.alamat = data.data.alamat
                form.kontak = data.data.kontak
            } else {
                throw new Error(data.message || 'Gagal mengambil data supplier')
            }
        } catch (e: any) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    }

    onMounted(() => {
    fetchSupplier()
    })

async function handleSubmit() {
    loading.value = true
    error.value = ''
    success.value = false

    try {
        const response = await fetch(`http://localhost:3000/api/supplier/${form.id}`, {
            method: 'PUT', // asumsi update pakai PUT
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ // parsing komponen form ke dalam request body
                nama: form.nama,
                alamat: form.alamat,
                kontak: form.kontak,
            }),
        })

        const data = await response.json()
        if (response.ok) {
            success.value = true
        } else {
            throw new Error(data.message || 'Gagal mengupdate supplier')
        }
    } catch (e: any) {
        error.value = e.message
    } finally {
        loading.value = false
    }
}
</script>
