<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 via-blue-50/30 to-indigo-50/20 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 transition-colors duration-500 overflow-x-hidden">
    <Navbar />

    <!-- Enhanced Hero Section with Advanced Animations -->
    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-800 text-white overflow-hidden py-16 lg:py-20">
      <!-- Advanced Background Animations -->
      <div class="absolute inset-0 overflow-hidden">
        <!-- Floating Geometric Shapes -->
        <div class="absolute top-1/4 left-1/4 w-32 h-32 bg-white/10 rounded-full blur-3xl animate-float-slow"></div>
        <div class="absolute top-1/3 right-1/4 w-24 h-24 bg-cyan-400/20 rounded-full blur-2xl animate-float-medium delay-1000"></div>
        <div class="absolute bottom-1/4 left-1/3 w-28 h-28 bg-purple-400/15 rounded-full blur-3xl animate-float-slow delay-2000"></div>
        <div class="absolute top-3/4 right-1/3 w-20 h-20 bg-blue-400/25 rounded-full blur-2xl animate-float-fast delay-1500"></div>

        <!-- Animated Grid Pattern -->
        <div class="absolute inset-0 opacity-20">
          <div class="w-full h-full bg-grid-white animate-grid-drift"></div>
        </div>

        <!-- Floating Particles -->
        <div 
          v-for="i in 12" 
          :key="i"
          class="absolute w-2 h-2 bg-white/40 rounded-full animate-particle-drift"
          :style="{
            top: Math.random() * 100 + '%',
            left: Math.random() * 100 + '%',
            animationDelay: Math.random() * 8 + 's',
            animationDuration: (15 + Math.random() * 10) + 's'
          }"
        ></div>

        <!-- Connection Lines -->
        <div class="absolute inset-0 opacity-30">
          <div v-for="i in 8" :key="i" 
               class="absolute w-0.5 h-32 bg-gradient-to-b from-white/60 to-transparent animate-connection-sweep"
               :style="{
                 left: (i * 12.5) + '%',
                 animationDelay: (i * 0.5) + 's',
                 animationDuration: (6 + Math.random() * 3) + 's'
               }">
          </div>
        </div>

        <!-- Pulse Rings -->
        <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <div class="w-64 h-64 border-2 border-white/20 rounded-full animate-pulse-ring"></div>
          <div class="w-96 h-96 border-2 border-white/15 rounded-full animate-pulse-ring delay-1000"></div>
          <div class="w-128 h-128 border-2 border-white/10 rounded-full animate-pulse-ring delay-2000"></div>
        </div>

        <!-- Floating Text Background -->
        <div class="absolute inset-0 opacity-5">
          <div class="absolute top-10 left-10 text-6xl font-black animate-text-float">EDU</div>
          <div class="absolute bottom-10 right-10 text-6xl font-black animate-text-float delay-3000">NEWS</div>
          <div class="absolute top-20 right-20 text-4xl font-black animate-text-float delay-1000">📚</div>
          <div class="absolute bottom-20 left-20 text-4xl font-black animate-text-float delay-2000">🎓</div>
        </div>
      </div>

      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <!-- Animated Badge -->
        <div class="inline-flex items-center px-4 py-2 bg-white/20 backdrop-blur-xl rounded-xl text-sm font-medium mb-6 border border-white/30 transform hover:scale-105 transition-all duration-300 group/badge animate-badge-float">
          <span class="w-2 h-2 bg-green-400 rounded-full mr-2 animate-pulse group-hover/badge:scale-150 transition-transform duration-300"></span>
          🎓 Portal Berita Pendidikan Terpercaya
          <div class="absolute inset-0 rounded-xl bg-gradient-to-r from-transparent via-white/10 to-transparent transform -translate-x-full group-hover/badge:translate-x-full transition-transform duration-1000"></div>
        </div>

        <!-- Enhanced Hero Text -->
        <div class="space-y-4 animate-text-slide-up">
          <h1 class="text-4xl md:text-5xl font-bold mb-4 leading-tight">
            Selamat Datang di
            <span class="block text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 via-orange-400 to-red-400 animate-gradient-x">
              EduNews
            </span>
          </h1>
          
          <p class="text-lg md:text-xl max-w-2xl mx-auto opacity-90 leading-relaxed animate-text-slide-up delay-200">
            Jelajahi berita pendidikan terkini dan perkembangan terbaru
          </p>
        </div>

        <!-- Enhanced CTA Buttons -->
        <div class="flex flex-col sm:flex-row gap-4 justify-center items-center mt-8 animate-text-slide-up delay-400">
          <button
            @click="scrollToNews"
            class="group bg-white/20 text-white px-6 py-3 rounded-xl font-medium hover:bg-white/30 
                   transition-all duration-300 transform hover:scale-110 backdrop-blur-xl border border-white/30
                   flex items-center space-x-2 shadow-lg hover:shadow-2xl relative overflow-hidden"
          >
            <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
            <span>📰 Jelajahi Berita</span>
            <svg class="w-4 h-4 group-hover:translate-y-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
        </div>

        <!-- Enhanced Stats -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-6 mt-12 max-w-md mx-auto animate-text-slide-up delay-600">
          <div 
            v-for="(stat, index) in statsArray" 
            :key="stat.label"
            class="text-center group/stat transform hover:scale-110 transition-all duration-300"
            :style="{ animationDelay: `${index * 100}ms` }"
          >
            <div class="text-2xl font-bold mb-1 animate-count-up" :class="stat.color">
              {{ stat.value.toLocaleString() }}+
            </div>
            <div class="text-white/70 text-xs font-medium group-hover/stat:text-white transition-colors duration-300">
              {{ stat.label }}
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Compact Featured Categories -->
    <section class="py-12 bg-white/50 dark:bg-gray-800/50 backdrop-blur-sm relative overflow-hidden">
      <!-- Background Elements -->
      <div class="absolute inset-0 opacity-10">
        <div class="w-full h-full bg-grid-gray animate-grid-move-slow"></div>
      </div>

      <!-- Floating Elements -->
      <div class="absolute top-10 right-10 w-16 h-16 bg-blue-500/10 rounded-full blur-xl animate-pulse-slow"></div>
      <div class="absolute bottom-10 left-10 w-12 h-12 bg-purple-500/10 rounded-full blur-xl animate-pulse-slow delay-1000"></div>

      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-10 animate-text-slide-up">
          <h2 class="text-2xl md:text-3xl font-bold text-gray-900 dark:text-white mb-3">
            Kategori <span class="text-blue-600 dark:text-blue-400">Populer</span>
          </h2>
          <p class="text-gray-600 dark:text-gray-400 text-sm">
            Temukan berita berdasarkan kategori yang diminati
          </p>
        </div>

        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div 
            v-for="(category, index) in featuredCategories" 
            :key="category.name"
            class="group bg-white dark:bg-gray-700 rounded-2xl p-4 text-center cursor-pointer hover:shadow-lg transition-all duration-300 
                   transform hover:scale-105 border border-gray-200 dark:border-gray-600 relative overflow-hidden backdrop-blur-sm animate-card-enter"
            @click="filterByCategory(category.name)"
            :style="{ animationDelay: `${index * 100}ms` }"
          >
            <!-- Hover Effect -->
            <div class="absolute inset-0 bg-gradient-to-br from-blue-500/5 to-purple-500/5 opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
            
            <!-- Animated Icon -->
            <div class="w-12 h-12 mx-auto mb-3 rounded-xl flex items-center justify-center text-xl 
                       group-hover:scale-110 group-hover:rotate-12 transition-all duration-500 relative"
                 :class="category.color">
              <div class="absolute inset-0 rounded-xl bg-gradient-to-br from-white/20 to-transparent"></div>
              {{ category.icon }}
            </div>
            
            <h3 class="font-semibold text-gray-900 dark:text-white mb-1 text-sm group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors duration-300">
              {{ category.name }}
            </h3>
            <p class="text-xs text-gray-600 dark:text-gray-400 font-medium">{{ category.count }} Artikel</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Enhanced News Section with Advanced Animations -->
    <section class="py-12 bg-gradient-to-br from-gray-50/50 to-blue-50/30 dark:from-gray-900 dark:to-gray-800 relative overflow-hidden">
      <!-- Advanced Background Animations -->
      <div class="absolute inset-0 overflow-hidden">
        <!-- Animated Grid -->
        <div class="absolute inset-0 opacity-10">
          <div class="w-full h-full bg-grid-blue animate-grid-move"></div>
        </div>

        <!-- Floating Dots -->
        <div 
          v-for="i in 15" 
          :key="i"
          class="absolute w-3 h-3 bg-blue-500/20 rounded-full animate-bounce-float"
          :style="{
            top: Math.random() * 100 + '%',
            left: Math.random() * 100 + '%',
            animationDelay: Math.random() * 5 + 's',
            animationDuration: (3 + Math.random() * 4) + 's'
          }"
        ></div>

        <!-- Pulse Waves -->
        <div class="absolute top-1/4 left-1/4">
          <div class="w-8 h-8 bg-blue-400/30 rounded-full animate-pulse-wave"></div>
        </div>
        <div class="absolute bottom-1/3 right-1/4">
          <div class="w-6 h-6 bg-purple-400/30 rounded-full animate-pulse-wave delay-1500"></div>
        </div>
        <div class="absolute top-2/3 left-1/2">
          <div class="w-10 h-10 bg-cyan-400/20 rounded-full animate-pulse-wave delay-3000"></div>
        </div>

        <!-- Floating Shapes -->
        <div class="absolute top-20 right-20 w-16 h-16 border-2 border-blue-300/20 rounded-lg rotate-45 animate-float-slow"></div>
        <div class="absolute bottom-20 left-20 w-12 h-12 border-2 border-purple-300/20 rounded-full animate-float-medium delay-2000"></div>
        <div class="absolute top-40 left-20 w-8 h-8 border-2 border-cyan-300/20 rotate-12 animate-float-fast delay-1000"></div>
      </div>

      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <!-- Enhanced Header -->
        <div class="flex flex-col lg:flex-row justify-between items-start lg:items-center mb-8 gap-4 animate-text-slide-up">
          <div class="space-y-2">
            <div class="flex items-center space-x-3">
              <div class="w-3 h-3 bg-blue-500 rounded-full animate-pulse"></div>
              <h2 class="text-2xl lg:text-3xl font-bold text-gray-900 dark:text-white">
                Berita <span class="text-blue-600 dark:text-blue-400">Terkini</span>
              </h2>
            </div>
            <p class="text-gray-600 dark:text-gray-400 text-sm">
              <span class="font-semibold text-blue-600 dark:text-blue-400">{{ filteredBerita.length }}</span> artikel ditemukan
              <span v-if="selectedCategory" class="text-green-600 dark:text-green-400 font-bold ml-2">
                • {{ selectedCategory }}
              </span>
            </p>
          </div>

          <div class="flex gap-3">
            <!-- View Toggle -->
            <div class="flex bg-white dark:bg-gray-700 rounded-xl p-1 border border-gray-200 dark:border-gray-600 shadow-sm">
              <button
                @click="viewMode = 'grid'"
                :class="[
                  'p-2 rounded-lg transition-all duration-300 transform hover:scale-110',
                  viewMode === 'grid' 
                    ? 'bg-blue-500 text-white shadow-md scale-105' 
                    : 'text-gray-600 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400'
                ]"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"/>
                </svg>
              </button>
              <button
                @click="viewMode = 'list'"
                :class="[
                  'p-2 rounded-lg transition-all duration-300 transform hover:scale-110',
                  viewMode === 'list' 
                    ? 'bg-blue-500 text-white shadow-md scale-105' 
                    : 'text-gray-600 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-400'
                ]"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
                </svg>
              </button>
            </div>

            <!-- Filter Button -->
            <button
              @click="isFilterPanelOpen = !isFilterPanelOpen"
              class="px-4 py-2 bg-white dark:bg-gray-700 border border-gray-200 dark:border-gray-600 
                     rounded-xl hover:bg-gray-50 dark:hover:bg-gray-600 transition-all duration-300 
                     transform hover:scale-105 flex items-center space-x-2 text-sm font-medium shadow-sm"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z"/>
              </svg>
              <span>Filter</span>
            </button>
          </div>
        </div>

        <!-- Compact Filter Panel -->
        <Transition
          enter-active-class="transition-all duration-500 ease-out"
          enter-from-class="transform opacity-0 -translate-y-4 scale-95"
          enter-to-class="transform opacity-100 translate-y-0 scale-100"
          leave-active-class="transition-all duration-300 ease-in"
          leave-from-class="transform opacity-100 translate-y-0 scale-100"
          leave-to-class="transform opacity-0 -translate-y-4 scale-95"
        >
          <div
            v-if="isFilterPanelOpen"
            class="bg-white dark:bg-gray-700 rounded-2xl shadow-lg border border-gray-200 dark:border-gray-600 p-6 mb-6"
          >
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <!-- Search -->
              <div>
                <label class="text-xs font-bold text-gray-700 dark:text-gray-300 mb-2 block">Cari Berita</label>
                <input
                  v-model="searchQuery"
                  type="text"
                  placeholder="Kata kunci..."
                  class="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 
                         focus:outline-none focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500
                         bg-white dark:bg-gray-600 text-gray-900 dark:text-white transition-all duration-300 text-sm"
                />
              </div>
              
              <!-- Category -->
              <div>
                <label class="text-xs font-bold text-gray-700 dark:text-gray-300 mb-2 block">Kategori</label>
                <select
                  v-model="selectedCategory"
                  class="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 
                         bg-white dark:bg-gray-600 text-gray-900 dark:text-white focus:outline-none 
                         focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all duration-300 text-sm"
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
                <label class="text-xs font-bold text-gray-700 dark:text-gray-300 mb-2 block">Urutkan</label>
                <select
                  v-model="sortBy"
                  class="w-full p-3 rounded-lg border border-gray-300 dark:border-gray-600 
                         bg-white dark:bg-gray-600 text-gray-900 dark:text-white focus:outline-none 
                         focus:ring-2 focus:ring-blue-500/50 focus:border-blue-500 transition-all duration-300 text-sm"
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

        <!-- Loading Animation -->
        <div v-if="loading" class="flex justify-center items-center py-16">
          <div class="text-center space-y-4">
            <div class="relative">
              <div class="animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent mx-auto mb-2"></div>
            </div>
            <p class="text-gray-600 dark:text-gray-400 text-sm">Memuat berita...</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-if="error" class="text-center py-16">
          <div class="w-16 h-16 mx-auto mb-4 text-red-400">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.082 16.5c-.77.833.192 2.5 1.732 2.5z"/>
            </svg>
          </div>
          <p class="text-red-500 dark:text-red-400 text-sm mb-4">
            Gagal memuat berita
          </p>
          <button 
            @click="fetchBerita"
            class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-all duration-300 transform hover:scale-105 text-sm font-medium"
          >
            🔄 Coba Lagi
          </button>
        </div>

        <!-- News Grid/List -->
        <div v-if="!loading && !error" id="news-section" :class="[
          'transition-all duration-500',
          viewMode === 'grid' 
            ? 'grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6' 
            : 'grid grid-cols-1 gap-4'
        ]">
          <NewsCard
            v-for="(b, index) in displayedBerita"
            :key="b.id"
            :berita="b"
            :view-mode="viewMode"
            @article-click="openArticle"
            :class="[
              'transition-all duration-500 transform hover:scale-105',
              showCards ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-4'
            ]"
            :style="{
              transitionDelay: `${index * 50}ms`
            }"
          />
          
          <!-- Empty State -->
          <div v-if="filteredBerita.length === 0" class="text-center py-12 col-span-full">
            <div class="w-12 h-12 mx-auto mb-4 text-gray-400">
              <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-2">Tidak ada berita</h3>
            <p class="text-gray-600 dark:text-gray-400 mb-4 text-sm">Coba ubah filter pencarian</p>
            <button 
              @click="clearFilters"
              class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-all duration-300 transform hover:scale-105 text-sm font-medium"
            >
              🔄 Reset Filter
            </button>
          </div>
        </div>

        <!-- Load More Button -->
        <div v-if="!loading && filteredBerita.length > visibleCount" class="text-center mt-8">
          <button
            @click="loadMore"
            class="px-8 py-3 bg-blue-500 text-white rounded-xl hover:bg-blue-600 transition-all duration-300 transform hover:scale-105 
                   font-medium shadow-lg hover:shadow-xl flex items-center space-x-2 mx-auto"
          >
            <span>📰 Muat Lebih Banyak</span>
            <svg class="w-4 h-4 transform group-hover:translate-y-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>
        </div>
      </div>
    </section>

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

// Reactive state
const searchQuery = ref('')
const selectedCategory = ref('')
const sortBy = ref('newest')
const loading = ref(true)
const error = ref<Error | null>(null)
const viewMode = ref('grid')
const isFilterPanelOpen = ref(false)
const berita = ref<Berita[]>([])
const showCards = ref(false)
const visibleCount = ref(6)

// Stats
const statsArray = ref([
  { value: 1250, label: 'Artikel', color: 'text-yellow-400' },
  { value: 15, label: 'Kategori', color: 'text-green-400' },
  { value: 25000, label: 'Pembaca', color: 'text-blue-400' },
  { value: 45, label: 'Ahli', color: 'text-purple-400' }
])

// Featured categories
const featuredCategories = ref([
  { name: 'Kurikulum', icon: '📚', count: 156, color: 'bg-blue-100 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400' },
  { name: 'Teknologi', icon: '💻', count: 89, color: 'bg-green-100 text-green-600 dark:bg-green-900/30 dark:text-green-400' },
  { name: 'Beasiswa', icon: '🎓', count: 67, color: 'bg-purple-100 text-purple-600 dark:bg-purple-900/30 dark:text-purple-400' },
  { name: 'Guru & Dosen', icon: '👨‍🏫', count: 134, color: 'bg-orange-100 text-orange-600 dark:bg-orange-900/30 dark:text-orange-400' }
])

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

const fetchBerita = async () => {
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
    setTimeout(() => {
      showCards.value = true
    }, 100)
  }
}

const filterByCategory = (category: string) => {
  selectedCategory.value = category
  isFilterPanelOpen.value = false
  scrollToNews()
}

const clearFilters = () => {
  searchQuery.value = ''
  selectedCategory.value = ''
  sortBy.value = 'newest'
  visibleCount.value = 6
}

const loadMore = () => {
  visibleCount.value += 6
}

const scrollToNews = () => {
  const newsSection = document.getElementById('news-section')
  if (newsSection) newsSection.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

onMounted(() => {
  fetchBerita()
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

const displayedBerita = computed(() => {
  return sortedBerita.value.slice(0, visibleCount.value)
})
</script>

<style scoped>
/* Advanced Animation System */
@keyframes gradient-x {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

@keyframes float-slow {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-20px) translateX(10px) rotate(120deg);
  }
  66% {
    transform: translateY(10px) translateX(-5px) rotate(240deg);
  }
}

@keyframes float-medium {
  0%, 100% {
    transform: translateY(0px) translateX(0px) scale(1);
  }
  50% {
    transform: translateY(-15px) translateX(8px) scale(1.05);
  }
}

@keyframes float-fast {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-10px) rotate(180deg);
  }
}

@keyframes particle-drift {
  0%, 100% {
    transform: translate(0, 0) rotate(0deg);
    opacity: 0.6;
  }
  25% {
    transform: translate(30px, -40px) rotate(90deg);
    opacity: 0.8;
  }
  50% {
    transform: translate(-20px, -60px) rotate(180deg);
    opacity: 1;
  }
  75% {
    transform: translate(-30px, -40px) rotate(270deg);
    opacity: 0.8;
  }
}

@keyframes connection-sweep {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  20% {
    opacity: 1;
  }
  80% {
    opacity: 1;
  }
  100% {
    transform: translateY(100vh);
    opacity: 0;
  }
}

@keyframes grid-drift {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 100px 100px;
  }
}

@keyframes grid-move {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 50px 50px;
  }
}

@keyframes grid-move-slow {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 80px 80px;
  }
}

@keyframes pulse-ring {
  0% {
    transform: scale(0.8);
    opacity: 1;
  }
  100% {
    transform: scale(2.5);
    opacity: 0;
  }
}

@keyframes pulse-wave {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  100% {
    transform: scale(3);
    opacity: 0;
  }
}

@keyframes text-float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.05;
  }
  50% {
    transform: translateY(-10px) rotate(5deg);
    opacity: 0.08;
  }
}

@keyframes badge-float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-5px);
  }
}

@keyframes text-slide-up {
  0% {
    opacity: 0;
    transform: translateY(30px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes card-enter {
  0% {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes count-up {
  0% {
    transform: scale(0.5);
    opacity: 0;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes bounce-float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-15px);
  }
}

/* Animation Classes */
.animate-gradient-x {
  background-size: 200% 200%;
  animation: gradient-x 3s ease infinite;
}

.animate-float-slow {
  animation: float-slow 20s ease-in-out infinite;
}

.animate-float-medium {
  animation: float-medium 15s ease-in-out infinite;
}

.animate-float-fast {
  animation: float-fast 10s ease-in-out infinite;
}

.animate-particle-drift {
  animation: particle-drift 25s ease-in-out infinite;
}

.animate-connection-sweep {
  animation: connection-sweep 9s linear infinite;
}

.animate-grid-drift {
  animation: grid-drift 40s linear infinite;
}

.animate-grid-move {
  animation: grid-move 30s linear infinite;
}

.animate-grid-move-slow {
  animation: grid-move-slow 50s linear infinite;
}

.animate-pulse-ring {
  animation: pulse-ring 3s ease-out infinite;
}

.animate-pulse-wave {
  animation: pulse-wave 4s ease-out infinite;
}

.animate-text-float {
  animation: text-float 15s ease-in-out infinite;
}

.animate-badge-float {
  animation: badge-float 4s ease-in-out infinite;
}

.animate-text-slide-up {
  animation: text-slide-up 1s ease-out forwards;
}

.animate-card-enter {
  animation: card-enter 0.6s ease-out forwards;
}

.animate-count-up {
  animation: count-up 0.8s ease-out forwards;
}

.animate-bounce-float {
  animation: bounce-float 3s ease-in-out infinite;
}

/* Grid Patterns */
.bg-grid-white {
  background-image: 
    linear-gradient(rgba(255, 255, 255, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 40px 40px;
}

.bg-grid-gray {
  background-image: 
    linear-gradient(rgba(156, 163, 175, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(156, 163, 175, 0.1) 1px, transparent 1px);
  background-size: 40px 40px;
}

.bg-grid-blue {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.05) 1px, transparent 1px);
  background-size: 40px 40px;
}

/* Enhanced backdrop blur */
.backdrop-blur-xl {
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
}

.backdrop-blur-sm {
  backdrop-filter: blur(8px) saturate(180%);
  -webkit-backdrop-filter: blur(8px) saturate(180%);
}

/* Smooth scroll behavior */
html {
  scroll-behavior: smooth;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #3b82f6, #8b5cf6);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #2563eb, #7c3aed);
}

.dark ::-webkit-scrollbar-track {
  background: #374151;
}

.dark ::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #60a5fa, #a78bfa);
}

.dark ::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #3b82f6, #8b5cf6);
}
</style>