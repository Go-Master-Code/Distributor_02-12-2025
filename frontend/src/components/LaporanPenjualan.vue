<template>
    <div class="container-fluid">
        <form @submit.prevent="handleSubmit">
            <div class="row mb-3 align-items-center">
                <label for="tgl_awal" class="col-sm-2 col-form-label">Tanggal Awal:</label>
                <div class="col-sm-2">
                    <input id="tgl_awal" v-model="form.tgl_awal" type="date" class="form-control" autocomplete="off" required />
                </div>
                <label for="tgl_akhir" class="col-sm-2 col-form-label">Tanggal Akhir:</label>
                <div class="col-sm-2">
                    <input id="tgl_akhir" v-model="form.tgl_akhir" type="date" class="form-control" autocomplete="off" required />
                </div>
            </div>

            <!-- Tombol Simpan -->
            <button
                type="submit"
                class="btn btn-primary mt-2"
            > Preview
            </button>
        </form>
    </div>
</template>

<script setup>
    import { reactive } from 'vue'

    const form = reactive({
        tgl_awal: '',
        tgl_akhir: ''
    })

    const handleSubmit = () => {
    const url = `http://localhost:3000/report/penjualan?awal=${form.tgl_awal}&akhir=${form.tgl_akhir}`

    // Buka PDF di tab baru
    window.open(url, "_blank")
}
</script>