<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-blue-50/20 dark:from-gray-900 dark:to-gray-800 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">

      <!-- Loading -->
      <div v-if="loading" class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border p-8 animate-pulse">
        <div class="h-6 bg-gray-200 dark:bg-gray-700 rounded w-1/3 mb-4"></div>
        <div class="h-4 bg-gray-200 dark:bg-gray-700 rounded w-1/2 mb-4"></div>
        <div class="h-32 bg-gray-200 dark:bg-gray-700 rounded"></div>
      </div>

      <!-- FORM -->
      <div v-else class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border overflow-hidden">
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-gray-700 dark:to-gray-800 px-6 py-4 border-b">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white">Formulir Edit Berita</h2>
          <p class="text-gray-600 dark:text-gray-400 text-sm mt-1">Perbarui informasi berita</p>
        </div>

        <div class="p-6">
          <ArticleForm
            :modelValue="article"
            :categories="categories"
            submitText="Simpan Perubahan"
            @save="update"
            @cancel="goBack"
          />

          <div class="flex gap-3 mt-6">
            <button @click="goBack" class="px-4 py-2 rounded bg-gray-100 dark:bg-gray-700">Kembali</button>
            <button @click="remove" class="px-4 py-2 rounded bg-red-100 text-red-700">Hapus</button>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from '#imports'
import ArticleForm from '~/components/ArticleForm.vue'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()

const loading = ref<boolean>(true)
const article = ref<any>(null)
const categories = ref<any[]>([])

function parseId(raw: any) {
  const n = Number(raw)
  return Number.isFinite(n) && n > 0 ? n : null
}

async function loadData() {
  loading.value = true

  const id = parseId(route.params.id)
  if (!id) {
    alert('ID tidak valid')
    loading.value = false
    return
  }

  try {
    const [rBerita, rCategories] = await Promise.all([
      fetch(`http://localhost:8080/api/berita/${id}`),
      fetch(`http://localhost:8080/api/categories`)
    ])

    if (!rBerita.ok) throw new Error(await rBerita.text())
    if (!rCategories.ok) throw new Error(await rCategories.text())

    const dataBerita = await rBerita.json()
    const dataCats = await rCategories.json()

    // 🔥 FIX UTAMA: CATEGORY ID TIDAK BOLEH 0
    const fixedCategoryId =
      Number(dataBerita.category_id) > 0
        ? Number(dataBerita.category_id)
        : ''

    article.value = {
      id: dataBerita.id,
      title: dataBerita.title ?? '',
      content: dataBerita.content ?? '',
      excerpt: dataBerita.excerpt ?? '',
      category_id: fixedCategoryId,
      cover_url: dataBerita.cover
        ? `http://localhost:8080/uploads/berita/${dataBerita.cover}`
        : ''
    }

    categories.value = Array.isArray(dataCats) ? dataCats : []

  } catch (err) {
    console.error('Gagal memuat data:', err)
    alert('Gagal memuat data. Cek console.')
    router.push('/admin/berita')
  } finally {
    loading.value = false
  }
}

async function update(payload: any) {
  const id = parseId(route.params.id)
  if (!id) return alert("ID tidak valid")

  try {
    const fd = new FormData()

    // append semua data
    Object.entries(payload.data).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        fd.append(key, String(value))
      }
    })

    // jika ada file, append juga
    if (payload.file) {
      fd.append("cover", payload.file)
    }

    const res = await fetch(`http://localhost:8080/api/berita/${id}`, {
      method: "PUT",
      body: fd
    })

    if (!res.ok) {
      throw new Error(await res.text())
    }

    alert("Berita berhasil disimpan")
    router.push("/admin/berita")

  } catch (err) {
    console.error("Gagal update berita:", err)
    alert("Gagal update berita, cek console")
  }
}


async function remove() {
  const id = parseId(route.params.id)
  if (!id) return alert('ID tidak valid')
  if (!confirm('Yakin ingin menghapus berita?')) return

  try {
    const res = await fetch(`http://localhost:8080/api/berita/${id}`, { method: 'DELETE' })
    if (!res.ok) throw new Error(await res.text())
    alert('Berita berhasil dihapus')
    router.push('/admin/berita')
  } catch (err) {
    console.error('Gagal menghapus berita:', err)
    alert('Gagal hapus berita')
  }
}

function goBack() {
  router.back()
}

onMounted(loadData)
</script>

<style scoped>
</style>
