<template>
    <div>
        <div class="container-fluid">
            
            <div class="mb-3 d-flex align-items-center gap-2" style="max-width: 500px;">
                <button @click="tambahSupplier()" class="btn btn-primary flex-shrink-0">
                    Tambah Supplier
                </button>

                <input
                    id="search"
                    autocomplete="off"
                    v-model="search"
                    type="text"
                    class="form-control"
                    placeholder="Cari supplier..."
                    aria-label="Search supplier"
                />

                <button
                    v-if="search"
                    class="btn btn-outline-secondary"
                    type="button"
                    @click="search = ''"
                >
                    Clear
                </button>
            </div>

            <table class="table table-bordered table-striped table-hover">
                <thead class="table-secondary">
                    <tr>
                        <th class="text-center" @click="sortBy('id')">
                            ID
                            <span v-if="sortKey === 'id'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th @click="sortBy('nama')">
                            Nama
                            <span v-if="sortKey === 'nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th @click="sortBy('alamat')">
                            Alamat
                            <span v-if="sortKey === 'alamat'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('kontak')">
                            Kontak
                            <span v-if="sortKey === 'kontak'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center">
                            Action
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="supplier in paginatedData" :key="supplier.id">
                        <td class="text-center">{{ supplier.id }}</td>
                        <td>{{ supplier.nama }}</td>
                        <td>{{ supplier.alamat }}</td>
                        <td class="text-center">{{ supplier.kontak }}</td>
                        <td class="text-center">
                            <button
                                class="btn btn-sm btn-warning"
                                @click="editSupplier(supplier.id)"
                            >
                            Edit
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <!-- Pagination -->
        <div class="d-flex align-items-center gap-3 justify-content-center">
            <button class="btn btn-secondary" @click="prevPage" :disabled="currentPage === 1">
                Prev
            </button>
            <span class="text-muted">Halaman {{ currentPage }} dari {{ totalPages }}</span>
            <button class="btn btn-secondary" @click="nextPage" :disabled="currentPage === totalPages">
                Next
            </button>
        </div>
    </div>
</template>

<script setup>
    import { useRouter } from 'vue-router' // wajib import router
    const router = useRouter()

    function tambahSupplier() { // func jika button tambah di klik, pindah ke halaman tambah
        router.push('/supplier/tambah')
    }

    function editSupplier(id) {
        router.push({ name: 'EditSupplier', params: { id } })
    }

    import { ref, computed, watch } from 'vue' // watch wajib di import di setup

    const props = defineProps({
        suppliers: {
            type: Array,
            required: true
        }
    })

    const search = ref('')

    // Watch search untuk reset halaman saat keyword berubah
    watch(search, (newVal, oldVal) => {
        currentPage.value = 1
    })

    const sortKey = ref('')
    const sortAsc = ref(true)

    function sortBy(key) {
        if (sortKey.value === key) {
            sortAsc.value = !sortAsc.value
        } else {
            sortKey.value = key
            sortAsc.value = true
        }
    }

    const filteredSuppliers = computed(() => {
    let result = props.suppliers

        // Search filter
        if (search.value.trim() !== '') {
            const keyword = search.value.toLowerCase()
            result = result.filter(
                s => // search berdasarkan apa? nama, alamat, dan kontak sesuai kode di bawah
                String(s.id).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                s.nama.toLowerCase().includes(keyword) ||
                s.alamat.toLowerCase().includes(keyword) ||
                s.kontak.toLowerCase().includes(keyword)
            )
        }

        if (sortKey.value) {
            result = [...result].sort((a, b) => {
                const aVal = a[sortKey.value]
                const bVal = b[sortKey.value]

                if (typeof aVal === 'number' && typeof bVal === 'number') {
                    return sortAsc.value ? aVal - bVal : bVal - aVal
                }

                return sortAsc.value
                ? String(aVal).localeCompare(String(bVal), undefined, { sensitivity: 'base' })
                : String(bVal).localeCompare(String(aVal), undefined, { sensitivity: 'base' })
            })
        }

    return result
    })

    // Pagination
    const itemsPerPage = 7
    const currentPage = ref(1)

    const totalPages = computed(() =>
        Math.ceil(filteredSuppliers.value.length / itemsPerPage)
    )

    const paginatedData = computed(() => {
        const start = (currentPage.value - 1) * itemsPerPage
        return filteredSuppliers.value.slice(start, start + itemsPerPage)
    })

    function nextPage() {
        if (currentPage.value < totalPages.value) currentPage.value++
    }

    function prevPage() {
        if (currentPage.value > 1) currentPage.value--
    }

</script>

<style>
    th {
      cursor: pointer;
    }
    .pagination {
        display: flex;
        justify-content: center;
        align-items: center;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
        font-size: 14px;
        color: #495057; /* abu gelap nyaman di mata */
        padding: 8px 0;
        gap: 8px; /* jarak antar tombol dan tulisan */
    }

    .pagination span {
        font-weight: 500;
        user-select: none;
    }
</style>
