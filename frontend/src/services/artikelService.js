import api from './api'

export async function getAllArtikel() {
    try {
        const response = await api.get('/artikel') // endpoint sekarang jadi /api/supplier
        return response.data.data // sesuai dengan response body json dari api (cek pakai postman)
    } catch (error) {
        console.error('Gagal mengambil data artikel:', error)
        throw error
    }
}