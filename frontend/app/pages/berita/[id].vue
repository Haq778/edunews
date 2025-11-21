<template>
  <div>
    <Navbar />
    <main class="max-w-3xl mx-auto p-6">
      <img :src="item.gambar" class="rounded-lg mb-4">
      <h1 class="text-3xl font-bold mb-2">{{ item.judul }}</h1>
      <p class="text-gray-600 mb-4">{{ item.penulis }} - {{ item.tanggal }}</p>
      <p class="text-lg leading-relaxed">{{ item.isi }}</p>
      <NuxtLink to="/" class="block mt-6 text-blue-600 hover:underline">← Kembali ke Beranda</NuxtLink>
    </main>
  </div>
</template>

<script setup>
import berita from '../berita'
import { useRoute } from 'vue-router'
import Navbar from '~/components/Navbar.vue'

const route = useRoute()
const item = berita.find(b => b.id === Number(route.params.id))
</script>
