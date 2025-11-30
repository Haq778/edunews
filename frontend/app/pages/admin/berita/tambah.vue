<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6">

      <!-- Header -->
      <div class="flex items-center justify-between mb-8">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white">Tambah Berita Baru</h1>
          <p class="text-gray-600 dark:text-gray-400">Buat artikel berita terbaru</p>
        </div>

        <button @click="goBack" class="px-4 py-2 text-gray-600 dark:text-gray-300">
          Kembali
        </button>
      </div>

      <!-- Card -->
      <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border">
        <div class="p-6">
          <ArticleForm
            :modelValue="{}"
            :categories="categories"
            submitText="Tambah Berita"
            :loading="loading"
            @save="save"
            @cancel="goBack"
          />
        </div>
      </div>

      <!-- Loading Overlay -->
      <div v-if="loading" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
        <div class="bg-white dark:bg-gray-800 p-8 rounded-2xl">
          <div class="animate-spin h-12 w-12 border-b-2 border-blue-600 rounded-full mx-auto"></div>
          <p class="text-center mt-4 text-gray-700 dark:text-gray-300">Menyimpan...</p>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import { useRouter } from "vue-router"
import ArticleForm from "~/components/ArticleForm.vue"

definePageMeta({ layout: "admin" })

const router = useRouter()
const categories = ref([])
const loading = ref(false)

async function loadCategories() {
  try {
    const res = await fetch("http://localhost:8080/api/categories")
    if (!res.ok) throw new Error("Gagal fetch kategori")

    categories.value = await res.json()
  } catch (e) {
    console.error(e)
    categories.value = [] // fallback
  }
}

async function save(payload) {
  try {
    loading.value = true

    const formData = new FormData()
    formData.append("title", payload.data.title)
    formData.append("content", payload.data.content)
    formData.append("category_id", payload.data.category_id)

    if (payload.file) {
      formData.append("cover", payload.file)
    }

    const res = await fetch("http://localhost:8080/api/berita", {
      method: "POST",
      body: formData
    })

    if (!res.ok) throw new Error("Gagal menyimpan berita")

    router.push("/admin/berita")
  } catch (e) {
    alert(e.message)
  } finally {
    loading.value = false
  }
}

function goBack() {
  router.push("/admin/berita")
}

onMounted(loadCategories)
</script>
