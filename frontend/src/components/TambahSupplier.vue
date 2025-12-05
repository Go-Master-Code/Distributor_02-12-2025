<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <label for="nama" class="col-sm-1 col-form-label">Nama:</label>
                <div class="col-sm-6">
                    <input id="nama" v-model="form.nama" type="text" class="form-control" autocomplete="off" autofocus required />
                </div>

                <label for="kontak" class="col-sm-1 col-form-label">Kontak:</label>
                <div class="col-sm-4">
                    <input id="kontak" v-model="form.kontak" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <div class="row mb-3 align-items-center">
                <label for="alamat" class="col-sm-1 col-form-label">Alamat:</label>
                <div class="col-sm-11">
                    <input id="alamat" v-model="form.alamat" type="text" class="form-control" autocomplete="off" required />
                </div>
            </div>

            
            <!-- Tombol Simpan -->
            <button
                type="submit"
                :disabled="submitting"
                class="btn btn-primary mt-2"
            >
                {{ submitting ? 'Menyimpan...' : 'Tambah Supplier' }}
            </button>
                
            <!-- Pesan sukses / error -->
            <p v-if="errorMessage" class="text-danger mt-3">⚠️ {{ errorMessage }}</p>
            <p v-if="successMessage" class="text-success mt-3">✔️ {{ successMessage }}</p>
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

    // State tombol submit & pesan
    const submitting = ref(false)
    const successMessage = ref('')
    const errorMessage = ref('')

    async function handleSubmit() {
        submitting.value = true
        successMessage.value = ''
        errorMessage.value = ''

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

            successMessage.value = 'Supplier berhasil disimpan!'
            form.nama = ''
            form.alamat = ''
            form.kontak = ''
        } catch (err) {
            console.error('Gagal simpan:', err)
            errorMessage.value = 'Gagal menyimpan supplier.'
        } finally {
            submitting.value = false
        }
    }
</script>