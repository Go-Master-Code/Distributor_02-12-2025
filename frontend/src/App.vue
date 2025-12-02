<template>
  <div>
    <nav>
      <router-link to="/">Home</router-link> |
      <router-link to="/supplier">Supplier</router-link> |
      <router-link to="/toko">Toko</router-link> |
      <router-link to="/artikel">Artikel</router-link> |
      <router-link to="/barang">Barang</router-link> |
      <router-link to="/kartu_stok">Kartu Stok</router-link> |
      <router-link to="/penjualan">Penjualan</router-link>
    </nav>
    <router-view></router-view>
  </div>
</template>
