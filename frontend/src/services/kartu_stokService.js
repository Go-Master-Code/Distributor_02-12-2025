import api from './api'

export async function getAllKartuStok() {
    try {
        const response = await api.get('/kartu_stok') // endpoint
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data kartu stok:', error)
        throw error
    }
}