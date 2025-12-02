import api from './api'

export async function getAllBarang() {
    try {
        const response = await api.get('/barang') // endpoint
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data barang:', error)
        throw error
    }
}