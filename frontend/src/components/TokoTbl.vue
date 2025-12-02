<template>
    <div>
        <div class="container-fluid">
            
            <div class="mb-3 d-flex align-items-center gap-2" style="max-width: 500px;">
                <button @click="tambahToko()" class="btn btn-primary flex-shrink-0">
                    Tambah Toko
                </button>

                <input
                    id="search"
                    autocomplete="off"
                    v-model="search"
                    type="text"
                    class="form-control"
                    placeholder="Cari Toko..."
                    aria-label="Search Toko"
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

            <table class="table table-bordered table-striped table-hover small">
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
                        <th class="text-center" @click="sortBy('nama')">
                            Nama
                            <span v-if="sortKey === 'nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('kategori_toko_nama')">
                            Kategori
                            <span v-if="sortKey === 'kategori_toko_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('kota_nama')">
                            Kota
                            <span v-if="sortKey === 'kota_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('alamat')">
                            Alamat
                            <span v-if="sortKey === 'alamat'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('disc_1')">
                            Disc 1
                            <span v-if="sortKey === 'disc_1'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('disc_2')">
                            Disc 2
                            <span v-if="sortKey === 'disc_2'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('disc_3')">
                            Disc 3
                            <span v-if="sortKey === 'disc_3'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('ekspedisi_nama')">
                            Ekspedisi
                            <span v-if="sortKey === 'ekspedisi_nama'">
                                <template v-if="sortAsc">▲</template>
                                <template v-else>▼</template>
                            </span>
                        </th>
                        <th class="text-center" @click="sortBy('ongkir_nama')">
                            Ongkir
                            <span v-if="sortKey === 'ongkir_nama'">
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
                    <tr v-for="toko in paginatedData" :key="toko.id">
                        <td class="text-center">{{ toko.id }}</td>
                        <td class="text-center">{{ toko.kode }}</td>
                        <td class="text-center">{{ toko.nama }}</td>
                        <td class="text-center">{{ toko.kategori_toko_nama }}</td>
                        <td class="text-center">{{ toko.kota_nama }}</td>
                        <td class="text-center">{{ toko.alamat }}</td>
                        <td class="text-center">{{ toko.disc_1*100 }}%</td>
                        <td class="text-center">{{ toko.disc_2*100 }}%</td>
                        <td class="text-center">{{ toko.disc_3*100 }}%</td>
                        <td class="text-center">{{ toko.ekspedisi_nama }}</td>
                        <td class="text-center">{{ toko.ongkir_nama }}</td>
                        <td class="text-center">
                            <button
                                class="btn btn-sm btn-warning"
                                @click="editToko(toko.id)"
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

    function tambahToko() { // func jika button tambah di klik, pindah ke halaman tambah
        router.push('/toko/tambah')
    }

    function editToko(id) {
        router.push({ name: 'EditToko', params: { id } })
    }

    import { ref, computed, watch } from 'vue' // watch wajib di import di setup

    const props = defineProps({
        tokos: {
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

    const filteredTokos = computed(() => {
    let result = props.tokos // pastikan nama var sama dengan const props di line 169

        // Search filter
        if (search.value.trim() !== '') {
            const keyword = search.value.toLowerCase()
            result = result.filter(
                t => // search berdasarkan apa? nama, alamat, dan kontak sesuai kode di bawah
                String(t.id).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                t.kode.toLowerCase().includes(keyword) ||
                t.nama.toLowerCase().includes(keyword) ||
                t.kategori_toko_nama.toLowerCase().includes(keyword) ||
                t.kota_nama.toLowerCase().includes(keyword) ||
                t.alamat.toLowerCase().includes(keyword) ||
                String(t.disc_1).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                String(t.disc_2).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                String(t.disc_3).includes(keyword) || // <-- convert dulu angka ke string agar bisa di filter
                t.ekspedisi_nama.toLowerCase().includes(keyword) ||
                t.ongkir_nama.toLowerCase().includes(keyword)
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
        Math.ceil(filteredTokos.value.length / itemsPerPage)
    )

    const paginatedData = computed(() => {
        const start = (currentPage.value - 1) * itemsPerPage
        return filteredTokos.value.slice(start, start + itemsPerPage)
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
