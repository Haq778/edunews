<!-- app/pages/disimpan.vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Artikel Disimpan
        </h1>
        <p class="text-xl text-gray-600 dark:text-gray-400">
          Kumpulan artikel yang telah Anda simpan
        </p>
      </div>
      
      <div class="text-center py-20">
        <div class="max-w-md mx-auto">
          <svg class="w-24 h-24 text-gray-300 dark:text-gray-600 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"/>
          </svg>
          <h3 class="text-2xl font-semibold text-gray-600 dark:text-gray-400 mb-3">Belum ada artikel disimpan</h3>
          <p class="text-gray-500 dark:text-gray-500 mb-6">Simpan artikel favorit Anda untuk dibaca nanti</p>
          <NuxtLink 
            to="/"
            class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-8 py-3 rounded-xl font-semibold hover:shadow-lg transform hover:scale-105 transition duration-300 inline-block"
          >
            Jelajahi Berita
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>