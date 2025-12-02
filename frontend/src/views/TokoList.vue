<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Daftar Toko</h1>
        </div>
        <TokoTbl :tokos="tokos" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllToko } from '../services/tokoService';
import TokoTbl from '../components/TokoTbl.vue';

const tokos = ref([])
onMounted(async () => {
    try {
        tokos.value = await getAllToko()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>