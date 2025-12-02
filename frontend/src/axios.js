// src/axios.js
import axios from 'axios'

const instance = axios.create({
  baseURL: 'http://localhost:3000', // ganti sesuai URL backend Gin kamu
  timeout: 5000, // optional, biar request nggak gantung lama
})

// Optional: interceptor untuk log error (bisa kamu hapus kalau belum perlu)
instance.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('API error:', error.response?.data || error.message)
    return Promise.reject(error)
  }
)

export default instance
