<template>
  <nav class="bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700 shadow-sm transition-all duration-300 sticky top-0 z-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">

        <!-- Logo -->
        <div class="flex items-center space-x-8">
          <NuxtLink to="/" class="flex items-center space-x-3 group">
            <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0 0h4m-4 0h4"/>
              </svg>
            </div>
            <span class="text-xl font-bold text-gray-900 dark:text-white">EduNews</span>
          </NuxtLink>

          <!-- Navigation -->
          <div class="hidden md:flex space-x-6">
            <NuxtLink 
              to="/" 
              class="text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 font-medium transition-colors duration-300 py-2 px-3 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800"
              :class="{ 'text-blue-600 dark:text-blue-400 bg-gray-100 dark:bg-gray-800': $route.path === '/' }"
            >
              Beranda
            </NuxtLink>

            <NuxtLink 
              to="/populer" 
              class="text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 font-medium transition-colors duration-300 py-2 px-3 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800"
              :class="{ 'text-blue-600 dark:text-blue-400 bg-gray-100 dark:bg-gray-800': $route.path === '/populer' }"
            >
              Populer
            </NuxtLink>

            <NuxtLink 
              to="/tentang" 
              class="text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 font-medium transition-colors duration-300 py-2 px-3 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-800"
              :class="{ 'text-blue-600 dark:text-blue-400 bg-gray-100 dark:bg-gray-800': $route.path === '/tentang' }"
            >
              Tentang
            </NuxtLink>
          </div>
        </div>

        <!-- Theme Toggle -->
        <button 
          @click="toggleTheme"
          class="p-2 rounded-xl bg-gray-100 dark:bg-gray-800 text-gray-600 dark:text-gray-400 hover:bg-gray-200 dark:hover:bg-gray-700 transition-all duration-300 transform hover:scale-105"
          :title="theme === 'light' ? 'Switch to dark mode' : 'Switch to light mode'"
        >
          <!-- Icon Light -->
          <svg v-if="theme === 'light'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"/>
          </svg>

          <!-- Icon Dark -->
          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"/>
          </svg>
        </button>

      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useTheme } from '@/composables/useTheme'

const { theme, toggleTheme } = useTheme()
</script>

<style scoped>
/* opsional: smooth transition untuk background dan text */
nav {
  transition: background-color 0.3s, color 0.3s;
}
</style>
