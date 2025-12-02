<template>
    <div>
        <div class="container-fluid">
            
            <div class="mb-3 d-flex align-items-center gap-2" style="max-width: 500px;">
                <input
                    id="search"
                    autocomplete="off"
                    v-model="search"
                    type="text"
                    class="form-control"
                    placeholder="Cari Barang..."
                    aria-label="Search Barang"
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
                        <th class="text-center" @click="sortBy('barang_id')">
                            ID Barang
                            <span v-if="sortKey === 'barang_id'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('barang_kode')">
                            Kode
                            <span v-if="sortKey === 'barang_kode'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('barang_merk')">
                            Merk
                            <span v-if="sortKey === 'barang_merk'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('barang_warna')">
                            Warna
                            <span v-if="sortKey === 'barang_warna'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <!-- nama field sort harus sesuai dengan nama var di tbody -->
                        <th class="text-center" @click="sortBy('size')">
                            Size
                            <span v-if="sortKey === 'size'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('stok')">
                            Stok
                            <span v-if="sortKey === 'stok'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="kartuStok in paginatedData" :key="kartuStok.id">
                        <td class="text-center">{{ kartuStok.barang_id }}</td>
                        <td class="text-center">{{ kartuStok.barang_kode }}</td>
                        <td class="text-center">{{ kartuStok.barang_merk }}</td>
                        <td class="text-center">{{ kartuStok.barang_warna }}</td>
                        <td class="text-center">{{ kartuStok.size }}</td>
                        <td class="text-center">{{ kartuStok.stok }}</td>
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

    import { ref, computed, watch } from 'vue' // watch wajib di import di setup

    const props = defineProps({
        kartustoks: {
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

    // sort column
    function sortBy(key) {
        if (sortKey.value === key) {
            sortAsc.value = !sortAsc.value
        } else {
            sortKey.value = key
            sortAsc.value = true
        }
    }

    const filteredKartuStok = computed(() => {
    let result = props.kartustoks

        // Search filter
        if (search.value.trim() !== '') {
            const keyword = search.value.toLowerCase()
            result = result.filter(
                ks => // search berdasarkan id, kode barang, size, atau stok
                String(ks.barang_id).includes(keyword) ||
                ks.barang_kode.toLowerCase().includes(keyword) ||
                ks.barang_merk.toLowerCase().includes(keyword) ||
                ks.barang_warna.toLowerCase().includes(keyword) ||
                String(ks.size).includes(keyword) ||// <-- convert dulu angka ke string agar bisa di filter
                String(ks.stok).includes(keyword) // <-- convert dulu angka ke string agar bisa di filter
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
        Math.ceil(filteredKartuStok.value.length / itemsPerPage)
    )

    const paginatedData = computed(() => {
        const start = (currentPage.value - 1) * itemsPerPage
        return filteredKartuStok.value.slice(start, start + itemsPerPage)
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
