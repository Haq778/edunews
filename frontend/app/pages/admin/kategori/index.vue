<template>
  <div class="p-6 max-w-3xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Kategori Berita</h1>

    <NuxtLink
      to="/admin/kategori/tambah"
      class="btn-primary mb-4 inline-block"
    >
      + Tambah Kategori
    </NuxtLink>

    <table class="w-full border">
      <tr class="bg-gray-100">
        <th class="border p-2">ID</th>
        <th class="border p-2">Nama</th>
        <th class="border p-2">Aksi</th>
      </tr>

      <tr v-for="kat in kategori" :key="kat.id">
        <td class="border p-2">{{ kat.id }}</td>
        <td class="border p-2">{{ kat.name }}</td>
        <td class="border p-2">
          <NuxtLink
            :to="`/admin/kategori/edit/${kat.id}`"
            class="btn-yellow mr-2"
          >
            Edit
          </NuxtLink>

          <button class="btn-red" @click="hapus(kat.id)">Hapus</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script setup>
definePageMeta({ layout: "admin" })

import { ref, onMounted } from 'vue'

const kategori = ref([])

onMounted(async () => {
  const res = await fetch("http://localhost:8080/api/kategori")
  kategori.value = await res.json()
})

const hapus = async (id) => {
  if (!confirm("Yakin ingin hapus kategori?")) return

  const res = await fetch(`http://localhost:8080/api/kategori/${id}`, {
    method: "DELETE"
  })
  if (!res.ok) return alert("Gagal menghapus")

  kategori.value = kategori.value.filter(k => k.id !== id)
}
</script>

<style scoped>
.btn-primary { @apply bg-blue-600 text-white px-3 py-1 rounded; }
.btn-yellow { @apply bg-yellow-500 text-white px-3 py-1 rounded; }
.btn-red { @apply bg-red-600 text-white px-3 py-1 rounded; }
</style>
