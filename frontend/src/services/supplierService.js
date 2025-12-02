import api from './api'

export async function getAllSuppliers() {
    try {
        const response = await api.get('/supplier') // endpoint sekarang jadi /api/supplier
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data supplier:', error)
        throw error
    }
}