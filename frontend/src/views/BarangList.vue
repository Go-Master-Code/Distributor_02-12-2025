<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Daftar Barang</h1>
        </div>
        <BarangTbl :barangs="barangs" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllBarang } from '../services/barangService';
import BarangTbl from '../components/BarangTbl.vue';

const barangs = ref([])
onMounted(async () => {
    try {
        barangs.value = await getAllBarang()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>