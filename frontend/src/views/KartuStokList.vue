<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Kartu Stok</h1>
        </div>
        <KartuStokTbl :kartustoks="kartustoks" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllKartuStok } from '../services/kartu_stokService';
import KartuStokTbl from '../components/KartuStokTbl.vue';

const kartustoks = ref([])
onMounted(async () => {
    try {
        kartustoks.value = await getAllKartuStok()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>