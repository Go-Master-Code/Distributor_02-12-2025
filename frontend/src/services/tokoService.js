import api from './api'

export async function getAllToko() {
    try {
        const response = await api.get('/toko') // endpoint
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data toko:', error)
        throw error
    }
}