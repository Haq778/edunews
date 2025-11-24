<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-300">
    <Navbar />

    <!-- Hero Section -->
    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700 text-white overflow-hidden py-20 lg:py-24">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-6">
          Selamat Datang di
          <span class="block text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 to-orange-400 animate-pulse">
            EduNews
          </span>
        </h1>
        <p class="text-xl md:text-2xl mb-12 max-w-3xl mx-auto opacity-90">
          Portal berita pendidikan terkini yang menghubungkan Anda dengan perkembangan terbaru di dunia pendidikan Indonesia
        </p>
        <button
          @click="scrollToNews"
          class="bg-white/20 text-white px-6 py-3 rounded-xl font-semibold hover:bg-white/30 transition"
        >
          Jelajahi Berita
        </button>
      </div>
    </section>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Filter Panel -->
      <div class="flex justify-between items-center mb-8">
        <h2 class="text-2xl lg:text-3xl font-bold text-gray-800 dark:text-white">
          Berita Terkini
          <span class="text-blue-600 dark:text-blue-400 block text-lg mt-1">
            {{ filteredBerita.length }} artikel ditemukan
          </span>
        </h2>
        <button
          @click="isFilterPanelOpen = !isFilterPanelOpen"
          class="px-4 py-2 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl hover:bg-gray-50 dark:hover:bg-gray-700 transition"
        >
          Filter
        </button>
      </div>

      <!-- Filter Panel Collapse -->
      <Transition
        enter-active-class="transition-all duration-300 ease-out"
        enter-from-class="transform opacity-0 -translate-y-4"
        enter-to-class="transform opacity-100 translate-y-0"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="transform opacity-100 translate-y-0"
        leave-to-class="transform opacity-0 -translate-y-4"
      >
        <div
          v-if="isFilterPanelOpen"
          class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border border-gray-200 dark:border-gray-700 p-6 mb-8"
        >
          <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
            <!-- Search -->
            <div class="lg:col-span-2">
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                🔍 Cari Berita
              </label>
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Ketik kata kunci..."
                class="w-full p-4 rounded-xl border border-gray-200 dark:border-gray-600 
                       focus:outline-none focus:ring-3 focus:ring-blue-500 focus:border-transparent
                       bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
              />
            </div>
            <!-- Category -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                📂 Kategori
              </label>
              <select
                v-model="selectedCategory"
                class="w-full p-4 rounded-xl border border-gray-200 dark:border-gray-600 
                       bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
              >
                <option value="">Semua Kategori</option>
                <option 
                  v-for="cat in categories" 
                  :key="cat" 
                  :value="cat"
                >
                  {{ cat }}
                </option>
              </select>
            </div>
            <!-- Sort -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">
                🔄 Urutkan
              </label>
              <select
                v-model="sortBy"
                class="w-full p-4 rounded-xl border border-gray-200 dark:border-gray-600 
                       bg-white dark:bg-gray-700 text-gray-900 dark:text-white"
              >
                <option value="newest">Terbaru</option>
                <option value="oldest">Terlama</option>
                <option value="title">Judul A-Z</option>
                <option value="views">Paling Populer</option>
              </select>
            </div>
          </div>
        </div>
      </Transition>

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center items-center py-20">
        <div class="text-center animate-pulse">
          <div class="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-600 mx-auto mb-4"></div>
          <p class="text-gray-600 dark:text-gray-400">Memuat berita...</p>
        </div>
      </div>

      <!-- Error -->
      <div v-if="error" class="text-center py-20">
        <p class="text-red-500 dark:text-red-400">
          Gagal memuat berita: {{ error.message }}
        </p>
      </div>

      <!-- News List -->
      <div v-if="!loading && !error" id="news-section" :class="viewMode === 'grid'
          ? 'grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6'
          : 'grid grid-cols-1 gap-6'">
        <NewsCard
          v-for="b in sortedBerita"
          :key="b.id"
          :berita="b"
          @article-click="openArticle"
        />
        <div v-if="filteredBerita.length === 0" class="text-center py-20 col-span-full">
          <p class="text-gray-500 dark:text-gray-400">Tidak ada berita yang cocok</p>
        </div>
      </div>
    </div>

    <!-- Article Modal -->
    <ArticleModal
      v-if="isModalOpen && selectedArticle"
      :article="selectedArticle"
      :is-open="isModalOpen"
      @close="closeModal"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import Navbar from '~/components/Navbar.vue'
import NewsCard from '~/components/NewsCard.vue'
import ArticleModal from '~/components/ArticleModal.vue'

interface Berita {
  id: number
  title: string
  content: string
  category: string
  date: string
  views: number
  cover: string
}

const searchQuery = ref('')
const selectedCategory = ref('')
const sortBy = ref('newest')
const loading = ref(true)
const error = ref<Error | null>(null)
const viewMode = ref('grid')
const isFilterPanelOpen = ref(false)
const berita = ref<Berita[]>([])

// Modal
const isModalOpen = ref(false)
const selectedArticle = ref<Berita | null>(null)

const openArticle = (article: Berita) => {
  selectedArticle.value = article
  isModalOpen.value = true
  document.body.style.overflow = 'hidden'
}

const closeModal = () => {
  selectedArticle.value = null
  isModalOpen.value = false
  document.body.style.overflow = 'auto'
}

onMounted(async () => {
  loading.value = true
  error.value = null
  try {
    const res = await fetch('http://localhost:8080/api/berita')
    if (!res.ok) throw new Error(`HTTP ${res.status}`)
    const data: Berita[] = await res.json()
    berita.value = data.map(b => ({
      id: b.id,
      title: b.title || 'Judul kosong',
      content: b.content || '',
      category: b.category || 'Umum',
      date: b.date || new Date().toISOString(),
      views: b.views || 0,
      cover: b.cover || ''
    }))
  } catch (err: unknown) {
    error.value = err as Error
  } finally {
    loading.value = false
  }
})

const categories = computed(() => [...new Set(berita.value.map(b => b.category))])

const filteredBerita = computed(() => {
  return berita.value.filter(b => {
    const searchTerm = searchQuery.value.toLowerCase()
    const matchSearch = b.title.toLowerCase().includes(searchTerm) || b.content.toLowerCase().includes(searchTerm)
    const matchCategory = selectedCategory.value === '' || b.category === selectedCategory.value
    return matchSearch && matchCategory
  })
})

const sortedBerita = computed<Berita[]>(() => {
  const articles = [...filteredBerita.value]
  switch (sortBy.value) {
    case 'newest':
      return articles.sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
    case 'oldest':
      return articles.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime())
    case 'title':
      return articles.sort((a, b) => a.title.localeCompare(b.title))
    case 'views':
      return articles.sort((a, b) => b.views - a.views)
    default:
      return articles
  }
})

const scrollToNews = () => {
  const newsSection = document.getElementById('news-section')
  if (newsSection) newsSection.scrollIntoView({ behavior: 'smooth', block: 'start' })
}
</script>
