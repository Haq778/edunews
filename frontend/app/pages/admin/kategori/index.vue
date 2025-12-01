<template>
  <div class="p-6 max-w-8xl mx-auto space-y-6">

    <!-- Header dengan Konsep Baru -->
    <div class="glass-card rounded-2xl p-6 border-l-4 border-l-blue-500">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between">
        <div class="mb-4 lg:mb-0">
          <div class="flex items-center space-x-3 mb-3">
            <div class="p-2 bg-gradient-to-br from-green-500/10 to-emerald-600/10 rounded-xl">
              <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
              </svg>
            </div>
            <div>
              <h1 class="text-2xl lg:text-3xl font-bold bg-gradient-to-r from-gray-900 to-green-900 dark:from-white dark:to-green-200 bg-clip-text text-transparent">
                Kategori Berita
              </h1>
              <p class="text-gray-600 dark:text-gray-300 mt-1 text-sm">Organisir berita Anda dengan kategori yang terstruktur</p>
            </div>
          </div>
          
          <!-- Quick Stats -->
          <div class="flex items-center space-x-4 text-sm text-gray-500 dark:text-gray-400">
            <div class="flex items-center space-x-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <span>{{ kategori.length }} Kategori</span>
            </div>
            <div class="flex items-center space-x-1">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
              </svg>
              <span>{{ totalUsage }} Penggunaan</span>
            </div>
          </div>
        </div>

        <NuxtLink
          to="/admin/kategori/tambah"
          class="group relative overflow-hidden bg-gradient-to-br from-green-600 via-emerald-600 to-teal-700 text-white px-6 py-3 rounded-xl hover:shadow-xl transition-all duration-500 hover:scale-105 transform shadow-lg"
        >
          <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent -skew-x-12 transform translate-x-[-100%] group-hover:translate-x-[100%] transition-transform duration-1000"></div>
          <div class="relative flex items-center space-x-2">
            <div class="p-1 bg-white/20 rounded-lg">
              <svg class="w-4 h-4 transform group-hover:rotate-90 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
              </svg>
            </div>
            <span class="font-semibold text-sm">Tambah Kategori</span>
          </div>
        </NuxtLink>
      </div>
    </div>

    <!-- Loading Animation -->
    <div v-if="loading" class="glass-card rounded-2xl p-8 text-center">
      <div class="flex flex-col items-center justify-center space-y-4">
        <div class="relative">
          <div class="w-16 h-16 border-4 border-green-200 dark:border-green-800 rounded-full"></div>
          <div class="absolute top-0 left-0 w-16 h-16 border-4 border-transparent border-t-green-600 rounded-full animate-spin"></div>
          <div class="absolute top-1 left-1 w-14 h-14 border-4 border-transparent border-b-emerald-600 rounded-full animate-spin animation-delay-500"></div>
        </div>
        <div>
          <p class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-1">Memuat Kategori</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">Menyiapkan kategori untuk Anda...</p>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="kategori.length === 0" class="glass-card rounded-2xl p-8 text-center">
      <div class="max-w-sm mx-auto">
        <div class="relative mb-6">
          <div class="w-24 h-24 mx-auto bg-gradient-to-br from-green-100 to-emerald-100 dark:from-green-900/20 dark:to-emerald-900/20 rounded-full flex items-center justify-center">
            <svg class="w-12 h-12 text-green-500 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"></path>
            </svg>
          </div>
        </div>
        <h3 class="text-xl font-bold text-gray-900 dark:text-white mb-3">Belum Ada Kategori</h3>
        <p class="text-gray-600 dark:text-gray-300 text-sm mb-6">Mulai organisir berita Anda dengan membuat kategori pertama</p>
        <NuxtLink 
          to="/admin/kategori/tambah" 
          class="inline-flex items-center space-x-2 bg-gradient-to-r from-green-600 to-emerald-600 text-white px-6 py-3 rounded-xl hover:shadow-lg transition-all duration-300 hover:scale-105 transform font-semibold text-sm"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
          </svg>
          <span>Buat Kategori Pertama</span>
        </NuxtLink>
      </div>
    </div>

    <!-- Content Grid -->
    <div v-else class="space-y-4">
      <!-- Search and Filter -->
      <div class="glass-card rounded-xl p-4">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between space-y-3 lg:space-y-0">
          <div class="flex-1 max-w-md">
            <div class="relative">
              <input 
                v-model="searchQuery"
                type="text" 
                placeholder="Cari kategori..." 
                class="w-full pl-10 pr-4 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-transparent transition-all duration-300 text-sm"
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
              v-model="sortBy"
              class="px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-sm focus:ring-2 focus:ring-green-500 focus:border-transparent"
            >
              <option value="name">Nama A-Z</option>
              <option value="name-desc">Nama Z-A</option>
              <option value="usage">Penggunaan Tertinggi</option>
              <option value="usage-low">Penggunaan Terendah</option>
              <option value="newest">Terbaru</option>
              <option value="oldest">Terlama</option>
            </select>
          </div>
        </div>
        
        <!-- Search Results Info -->
        <div v-if="searchQuery" class="mt-3 flex items-center justify-between text-xs">
          <div class="text-gray-600 dark:text-gray-400">
            Menampilkan {{ filteredKategori.length }} dari {{ kategori.length }} kategori
            <span v-if="searchQuery"> untuk "<span class="font-semibold">{{ searchQuery }}</span>"</span>
          </div>
          <button 
            v-if="searchQuery"
            @click="searchQuery = ''"
            class="text-green-600 dark:text-green-400 hover:text-green-700 dark:hover:text-green-300 transition-colors flex items-center space-x-1"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Reset Pencarian</span>
          </button>
        </div>
      </div>

      <!-- No Search Results -->
      <div v-if="!loading && kategori.length > 0 && filteredKategori.length === 0" class="glass-card rounded-xl p-8 text-center">
        <div class="max-w-sm mx-auto">
          <div class="w-16 h-16 mx-auto bg-gradient-to-br from-gray-100 to-gray-200 dark:from-gray-800 dark:to-gray-700 rounded-full flex items-center justify-center mb-4">
            <svg class="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">Kategori Tidak Ditemukan</h3>
          <p class="text-gray-600 dark:text-gray-300 text-sm mb-4">
            Tidak ada kategori yang cocok dengan "<span class="font-semibold">{{ searchQuery }}</span>"
          </p>
          <button 
            @click="searchQuery = ''"
            class="inline-flex items-center space-x-2 bg-gradient-to-r from-green-600 to-emerald-600 text-white px-4 py-2 rounded-lg hover:shadow-lg transition-all duration-300 transform font-semibold text-sm"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Tampilkan Semua Kategori</span>
          </button>
        </div>
      </div>

      <!-- Categories Grid - Tag Cloud Style -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
        <div 
          v-for="kat in filteredKategori" 
          :key="kat.id"
          class="glass-card rounded-xl p-4 hover:shadow-lg transition-all duration-300 group hover:scale-102 transform border border-gray-200/50 dark:border-gray-700/50"
          :class="getCategoryColor(kat.id)"
        >
          <div class="flex items-start justify-between mb-3">
            <div class="flex-1 min-w-0">
              <h3 class="font-semibold text-gray-900 dark:text-white text-sm truncate group-hover:text-green-600 dark:group-hover:text-green-400 transition-colors duration-300">
                {{ kat.name }}
              </h3>
              <p class="text-xs text-gray-500 dark:text-gray-400 mt-1">ID: {{ kat.id }}</p>
            </div>
            
            <!-- Tombol Action yang lebih kecil dan jelas -->
            <div class="flex space-x-1 ml-2">
              <!-- Edit Button -->
              <NuxtLink
                :to="`/admin/kategori/edit/${kat.id}`"
                class="inline-flex items-center gap-1 px-2 py-1 bg-gradient-to-r from-yellow-500/10 to-yellow-600/10 hover:from-yellow-500/20 hover:to-yellow-600/20 
                       text-yellow-600 dark:text-yellow-400 text-[10px] font-medium rounded-md transition-all duration-200 
                       border border-yellow-200 dark:border-yellow-800 hover:shadow-sm group/edit"
                title="Edit Kategori"
              >
                <svg class="w-3 h-3 group-hover/edit:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                </svg>
                <span class="font-semibold">Edit</span>
              </NuxtLink>
              
              <!-- Delete Button - TIDAK DISABLED -->
              <button
                @click="hapus(kat.id, kat.name, kat.usage_count)"
                class="inline-flex items-center gap-1 px-2 py-1 bg-gradient-to-r from-red-500/10 to-red-600/10 hover:from-red-500/20 hover:to-red-600/20 
                       text-red-600 dark:text-red-400 text-[10px] font-medium rounded-md transition-all duration-200 
                       border border-red-200 dark:border-red-800 hover:shadow-sm group/delete"
                title="Hapus Kategori"
              >
                <svg class="w-3 h-3 group-hover/delete:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
                <span class="font-semibold">Hapus</span>
              </button>
            </div>
          </div>

          <!-- Usage Stats -->
          <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-400">
            <div class="flex items-center space-x-1" :class="getUsageColor(kat.usage_count)">
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
              <span>{{ kat.usage_count || 0 }} berita</span>
            </div>
            <div class="text-[10px] px-1.5 py-0.5 rounded bg-gray-100 dark:bg-gray-800">
              {{ formatDate(kat.created_at) }}
            </div>
          </div>

          <!-- Color Indicator -->
          <div class="mt-3 h-1 rounded-full bg-gradient-to-r from-green-400 to-emerald-500 opacity-80 group-hover:opacity-100 transition-opacity duration-300"></div>
        </div>
      </div>

      <!-- Quick Actions Footer -->
      <div v-if="kategori.length > 0" class="glass-card rounded-xl p-4 mt-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between space-y-3 sm:space-y-0">
          <div class="text-sm text-gray-600 dark:text-gray-400">
            Total <span class="font-semibold text-green-600 dark:text-green-400">{{ kategori.length }}</span> kategori terdaftar • 
            <span class="font-semibold text-blue-600 dark:text-blue-400">{{ totalUsage }}</span> total penggunaan
          </div>
          <div class="flex space-x-2">
            <button 
              @click="exportCategories"
              class="inline-flex items-center gap-1 px-2 py-1.5 text-[10px] bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 rounded-md hover:bg-blue-100 dark:hover:bg-blue-900/30 transition-all duration-300 border border-blue-200 dark:border-blue-800"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
              </svg>
              <span>Export</span>
            </button>
            <NuxtLink
              to="/admin/kategori/tambah"
              class="inline-flex items-center gap-1 px-2 py-1.5 text-[10px] bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400 rounded-md hover:bg-green-100 dark:hover:bg-green-900/30 transition-all duration-300 border border-green-200 dark:border-green-800"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
              </svg>
              <span>Tambah</span>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup>
definePageMeta({ layout: "admin" });

import { ref, onMounted, computed } from "vue";

const kategori = ref([]);
const loading = ref(true);
const searchQuery = ref('');
const sortBy = ref('name');

// Computed properties
const filteredKategori = computed(() => {
  let filtered = kategori.value;

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(item => 
      item.name?.toLowerCase().includes(query)
    );
  }

  // Sort results
  switch (sortBy.value) {
    case 'name':
      filtered = [...filtered].sort((a, b) => a.name?.localeCompare(b.name));
      break;
    case 'name-desc':
      filtered = [...filtered].sort((a, b) => b.name?.localeCompare(a.name));
      break;
    case 'usage':
      filtered = [...filtered].sort((a, b) => (b.usage_count || 0) - (a.usage_count || 0));
      break;
    case 'usage-low':
      filtered = [...filtered].sort((a, b) => (a.usage_count || 0) - (b.usage_count || 0));
      break;
    case 'newest':
      filtered = [...filtered].sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime());
      break;
    case 'oldest':
      filtered = [...filtered].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime());
      break;
  }

  return filtered;
});

const totalUsage = computed(() => {
  return kategori.value.reduce((total, kat) => total + (kat.usage_count || 0), 0);
});

// Helper functions
const getCategoryColor = (id) => {
  const colors = [
    'border-l-green-400',
    'border-l-blue-400', 
    'border-l-purple-400',
    'border-l-orange-400',
    'border-l-pink-400',
    'border-l-indigo-400',
    'border-l-teal-400',
    'border-l-cyan-400'
  ];
  return colors[id % colors.length];
};

const getUsageColor = (count) => {
  if (count === 0) return 'text-gray-500';
  if (count <= 5) return 'text-green-600 dark:text-green-400';
  if (count <= 20) return 'text-blue-600 dark:text-blue-400';
  return 'text-purple-600 dark:text-purple-400';
};

const formatDate = (dateString) => {
  if (!dateString) return '-';
  return new Date(dateString).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'short'
  });
};

const exportCategories = () => {
  const data = JSON.stringify(kategori.value, null, 2);
  const blob = new Blob([data], { type: 'application/json' });
  const url = URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  a.download = 'kategori-berita.json';
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
};

// 🟢 LOAD KATEGORI DENGAN USAGE COUNT
async function loadKategori() {
  loading.value = true;
  try {
    const res = await $fetch("http://localhost:8080/api/categories");
    
    // Load usage count untuk setiap kategori
    const categoriesWithUsage = await Promise.all(
      res.map(async (category) => {
        try {
          // 🟢 AMBIL DATA PENGGUNAAN KATEGORI DARI BERITA
          const usageRes = await $fetch(`http://localhost:8080/api/berita/count-by-category/${category.id}`);
          return {
            ...category,
            usage_count: usageRes.count || 0
          };
        } catch (error) {
          console.error(`Gagal load usage untuk kategori ${category.id}:`, error);
          return {
            ...category,
            usage_count: 0
          };
        }
      })
    );
    
    kategori.value = categoriesWithUsage;
  } catch (err) {
    console.error("Gagal load kategori:", err);
  } finally {
    loading.value = false;
  }
}

onMounted(loadKategori);

// 🟢 DELETE DENGAN PESAN KONFIRMASI YANG INFORMATIF
const hapus = async (id, name, usageCount) => {
  if (!confirm(`Yakin ingin menghapus kategori "${name}"? ${usageCount > 0 ? `\n\n⚠️ PERHATIAN: Kategori ini masih digunakan oleh ${usageCount} berita.` : ''}`)) return;

  try {
    await $fetch(`http://localhost:8080/api/categories/${id}`, {
      method: "DELETE",
    });

    kategori.value = kategori.value.filter((k) => k.id !== id);
  } catch (err) {
    alert("Gagal menghapus kategori");
    console.error("Error deleting category:", err);
  }
};
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

/* Smooth scrollbar styling */
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