<template>
  <nav 
    class="bg-white/90 dark:bg-gray-900/90 backdrop-blur-2xl border-b border-gray-200/60 dark:border-gray-700/60 shadow-2xl sticky top-0 z-50 transition-all duration-500"
    :class="{
      'h-20': !scrolled,
      'h-16': scrolled,
      'shadow-lg': scrolled
    }"
  >
    <!-- Animated Background -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
      <div class="absolute inset-0 opacity-20">
        <div class="w-full h-full bg-grid-pattern animate-grid-move"></div>
      </div>
      
      <div 
        v-for="i in 12" 
        :key="i"
        class="absolute w-2 h-2 rounded-full animate-float-particle"
        :class="particleClasses[i % particleClasses.length]"
        :style="{
          left: Math.random() * 100 + '%',
          top: Math.random() * 100 + '%',
          animationDelay: Math.random() * 8 + 's',
          animationDuration: (4 + Math.random() * 6) + 's'
        }"
      ></div>
      
      <div class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-blue-500/30 to-transparent animate-scan-line"></div>
      <div class="absolute bottom-0 left-0 w-full h-0.5 bg-gradient-to-r from-transparent via-purple-500/20 to-transparent animate-scan-line delay-2000"></div>
      
      <div class="absolute -top-10 -left-10 w-20 h-20 bg-blue-500/10 rounded-full blur-xl animate-orb-pulse"></div>
      <div class="absolute -top-5 -right-5 w-16 h-16 bg-purple-500/10 rounded-full blur-xl animate-orb-pulse delay-1500"></div>
      <div class="absolute bottom-0 left-1/4 w-12 h-12 bg-cyan-500/10 rounded-full blur-lg animate-orb-pulse delay-3000"></div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-full">
        <!-- Logo Section -->
        <div class="flex items-center">
          <NuxtLink 
            to="/" 
            class="flex items-center space-x-3 group relative overflow-hidden py-2 px-1"
            @mouseenter="startLogoAnimation"
            @mouseleave="stopLogoAnimation"
          >
            <div class="absolute inset-0 bg-gradient-to-r from-blue-500/0 via-blue-500/15 to-purple-500/0 transform -translate-x-full group-hover:translate-x-full transition-transform duration-1000 rounded-2xl"></div>
            
            <div
              class="relative w-12 h-12 bg-gradient-to-br from-blue-500 via-purple-500 to-cyan-500 rounded-2xl flex items-center justify-center transition-all duration-700 shadow-2xl group-hover:shadow-3xl group/logo overflow-hidden"
              :class="{
                'scale-110 rotate-3': logoAnimating,
                'animate-logo-glow': logoAnimating
              }"
            >
              <div class="absolute inset-0 border-2 border-white/30 rounded-2xl animate-ping-slow"></div>
              <div class="absolute inset-1 border border-white/20 rounded-xl animate-ping-slower"></div>
              
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/40 to-transparent transform -translate-x-full group-hover/logo:translate-x-full transition-transform duration-1000"></div>
              
              <div class="absolute inset-1 bg-gradient-to-br from-white/30 to-transparent rounded-xl"></div>
              
              <svg class="w-6 h-6 text-white relative z-10 transition-all duration-500 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0 0h4m-4 0h4"/>
              </svg>
              
              <div class="absolute -top-1 -right-1 w-2 h-2 bg-green-400 rounded-full animate-ping"></div>
              <div class="absolute -bottom-1 -left-1 w-1.5 h-1.5 bg-yellow-400 rounded-full animate-ping delay-1000"></div>
            </div>
            
            <div class="flex flex-col relative z-10">
              <span class="text-2xl font-black bg-gradient-to-r from-gray-900 via-blue-600 to-purple-600 dark:from-white dark:via-blue-400 dark:to-purple-400 bg-clip-text text-transparent transition-all duration-500 group-hover:scale-105 transform origin-left">
                EduNews
              </span>
              <span class="text-xs text-gray-500 dark:text-gray-400 font-medium -mt-1 transition-all duration-300 delay-75 group-hover:translate-x-1">
                Stay Informed
              </span>
            </div>
            
            <div class="absolute -top-1 -right-1 flex space-x-1">
              <div class="w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse"></div>
              <div class="w-1.5 h-1.5 bg-blue-500 rounded-full animate-pulse delay-300"></div>
            </div>
          </NuxtLink>
        </div>

        <!-- Navigation Links - Center -->
        <div class="absolute left-1/2 transform -translate-x-1/2">
          <div class="hidden lg:flex items-center space-x-1 bg-white/60 dark:bg-gray-800/60 rounded-3xl p-2 backdrop-blur-xl border border-gray-200/50 dark:border-gray-700/50 shadow-lg">
            <NuxtLink
              v-for="link in navigationLinks"
              :key="link.to"
              :to="link.to"
              class="relative px-6 py-3 rounded-2xl font-semibold transition-all duration-500 group/nav overflow-hidden"
              :class="{
                'text-white': $route.path === link.to,
                'text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400': $route.path !== link.to
              }"
              @mouseenter="setActiveHover(link.to)"
              @mouseleave="activeHover = null"
            >
              <div 
                class="absolute inset-0 rounded-2xl transition-all duration-500 transform"
                :class="{
                  'bg-gradient-to-r from-blue-500 to-purple-600 scale-100': $route.path === link.to,
                  'bg-gradient-to-r from-blue-500/20 to-purple-500/20 scale-105': activeHover === link.to && $route.path !== link.to,
                  'scale-0': $route.path !== link.to && activeHover !== link.to
                }"
              ></div>

              <div 
                class="absolute inset-0 rounded-2xl border-2 border-transparent transition-all duration-300"
                :class="{
                  'border-white/30': $route.path === link.to,
                  'border-blue-500/30': activeHover === link.to && $route.path !== link.to
                }"
              ></div>

              <span class="relative z-10 flex items-center space-x-2">
                <component 
                  :is="link.icon" 
                  class="w-4 h-4 transition-all duration-300"
                  :class="{
                    'scale-125': $route.path === link.to || activeHover === link.to
                  }"
                />
                <span class="transition-all duration-300"
                      :class="{
                        'translate-x-0.5': $route.path === link.to || activeHover === link.to
                      }">
                  {{ link.name }}
                </span>
              </span>
            </NuxtLink>
          </div>
        </div>

        <!-- Right Section -->
        <div class="flex items-center space-x-2">
          <!-- Search Button -->
          <button
            class="p-3 rounded-2xl bg-white/60 dark:bg-gray-800/60 backdrop-blur-xl border border-gray-200/50 dark:border-gray-700/50 text-gray-600 dark:text-gray-300 transition-all duration-500 hover:scale-110 hover:shadow-2xl group/search"
          >
            <div class="relative z-10 transform transition-transform duration-300 group-hover/search:scale-110">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"/>
              </svg>
            </div>
            <div class="absolute inset-0 rounded-2xl border-2 border-transparent group-hover/search:border-blue-500/20 transition-all duration-500 animate-pulse"></div>
          </button>

          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="relative p-3 rounded-2xl bg-white/60 dark:bg-gray-800/60 backdrop-blur-xl border border-gray-200/50 dark:border-gray-700/50 transition-all duration-700 hover:scale-110 hover:shadow-2xl group/toggle overflow-hidden"
            :class="{
              'rotate-180': colorMode === 'dark',
              'text-yellow-500': colorMode === 'light',
              'text-blue-300': colorMode === 'dark'
            }"
            :title="colorMode === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode'"
          >
            <div class="absolute inset-0 bg-gradient-to-br from-blue-500/10 to-purple-500/10 opacity-0 group-hover/toggle:opacity-100 rounded-2xl transition-all duration-500"></div>
            
            <div class="relative z-10">
              <svg v-if="colorMode === 'light'" class="w-5 h-5 transform transition-all duration-700 group-hover/toggle:rotate-180 group-hover/toggle:scale-125" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
              </svg>
              <svg v-else class="w-5 h-5 transform transition-all duration-700 group-hover/toggle:rotate-180 group-hover/toggle:scale-125" fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
              </svg>
            </div>
            
            <div v-if="colorMode === 'light'" class="absolute inset-0">
              <div class="absolute top-1 left-1 w-1 h-1 bg-yellow-400 rounded-full animate-orbit-sun"></div>
            </div>
          </button>

          <!-- Auth Buttons -->
          <ClientOnly>
            <template v-if="auth.isLoggedIn.value">
              <div class="relative group/user">
                <button
                  class="flex items-center space-x-2 px-4 py-2.5 rounded-2xl bg-gradient-to-r from-blue-500 to-purple-600 text-white transition-all duration-500 shadow-2xl hover:shadow-3xl transform hover:scale-105 group/btn relative overflow-hidden"
                >
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -translate-x-full group-hover/btn:translate-x-full transition-transform duration-1000"></div>
                  
                  <div class="w-8 h-8 bg-white/20 rounded-2xl flex items-center justify-center backdrop-blur-sm border-2 border-white/30 transition-all duration-500 group-hover/btn:rotate-12 relative overflow-hidden">
                    <span class="text-sm font-bold relative z-10">{{ userInitial }}</span>
                    <div class="absolute inset-0 bg-gradient-to-br from-white/10 to-transparent"></div>
                  </div>
                  
                  <span class="font-semibold text-sm hidden sm:block transition-all duration-300">
                    {{ userName }}
                  </span>
                  
                  <svg class="w-4 h-4 transition-all duration-500 transform group-hover/user:rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                  </svg>
                </button>
                
                <div class="absolute right-0 top-full mt-2 w-56 bg-white/95 dark:bg-gray-800/95 backdrop-blur-xl rounded-2xl shadow-2xl border border-gray-200/50 dark:border-gray-700/50 opacity-0 invisible group-hover/user:opacity-100 group-hover/user:visible transition-all duration-300 transform translate-y-2 group-hover/user:translate-y-0 overflow-hidden">
                  <div class="absolute inset-0 bg-gradient-to-br from-blue-500/5 to-purple-500/5"></div>
                  
                  <div class="relative z-10 p-2">
                    <div class="px-3 py-2 border-b border-gray-200/30 dark:border-gray-700/30 mb-1">
                      <p class="text-sm font-bold text-gray-900 dark:text-white truncate">{{ userName }}</p>
                      <p class="text-xs text-gray-500 dark:text-gray-400 truncate mt-1">{{ userEmail }}</p>
                    </div>
                    
                    <button
                      @click="handleLogout"
                      class="w-full flex items-center space-x-2 px-3 py-2.5 text-sm text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-xl transition-all duration-300 group/logout"
                    >
                      <svg class="w-4 h-4 transform group-hover/logout:rotate-12 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
                class="px-5 py-2.5 rounded-2xl bg-gradient-to-r from-blue-500 to-purple-600 text-white transition-all duration-500 shadow-2xl hover:shadow-3xl transform hover:scale-105 font-semibold flex items-center space-x-2 group/login relative overflow-hidden"
              >
                <div class="absolute inset-0 bg-gradient-to-r from-blue-600 to-purple-700 opacity-0 group-hover/login:opacity-100 transition-opacity duration-500"></div>
                
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -translate-x-full group-hover/login:translate-x-full transition-transform duration-1000"></div>
                
                <svg class="w-4 h-4 transform group-hover/login:scale-110 transition-transform duration-300 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"/>
                </svg>
                <span class="relative z-10 text-sm transform transition-transform duration-300 group-hover/login:translate-x-0.5">
                  Login
                </span>
              </NuxtLink>
            </template>
          </ClientOnly>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
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
let logoAnimationTimer: NodeJS.Timeout | null = null

// Data
const particleClasses = [
  'bg-blue-400/40',
  'bg-purple-400/40', 
  'bg-cyan-400/40',
  'bg-pink-400/40',
  'bg-yellow-400/40',
  'bg-green-400/40'
]

const navigationLinks = [
  {
    to: '/',
    name: 'Beranda',
    icon: 'HomeIcon'
  },
  {
    to: '/populer',
    name: 'Populer',
    icon: 'TrendingUpIcon'
  },
  {
    to: '/tentang',
    name: 'Tentang',
    icon: 'InformationCircleIcon'
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

const setActiveHover = (route: string) => {
  activeHover.value = route
}

const startLogoAnimation = () => {
  logoAnimating.value = true
  if (logoAnimationTimer) clearTimeout(logoAnimationTimer)
}

const stopLogoAnimation = () => {
  logoAnimationTimer = setTimeout(() => {
    logoAnimating.value = false
  }, 300)
}

const handleLogout = () => {
  auth.logout()
  router.push('/')
}

const handleScroll = () => {
  scrolled.value = window.scrollY > 10
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
    handleScroll()
  }
})

onUnmounted(() => {
  if (process.client) {
    window.removeEventListener('scroll', handleScroll)
  }
  if (logoAnimationTimer) clearTimeout(logoAnimationTimer)
})
</script>

<style scoped>
/* Animations */
@keyframes float-particle {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
    opacity: 0.3;
  }
  25% {
    transform: translateY(-15px) translateX(10px) rotate(90deg);
    opacity: 0.6;
  }
  50% {
    transform: translateY(-25px) translateX(-5px) rotate(180deg);
    opacity: 0.8;
  }
  75% {
    transform: translateY(-15px) translateX(-10px) rotate(270deg);
    opacity: 0.6;
  }
}

@keyframes orb-pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.1;
  }
  50% {
    transform: scale(1.2);
    opacity: 0.2;
  }
}

@keyframes scan-line {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

@keyframes grid-move {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 50px 50px;
  }
}

@keyframes logo-glow {
  0%, 100% {
    box-shadow: 0 0 20px rgba(59, 130, 246, 0.3);
  }
  50% {
    box-shadow: 0 0 30px rgba(59, 130, 246, 0.6), 0 0 40px rgba(139, 92, 246, 0.4);
  }
}

@keyframes ping-slow {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.5;
  }
}

@keyframes ping-slower {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.3;
  }
}

@keyframes orbit-sun {
  0% {
    transform: rotate(0deg) translateX(8px) rotate(0deg);
  }
  100% {
    transform: rotate(360deg) translateX(8px) rotate(-360deg);
  }
}

.animate-float-particle {
  animation: float-particle 8s ease-in-out infinite;
}

.animate-orb-pulse {
  animation: orb-pulse 6s ease-in-out infinite;
}

.animate-scan-line {
  animation: scan-line 3s linear infinite;
}

.animate-grid-move {
  animation: grid-move 20s linear infinite;
}

.animate-logo-glow {
  animation: logo-glow 2s ease-in-out infinite;
}

.animate-ping-slow {
  animation: ping-slow 3s ease-in-out infinite;
}

.animate-ping-slower {
  animation: ping-slower 4s ease-in-out infinite;
}

.animate-orbit-sun {
  animation: orbit-sun 4s linear infinite;
}

.bg-grid-pattern {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
}

.backdrop-blur-2xl {
  backdrop-filter: blur(24px) saturate(200%);
  -webkit-backdrop-filter: blur(24px) saturate(200%);
}

/* Icon Components */
.HomeIcon,
.TrendingUpIcon, 
.InformationCircleIcon {
  fill: none;
  stroke: currentColor;
}
</style>