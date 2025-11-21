<!-- app/pages/login.vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-md mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Masuk ke EduNews
        </h1>
        <p class="text-xl text-gray-600 dark:text-gray-400">
          Akses konten eksklusif dan simpan artikel favorit
        </p>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg p-8">
        <div class="text-center py-12">
          <svg class="w-24 h-24 text-gray-300 dark:text-gray-600 mx-auto mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
          </svg>
          <h3 class="text-2xl font-semibold text-gray-600 dark:text-gray-400 mb-3">Fitur Login Segera Hadir</h3>
          <p class="text-gray-500 dark:text-gray-500 mb-6">Sistem autentikasi sedang dalam pengembangan</p>
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