<template>
  <nav 
    class="bg-white/95 dark:bg-gray-900/95 backdrop-blur-2xl border-b border-gray-200/40 dark:border-gray-700/40 shadow-lg sticky top-0 z-50 transition-all duration-300"
    :class="{
      'h-20': !scrolled,
      'h-16': scrolled,
      'shadow-md': scrolled
    }"
  >
    <!-- Modern Background Effects -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <!-- Gradient Overlay -->
      <div class="absolute inset-0 bg-gradient-to-b from-white/30 to-transparent dark:from-gray-900/30"></div>
      
      <!-- Subtle Grid Pattern -->
      <div class="absolute inset-0 opacity-[0.02] dark:opacity-[0.03]">
        <div class="w-full h-full bg-grid-pattern"></div>
      </div>
      
      <!-- Floating Particles -->
      <div 
        v-for="i in 8" 
        :key="i"
        class="absolute w-1.5 h-1.5 rounded-full bg-blue-400/20 dark:bg-blue-400/10 animate-float-particle"
        :style="{
          left: Math.random() * 100 + '%',
          top: Math.random() * 100 + '%',
          animationDelay: Math.random() * 4 + 's',
          animationDuration: (6 + Math.random() * 4) + 's'
        }"
      ></div>
      
      <!-- Accent Glow Effects -->
      <div class="absolute -top-10 -left-10 w-24 h-24 bg-blue-500/5 dark:bg-blue-500/3 rounded-full blur-2xl"></div>
      <div class="absolute -top-5 -right-5 w-16 h-16 bg-purple-500/5 dark:bg-purple-500/3 rounded-full blur-2xl"></div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-full">
      <div class="flex items-center justify-between h-full">
        <!-- Logo Section - Compact -->
        <div class="flex items-center flex-shrink-0">
          <NuxtLink 
            to="/" 
            class="flex items-center space-x-2 group relative py-1.5 px-1.5"
            @mouseenter="startLogoAnimation"
            @mouseleave="stopLogoAnimation"
          >
            <!-- Logo Container -->
            <div
              class="relative w-10 h-10 bg-gradient-to-br from-blue-500 via-purple-500 to-cyan-500 rounded-xl flex items-center justify-center transition-all duration-500 shadow-lg group-hover:shadow-xl group-hover:scale-105 overflow-hidden group/logo"
              :class="{
                'animate-logo-glow': logoAnimating
              }"
            >
              <!-- Glow Border -->
              <div class="absolute inset-0 border-2 border-white/20 rounded-xl animate-ping-slow"></div>
              
              <!-- Shine Effect -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -translate-x-full group-hover/logo:translate-x-full transition-transform duration-700"></div>
              
              <!-- Icon -->
              <svg class="w-5 h-5 text-white relative z-10 transition-transform duration-300 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0 0h4m-4 0h4"/>
              </svg>
              
              <!-- Status Dots -->
              <div class="absolute -top-0.5 -right-0.5 w-1.5 h-1.5 bg-green-400 rounded-full animate-pulse"></div>
            </div>
            
            <!-- Brand Text -->
            <div class="flex flex-col">
              <span class="text-xl font-bold bg-gradient-to-r from-gray-800 via-blue-600 to-purple-600 dark:from-white dark:via-blue-400 dark:to-purple-400 bg-clip-text text-transparent transition-all duration-300 group-hover:scale-105 transform origin-left">
                EduNews
              </span>
              <span class="text-xs text-gray-500 dark:text-gray-400 font-medium -mt-0.5 transition-all duration-300 group-hover:translate-x-0.5">
                Stay Informed
              </span>
            </div>
          </NuxtLink>
        </div>

        <!-- Navigation Cards - Modern Compact Design -->
        <div class="flex-1 flex justify-center px-4">
          <div class="hidden lg:flex items-center space-x-2 bg-white/80 dark:bg-gray-800/80 backdrop-blur-xl rounded-2xl p-1.5 shadow-lg border border-gray-200/30 dark:border-gray-700/30">
            <NuxtLink
              v-for="link in navigationLinks"
              :key="link.to"
              :to="link.to"
              class="relative px-4 py-2.5 rounded-xl font-medium text-sm transition-all duration-400 group/nav"
              :class="{
                'text-white': $route.path === link.to,
                'text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400': $route.path !== link.to
              }"
              @mouseenter="activeHover = link.to"
              @mouseleave="activeHover = null"
            >
              <!-- Active/Hover Background -->
              <div 
                class="absolute inset-0 rounded-xl transition-all duration-400"
                :class="{
                  'bg-gradient-to-r from-blue-500 to-purple-600': $route.path === link.to,
                  'bg-gradient-to-r from-blue-100/80 to-purple-100/80 dark:from-blue-900/30 dark:to-purple-900/30 scale-105': activeHover === link.to && $route.path !== link.to,
                  'group-hover/nav:bg-gray-100/50 dark:group-hover/nav:bg-gray-700/30': $route.path !== link.to
                }"
              ></div>

              <!-- Border Effect -->
              <div 
                class="absolute inset-0 rounded-xl border transition-all duration-300"
                :class="{
                  'border-white/30': $route.path === link.to,
                  'border-blue-400/30 dark:border-blue-500/30': activeHover === link.to && $route.path !== link.to,
                  'border-transparent': $route.path !== link.to && activeHover !== link.to
                }"
              ></div>

              <!-- Content -->
              <span class="relative z-10 flex items-center space-x-1.5">
                <component 
                  :is="link.icon" 
                  class="w-3.5 h-3.5 transition-all duration-300"
                  :class="{
                    'scale-110': $route.path === link.to || activeHover === link.to
                  }"
                />
                <span class="font-semibold tracking-wide transition-all duration-300"
                      :class="{
                        'scale-105': $route.path === link.to || activeHover === link.to
                      }">
                  {{ link.name }}
                </span>
              </span>
              
              <!-- Active Indicator -->
              <div 
                v-if="$route.path === link.to"
                class="absolute -bottom-1 left-1/2 transform -translate-x-1/2 w-6 h-0.5 bg-gradient-to-r from-blue-400 to-purple-500 rounded-full"
              ></div>
            </NuxtLink>
          </div>
        </div>

        <!-- Mobile Menu Button -->
        <div class="lg:hidden flex items-center">
          <button
            @click="isMobileMenuOpen = !isMobileMenuOpen"
            class="p-2 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-xl border border-gray-200/30 dark:border-gray-700/30 text-gray-600 dark:text-gray-300 transition-all duration-300 hover:scale-105"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path 
                stroke-linecap="round" 
                stroke-linejoin="round" 
                stroke-width="2" 
                :d="isMobileMenuOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'"
              />
            </svg>
          </button>
        </div>

        <!-- Right Section - Compact -->
        <div class="hidden lg:flex items-center space-x-1.5 flex-shrink-0">
          <!-- Search Button -->
          <button
            class="p-2.5 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-xl border border-gray-200/30 dark:border-gray-700/30 text-gray-600 dark:text-gray-300 transition-all duration-300 hover:scale-105 hover:shadow-md group/search"
            title="Search"
          >
            <div class="relative transform transition-transform duration-200 group-hover/search:scale-110">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
            </div>
            <div class="absolute inset-0 rounded-xl group-hover/search:bg-blue-500/5 dark:group-hover/search:bg-blue-500/10 transition-all duration-300"></div>
          </button>

          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="relative p-2.5 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-xl border border-gray-200/30 dark:border-gray-700/30 transition-all duration-500 hover:scale-105 group/toggle"
            :title="colorMode === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode'"
          >
            <div class="absolute inset-0 rounded-xl bg-gradient-to-br from-blue-500/5 to-purple-500/5 opacity-0 group-hover/toggle:opacity-100 transition-all duration-300"></div>
            
            <div class="relative">
              <svg v-if="colorMode === 'light'" class="w-4 h-4 text-yellow-500 transform transition-all duration-500 group-hover/toggle:rotate-90" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
              </svg>
              <svg v-else class="w-4 h-4 text-blue-300 transform transition-all duration-500 group-hover/toggle:rotate-90" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
              </svg>
            </div>
          </button>

          <!-- Auth Button -->
          <ClientOnly>
            <template v-if="auth.isLoggedIn.value">
              <div class="relative group/user">
                <button
                  class="flex items-center space-x-1.5 px-3 py-2 rounded-xl bg-gradient-to-r from-blue-500 to-purple-600 text-white transition-all duration-300 shadow-md hover:shadow-lg transform hover:scale-105 group/btn overflow-hidden"
                >
                  <!-- Shine Effect -->
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/15 to-transparent transform -translate-x-full group-hover/btn:translate-x-full transition-transform duration-700"></div>
                  
                  <!-- Avatar -->
                  <div class="w-7 h-7 bg-white/20 rounded-lg flex items-center justify-center backdrop-blur-sm border border-white/30 transition-all duration-300 group-hover/btn:rotate-6 relative overflow-hidden">
                    <span class="text-xs font-bold relative z-10">{{ userInitial }}</span>
                    <div class="absolute inset-0 bg-gradient-to-br from-white/10 to-transparent"></div>
                  </div>
                  
                  <!-- Name & Dropdown Icon -->
                  <span class="font-semibold text-xs hidden sm:block transition-all duration-300">
                    {{ userName }}
                  </span>
                  <svg class="w-3 h-3 transition-all duration-300 transform group-hover/user:rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                  </svg>
                </button>
                
                <!-- Dropdown Menu -->
                <div class="absolute right-0 top-full mt-1.5 w-48 bg-white/95 dark:bg-gray-800/95 backdrop-blur-xl rounded-xl shadow-lg border border-gray-200/50 dark:border-gray-700/50 opacity-0 invisible group-hover/user:opacity-100 group-hover/user:visible transition-all duration-200 transform translate-y-1 group-hover/user:translate-y-0 overflow-hidden">
                  <div class="relative z-10 p-1.5">
                    <!-- User Info -->
                    <div class="px-2.5 py-2 border-b border-gray-200/30 dark:border-gray-700/30 mb-1">
                      <p class="text-sm font-bold text-gray-900 dark:text-white truncate">{{ userName }}</p>
                      <p class="text-xs text-gray-500 dark:text-gray-400 truncate mt-0.5">{{ userEmail }}</p>
                    </div>
                    
                    <!-- Logout Button -->
                    <button
                      @click="handleLogout"
                      class="w-full flex items-center space-x-2 px-2.5 py-2 text-xs text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-all duration-200 group/logout"
                    >
                      <svg class="w-3.5 h-3.5 transform group-hover/logout:rotate-12 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
                      </svg>
                      <span class="font-semibold">Logout</span>
                    </button>
                  </div>
                </div>
              </div>
            </template>
            <template v-else>
              <NuxtLink
                to="/auth"
                class="px-3.5 py-2 rounded-xl bg-gradient-to-r from-blue-500 to-purple-600 text-white transition-all duration-300 shadow-md hover:shadow-lg transform hover:scale-105 font-semibold flex items-center space-x-1.5 group/login overflow-hidden text-sm"
              >
                <!-- Hover Effect -->
                <div class="absolute inset-0 bg-gradient-to-r from-blue-600 to-purple-700 opacity-0 group-hover/login:opacity-100 transition-opacity duration-300"></div>
                
                <!-- Shine Effect -->
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/15 to-transparent transform -translate-x-full group-hover/login:translate-x-full transition-transform duration-700"></div>
                
                <!-- Icon -->
                <svg class="w-3.5 h-3.5 transform group-hover/login:scale-110 transition-transform duration-200 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"/>
                </svg>
                
                <!-- Text -->
                <span class="relative z-10 transform transition-transform duration-200 group-hover/login:translate-x-0.5">
                  Login
                </span>
              </NuxtLink>
            </template>
          </ClientOnly>
        </div>
      </div>

      <!-- Mobile Menu -->
      <div
        v-if="isMobileMenuOpen"
        class="lg:hidden mt-4 pb-4 border-t border-gray-200/50 dark:border-gray-700/50 pt-4"
      >
        <div class="space-y-2">
          <NuxtLink
            v-for="link in navigationLinks"
            :key="link.to"
            :to="link.to"
            @click="isMobileMenuOpen = false"
            class="block px-4 py-3 rounded-xl font-medium transition-all duration-300"
            :class="{
              'bg-gradient-to-r from-blue-500 to-purple-600 text-white': $route.path === link.to,
              'text-gray-700 dark:text-gray-300 hover:bg-gray-100/50 dark:hover:bg-gray-700/30': $route.path !== link.to
            }"
          >
            <div class="flex items-center space-x-3">
              <component :is="link.icon" class="w-5 h-5" />
              <span>{{ link.name }}</span>
            </div>
          </NuxtLink>
          
          <!-- Mobile Search and Theme -->
          <div class="flex items-center space-x-2 px-4 pt-2">
            <button
              class="flex-1 px-4 py-3 rounded-xl bg-gray-100/50 dark:bg-gray-800/50 text-gray-600 dark:text-gray-300 flex items-center justify-center space-x-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
              <span>Search</span>
            </button>
            
            <button
              @click="toggleTheme"
              class="px-4 py-3 rounded-xl bg-gray-100/50 dark:bg-gray-800/50 text-gray-600 dark:text-gray-300"
            >
              <svg v-if="colorMode === 'light'" class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
              </svg>
              <svg v-else class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
              </svg>
            </button>
          </div>
          
          <!-- Mobile Auth -->
          <div class="px-4 pt-2">
            <ClientOnly>
              <template v-if="auth.isLoggedIn.value">
                <div class="px-4 py-3 rounded-xl bg-gradient-to-r from-blue-500 to-purple-600 text-white">
                  <p class="font-semibold">{{ userName }}</p>
                  <p class="text-sm opacity-90">{{ userEmail }}</p>
                  <button
                    @click="handleLogout"
                    class="mt-3 w-full py-2 rounded-lg bg-white/20 flex items-center justify-center space-x-2"
                  >
                    <span>Logout</span>
                  </button>
                </div>
              </template>
              <template v-else>
                <NuxtLink
                  to="/auth"
                  @click="isMobileMenuOpen = false"
                  class="block px-4 py-3 rounded-xl bg-gradient-to-r from-blue-500 to-purple-600 text-white text-center font-semibold"
                >
                  Login
                </NuxtLink>
              </template>
            </ClientOnly>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, h } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuth } from '~/composables/useAuth'

const router = useRouter()
const route = useRoute()
const auth = useAuth()

// States
const colorMode = ref<'light' | 'dark'>('light')
const scrolled = ref(false)
const logoAnimating = ref(false)
const activeHover = ref<string | null>(null)
const isMobileMenuOpen = ref(false)
let logoAnimationTimer: NodeJS.Timeout | null = null

// Icon Components
const HomeIcon = () => h('svg', {
  class: 'w-3.5 h-3.5',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2',
    'd': 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6'
  })
])

const TrendingUpIcon = () => h('svg', {
  class: 'w-3.5 h-3.5',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2',
    'd': 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6'
  })
])

const InformationCircleIcon = () => h('svg', {
  class: 'w-3.5 h-3.5',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2',
    'd': 'M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z'
  })
])

// Navigation Data dengan ikon yang sudah didefinisikan
const navigationLinks = [
  {
    to: '/',
    name: 'Beranda',
    icon: HomeIcon
  },
  {
    to: '/populer',
    name: 'Populer',
    icon: TrendingUpIcon
  },
  {
    to: '/tentang',
    name: 'Tentang',
    icon: InformationCircleIcon
  }
]

// Computed
const userName = computed(() => auth.user.value?.name || 'User')
const userEmail = computed(() => auth.user.value?.email || '')
const userInitial = computed(() => {
  const name = auth.user.value?.name || 'U'
  return name.charAt(0).toUpperCase()
})

// Methods
const toggleTheme = () => {
  colorMode.value = colorMode.value === 'dark' ? 'light' : 'dark'
  document.documentElement.classList.toggle('dark', colorMode.value === 'dark')
  if (process.client) {
    localStorage.setItem('colorMode', colorMode.value)
  }
}

const startLogoAnimation = () => {
  logoAnimating.value = true
  if (logoAnimationTimer) clearTimeout(logoAnimationTimer)
}

const stopLogoAnimation = () => {
  logoAnimationTimer = setTimeout(() => {
    logoAnimating.value = false
  }, 200)
}

const handleLogout = () => {
  auth.logout()
  router.push('/')
  isMobileMenuOpen.value = false
}

const handleScroll = () => {
  scrolled.value = window.scrollY > 10
  // Close mobile menu on scroll
  if (isMobileMenuOpen.value) {
    isMobileMenuOpen.value = false
  }
}

// Close mobile menu when clicking outside
const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as HTMLElement
  if (isMobileMenuOpen.value && !target.closest('nav')) {
    isMobileMenuOpen.value = false
  }
}

// Lifecycle
onMounted(() => {
  if ('checkToken' in auth && typeof (auth as any).checkToken === 'function') {
    (auth as any).checkToken()
  }
  
  if (process.client) {
    const savedMode = localStorage.getItem('colorMode') as 'light' | 'dark' | null
    if (savedMode) {
      colorMode.value = savedMode
      document.documentElement.classList.toggle('dark', colorMode.value === 'dark')
    }
    
    window.addEventListener('scroll', handleScroll)
    document.addEventListener('click', handleClickOutside)
    handleScroll()
  }
})

onUnmounted(() => {
  if (process.client) {
    window.removeEventListener('scroll', handleScroll)
    document.removeEventListener('click', handleClickOutside)
  }
  if (logoAnimationTimer) clearTimeout(logoAnimationTimer)
})
</script>

<style scoped>
/* Optimized Animations */
@keyframes float-particle {
  0%, 100% {
    transform: translateY(0) translateX(0) rotate(0deg);
    opacity: 0.3;
  }
  50% {
    transform: translateY(-10px) translateX(5px) rotate(180deg);
    opacity: 0.8;
  }
}

@keyframes logo-glow {
  0%, 100% {
    box-shadow: 0 0 15px rgba(59, 130, 246, 0.3);
  }
  50% {
    box-shadow: 0 0 20px rgba(59, 130, 246, 0.5), 0 0 25px rgba(139, 92, 246, 0.3);
  }
}

@keyframes ping-slow {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.5;
  }
}

/* Utility Classes */
.animate-float-particle {
  animation: float-particle 8s ease-in-out infinite;
}

.animate-logo-glow {
  animation: logo-glow 2s ease-in-out infinite;
}

.animate-ping-slow {
  animation: ping-slow 2s ease-in-out infinite;
}

.bg-grid-pattern {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.05) 1px, transparent 1px);
  background-size: 30px 30px;
}

/* Enhanced Backdrop Blur */
.backdrop-blur-2xl {
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
}

/* Smooth Transitions */
* {
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}
</style>