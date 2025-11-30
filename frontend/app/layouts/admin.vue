<template>
  <div class="min-h-screen flex bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-950 relative transition-all duration-500 overflow-hidden">

    <!-- Background Effects -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <!-- Light Mode Background Effects -->
      <div v-if="colorMode === 'light'" class="absolute inset-0">
        <div class="absolute top-1/4 left-1/4 w-72 h-72 bg-blue-100 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-pulse-slow"></div>
        <div class="absolute top-1/3 right-1/4 w-72 h-72 bg-cyan-100 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-pulse-slow delay-1000"></div>
        <div class="absolute bottom-1/4 left-1/2 w-72 h-72 bg-indigo-100 rounded-full mix-blend-multiply filter blur-3xl opacity-30 animate-pulse-slow delay-2000"></div>
        
        <!-- Light Mode Border Effects -->
        <div class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-blue-200 to-transparent opacity-40"></div>
        <div class="absolute top-0 right-0 w-1 h-full bg-gradient-to-b from-transparent via-cyan-200 to-transparent opacity-40"></div>
        <div class="absolute bottom-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-indigo-200 to-transparent opacity-40"></div>
        <div class="absolute top-0 left-0 w-1 h-full bg-gradient-to-b from-transparent via-purple-200 to-transparent opacity-40"></div>
      </div>
      
      <!-- Dark Mode Neon Effects -->
      <div v-else class="absolute inset-0">
        <!-- Main Neon Borders -->
        <div class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-cyan-500 to-transparent opacity-60 animate-pulse"></div>
        <div class="absolute top-0 right-0 w-1 h-full bg-gradient-to-b from-transparent via-blue-500 to-transparent opacity-60 animate-pulse delay-500"></div>
        <div class="absolute bottom-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-purple-500 to-transparent opacity-60 animate-pulse delay-1000"></div>
        <div class="absolute top-0 left-0 w-1 h-full bg-gradient-to-b from-transparent via-indigo-500 to-transparent opacity-60 animate-pulse delay-1500"></div>
        
        <!-- Corner Neon Effects -->
        <div class="absolute top-0 left-0 w-4 h-4">
          <div class="absolute top-0 left-0 w-4 h-1 bg-cyan-400 opacity-80 animate-pulse"></div>
          <div class="absolute top-0 left-0 w-1 h-4 bg-cyan-400 opacity-80 animate-pulse delay-300"></div>
        </div>
        <div class="absolute top-0 right-0 w-4 h-4">
          <div class="absolute top-0 right-0 w-4 h-1 bg-blue-400 opacity-80 animate-pulse"></div>
          <div class="absolute top-0 right-0 w-1 h-4 bg-blue-400 opacity-80 animate-pulse delay-300"></div>
        </div>
        <div class="absolute bottom-0 left-0 w-4 h-4">
          <div class="absolute bottom-0 left-0 w-4 h-1 bg-purple-400 opacity-80 animate-pulse"></div>
          <div class="absolute bottom-0 left-0 w-1 h-4 bg-purple-400 opacity-80 animate-pulse delay-300"></div>
        </div>
        <div class="absolute bottom-0 right-0 w-4 h-4">
          <div class="absolute bottom-0 right-0 w-4 h-1 bg-indigo-400 opacity-80 animate-pulse"></div>
          <div class="absolute bottom-0 right-0 w-1 h-4 bg-indigo-400 opacity-80 animate-pulse delay-300"></div>
        </div>
        
        <!-- Floating particles -->
        <div v-for="i in 20" :key="i" 
             class="absolute rounded-full opacity-60 animate-float"
             :class="getParticleClass(i)"
             :style="{
               top: Math.random() * 100 + '%',
               left: Math.random() * 100 + '%',
               animationDelay: Math.random() * 5 + 's',
               animationDuration: 4 + Math.random() * 3 + 's'
             }">
        </div>

        <!-- Glowing Orbs -->
        <div class="absolute top-1/4 left-1/4 w-4 h-4 bg-cyan-400 rounded-full blur-md opacity-40 animate-pulse-slow"></div>
        <div class="absolute top-3/4 right-1/4 w-3 h-3 bg-blue-400 rounded-full blur-md opacity-50 animate-pulse-slow delay-2000"></div>
        <div class="absolute bottom-1/3 left-1/3 w-5 h-5 bg-purple-400 rounded-full blur-md opacity-30 animate-pulse-slow delay-4000"></div>
      </div>
    </div>

    <!-- SIDEBAR -->
    <aside
      :class="[
        'fixed top-0 left-0 h-screen w-80 z-50 transition-all duration-700 ease-out backdrop-blur-xl transform',
        isSidebarOpen ? 'translate-x-0 shadow-2xl' : '-translate-x-full shadow-none',
        'md:translate-x-0 md:shadow-2xl',
        colorMode === 'light' 
          ? 'bg-white/90 border-r border-gray-200/60 sidebar-light-glow' 
          : 'bg-gray-900/90 border-r border-cyan-500/30 sidebar-neon-glow'
      ]"
    >
      <!-- SIDEBAR CONTENT -->
      <div class="flex flex-col h-full px-6 py-8 overflow-y-auto relative">

        <!-- Animated Background Pattern -->
        <div class="absolute inset-0 opacity-[0.03] pointer-events-none">
          <div class="absolute inset-0" :class="colorMode === 'light' ? 'bg-grid-black' : 'bg-grid-white'"></div>
        </div>

        <!-- Logo Section -->
        <div class="flex items-center justify-between mb-12 px-2 relative z-10">
          <div class="flex items-center space-x-4 group cursor-pointer">
            <div class="relative">
              <div class="w-12 h-12 bg-gradient-to-br from-blue-500 via-cyan-500 to-indigo-600 rounded-2xl flex items-center justify-center shadow-2xl transform group-hover:scale-110 transition-all duration-500 animate-pulse-glow">
                <span class="text-white font-bold text-sm">EN</span>
              </div>
              <div class="absolute -inset-1 bg-gradient-to-r from-blue-500 to-cyan-400 rounded-2xl blur opacity-30 group-hover:opacity-50 transition-all duration-500 animate-pulse-slow"></div>
            </div>
            <div class="transform group-hover:translate-x-1 transition-transform duration-500">
              <h2 class="text-2xl font-bold bg-gradient-to-r from-blue-600 to-cyan-500 bg-clip-text text-transparent">Edunews</h2>
              <p class="text-gray-500 dark:text-cyan-300 text-xs font-medium mt-1 transform group-hover:scale-105 transition-transform duration-500">Admin Panel</p>
            </div>
          </div>

          <!-- Close Button Mobile -->
          <button
            @click="isSidebarOpen = false"
            class="md:hidden p-2 rounded-xl hover:bg-gray-200/50 dark:hover:bg-gray-700/50 transition-all duration-300 group"
          >
            <div class="w-6 h-6 relative">
              <div class="absolute inset-0 bg-gradient-to-r from-red-400 to-pink-500 rounded-lg opacity-0 group-hover:opacity-100 transition-opacity duration-300 blur-sm"></div>
              <svg class="w-6 h-6 text-gray-600 dark:text-cyan-200 relative z-10 transform group-hover:rotate-90 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </div>
          </button>
        </div>

        <!-- Navigation -->
        <nav class="space-y-2 flex-1 relative z-10">
          <div class="px-3 mb-6 transform hover:translate-x-1 transition-transform duration-500">
            <p class="text-gray-400 dark:text-cyan-400 text-xs font-semibold uppercase tracking-wider">Main Menu</p>
          </div>
          
          <NuxtLink
            v-for="(item, index) in navigation"
            :key="item.path"
            :to="item.path"
            class="flex items-center px-4 py-4 rounded-2xl transition-all duration-500 group relative overflow-hidden border transform hover:scale-[1.02]"
            :class="[
              navClass(item.path),
              colorMode === 'light' 
                ? 'hover:shadow-lg hover:border-gray-300/50' 
                : 'hover:shadow-cyan-500/10 hover:border-cyan-500/30',
              'hover:translate-x-2'
            ]"
            @click="isSidebarOpen = false"
            :style="{
              animationDelay: `${index * 100}ms`
            }"
          >
            <!-- Animated Background -->
            <div class="absolute inset-0 bg-gradient-to-r from-blue-500/0 via-cyan-500/5 to-blue-500/0 opacity-0 group-hover:opacity-100 transition-all duration-500 transform -translate-x-full group-hover:translate-x-0"></div>
            
            <!-- Active Indicator Bar -->
            <div v-if="route.path.includes(item.path)" 
                 class="absolute left-0 top-1/2 transform -translate-y-1/2 w-1 h-8 bg-gradient-to-b from-cyan-400 to-blue-500 rounded-r-full animate-pulse-glow">
            </div>
            
            <!-- Icon Container -->
            <div class="relative transform group-hover:scale-110 transition-transform duration-500">
              <div class="w-10 h-10 rounded-xl bg-gradient-to-br from-blue-500 to-cyan-400 flex items-center justify-center mr-4 shadow-lg z-10 relative overflow-hidden">
                <component :is="item.icon" class="w-5 h-5 text-white relative z-10" />
                <div class="absolute inset-0 bg-gradient-to-br from-white/20 to-transparent"></div>
              </div>
              <div class="absolute inset-0 bg-gradient-to-br from-blue-500 to-cyan-400 rounded-xl blur opacity-50 group-hover:opacity-75 transition-opacity duration-500 group-hover:scale-110"></div>
            </div>
            
            <span class="font-semibold z-10 transform transition-all duration-500 group-hover:translate-x-1"
                  :class="route.path.includes(item.path) ? 'text-gray-800 dark:text-white' : 'text-gray-600 dark:text-cyan-200 group-hover:text-gray-800 dark:group-hover:text-white'">
              {{ item.name }}
            </span>
            
            <!-- Hover Arrow -->
            <div class="ml-auto opacity-0 group-hover:opacity-100 transform translate-x-2 group-hover:translate-x-0 transition-all duration-500">
              <svg class="w-4 h-4 text-cyan-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </div>
          </NuxtLink>
        </nav>

        <!-- Footer -->
        <div class="pt-6 mt-8 border-t border-gray-300/50 dark:border-cyan-500/20 relative z-10">
          <div class="flex flex-col space-y-4 px-2">

            <!-- Status & Version -->
            <div class="flex items-center justify-between transform hover:scale-105 transition-transform duration-500">
              <span class="text-gray-500 dark:text-cyan-300 text-sm font-medium">v1.0.0</span>
              <div class="flex items-center space-x-2 group cursor-pointer">
                <div class="w-2 h-2 bg-green-400 rounded-full animate-pulse-glow group-hover:scale-150 transition-transform duration-300"></div>
                <span class="text-green-500 dark:text-green-400 text-xs font-medium group-hover:text-green-600 dark:group-hover:text-green-300 transition-colors duration-300">Online</span>
              </div>
            </div>

            <!-- Logout Button -->
            <button
              @click="handleLogout"
              class="flex items-center justify-center w-full px-4 py-3.5 text-white bg-gradient-to-r from-red-500 to-pink-600 rounded-2xl hover:from-red-600 hover:to-pink-700 shadow-lg transition-all duration-500 transform hover:scale-[1.02] group relative overflow-hidden"
            >
              <!-- Animated background -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/10 to-transparent transform -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
              
              <!-- Ripple effect -->
              <div class="absolute inset-0 overflow-hidden rounded-2xl">
                <div class="absolute -inset-1 bg-gradient-to-r from-red-400 to-pink-500 opacity-0 group-hover:opacity-20 blur transition-opacity duration-500"></div>
              </div>
              
              <svg class="w-5 h-5 mr-3 group-hover:rotate-12 transition-transform duration-500 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
              </svg>
              <span class="font-medium relative z-10 transform group-hover:translate-x-1 transition-transform duration-500">Logout</span>
            </button>

          </div>
        </div>
      </div>
    </aside>

    <!-- OVERLAY MOBILE -->
    <div
      v-if="isSidebarOpen"
      @click="isSidebarOpen = false"
      class="fixed inset-0 bg-black/40 backdrop-blur-sm z-40 md:hidden transition-all duration-500 animate-fade-in"
    ></div>

    <!-- MAIN CONTENT -->
    <main class="flex-1 min-w-0 transition-all duration-500 md:ml-80 relative z-10">
      <div class="p-6 md:p-8">

        <!-- Topbar -->
        <div class="flex items-center justify-between mb-8">
          <div class="flex items-center space-x-4">
            <button
              @click="isSidebarOpen = !isSidebarOpen"
              class="md:hidden p-3 rounded-2xl border backdrop-blur-xl shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105 group"
              :class="colorMode === 'light' 
                ? 'bg-white/80 border-gray-200/50 hover:border-gray-300/50' 
                : 'bg-gray-800/80 border-cyan-500/20 hover:border-cyan-500/30'"
            >
              <div class="w-5 h-5 relative">
                <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-cyan-400 rounded opacity-0 group-hover:opacity-100 blur transition-opacity duration-300"></div>
                <svg class="w-5 h-5 text-gray-700 dark:text-cyan-200 relative z-10 transform group-hover:rotate-180 transition-transform duration-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-width="2" stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16"/>
                </svg>
              </div>
            </button>

            <div class="transform hover:translate-x-1 transition-transform duration-500">
              <h1 class="text-3xl md:text-4xl font-bold bg-gradient-to-r from-gray-800 to-gray-600 dark:from-white dark:to-cyan-200 bg-clip-text text-transparent transition-all duration-500">
                <slot name="title">Dashboard</slot>
              </h1>
              <p class="text-gray-500 dark:text-cyan-300 text-sm mt-2 transition-colors duration-500">Kelola konten dan pengguna dengan mudah</p>
            </div>
          </div>

          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="p-3.5 rounded-2xl border backdrop-blur-xl shadow-lg hover:shadow-xl transition-all duration-500 group hover:scale-105 transform"
            :class="colorMode === 'light' 
              ? 'bg-white/80 border-gray-200/50 hover:border-gray-300/50' 
              : 'bg-gray-800/80 border-cyan-500/20 hover:border-cyan-500/30 neon-glow'"
            :title="colorMode === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode'"
          >
            <div class="flex items-center space-x-3">
              <div class="w-5 h-5 transform transition-all duration-500 relative" :class="colorMode === 'light' ? 'rotate-0 scale-100' : 'rotate-180 scale-110'">
                <div class="absolute inset-0 bg-gradient-to-br from-amber-400 to-orange-500 rounded opacity-0 group-hover:opacity-20 transition-opacity duration-300" v-if="colorMode === 'light'"></div>
                <div class="absolute inset-0 bg-gradient-to-br from-blue-400 to-cyan-500 rounded opacity-0 group-hover:opacity-20 transition-opacity duration-300" v-else></div>
                <svg v-if="colorMode === 'light'" class="w-5 h-5 text-amber-500 relative z-10" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
                </svg>
                <svg v-else class="w-5 h-5 text-blue-400 relative z-10" fill="currentColor" viewBox="0 0 24 24">
                  <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-700 dark:text-cyan-200 transition-colors duration-500 group-hover:scale-105 transform">
                {{ colorMode === 'light' ? 'Dark' : 'Light' }}
              </span>
            </div>
          </button>
        </div>

        <!-- Content Box -->
        <div class="backdrop-blur-xl p-6 md:p-8 rounded-3xl shadow-xl border transition-all duration-500 hover:shadow-2xl transform hover:scale-[1.005]"
             :class="colorMode === 'light' 
               ? 'bg-white/80 border-gray-200/50 hover:border-gray-300/50 content-light-glow' 
               : 'bg-gray-800/80 border-cyan-500/20 hover:border-cyan-500/30 content-neon-glow'">
          <slot />
        </div>

      </div>
    </main>

  </div>
</template>

<script setup>
import { ref, watch, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const isSidebarOpen = ref(false)
const colorMode = ref('light')
const route = useRoute()
const router = useRouter()
const auth = useAuth()

// Navigation items with icons
const navigation = [
  {
    name: 'Manajemen Berita',
    path: '/admin/berita',
    icon: 'NewsIcon'
  },
  {
    name: 'Manajemen Kategori',
    path: '/admin/kategori',
    icon: 'CategoryIcon'
  },
  {
    name: 'Manajemen User',
    path: '/admin/users',
    icon: 'UsersIcon'
  }
]

// Icon components
const NewsIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v3m0-3a2 2 0 012-2h2a2 2 0 012 2m-6 5v6m4-3H9"></path>
  </svg>`
}

const CategoryIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-width="2" stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 10h16M4 14h16M4 18h16"></path>
  </svg>`
}

const UsersIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
    <path stroke-width="2" stroke-linecap="round" stroke-linejoin="round" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
  </svg>`
}

const isAdmin = computed(() => {
  return auth.user.value?.role === 'admin'
})

function navClass(path) {
  const active = route.path.includes(path)
  if (active) {
    return colorMode.value === 'light'
      ? 'bg-white/20 text-gray-800 shadow-lg border-gray-300/50 backdrop-blur-sm'
      : 'bg-cyan-500/10 text-white shadow-cyan-500/20 border-cyan-500/30 backdrop-blur-sm'
  }
  return colorMode.value === 'light'
    ? 'text-gray-600 hover:text-gray-800 border-transparent'
    : 'text-cyan-200 hover:text-white border-transparent'
}

function getParticleClass(index) {
  const sizes = ['w-1 h-1', 'w-2 h-2', 'w-1.5 h-1.5']
  const colors = ['bg-cyan-400', 'bg-blue-400', 'bg-purple-400', 'bg-indigo-400']
  const size = sizes[index % sizes.length]
  const color = colors[index % colors.length]
  return `${size} ${color}`
}

function toggleTheme() {
  colorMode.value = colorMode.value === 'dark' ? 'light' : 'dark'
  document.documentElement.classList.toggle('dark', colorMode.value === 'dark')
  localStorage.setItem('colorMode', colorMode.value)
}

function handleLogout() {
  auth.logout()
  router.push('/')
}

onMounted(() => {
  if (!auth.isLoggedIn.value || auth.user.value?.role !== 'admin') {
    console.warn('User tidak memiliki akses admin, redirecting...')
    router.push('/auth')
    return
  }

  const saved = localStorage.getItem('colorMode')
  if (saved) {
    colorMode.value = saved
    document.documentElement.classList.toggle('dark', saved === 'dark')
  } else {
    colorMode.value = 'light'
    document.documentElement.classList.remove('dark')
  }
})

watch(() => route.path, () => {
  isSidebarOpen.value = false
  if (!auth.isLoggedIn.value || auth.user.value?.role !== 'admin') {
    router.push('/auth')
  }
})

watch(() => auth.isLoggedIn.value, (newVal) => {
  if (!newVal) {
    router.push('/auth')
  }
})
</script>

<style scoped>
/* Enhanced Custom Animations */
@keyframes float {
  0%, 100% { 
    transform: translateY(0px) rotate(0deg) scale(1);
    opacity: 0.6;
  }
  33% { 
    transform: translateY(-15px) rotate(120deg) scale(1.1);
    opacity: 0.8;
  }
  66% { 
    transform: translateY(-25px) rotate(240deg) scale(0.9);
    opacity: 0.7;
  }
}

@keyframes pulse-glow {
  0%, 100% { 
    box-shadow: 
      0 0 5px rgba(34, 211, 238, 0.5),
      0 0 10px rgba(34, 211, 238, 0.3);
    opacity: 1;
  }
  50% { 
    box-shadow: 
      0 0 15px rgba(34, 211, 238, 0.8),
      0 0 25px rgba(34, 211, 238, 0.6),
      0 0 35px rgba(34, 211, 238, 0.4);
    opacity: 0.8;
  }
}

@keyframes pulse-slow {
  0%, 100% { 
    opacity: 0.3; 
    transform: scale(1) rotate(0deg);
  }
  50% { 
    opacity: 0.5; 
    transform: scale(1.05) rotate(180deg);
  }
}

@keyframes fade-in {
  from { 
    opacity: 0; 
    transform: translateY(-10px);
  }
  to { 
    opacity: 1; 
    transform: translateY(0);
  }
}

@keyframes slide-in {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animate-pulse-glow {
  animation: pulse-glow 2s ease-in-out infinite;
}

.animate-pulse-slow {
  animation: pulse-slow 8s ease-in-out infinite;
}

.animate-fade-in {
  animation: fade-in 0.5s ease-out;
}

.animate-slide-in {
  animation: slide-in 0.6s ease-out;
}

/* Enhanced Neon Glow Effects */
.sidebar-neon-glow {
  box-shadow: 
    0 0 20px rgba(34, 211, 238, 0.1),
    0 0 40px rgba(34, 211, 238, 0.05),
    0 0 60px rgba(34, 211, 238, 0.025),
    inset 0 0 20px rgba(34, 211, 238, 0.05);
  border-right: 1px solid rgba(34, 211, 238, 0.2) !important;
}

.sidebar-light-glow {
  box-shadow: 
    0 0 30px rgba(59, 130, 246, 0.1),
    0 0 50px rgba(59, 130, 246, 0.05),
    inset 0 0 20px rgba(59, 130, 246, 0.05);
  border-right: 1px solid rgba(59, 130, 246, 0.2) !important;
}

.content-neon-glow {
  box-shadow: 
    0 0 30px rgba(34, 211, 238, 0.15),
    0 0 50px rgba(34, 211, 238, 0.1),
    0 0 70px rgba(34, 211, 238, 0.05);
}

.content-light-glow {
  box-shadow: 
    0 0 40px rgba(59, 130, 246, 0.15),
    0 0 60px rgba(59, 130, 246, 0.1);
}

.neon-glow {
  box-shadow: 
    0 0 10px rgba(34, 211, 238, 0.3),
    0 0 20px rgba(34, 211, 238, 0.2),
    0 0 30px rgba(34, 211, 238, 0.1);
}

/* Background Grid Pattern */
.bg-grid-black {
  background-image: 
    linear-gradient(rgba(0, 0, 0, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0, 0, 0, 0.1) 1px, transparent 1px);
  background-size: 25px 25px;
}

.bg-grid-white {
  background-image: 
    linear-gradient(rgba(255, 255, 255, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 25px 25px;
}

/* Enhanced transitions */
* {
  transition: 
    background-color 0.5s cubic-bezier(0.4, 0, 0.2, 1),
    border-color 0.5s cubic-bezier(0.4, 0, 0.2, 1),
    color 0.5s cubic-bezier(0.4, 0, 0.2, 1),
    transform 0.3s cubic-bezier(0.4, 0, 0.2, 1),
    box-shadow 0.3s cubic-bezier(0.4, 0, 0.2, 1),
    opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Custom scrollbar */
aside ::-webkit-scrollbar {
  width: 8px;
}

aside ::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  margin: 4px 0;
}

aside ::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #3b82f6, #06b6d4);
  border-radius: 10px;
  transition: all 0.3s ease;
}

aside ::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #2563eb, #0891b2);
  transform: scale(1.1);
}

/* Glass morphism enhancement */
.glass-effect {
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
}

/* Smooth page transitions */
.page-enter-active,
.page-leave-active {
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.page-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.98);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-20px) scale(0.98);
}

/* Delay utilities */
.delay-500 {
  animation-delay: 500ms;
}

.delay-1000 {
  animation-delay: 1000ms;
}

.delay-1500 {
  animation-delay: 1500ms;
}

.delay-2000 {
  animation-delay: 2000ms;
}

.delay-4000 {
  animation-delay: 4000ms;
}

/* Hover scale utilities */
.hover-scale-102 {
  transform: scale(1);
  transition: transform 0.3s ease;
}

.hover-scale-102:hover {
  transform: scale(1.02);
}
</style>  