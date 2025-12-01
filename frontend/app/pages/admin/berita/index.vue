<template>
  <div class="p-6 max-w-8xl mx-auto space-y-6">

    <!-- Header dengan Glass Effect -->
    <div class="glass-card rounded-2xl p-6">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between">
        <div class="mb-4 lg:mb-0">
          <div class="flex items-center space-x-3">
            <div class="p-2 bg-gradient-to-br from-blue-500/10 to-purple-600/10 rounded-xl">
              <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v3m0-3a2 2 0 012-2h2a2 2 0 012 2m-6 5v6m4-3H9" />
              </svg>
            </div>
            <div>
              <h1 class="text-2xl lg:text-3xl font-bold bg-gradient-to-r from-gray-900 to-blue-900 dark:from-white dark:to-blue-200 bg-clip-text text-transparent">
                Daftar Berita
              </h1>
              <p class="text-gray-600 dark:text-gray-300 mt-1 text-sm">Kelola konten berita dan artikel dengan mudah</p>
            </div>
          </div>
        </div>

        <NuxtLink
          to="/admin/berita/tambah"
          class="group relative overflow-hidden bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700 text-white px-6 py-3 rounded-xl hover:shadow-xl transition-all duration-500 hover:scale-105 transform shadow-lg"
        >
          <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -skew-x-12 transform translate-x-[-100%] group-hover:translate-x-[100%] transition-transform duration-1000"></div>
          <div class="relative flex items-center space-x-2">
            <div class="p-1 bg-white/20 rounded-lg">
              <svg class="w-4 h-4 transform group-hover:rotate-90 transition-transform duration-300" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
            </div>
            <span class="font-semibold text-sm">Tambah Berita</span>
          </div>
        </NuxtLink>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-2 md:grid-cols-4 gap-4" v-if="!loading && filteredBerita.length > 0">
      <div class="glass-card rounded-xl p-4 text-center">
        <div class="text-2xl font-bold text-blue-600 dark:text-blue-400 mb-1">{{ filteredBerita.length }}</div>
        <div class="text-xs text-gray-600 dark:text-gray-300">Total Berita</div>
      </div>
      <div class="glass-card rounded-xl p-4 text-center">
        <div class="text-2xl font-bold text-green-600 dark:text-green-400 mb-1">
          {{ filteredBerita.filter(item => item.cover).length }}
        </div>
        <div class="text-xs text-gray-600 dark:text-gray-300">Dengan Cover</div>
      </div>
      <div class="glass-card rounded-xl p-4 text-center">
        <div class="text-2xl font-bold text-purple-600 dark:text-purple-400 mb-1">
          {{ new Set(filteredBerita.map(item => item.category?.name)).size }}
        </div>
        <div class="text-xs text-gray-600 dark:text-gray-300">Kategori</div>
      </div>
      <div class="glass-card rounded-xl p-4 text-center">
        <div class="text-2xl font-bold text-orange-600 dark:text-orange-400 mb-1">
          {{ filteredBerita.filter(item => !item.category).length }}
        </div>
        <div class="text-xs text-gray-600 dark:text-gray-300">Tanpa Kategori</div>
      </div>
    </div>

    <!-- Loading Animation -->
    <div v-if="loading" class="glass-card rounded-2xl p-8 text-center">
      <div class="flex flex-col items-center justify-center space-y-4">
        <div class="relative">
          <div class="w-16 h-16 border-4 border-blue-200 dark:border-blue-800 rounded-full"></div>
          <div class="absolute top-0 left-0 w-16 h-16 border-4 border-transparent border-t-blue-600 rounded-full animate-spin"></div>
          <div class="absolute top-1 left-1 w-14 h-14 border-4 border-transparent border-b-purple-600 rounded-full animate-spin animation-delay-500"></div>
        </div>
        <div>
          <p class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-1">Memuat Berita</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">Menyiapkan konten untuk Anda...</p>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="beritaList.length === 0" class="glass-card rounded-2xl p-8 text-center">
      <div class="max-w-sm mx-auto">
        <div class="relative mb-6">
          <div class="w-24 h-24 mx-auto bg-gradient-to-br from-blue-100 to-purple-100 dark:from-blue-900/20 dark:to-purple-900/20 rounded-full flex items-center justify-center">
            <svg class="w-12 h-12 text-blue-500 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v3m0-3a2 2 0 012-2h2a2 2 0 012 2m-6 5v6m4-3H9" />
            </svg>
          </div>
        </div>
        <h3 class="text-xl font-bold text-gray-900 dark:text-white mb-3">Belum Ada Berita</h3>
        <p class="text-gray-600 dark:text-gray-300 text-sm mb-6">Mulai perjalanan konten Anda dengan membuat berita pertama</p>
        <NuxtLink 
          to="/admin/berita/tambah" 
          class="inline-flex items-center space-x-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-6 py-3 rounded-xl hover:shadow-lg transition-all duration-300 hover:scale-105 transform font-semibold text-sm"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          <span>Buat Berita Pertama</span>
        </NuxtLink>
      </div>
    </div>

    <!-- Content Grid -->
    <div v-else class="space-y-4">
      <!-- Search and Filters -->
      <div class="glass-card rounded-xl p-4">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between space-y-3 lg:space-y-0">
          <div class="flex-1 max-w-md">
            <div class="relative">
              <input 
                v-model="searchQuery"
                type="text" 
                placeholder="Cari berita berdasarkan judul atau kategori..." 
                class="w-full pl-10 pr-4 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-300 text-sm"
              >
              <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <button 
                v-if="searchQuery" 
                @click="searchQuery = ''"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
          <div class="flex space-x-2">
            <select 
              v-model="categoryFilter"
              class="px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="">Semua Kategori</option>
              <option 
                v-for="category in availableCategories" 
                :key="category" 
                :value="category"
              >
                {{ category }}
              </option>
            </select>
            <select 
              v-model="sortBy"
              class="px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-sm focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            >
              <option value="newest">Terbaru</option>
              <option value="oldest">Terlama</option>
              <option value="title">Judul A-Z</option>
            </select>
          </div>
        </div>
        
        <!-- Search Results Info -->
        <div v-if="searchQuery || categoryFilter" class="mt-3 flex items-center justify-between text-xs">
          <div class="text-gray-600 dark:text-gray-400">
            Menampilkan {{ filteredBerita.length }} dari {{ beritaList.length }} berita
            <span v-if="searchQuery"> untuk "<span class="font-semibold">{{ searchQuery }}</span>"</span>
            <span v-if="categoryFilter"> dalam kategori "<span class="font-semibold">{{ categoryFilter }}</span>"</span>
          </div>
          <button 
            v-if="searchQuery || categoryFilter"
            @click="clearFilters"
            class="text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 transition-colors flex items-center space-x-1"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Reset Filter</span>
          </button>
        </div>
      </div>

      <!-- No Search Results -->
      <div v-if="!loading && beritaList.length > 0 && filteredBerita.length === 0" class="glass-card rounded-xl p-8 text-center">
        <div class="max-w-sm mx-auto">
          <div class="w-16 h-16 mx-auto bg-gradient-to-br from-gray-100 to-gray-200 dark:from-gray-800 dark:to-gray-700 rounded-full flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">Tidak Ditemukan</h3>
          <p class="text-gray-600 dark:text-gray-300 text-sm mb-4">
            Tidak ada berita yang cocok dengan pencarian "<span class="font-semibold">{{ searchQuery }}</span>"
            <span v-if="categoryFilter"> dalam kategori "<span class="font-semibold">{{ categoryFilter }}</span>"</span>
          </p>
          <button 
            @click="clearFilters"
            class="inline-flex items-center space-x-2 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-4 py-2 rounded-lg hover:shadow-lg transition-all duration-300 transform font-semibold text-sm"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Tampilkan Semua Berita</span>
          </button>
        </div>
      </div>

      <!-- News Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4">
        <div 
          v-for="item in filteredBerita" 
          :key="item.id"
          class="glass-card rounded-xl overflow-hidden hover:shadow-lg transition-all duration-300 group hover:scale-102 transform border border-gray-200/50 dark:border-gray-700/50"
        >
          <!-- Cover Image -->
          <div class="relative h-32 bg-gradient-to-br from-blue-500/20 to-purple-600/20 overflow-hidden">
            <img
              v-if="item.cover"
              :src="coverUrl(item.cover)"
              class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
              alt="Cover Berita"
              @error="hideBrokenImage"
            />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-10 h-10 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </div>
            <div class="absolute top-2 left-2">
              <span class="bg-black/70 text-white px-2 py-1 rounded text-xs font-medium backdrop-blur-sm">
                #{{ item.id }}
              </span>
            </div>
            <div class="absolute top-2 right-2">
              <span 
                v-if="item.category?.name"
                class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-2 py-1 rounded text-xs font-medium backdrop-blur-sm max-w-20 truncate"
                :title="item.category.name"
              >
                {{ item.category.name }}
              </span>
            </div>
          </div>

          <!-- Content -->
          <div class="p-3">
            <h3 class="font-semibold text-sm text-gray-900 dark:text-white mb-2 line-clamp-2 group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors duration-300 leading-tight">
              {{ item.title }}
            </h3>
            
            <!-- Tombol Action yang lebih kecil dan rapi -->
            <div class="flex items-center justify-between pt-2 border-t border-gray-200 dark:border-gray-700">
              <div class="flex space-x-1.5">
                <!-- Edit Button yang lebih kecil -->
                <NuxtLink 
                  :to="`/admin/berita/edit/${item.id}`"
                  class="inline-flex items-center gap-1 px-2 py-1 bg-gradient-to-r from-blue-500/10 to-blue-600/10 hover:from-blue-500/20 hover:to-blue-600/20 
                         text-blue-600 dark:text-blue-400 text-[10px] font-medium rounded-md transition-all duration-200 
                         border border-blue-200 dark:border-blue-800 hover:shadow-sm group/edit"
                  title="Edit Berita"
                >
                  <svg class="w-3 h-3 group-hover/edit:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                  </svg>
                  <span class="font-semibold">Edit</span>
                </NuxtLink>
                
                <!-- Delete Button yang lebih kecil -->
                <button
                  @click="hapus(item.id)"
                  class="inline-flex items-center gap-1 px-2 py-1 bg-gradient-to-r from-red-500/10 to-red-600/10 hover:from-red-500/20 hover:to-red-600/20 
                         text-red-600 dark:text-red-400 text-[10px] font-medium rounded-md transition-all duration-200 
                         border border-red-200 dark:border-red-800 hover:shadow-sm group/delete"
                  title="Hapus Berita"
                >
                  <svg class="w-3 h-3 group-hover/delete:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  <span class="font-semibold">Hapus</span>
                </button>
              </div>
              
              <!-- Date -->
              <div class="text-[10px] text-gray-500 dark:text-gray-400 bg-gray-100 dark:bg-gray-800 px-1.5 py-0.5 rounded">
                {{ formatDate(item.created_at) }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'

definePageMeta({ layout: 'admin' })

const beritaList = ref<any[]>([])
const loading = ref(true)
const searchQuery = ref('')
const categoryFilter = ref('')
const sortBy = ref('newest')

// Computed properties
const availableCategories = computed(() => {
  const categories = beritaList.value
    .map(item => item.category?.name)
    .filter(Boolean)
    .filter((value, index, self) => self.indexOf(value) === index)
    .sort()
  return categories
})

const filteredBerita = computed(() => {
  let filtered = beritaList.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(item => 
      item.title?.toLowerCase().includes(query) ||
      item.category?.name?.toLowerCase().includes(query)
    )
  }

  // Filter by category
  if (categoryFilter.value) {
    filtered = filtered.filter(item => 
      item.category?.name === categoryFilter.value
    )
  }

  // Sort results
  switch (sortBy.value) {
    case 'newest':
      filtered = [...filtered].sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
      break
    case 'oldest':
      filtered = [...filtered].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime())
      break
    case 'title':
      filtered = [...filtered].sort((a, b) => a.title?.localeCompare(b.title))
      break
  }

  return filtered
})

/* FIX: always return a string */
const coverUrl = (cover: string): string =>
  `http://localhost:8080/uploads/berita/${cover}`

const hideBrokenImage = (e: Event) => {
  const img = e.target as HTMLImageElement
  img.style.display = 'none'
}

const formatDate = (dateString: string) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  })
}

const clearFilters = () => {
  searchQuery.value = ''
  categoryFilter.value = ''
  sortBy.value = 'newest'
}

async function loadBerita() {
  loading.value = true
  try {
    const res = await fetch("http://localhost:8080/api/berita")
    beritaList.value = await res.json()
  } catch (err) {
    console.error("Gagal memuat berita:", err)
  } finally {
    loading.value = false
  }
}

async function hapus(id: number) {
  if (!confirm("Yakin ingin menghapus berita ini?")) return

  try {
    const res = await fetch(`http://localhost:8080/api/berita/${id}`, {
      method: "DELETE"
    })

    if (!res.ok) throw new Error(await res.text())

    beritaList.value = beritaList.value.filter(b => b.id !== id)
  } catch (err) {
    console.error("Gagal menghapus berita:", err)
    alert("Gagal menghapus berita, cek console")
  }
}

onMounted(loadBerita)
</script>

<style scoped>
.glass-card {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dark .glass-card {
  background: rgba(17, 24, 39, 0.7);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.animation-delay-500 {
  animation-delay: 0.5s;
}

.hover-scale-102 {
  transform: scale(1);
  transition: transform 0.3s ease;
}

.hover-scale-102:hover {
  transform: scale(1.02);
}

/* Smooth scrolling */
.overflow-x-auto {
  scrollbar-width: thin;
  scrollbar-color: rgba(156, 163, 175, 0.5) transparent;
}

.overflow-x-auto::-webkit-scrollbar {
  height: 6px;
}

.overflow-x-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-x-auto::-webkit-scrollbar-thumb {
  background-color: rgba(156, 163, 175, 0.5);
  border-radius: 3px;
}

.overflow-x-auto::-webkit-scrollbar-thumb:hover {
  background-color: rgba(156, 163, 175, 0.7);
}
</style>