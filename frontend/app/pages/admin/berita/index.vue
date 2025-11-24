<template>
  <div class="p-6 max-w-6xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Daftar Berita</h1>

    <!-- Tombol tambah -->
    <div class="mb-4">
      <NuxtLink
        to="/admin/berita/tambah"
        class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
      >
        + Tambah Berita
      </NuxtLink>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="text-gray-500">Memuat berita...</div>

    <!-- Tidak ada berita -->
    <div v-else-if="beritaList.length === 0" class="text-gray-500">
      Belum ada berita.
    </div>

    <!-- Tabel berita -->
    <table v-else class="w-full border-collapse border">
      <thead>
        <tr class="bg-gray-100">
          <th class="border p-2">ID</th>
          <th class="border p-2">Judul</th>
          <th class="border p-2">Kategori</th>
          <th class="border p-2">Cover</th>
          <th class="border p-2">Aksi</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="item in beritaList"
          :key="item.id"
          class="hover:bg-gray-50"
        >
          <td class="border p-2">{{ item.id }}</td>
          <td class="border p-2">{{ item.title }}</td>
          <td class="border p-2">{{ item.category }}</td>

          <!-- Cover -->
          <td class="border p-2">
            <img
              v-if="item.cover"
              :src="coverUrl(item.cover)"
              class="h-20 w-20 object-cover border rounded"
              alt="Cover Berita"
              @error="hideBrokenImage"
            />
            <span v-else class="text-gray-400">Tidak ada</span>
          </td>

          <!-- Aksi -->
          <td class="border p-2 flex gap-2">
            <NuxtLink
              :to="`/admin/berita/edit/${item.id}`"
              class="px-3 py-1 bg-yellow-500 text-white rounded hover:bg-yellow-600"
            >
              Edit
            </NuxtLink>

            <button
              class="px-3 py-1 bg-red-600 text-white rounded hover:bg-red-700"
              @click="hapus(item.id)"
            >
              Hapus
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

definePageMeta({ layout: 'admin' })

const beritaList = ref<any[]>([])
const loading = ref(true)

/*  
  FIX: Menghapus /default-cover.jpg untuk menghindari warning router.
  Jika cover tidak ada → tampilkan "Tidak ada" saja.
*/
const coverUrl = (cover: string | null) => {
  if (!cover) return undefined
  return `http://localhost:8080/uploads/berita/${cover}`
}

// Hides broken image without using default-cover.jpg
const hideBrokenImage = (e: Event) => {
  const target = e.target as HTMLImageElement
  target.style.display = 'none'
}

async function loadBerita() {
  try {
    const res = await fetch('http://localhost:8080/api/berita')
    if (!res.ok) throw new Error(await res.text())
    beritaList.value = await res.json()
  } catch (err: any) {
    console.error('Gagal memuat berita:', err)
    alert('Gagal memuat berita, cek console')
  } finally {
    loading.value = false
  }
}

onMounted(loadBerita)

const hapus = async (id: number) => {
  if (!confirm('Yakin ingin menghapus berita ini?')) return

  try {
    const res = await fetch(`http://localhost:8080/api/berita/${id}`, {
      method: 'DELETE'
    })
    if (!res.ok) throw new Error(await res.text())

    alert('Berita berhasil dihapus.')
    loadBerita()
  } catch (err: any) {
    console.error('Gagal menghapus berita:', err)
    alert('Gagal menghapus berita, cek console')
  }
}
</script>
