import api from './api'

export async function getAllHargaBarang() {
    try {
        const response = await api.get('/harga_barang') // endpoint
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data harga barang:', error)
        throw error
    }
}