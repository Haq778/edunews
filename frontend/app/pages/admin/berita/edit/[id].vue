<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-3xl mx-auto px-4 sm:px-6">

      <!-- Header minimalis -->
      <div class="mb-10">
        <div class="flex items-center gap-4 mb-2">
          <button 
            @click="goBack"
            class="p-2 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 transition-colors rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800"
            aria-label="Kembali"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
          </button>
          <div>
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">Edit Berita</h1>
            <p class="text-gray-600 dark:text-gray-400 text-sm mt-1">Perbarui informasi berita</p>
          </div>
        </div>
      </div>

      <!-- Loading Skeleton -->
      <div v-if="loading" class="space-y-4">
        <div class="h-6 bg-gray-200 dark:bg-gray-700 rounded w-1/3"></div>
        <div class="h-4 bg-gray-200 dark:bg-gray-700 rounded w-1/2"></div>
        <div class="h-32 bg-gray-200 dark:bg-gray-700 rounded"></div>
      </div>

      <!-- Form Card -->
      <div v-else class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700">
        <div class="p-6 md:p-8">
          <ArticleForm
            :modelValue="article"
            :categories="categories"
            submitText="Simpan Perubahan"
            :loading="false"
            @save="update"
            @cancel="goBack"
          />

          <!-- Action Buttons - minimal dan rapi -->
          <div class="flex justify-between items-center mt-8 pt-6 border-t border-gray-100 dark:border-gray-700">
            <button 
              @click="remove"
              class="px-4 py-2 text-sm text-red-600 hover:text-red-700 hover:bg-red-50 dark:text-red-400 dark:hover:text-red-300 dark:hover:bg-red-900/30 rounded-lg transition-colors"
            >
              Hapus Berita
            </button>
            <button 
              @click="goBack"
              class="px-4 py-2 text-sm text-gray-600 hover:text-gray-700 hover:bg-gray-50 dark:text-gray-400 dark:hover:text-gray-300 dark:hover:bg-gray-700 rounded-lg transition-colors"
            >
              Batal
            </button>
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