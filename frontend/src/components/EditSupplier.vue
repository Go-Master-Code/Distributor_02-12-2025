<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
        
        <div class="row mb-3 align-items-center">
            <label for="id" class="col-sm-1 col-form-label">ID:</label>
            <div class="col-sm-2">
                <input id="id" v-model="form.id" type="text" class="form-control" disabled />
            </div>

            <label for="nama" class="col-sm-1 col-form-label">Nama:</label>
            <div class="col-sm-3">
                <input id="nama" v-model="form.nama" type="text" class="form-control" autofocus required />
            </div>

            <label for="kontak" class="col-sm-1 col-form-label">Kontak:</label>
            <div class="col-sm-4">
                <input id="kontak" v-model="form.kontak" type="text" class="form-control" required />
            </div>
        </div>

        <div class="row mb-3 align-items-center">
            <label for="alamat" class="col-sm-1 col-form-label">Alamat:</label>
            <div class="col-sm-11">
                <input id="alamat" v-model="form.alamat" type="text" class="form-control" required />
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
