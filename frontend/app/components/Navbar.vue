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
        <!-- Logo Section - FIXED -->
        <div class="flex items-center flex-shrink-0">
          <NuxtLink 
            to="/" 
            class="flex items-center space-x-3 group relative py-1.5 px-2"
            @mouseenter="startLogoAnimation"
            @mouseleave="stopLogoAnimation"
          >
            <!-- Logo Container - FIXED -->
            <div
              class="relative w-12 h-12 bg-gradient-to-br from-blue-500 via-purple-500 to-cyan-500 rounded-2xl flex items-center justify-center transition-all duration-500 shadow-xl group-hover:shadow-2xl overflow-hidden group/logo"
              :class="{
                'animate-logo-glow': logoAnimating,
                'group-hover:scale-110': true
              }"
            >
              <!-- SIMPLIFIED: Hanya gradient background tanpa overlay putih -->
              <div class="absolute inset-0 bg-gradient-to-br from-blue-500 via-purple-500 to-cyan-500 rounded-2xl"></div>
              
              <!-- Shine Effect -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent transform -translate-x-full group-hover/logo:translate-x-full transition-transform duration-1000"></div>
              
              <!-- Icon with Zoom/Fade Animation -->
              <svg class="w-6 h-6 text-white relative z-10 transition-all duration-500 group-hover:scale-125 group-hover:opacity-90" 
                   fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" 
                      d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0 0h4m-4 0h4"/>
              </svg>
              
              <!-- Animated Rings -->
              <div class="absolute inset-0 rounded-2xl border-2 border-white/30 animate-ping-slow"></div>
              
              <!-- Status Dots - Enhanced -->
              <div class="absolute -top-1 -right-1 w-3 h-3 bg-green-400 rounded-full animate-pulse shadow-lg shadow-green-400/50"></div>
            </div>
            
            <!-- Brand Text - Enhanced -->
            <div class="flex flex-col space-y-1">
              <span class="text-2xl font-black bg-gradient-to-r from-blue-600 via-purple-600 to-cyan-500 dark:from-blue-400 dark:via-purple-400 dark:to-cyan-300 bg-clip-text text-transparent transition-all duration-500 group-hover:scale-105 transform origin-left tracking-tight">
                EduNews
              </span>
              <div class="flex items-center space-x-2">
                <span class="text-xs font-semibold text-gray-600 dark:text-gray-400 transition-all duration-500 group-hover:translate-x-1">
                  Stay Informed
                </span>
                <div class="w-1 h-1 rounded-full bg-gradient-to-r from-blue-400 to-purple-400 animate-pulse"></div>
              </div>
            </div>
            
            <!-- Hover Effect Line -->
            <div class="absolute -bottom-1 left-0 right-0 h-0.5 bg-gradient-to-r from-blue-400 via-purple-400 to-cyan-400 transform scale-x-0 group-hover:scale-x-100 transition-transform duration-500 origin-left rounded-full"></div>
          </NuxtLink>
        </div>

        <!-- Navigation Cards - Modern Compact Design -->
        <div class="flex-1 flex justify-center px-6">
          <div class="hidden lg:flex items-center justify-center space-x-2 bg-white/90 dark:bg-gray-800/90 backdrop-blur-xl rounded-2xl p-2 shadow-2xl border border-gray-200/40 dark:border-gray-700/40 mx-auto">
            <NuxtLink
              v-for="link in navigationLinks"
              :key="link.to"
              :to="link.to"
              class="relative px-5 py-3 rounded-xl font-medium text-sm transition-all duration-500 group/nav flex flex-col items-center"
              :class="{
                'text-white shadow-lg': $route.path === link.to,
                'text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400': $route.path !== link.to
              }"
              @mouseenter="activeHover = link.to"
              @mouseleave="activeHover = null"
            >
              <!-- Animated Gradient Background -->
              <div 
                class="absolute inset-0 rounded-xl transition-all duration-500 overflow-hidden"
                :class="{
                  'bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 shadow-lg shadow-blue-500/30': $route.path === link.to,
                  'bg-gradient-to-r from-blue-100/50 to-purple-100/50 dark:from-blue-900/30 dark:to-purple-900/30 scale-105 shadow-md': activeHover === link.to && $route.path !== link.to,
                  'group-hover/nav:bg-gradient-to-r group-hover/nav:from-gray-100/50 group-hover/nav:to-gray-200/50 dark:group-hover/nav:from-gray-700/30 dark:group-hover/nav:to-gray-800/30': $route.path !== link.to
                }"
              >
                <!-- Animated Background Shine -->
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent transform -translate-x-full group-hover/nav:translate-x-full transition-transform duration-1000"></div>
              </div>

              <!-- Animated Border -->
              <div 
                class="absolute inset-0 rounded-xl transition-all duration-500"
                :class="{
                  'border-2 border-white/40 shadow-inner shadow-white/20': $route.path === link.to,
                  'border-2 border-blue-400/40 dark:border-blue-500/40 shadow-inner shadow-blue-400/20': activeHover === link.to && $route.path !== link.to,
                  'border border-gray-200/50 dark:border-gray-700/50': $route.path !== link.to && activeHover !== link.to
                }"
              ></div>

              <!-- Content -->
              <div class="relative z-10 flex items-center space-x-2.5">
                <component 
                  :is="link.icon" 
                  class="w-4 h-4 transition-all duration-500"
                  :class="{
                    'scale-125 opacity-100': $route.path === link.to || activeHover === link.to,
                    'group-hover/nav:scale-110 group-hover/nav:opacity-80': $route.path !== link.to
                  }"
                />
                <span class="font-bold tracking-wide transition-all duration-500"
                      :class="{
                        'scale-110': $route.path === link.to || activeHover === link.to,
                        'group-hover/nav:scale-105': $route.path !== link.to
                      }">
                  {{ link.name }}
                </span>
              </div>
              
              <!-- Active Indicator - FIXED -->
              <div 
                v-if="$route.path === link.to"
                class="absolute -bottom-2 left-0 right-0 mx-auto w-8 h-1.5 bg-gradient-to-r from-blue-400 via-purple-400 to-cyan-400 rounded-full shadow-lg shadow-blue-400/30 animate-pulse-slow"
              ></div>
            </NuxtLink>
          </div>
        </div>

        <!-- Mobile Menu Button -->
        <div class="lg:hidden flex items-center">
          <button
            @click="isMobileMenuOpen = !isMobileMenuOpen"
            class="p-3 rounded-2xl bg-white/90 dark:bg-gray-800/90 backdrop-blur-xl border border-gray-200/40 dark:border-gray-700/40 text-gray-600 dark:text-gray-300 transition-all duration-500 hover:scale-110 hover:shadow-lg group/mobile"
          >
            <svg class="w-5 h-5 transition-transform duration-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path 
                stroke-linecap="round" 
                stroke-linejoin="round" 
                stroke-width="2" 
                :d="isMobileMenuOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'"
              />
            </svg>
          </button>
        </div>

        <!-- Right Section - Enhanced -->
        <div class="hidden lg:flex items-center space-x-3 flex-shrink-0">
          <!-- Theme Toggle - Enhanced -->
          <button
            @click="toggleTheme"
            class="relative p-3 rounded-2xl bg-white/90 dark:bg-gray-800/90 backdrop-blur-xl border border-gray-200/40 dark:border-gray-700/40 transition-all duration-500 hover:scale-110 hover:shadow-lg group/toggle"
            :title="colorMode === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode'"
          >
            <!-- Animated Background -->
            <div class="absolute inset-0 rounded-2xl bg-gradient-to-br from-blue-500/10 to-purple-500/10 opacity-0 group-hover/toggle:opacity-100 transition-all duration-500"></div>
            
            <!-- Sun/Moon Icon with Fade/Zoom Animation -->
            <div class="relative">
              <svg v-if="colorMode === 'light'" 
                   class="w-5 h-5 text-yellow-500 transition-all duration-500 group-hover/toggle:scale-125 group-hover/toggle:opacity-80" 
                   fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
              </svg>
              <svg v-else 
                   class="w-5 h-5 text-blue-300 transition-all duration-500 group-hover/toggle:scale-125 group-hover/toggle:opacity-80" 
                   fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
              </svg>
            </div>
            
            <!-- Particle Effect -->
            <div class="absolute inset-0 rounded-2xl opacity-0 group-hover/toggle:opacity-100 transition-opacity duration-500">
              <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-12 h-12 bg-gradient-to-r from-blue-400/10 to-purple-400/10 rounded-full blur-md"></div>
            </div>
          </button>

          <!-- Auth Button - Enhanced -->
          <ClientOnly>
            <template v-if="auth.isLoggedIn.value">
              <div class="relative group/user">
                <button
                  class="flex items-center space-x-3 px-4 py-3 rounded-2xl bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 text-white transition-all duration-500 shadow-xl hover:shadow-2xl group/btn overflow-hidden"
                >
                  <!-- Animated Gradient Background -->
                  <div class="absolute inset-0 bg-gradient-to-r from-blue-600 via-purple-600 to-cyan-600 opacity-0 group-hover/btn:opacity-100 transition-opacity duration-500"></div>
                  
                  <!-- Shine Effect -->
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent transform -translate-x-full group-hover/btn:translate-x-full transition-transform duration-1000"></div>
                  
                  <!-- Avatar - Enhanced with Zoom/Fade -->
                  <div class="w-8 h-8 bg-white/30 rounded-xl flex items-center justify-center backdrop-blur-sm border-2 border-white/40 transition-all duration-500 group-hover/btn:scale-110 group-hover/btn:opacity-90 relative overflow-hidden">
                    <span class="text-sm font-black relative z-10 transition-transform duration-500 group-hover/btn:scale-110">{{ userInitial }}</span>
                    <div class="absolute inset-0 bg-gradient-to-br from-white/30 to-transparent"></div>
                  </div>
                  
                  <!-- Name & Dropdown Icon -->
                  <span class="font-bold text-sm transition-all duration-500 group-hover/btn:translate-x-1">
                    {{ userName }}
                  </span>
                  <svg class="w-3.5 h-3.5 transition-all duration-500 transform group-hover/user:translate-y-1" 
                       fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                  </svg>
                </button>
                
                <!-- Dropdown Menu - Enhanced -->
                <div class="absolute right-0 top-full mt-2 w-56 bg-white/95 dark:bg-gray-800/95 backdrop-blur-xl rounded-2xl shadow-2xl border border-gray-200/50 dark:border-gray-700/50 opacity-0 invisible group-hover/user:opacity-100 group-hover/user:visible transition-all duration-300 transform translate-y-2 group-hover/user:translate-y-0 overflow-hidden">
                  <!-- Gradient Border -->
                  <div class="absolute inset-0 rounded-2xl p-[1.5px] bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500">
                    <div class="w-full h-full rounded-2xl bg-white/95 dark:bg-gray-800/95 backdrop-blur-xl"></div>
                  </div>
                  
                  <div class="relative z-10 p-2">
                    <!-- User Info -->
                    <div class="px-3 py-3 border-b border-gray-200/30 dark:border-gray-700/30 mb-1">
                      <p class="text-sm font-black text-gray-900 dark:text-white truncate">{{ userName }}</p>
                      <p class="text-xs text-gray-600 dark:text-gray-400 truncate mt-1">{{ userEmail }}</p>
                    </div>
                    
                    <!-- Logout Button -->
                    <button
                      @click="handleLogout"
                      class="w-full flex items-center space-x-3 px-3 py-3 text-sm font-bold text-red-600 dark:text-red-400 hover:bg-red-50 dark:hover:bg-red-900/30 rounded-xl transition-all duration-300 group/logout"
                    >
                      <svg class="w-4 h-4 transition-transform duration-300 group-hover/logout:scale-110 group-hover/logout:opacity-80" 
                           fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
                      </svg>
                      <span class="transition-all duration-300 group-hover/logout:scale-105">Logout</span>
                    </button>
                  </div>
                </div>
              </div>
            </template>
            <template v-else>
              <NuxtLink
                to="/auth"
                class="relative px-5 py-3 rounded-2xl bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 text-white transition-all duration-500 shadow-xl hover:shadow-2xl font-black flex items-center space-x-3 group/login overflow-hidden"
              >
                <!-- Hover Gradient -->
                <div class="absolute inset-0 bg-gradient-to-r from-blue-600 via-purple-600 to-cyan-600 opacity-0 group-hover/login:opacity-100 transition-opacity duration-500"></div>
                
                <!-- Animated Shine Effect -->
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/40 to-transparent transform -translate-x-full group-hover/login:translate-x-full transition-transform duration-1000"></div>
                
                <!-- Icon with Zoom/Fade Animation -->
                <svg class="w-4 h-4 transition-all duration-500 group-hover/login:scale-125 group-hover/login:opacity-80 relative z-10" 
                     fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
                </svg>
                
                <!-- Text -->
                <span class="relative z-10 transition-all duration-500 group-hover/login:scale-105">
                  Login
                </span>
              </NuxtLink>
            </template>
          </ClientOnly>
        </div>
      </div>

      <!-- Mobile Menu - Enhanced -->
      <Transition
        enter-active-class="transition-all duration-500 ease-out"
        enter-from-class="transform opacity-0 -translate-y-4 scale-95"
        enter-to-class="transform opacity-100 translate-y-0 scale-100"
        leave-active-class="transition-all duration-300 ease-in"
        leave-from-class="transform opacity-100 translate-y-0 scale-100"
        leave-to-class="transform opacity-0 -translate-y-4 scale-95"
      >
        <div
          v-if="isMobileMenuOpen"
          class="lg:hidden mt-4 pb-4 border-t border-gray-200/50 dark:border-gray-700/50 pt-4"
        >
          <div class="space-y-3">
            <NuxtLink
              v-for="link in navigationLinks"
              :key="link.to"
              :to="link.to"
              @click="isMobileMenuOpen = false"
              class="block px-4 py-4 rounded-2xl font-medium transition-all duration-500 group/mobile-link"
              :class="{
                'bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 text-white shadow-lg': $route.path === link.to,
                'bg-white/90 dark:bg-gray-800/90 text-gray-700 dark:text-gray-300 hover:bg-gradient-to-r hover:from-blue-50/50 hover:to-purple-50/50 dark:hover:from-blue-900/20 dark:hover:to-purple-900/20': $route.path !== link.to
              }"
            >
              <div class="flex items-center space-x-4">
                <component :is="link.icon" 
                  class="w-5 h-5 transition-all duration-500 group-hover/mobile-link:scale-110 group-hover/mobile-link:opacity-80"
                  :class="{
                    'text-white': $route.path === link.to,
                    'text-blue-500 dark:text-blue-400': $route.path !== link.to
                  }"
                />
                <span class="font-bold text-base transition-all duration-500 group-hover/mobile-link:scale-105">
                  {{ link.name }}
                </span>
              </div>
            </NuxtLink>
            
            <!-- Mobile Theme -->
            <button
              @click="toggleTheme"
              class="w-full px-4 py-4 rounded-2xl bg-white/90 dark:bg-gray-800/90 text-gray-700 dark:text-gray-300 flex items-center justify-center space-x-3 hover:scale-[1.02] transition-all duration-500 font-bold group/mobile-theme"
            >
              <svg v-if="colorMode === 'light'" 
                   class="w-5 h-5 text-yellow-500 transition-all duration-500 group-hover/mobile-theme:scale-110 group-hover/mobile-theme:opacity-80" 
                   fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.25a.75.75 0 01.75.75v2.25a.75.75 0 01-1.5 0V3a.75.75 0 01.75-.75zM7.5 12a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM18.894 6.166a.75.75 0 00-1.06-1.06l-1.591 1.59a.75.75 0 101.06 1.061l1.591-1.59zM21.75 12a.75.75 0 01-.75.75h-2.25a.75.75 0 010-1.5H21a.75.75 0 01.75.75zM17.834 18.894a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 10-1.061 1.06l1.59 1.591zM12 18a.75.75 0 01.75.75V21a.75.75 0 01-1.5 0v-2.25A.75.75 0 0112 18zM7.758 17.303a.75.75 0 00-1.061-1.06l-1.591 1.59a.75.75 0 001.06 1.061l1.591-1.59zM6 12a.75.75 0 01-.75.75H3a.75.75 0 010-1.5h2.25A.75.75 0 016 12zM6.697 7.757a.75.75 0 001.06-1.06l-1.59-1.591a.75.75 0 00-1.061 1.06l1.59 1.591z"/>
              </svg>
              <svg v-else 
                   class="w-5 h-5 text-blue-300 transition-all duration-500 group-hover/mobile-theme:scale-110 group-hover/mobile-theme:opacity-80" 
                   fill="currentColor" viewBox="0 0 24 24">
                <path fill-rule="evenodd" d="M9.528 1.718a.75.75 0 01.162.819A8.97 8.97 0 009 6a9 9 0 009 9 8.97 8.97 0 003.463-.69.75.75 0 01.981.98 10.503 10.503 0 01-9.694 6.46c-5.799 0-10.5-4.701-10.5-10.5 0-4.368 2.667-8.112 6.46-9.694a.75.75 0 01.818.162z" clip-rule="evenodd"/>
              </svg>
              <span class="transition-all duration-500 group-hover/mobile-theme:scale-105">
                {{ colorMode === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode' }}
              </span>
            </button>
            
            <!-- Mobile Auth -->
            <div class="pt-2">
              <ClientOnly>
                <template v-if="auth.isLoggedIn.value">
                  <div class="px-4 py-4 rounded-2xl bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 text-white shadow-lg group/mobile-auth">
                    <p class="font-black text-base transition-all duration-500 group-hover/mobile-auth:scale-105">
                      {{ userName }}
                    </p>
                    <p class="text-sm opacity-90 mt-1">{{ userEmail }}</p>
                    <button
                      @click="handleLogout"
                      class="mt-4 w-full py-3 rounded-xl bg-white/30 flex items-center justify-center space-x-2 hover:bg-white/40 transition-all duration-500 font-bold group/mobile-logout"
                    >
                      <svg class="w-4 h-4 transition-all duration-500 group-hover/mobile-logout:scale-110 group-hover/mobile-logout:opacity-80" 
                           fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
                      </svg>
                      <span class="transition-all duration-500 group-hover/mobile-logout:scale-105">Logout</span>
                    </button>
                  </div>
                </template>
                <template v-else>
                  <NuxtLink
                    to="/auth"
                    @click="isMobileMenuOpen = false"
                    class="block px-4 py-4 rounded-2xl bg-gradient-to-r from-blue-500 via-purple-500 to-cyan-500 text-white text-center font-black shadow-lg hover:scale-[1.02] transition-all duration-500 group/mobile-login"
                  >
                    <span class="transition-all duration-500 group-hover/mobile-login:scale-105">Login</span>
                  </NuxtLink>
                </template>
              </ClientOnly>
            </div>
          </div>
        </div>
      </Transition>
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

// Icon Components - tanpa rotasi
const HomeIcon = () => h('svg', {
  class: 'w-4 h-4',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2.5',
    'd': 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6'
  })
])

const TrendingUpIcon = () => h('svg', {
  class: 'w-4 h-4',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2.5',
    'd': 'M13 7h8m0 0v8m0-8l-8 8-4-4-6 6'
  })
])

const InformationCircleIcon = () => h('svg', {
  class: 'w-4 h-4',
  fill: 'none',
  stroke: 'currentColor',
  viewBox: '0 0 24 24'
}, [
  h('path', {
    'stroke-linecap': 'round',
    'stroke-linejoin': 'round',
    'stroke-width': '2.5',
    'd': 'M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z'
  })
])

// Navigation Data
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
/* Enhanced Animations */
@keyframes float-particle {
  0%, 100% {
    transform: translateY(0) translateX(0);
    opacity: 0.3;
  }
  50% {
    transform: translateY(-15px) translateX(10px);
    opacity: 0.8;
  }
}

@keyframes logo-glow {
  0%, 100% {
    box-shadow: 
      0 0 20px rgba(59, 130, 246, 0.4),
      0 0 30px rgba(139, 92, 246, 0.3),
      0 0 40px rgba(6, 182, 212, 0.2);
  }
  50% {
    box-shadow: 
      0 0 30px rgba(59, 130, 246, 0.6),
      0 0 40px rgba(139, 92, 246, 0.4),
      0 0 50px rgba(6, 182, 212, 0.3);
  }
}

@keyframes ping-slow {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  70% {
    transform: scale(1.1);
    opacity: 0.7;
  }
  100% {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes gradient-border {
  0%, 100% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
}

@keyframes pulse-slow {
  0%, 100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.8;
    transform: scale(0.95);
  }
}

@keyframes zoom-fade {
  0% {
    transform: scale(1);
    opacity: 1;
  }
  100% {
    transform: scale(1.15);
    opacity: 0.9;
  }
}

/* Utility Classes */
.animate-float-particle {
  animation: float-particle 10s ease-in-out infinite;
}

.animate-logo-glow {
  animation: logo-glow 3s ease-in-out infinite;
}

.animate-ping-slow {
  animation: ping-slow 2.5s ease-in-out infinite;
}

.animate-gradient-border {
  background-size: 200% 200%;
  animation: gradient-border 3s ease infinite;
}

.animate-pulse-slow {
  animation: pulse-slow 2s ease-in-out infinite;
}

.animate-zoom-fade {
  animation: zoom-fade 0.3s ease-out forwards;
}

.bg-grid-pattern {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.05) 1px, transparent 1px);
  background-size: 30px 30px;
}

/* Enhanced Backdrop Blur */
.backdrop-blur-2xl {
  backdrop-filter: blur(24px) saturate(200%);
  -webkit-backdrop-filter: blur(24px) saturate(200%);
}

/* Smooth Transitions */
* {
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Gradient Text Shine */
.gradient-text-shine {
  background: linear-gradient(
    90deg,
    #3b82f6 0%,
    #8b5cf6 25%,
    #06b6d4 50%,
    #8b5cf6 75%,
    #3b82f6 100%
  );
  background-size: 200% auto;
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: shine 3s linear infinite;
}

@keyframes shine {
  to {
    background-position: 200% center;
  }
}

/* Glass Effect */
.glass-effect {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dark .glass-effect {
  background: rgba(17, 24, 39, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

/* Enhanced Hover Effects */
.hover-zoom {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.hover-zoom:hover {
  transform: scale(1.1);
  opacity: 0.9;
}

.hover-fade {
  transition: opacity 0.3s ease;
}

.hover-fade:hover {
  opacity: 0.8;
}

/* Gradient Shadow */
.gradient-shadow {
  position: relative;
}

.gradient-shadow::after {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(45deg, #3b82f6, #8b5cf6, #06b6d4);
  border-radius: inherit;
  z-index: -1;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.gradient-shadow:hover::after {
  opacity: 0.3;
}

/* Scale and Fade Animation */
.scale-fade-enter-active,
.scale-fade-leave-active {
  transition: all 0.3s ease;
}

.scale-fade-enter-from,
.scale-fade-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

/* Active Indicator Centering Fix */
.active-indicator-fix {
  position: absolute;
  bottom: -0.5rem;
  left: 0;
  right: 0;
  margin: 0 auto;
  width: 2rem;
}
</style>