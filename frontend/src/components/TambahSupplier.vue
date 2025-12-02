<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <label for="nama" class="col-sm-2 col-form-label">Nama Supplier:</label>
                <div class="col-sm-10">
                    <input id="nama" v-model="form.nama" type="text" class="form-control" autocomplete="off" autofocus required />
                </div>
            </div>

            <div class="row mb-3 align-items-center">
                <label for="alamat" class="col-sm-2 col-form-label">Alamat Supplier:</label>
                <div class="col-sm-10">
                    <input id="alamat" v-model="form.alamat" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <div class="row mb-3 align-items-center">
                <label for="kontak" class="col-sm-2 col-form-label">Kontak Supplier:</label>
                <div class="col-sm-10">
                    <input id="kontak" v-model="form.kontak" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <div class="row">
                <div class="col-sm-6 offset-sm-2">
                <button
                    class="btn btn-primary mt-2"
                    type="submit"
                    :disabled="loading"
                >
                    Tambah Supplier
                </button>
                </div>
            </div>
            <p v-if="error" class="text-danger mt-3">⚠️ {{ error }}</p>
            <p v-if="success" class="text-success mt-3">✔️ Supplier berhasil ditambahkan!</p>
        </form>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'

const form = reactive ({
    nama: '',
    alamat: '',
    kontak: '',
})

const loading = ref(false)
const error = ref('')
const success = ref(false)

async function handleSubmit() {
    loading.value = true
    error.value = ''
    success.value = false

    try {
        const response = await fetch('http://localhost:3000/api/supplier', {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(form),
        })
    

        if (!response.ok) {
            const data = await response.json()
            throw new Error(data.message || 'Gagal menambah data supplier')
        }

        success.value = true
        form.nama = ''
        form.alamat = ''
        form.kontak = ''
    } catch (e: any) {
        error.value = e.message
    } finally {
        loading.value = false
    }
}
</script>