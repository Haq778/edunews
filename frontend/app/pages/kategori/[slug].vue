<!-- app/pages/kategori/[slug].vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Kategori: {{ categoryName }}
        </h1>
        <p class="text-xl text-gray-600 dark:text-gray-400">
          Berita terkait {{ categoryName.toLowerCase() }}
        </p>
      </div>
      
      <div class="text-center py-20">
        <div class="max-w-md mx-auto">
          <svg class="w-24 h-24 text-gray-300 dark:text-gray-600 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M19 9l-7 7-7-7"/>
          </svg>
          <h3 class="text-2xl font-semibold text-gray-600 dark:text-gray-400 mb-3">Konten Kategori</h3>
          <p class="text-gray-500 dark:text-gray-500 mb-6">Berita untuk kategori ini akan segera tersedia</p>
          <NuxtLink 
            to="/"
            class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-8 py-3 rounded-xl font-semibold hover:shadow-lg transform hover:scale-105 transition duration-300 inline-block"
          >
            Kembali ke Beranda
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const route = useRoute()
const slug = route.params.slug

// Map slug ke nama kategori
const categoryMap = {
  'teknologi-pendidikan': 'Teknologi Pendidikan',
  'kurikulum': 'Kurikulum',
  'guru-pengajar': 'Guru & Pengajar',
  'siswa-mahasiswa': 'Siswa & Mahasiswa',
  'literasi': 'Literasi',
  'penelitian': 'Penelitian'
}

const categoryName = categoryMap[slug] || 'Kategori'
</script>