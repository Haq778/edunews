<!-- app/pages/profil.vue -->
<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          Profil Saya
        </h1>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg p-8">
        <div class="text-center">
          <div class="w-24 h-24 bg-gradient-to-br from-blue-500 to-purple-600 rounded-full flex items-center justify-center text-white text-2xl font-bold mx-auto mb-4">
            U
          </div>
          <h2 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">User Guest</h2>
          <p class="text-gray-600 dark:text-gray-400 mb-6">guest@edunews.id</p>
          
          <div class="max-w-md mx-auto text-left space-y-4">
            <div class="flex justify-between py-3 border-b border-gray-200 dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-400">Status</span>
              <span class="text-gray-900 dark:text-white font-medium">Pengunjung</span>
            </div>
            <div class="flex justify-between py-3 border-b border-gray-200 dark:border-gray-700">
              <span class="text-gray-600 dark:text-gray-400">Bergabung</span>
              <span class="text-gray-900 dark:text-white font-medium">Hari ini</span>
            </div>
            <div class="flex justify-between py-3">
              <span class="text-gray-600 dark:text-gray-400">Artikel Disimpan</span>
              <span class="text-gray-900 dark:text-white font-medium">3</span>
            </div>
          </div>
        </div>
      </div>

      <div class="text-center mt-8">
        <NuxtLink 
          to="/"
          class="bg-gradient-to-r from-blue-500 to-purple-600 text-white px-8 py-3 rounded-xl font-semibold hover:shadow-lg transform hover:scale-105 transition duration-300 inline-block"
        >
          Kembali ke Beranda
        </NuxtLink>
      </div>
    </div>
  </div>
</template>