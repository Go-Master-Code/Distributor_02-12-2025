<template>
    <div>
        <div class="container-fluid">
            
            <div class="mb-3 d-flex align-items-center gap-2" style="max-width: 500px;">
                <button @click="tambahBarang()" class="btn btn-primary flex-shrink-0">
                    Tambah Barang
                </button>

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
                        <th class="text-center" @click="sortBy('id')">
                            ID
                            <span v-if="sortKey === 'id'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('kode')">
                            Kode
                            <span v-if="sortKey === 'kode'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <!-- nama field sort harus sesuai dengan nama var di tbody -->
                        <th class="text-center" @click="sortBy('merk_nama')">
                            Merk
                            <span v-if="sortKey === 'merk_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('artikel_nama')">
                            Artikel
                            <span v-if="sortKey === 'artikel_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('warna_nama')">
                            Warna
                            <span v-if="sortKey === 'warna_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('kategori_barang_nama')">
                            Kategori
                            <span v-if="sortKey === 'kategori_barang_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('jenis_barang_nama')">
                            Jenis
                            <span v-if="sortKey === 'jenis_barang_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('ukuran_nama')">
                            Ukuran
                            <span v-if="sortKey === 'ukuran_nama'">
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
                        <th class="text-center">
                            Action
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="barang in paginatedData" :key="barang.id">
                        <td class="text-center">{{ barang.id }}</td>
                        <td class="text-center">{{ barang.kode }}</td>
                        <td class="text-center">{{ barang.merk_nama }}</td>
                        <td class="text-center">{{ barang.artikel_nama }}</td>
                        <td class="text-center">{{ barang.warna_nama }}</td>
                        <td class="text-center">{{ barang.kategori_barang_nama }}</td>
                        <td class="text-center">{{ barang.jenis_barang_nama }}</td>
                        <td class="text-center">{{ barang.ukuran_nama }}</td>
                        <td class="text-center">{{ barang.stok }}</td>
                        <td class="text-center">
                            <button
                                class="btn btn-sm btn-warning"
                                @click="editBarang(barang.id)"
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

    function tambahBarang() { // func jika button tambah di klik, pindah ke halaman tambah
        router.push('/barang/tambah')
    }

    function editBarang(id) {
        router.push({ name: 'EditBarang', params: { id } })
    }

    import { ref, computed, watch } from 'vue' // watch wajib di import di setup

    const props = defineProps({
        barangs: {
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

    const filteredBarangs = computed(() => {
    let result = props.barangs

        // Search filter
        if (search.value.trim() !== '') {
            const keyword = search.value.toLowerCase()
            result = result.filter(
                b => // search berdasarkan apa? nama, alamat, dan kontak sesuai kode di bawah
                String(b.id).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                b.kode.toLowerCase().includes(keyword) ||
                b.merk_nama.toLowerCase().includes(keyword) ||
                b.artikel_nama.toLowerCase().includes(keyword) ||
                b.warna_nama.toLowerCase().includes(keyword) ||
                b.kategori_barang_nama.toLowerCase().includes(keyword) ||
                b.jenis_barang_nama.toLowerCase().includes(keyword) ||
                b.ukuran_nama.toLowerCase().includes(keyword) ||
                String(b.stok).includes(keyword) // <-- convert dulu angka ke string agar bisa di filter
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
        Math.ceil(filteredBarangs.value.length / itemsPerPage)
    )

    const paginatedData = computed(() => {
        const start = (currentPage.value - 1) * itemsPerPage
        return filteredBarangs.value.slice(start, start + itemsPerPage)
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
