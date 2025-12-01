<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-3xl mx-auto px-4 sm:px-6">

      <!-- Header yang lebih minimalis -->
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
          <h1 class="text-2xl font-bold text-gray-900 dark:text-white">Tambah Berita Baru</h1>
        </div>
        <p class="text-gray-600 dark:text-gray-400 ml-12">Buat artikel berita terbaru untuk ditampilkan</p>
      </div>

      <!-- Form Card dengan desain lebih bersih -->
      <div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700">
        <div class="p-6 md:p-8">
          <ArticleForm
            :modelValue="{}"
            :categories="categories"
            submitText="Simpan Berita"
            :loading="loading"
            @save="save"
            @cancel="goBack"
          />
        </div>
      </div>

      <!-- Loading Overlay yang lebih halus -->
      <div v-if="loading" class="fixed inset-0 bg-black/40 backdrop-blur-sm flex items-center justify-center z-50">
        <div class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-xl min-w-[200px]">
          <div class="flex flex-col items-center gap-3">
            <div class="animate-spin h-8 w-8 border-2 border-blue-500 border-t-transparent rounded-full"></div>
            <p class="text-sm text-gray-600 dark:text-gray-300">Menyimpan berita...</p>
          </div>
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