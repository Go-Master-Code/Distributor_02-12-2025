<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Daftar Artikel</h1>
        </div>
        <ArtikelTbl :artikels="artikels" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllArtikel } from '../services/ArtikelService';
import ArtikelTbl from '../components/ArtikelTbl.vue';

const artikels = ref([])
onMounted(async () => {
    try {
        artikels.value = await getAllArtikel()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>