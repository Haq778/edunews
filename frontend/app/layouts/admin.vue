<template>
  <div class="flex min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-300">
    <!-- SIDEBAR -->
    <aside
      :class="[
        'bg-white dark:bg-gray-800 shadow-lg px-4 py-6 flex flex-col transition-all duration-300',
        isSidebarOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
        'fixed md:static z-50 h-full w-64'
      ]"
    >
      <!-- Logo -->
      <div class="flex items-center mb-8 px-2">
        <div class="w-8 h-8 bg-indigo-600 rounded-lg flex items-center justify-center mr-3">
          <span class="text-white font-bold text-sm">EN</span>
        </div>
        <h2 class="text-xl font-bold dark:text-white">Edunews Admin</h2>
      </div>

      <!-- Navigation -->
      <nav class="space-y-1 flex-1">
        <NuxtLink
          to="/admin/berita"
          class="flex items-center px-3 py-3 rounded-lg transition-colors duration-200 group"
          :class="{
            'bg-indigo-50 text-indigo-700 dark:bg-indigo-900/30 dark:text-indigo-300': $route.path.includes('/admin/berita'),
            'text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700': !$route.path.includes('/admin/berita')
          }"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0-12a2 2 0 012-2h2a2 2 0 012 2M9 6a2 2 0 012-2h2a2 2 0 012 2"></path>
          </svg>
          Manajemen Berita
        </NuxtLink>

        <NuxtLink
          to="/admin/kategori"
          class="flex items-center px-3 py-3 rounded-lg transition-colors duration-200 group"
          :class="{
            'bg-indigo-50 text-indigo-700 dark:bg-indigo-900/30 dark:text-indigo-300': $route.path.includes('/admin/kategori'),
            'text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700': !$route.path.includes('/admin/kategori')
          }"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
          </svg>
          Manajemen Kategori
        </NuxtLink>

        <NuxtLink
          to="/admin/users"
          class="flex items-center px-3 py-3 rounded-lg transition-colors duration-200 group"
          :class="{
            'bg-indigo-50 text-indigo-700 dark:bg-indigo-900/30 dark:text-indigo-300': $route.path.includes('/admin/users'),
            'text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700': !$route.path.includes('/admin/users')
          }"
        >
          <svg class="w-5 h-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path>
          </svg>
          Manajemen User
        </NuxtLink>
      </nav>

      <!-- Logout -->
      <div class="pt-4 mt-4 border-t border-gray-200 dark:border-gray-700">
        <div class="flex items-center justify-between px-3 py-2 text-sm text-gray-600 dark:text-gray-400">
          <span>Admin Panel v1.0</span>
          <button
            @click="handleLogout"
            class="px-2 py-1 text-white bg-red-500 rounded hover:bg-red-400 transition"
          >
            Logout
          </button>
        </div>
      </div>
    </aside>

    <!-- MAIN CONTENT -->
    <main class="flex-1 md:ml-64 p-6 transition-colors duration-300">
      <!-- Topbar -->
      <div class="flex justify-between items-center mb-8">
        <div class="flex items-center space-x-3">
          <!-- Sidebar toggle for mobile -->
          <button
            @click="isSidebarOpen = !isSidebarOpen"
            class="md:hidden p-2 bg-gray-200 dark:bg-gray-700 rounded-md hover:bg-gray-300 dark:hover:bg-gray-600 transition"
          >
            <svg class="w-5 h-5 text-gray-700 dark:text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>

          <div>
            <h1 class="text-2xl font-bold dark:text-white mb-1">
              <slot name="title"></slot>
            </h1>
            <p class="text-gray-500 dark:text-gray-400 text-sm">Kelola konten dan pengguna dengan mudah</p>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center space-x-4">
          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="flex items-center px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-200 shadow-sm"
          >
            <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"></path>
            </svg>
            {{ colorMode === 'light' ? 'Dark' : 'Light' }} Mode
          </button>
        </div>
      </div>

      <!-- Content Slot -->
      <div class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-sm border border-gray-100 dark:border-gray-700 transition-colors duration-300">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const router = useRouter()
const auth = useAuth()
const colorMode = ref<'light' | 'dark'>('light')
const isSidebarOpen = ref(false)

// Toggle theme
function toggleTheme() {
  colorMode.value = colorMode.value === 'dark' ? 'light' : 'dark'
  document.documentElement.classList.toggle('dark', colorMode.value === 'dark')
  if (process.client) localStorage.setItem('colorMode', colorMode.value)
}

// Move theme toggle button to top-right of the viewport
onMounted(() => {
  // ensure DOM is rendered
  requestAnimationFrame(() => {
    const buttons = Array.from(document.querySelectorAll('button'))
    const themeButton = buttons.find(b => {
      const txt = (b.textContent || '').trim()
      return txt.endsWith('Mode') && !!b.querySelector('svg')
    })
    if (themeButton) {
      themeButton.style.position = 'fixed'
      themeButton.style.top = '1rem'
      themeButton.style.right = '1rem'
      themeButton.style.zIndex = '60'
      themeButton.style.margin = '0'
    }
  })
})


// Logout
function handleLogout() {
  auth.logout()
  router.push('/')
}

// Set theme on mounted
onMounted(() => {
  const savedMode = localStorage.getItem('colorMode') as 'light' | 'dark' | null
  if (savedMode) {
    colorMode.value = savedMode
    document.documentElement.classList.toggle('dark', colorMode.value === 'dark')
  }
})
</script>
