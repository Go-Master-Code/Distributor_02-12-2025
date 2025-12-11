<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Daftar Harga Barang</h1>
        </div>
        <HargaBarangTbl :barangs="barangs" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllHargaBarang } from '../services/harga_barangService';
import HargaBarangTbl from '../components/HargaBarangTbl.vue';

const barangs = ref([])
onMounted(async () => {
    try {
        barangs.value = await getAllHargaBarang()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>