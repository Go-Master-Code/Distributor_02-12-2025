<template>
    <div>
        <div class="container-fluid">
            <h1 class="my-4">Daftar Supplier</h1>
        </div>
        <SupplierTbl :suppliers="suppliers" />
    </div>
</template>

<script setup>
import { ref, onMounted} from 'vue'
import { getAllSuppliers } from '../services/supplierService';
import SupplierTbl from '../components/SupplierTbl.vue';

const suppliers = ref([])
onMounted(async () => {
    try {
        suppliers.value = await getAllSuppliers()
    } catch (err) {
        console.error('Error saat load data:', err)
    }
})
</script>