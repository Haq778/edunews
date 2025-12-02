<template>
  <div class="p-4 sm:p-6 max-w-8xl mx-auto space-y-4 sm:space-y-6">

    <!-- Header dengan Glass Effect - Responsif -->
    <div class="glass-card rounded-xl sm:rounded-2xl p-4 sm:p-6 border-l-4 border-l-purple-500">
      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <div class="flex-1 min-w-0">
          <div class="flex flex-col sm:flex-row sm:items-center sm:space-x-3 space-y-2 sm:space-y-0 mb-3">
            <div class="p-2 bg-gradient-to-br from-purple-500/10 to-pink-600/10 rounded-xl self-start sm:self-auto">
              <svg class="w-5 h-5 sm:w-6 sm:h-6 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
              </svg>
            </div>
            <div class="flex-1 min-w-0">
              <h1 class="text-xl sm:text-2xl lg:text-3xl font-bold bg-gradient-to-r from-gray-900 to-purple-900 dark:from-white dark:to-purple-200 bg-clip-text text-transparent leading-tight">
                Manajemen Users
              </h1>
              <p class="text-gray-600 dark:text-gray-300 mt-1 text-xs sm:text-sm truncate">Kelola pengguna dan akses sistem dengan mudah</p>
            </div>
          </div>
          
          <!-- Quick Stats - Responsif -->
          <div class="flex flex-wrap items-center gap-2 sm:gap-4 text-xs sm:text-sm text-gray-500 dark:text-gray-400">
            <div class="flex items-center space-x-1">
              <svg class="w-3 h-3 sm:w-4 sm:h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
              <span>{{ users.length }} Total Users</span>
            </div>
            <div class="flex items-center space-x-1">
              <svg class="w-3 h-3 sm:w-4 sm:h-4 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"></path>
              </svg>
              <span>{{ adminCount }} Admin</span>
            </div>
            <div class="flex items-center space-x-1">
              <svg class="w-3 h-3 sm:w-4 sm:h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
              <span>{{ editorCount }} Editor</span>
            </div>
            
          </div>
        </div>

        <div class="flex flex-col sm:flex-row sm:space-x-3 space-y-2 sm:space-y-0 w-full sm:w-auto">
          <button 
            @click="exportUsers"
            class="group inline-flex items-center justify-center sm:justify-start space-x-2 px-3 py-2 sm:px-4 sm:py-2.5 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg sm:rounded-xl hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-xs sm:text-sm font-medium text-gray-700 dark:text-gray-300 w-full sm:w-auto"
          >
            <svg class="w-3 h-3 sm:w-4 sm:h-4 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
            </svg>
            <span>Export</span>
          </button>
          
          
        </div>
      </div>
    </div>

    <!-- Mobile Search and Filter -->
    <div v-if="users.length > 0" class="lg:hidden glass-card rounded-xl p-3 space-y-3">
      <div class="relative">
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Cari user..." 
          class="w-full pl-9 pr-4 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all duration-300 text-sm"
        >
        <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-3.5 h-3.5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
        <button 
          v-if="searchQuery" 
          @click="searchQuery = ''"
          class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
        >
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
      
      <div class="flex space-x-2">
        <button
          @click="isFilterOpen = !isFilterOpen"
          class="flex-1 inline-flex items-center justify-center space-x-1 px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-xs font-medium text-gray-700 dark:text-gray-300"
        >
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 00-.293-.707L3.293 7.293A1 1 0 013 6.586V4z"></path>
          </svg>
          <span>Filter</span>
        </button>
        
        <!-- Mobile View Toggle -->
        <button
          @click="toggleViewMode"
          class="inline-flex items-center justify-center px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-xs font-medium text-gray-700 dark:text-gray-300"
        >
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path v-if="viewMode === 'grid'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
            <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Loading Animation - Responsif -->
    <div v-if="loading" class="glass-card rounded-xl sm:rounded-2xl p-6 sm:p-8 text-center">
      <div class="flex flex-col items-center justify-center space-y-3 sm:space-y-4">
        <div class="relative">
          <div class="w-12 h-12 sm:w-16 sm:h-16 border-4 border-purple-200 dark:border-purple-800 rounded-full"></div>
          <div class="absolute top-0 left-0 w-12 h-12 sm:w-16 sm:h-16 border-4 border-transparent border-t-purple-600 rounded-full animate-spin"></div>
          <div class="absolute top-1 left-1 w-10 h-10 sm:w-14 sm:h-14 border-4 border-transparent border-b-pink-600 rounded-full animate-spin animation-delay-500"></div>
        </div>
        <div>
          <p class="text-base sm:text-lg font-semibold text-gray-700 dark:text-gray-300 mb-1">Memuat Data User</p>
          <p class="text-xs sm:text-sm text-gray-500 dark:text-gray-400">Menyiapkan informasi pengguna...</p>
        </div>
      </div>
    </div>

    <!-- Content -->
    <div v-else class="space-y-3 sm:space-y-4">
      <!-- Search and Filter Desktop -->
      <div class="hidden lg:block glass-card rounded-xl p-4">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between space-y-3 lg:space-y-0">
          <div class="flex-1 max-w-md">
            <div class="relative">
              <input 
                v-model="searchQuery"
                type="text" 
                placeholder="Cari user berdasarkan nama atau email..." 
                class="w-full pl-10 pr-4 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all duration-300 text-sm"
              >
              <svg class="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              <button 
                v-if="searchQuery" 
                @click="searchQuery = ''"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
          <div class="flex space-x-2">
            <select 
              v-model="roleFilter"
              class="px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-sm focus:ring-2 focus:ring-purple-500 focus:border-transparent"
            >
              <option value="">Semua Role</option>
              <option value="admin">Admin</option>
              <option value="editor">Editor</option>
              <option value="user">User</option>
            </select>
            <select 
              v-model="sortBy"
              class="px-3 py-2 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-white dark:hover:bg-gray-800 transition-all duration-300 text-sm focus:ring-2 focus:ring-purple-500 focus:border-transparent"
            >
              <option value="newest">Terbaru</option>
              <option value="oldest">Terlama</option>
              <option value="name">Nama A-Z</option>
            </select>
          </div>
        </div>
        
        <!-- Search Results Info -->
        <div v-if="searchQuery || roleFilter" class="mt-3 flex items-center justify-between text-xs">
          <div class="text-gray-600 dark:text-gray-400">
            Menampilkan {{ filteredUsers.length }} dari {{ users.length }} user
            <span v-if="searchQuery"> untuk "<span class="font-semibold">{{ searchQuery }}</span>"</span>
            <span v-if="roleFilter"> dengan role "<span class="font-semibold">{{ roleFilter }}</span>"</span>
          </div>
          <button 
            v-if="searchQuery || roleFilter"
            @click="clearFilters"
            class="text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300 transition-colors flex items-center space-x-1"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Reset Filter</span>
          </button>
        </div>
      </div>

      <!-- Mobile Filter Panel -->
      <Transition
        enter-active-class="transition-all duration-300 ease-out"
        enter-from-class="opacity-0 transform -translate-y-2"
        enter-to-class="opacity-100 transform translate-y-0"
        leave-active-class="transition-all duration-300 ease-in"
        leave-from-class="opacity-100 transform translate-y-0"
        leave-to-class="opacity-0 transform -translate-y-2"
      >
        <div 
          v-if="isFilterOpen"
          class="lg:hidden glass-card rounded-xl p-4 mb-3 space-y-3"
        >
          <div class="grid grid-cols-2 gap-2">
            <div>
              <label class="text-xs font-medium text-gray-700 dark:text-gray-300 mb-1 block">Role</label>
              <select 
                v-model="roleFilter"
                class="w-full px-2 py-1.5 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg text-xs"
              >
                <option value="">Semua</option>
                <option value="admin">Admin</option>
                <option value="editor">Editor</option>
                <option value="user">User</option>
              </select>
            </div>
            <div>
              <label class="text-xs font-medium text-gray-700 dark:text-gray-300 mb-1 block">Urutkan</label>
              <select 
                v-model="sortBy"
                class="w-full px-2 py-1.5 bg-white/50 dark:bg-gray-800/50 border border-gray-200 dark:border-gray-700 rounded-lg text-xs"
              >
                <option value="newest">Terbaru</option>
                <option value="oldest">Terlama</option>
                <option value="name">Nama A-Z</option>
              </select>
            </div>
          </div>
          
          <!-- Mobile Search Results Info -->
          <div v-if="searchQuery || roleFilter" class="pt-2 border-t border-gray-200 dark:border-gray-700">
            <div class="text-xs text-gray-600 dark:text-gray-400">
              <div class="mb-1">
                Menampilkan {{ filteredUsers.length }} dari {{ users.length }} user
              </div>
              <button 
                v-if="searchQuery || roleFilter"
                @click="clearFilters"
                class="w-full inline-flex items-center justify-center space-x-1 text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300 transition-colors"
              >
                <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
                <span>Reset Filter</span>
              </button>
            </div>
          </div>
        </div>
      </Transition>

      <!-- No Search Results - Responsif -->
      <div v-if="!loading && users.length > 0 && filteredUsers.length === 0" class="glass-card rounded-xl sm:rounded-2xl p-6 sm:p-8 text-center">
        <div class="max-w-sm mx-auto">
          <div class="w-12 h-12 sm:w-16 sm:h-16 mx-auto bg-gradient-to-br from-gray-100 to-gray-200 dark:from-gray-800 dark:to-gray-700 rounded-full flex items-center justify-center mb-3 sm:mb-4">
            <svg class="w-6 h-6 sm:w-8 sm:h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="text-base sm:text-lg font-semibold text-gray-900 dark:text-white mb-2">User Tidak Ditemukan</h3>
          <p class="text-gray-600 dark:text-gray-300 text-xs sm:text-sm mb-4">
            Tidak ada user yang cocok dengan pencarian Anda
          </p>
          <button 
            @click="clearFilters"
            class="inline-flex items-center justify-center space-x-2 bg-gradient-to-r from-purple-600 to-pink-600 text-white px-4 py-2 rounded-lg hover:shadow-lg transition-all duration-300 transform font-semibold text-xs sm:text-sm w-full sm:w-auto"
          >
            <svg class="w-3 h-3 sm:w-4 sm:h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
            <span>Tampilkan Semua User</span>
          </button>
        </div>
      </div>

      <!-- Empty State - Responsif -->
      <div v-else-if="users.length === 0" class="glass-card rounded-xl sm:rounded-2xl p-6 sm:p-8 text-center">
        <div class="max-w-sm mx-auto">
          <div class="relative mb-4 sm:mb-6">
            <div class="w-16 h-16 sm:w-24 sm:h-24 mx-auto bg-gradient-to-br from-purple-100 to-pink-100 dark:from-purple-900/20 dark:to-pink-900/20 rounded-full flex items-center justify-center">
              <svg class="w-8 h-8 sm:w-12 sm:h-12 text-purple-500 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"></path>
              </svg>
            </div>
          </div>
          <h3 class="text-lg sm:text-xl font-bold text-gray-900 dark:text-white mb-2 sm:mb-3">Belum Ada User</h3>
          <p class="text-gray-600 dark:text-gray-300 text-xs sm:text-sm mb-4 sm:mb-6">Mulai dengan menambahkan user pertama ke sistem</p>
          
        </div>
      </div>

      <!-- Users Grid - Responsif dengan View Mode -->
      <div :class="[
        'gap-3 sm:gap-4',
        viewMode === 'grid' 
          ? 'grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4'
          : 'flex flex-col space-y-3'
      ]">
        <div 
          v-for="user in filteredUsers" 
          :key="user.id"
          :class="[
            'glass-card rounded-xl p-3 sm:p-4 hover:shadow-lg transition-all duration-300 group hover:scale-102 transform border border-gray-200/50 dark:border-gray-700/50',
            viewMode === 'list' ? 'flex flex-col sm:flex-row sm:items-center sm:justify-between' : ''
          ]"
        >
          <!-- User Header -->
          <div :class="viewMode === 'list' ? 'flex-1 min-w-0 pr-4' : ''">
            <div class="flex items-start justify-between mb-3">
              <div class="flex items-center space-x-3">
                <div 
                  class="w-10 h-10 sm:w-12 sm:h-12 rounded-lg sm:rounded-xl flex items-center justify-center text-white font-bold text-base sm:text-lg shadow-lg"
                  :class="getUserAvatarColor(user.id)"
                >
                  {{ user.name?.charAt(0)?.toUpperCase() || 'U' }}
                </div>
                <div class="flex-1 min-w-0">
                  <h3 class="font-semibold text-gray-900 dark:text-white text-sm truncate group-hover:text-purple-600 dark:group-hover:text-purple-400 transition-colors duration-300">
                    {{ user.name }}
                  </h3>
                  <p class="text-xs text-gray-500 dark:text-gray-400 mt-1 truncate">{{ user.email }}</p>
                </div>
              </div>
              <!-- Action Buttons untuk Desktop Grid View -->
              <div v-if="viewMode === 'grid'" class="flex space-x-1 sm:hidden lg:flex">
                <NuxtLink
                  :to="`/admin/users/edit/${user.id}`"
                  class="p-1.5 bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 rounded-lg hover:bg-blue-100 dark:hover:bg-blue-900/30 transition-all duration-300 hover:scale-110 transform group/edit"
                  title="Edit User"
                >
                  <svg class="w-3 h-3 sm:w-3.5 sm:h-3.5 group-hover/edit:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                  </svg>
                </NuxtLink>
                <button
                  class="p-1.5 bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400 rounded-lg hover:bg-red-100 dark:hover:bg-red-900/30 transition-all duration-300 hover:scale-110 transform group/delete"
                  @click="hapus(user.id)"
                  title="Hapus User"
                >
                  <svg class="w-3 h-3 sm:w-3.5 sm:h-3.5 group-hover/delete:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- User Details -->
            <div class="space-y-2 mb-3">
              <div class="flex items-center space-x-2 text-xs">
                <span 
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium border"
                  :class="getRoleBadgeClass(user.role)"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-1" :class="getRoleDotClass(user.role)"></span>
                  {{ user.role || 'user' }}
                </span>
                <span class="text-gray-500 dark:text-gray-400 text-xs">
                  {{ formatDate(user.created_at) }}
                </span>
              </div>
            </div>
          </div>

          <!-- Action Buttons untuk Mobile Grid dan List View -->
          <div :class="[
            'flex space-x-2 pt-3 border-t border-gray-200 dark:border-gray-700',
            viewMode === 'list' ? 'sm:border-0 sm:pt-0' : ''
          ]">
            <NuxtLink
              :to="`/admin/users/edit/${user.id}`"
              :class="[
                'flex-1 inline-flex items-center justify-center space-x-1 px-2 py-1.5 text-xs bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 rounded-lg hover:bg-blue-100 dark:hover:bg-blue-900/30 transition-all duration-300 font-medium',
                viewMode === 'list' ? 'sm:w-auto sm:px-3' : ''
              ]"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
              </svg>
              <span class="hidden sm:inline">Edit</span>
              <span class="sm:hidden">Edit</span>
            </NuxtLink>
            <button
              @click="hapus(user.id)"
              :class="[
                'flex-1 inline-flex items-center justify-center space-x-1 px-2 py-1.5 text-xs bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400 rounded-lg hover:bg-red-100 dark:hover:bg-red-900/30 transition-all duration-300 font-medium',
                viewMode === 'list' ? 'sm:w-auto sm:px-3' : ''
              ]"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
              </svg>
              <span class="hidden sm:inline">Hapus</span>
              <span class="sm:hidden">Hapus</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Quick Actions Footer - Responsif -->
      <div v-if="users.length > 0" class="glass-card rounded-xl p-3 sm:p-4 mt-4 sm:mt-6">
        <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between space-y-3 sm:space-y-0">
          <div class="text-xs sm:text-sm text-gray-600 dark:text-gray-400">
            Total <span class="font-semibold text-purple-600 dark:text-purple-400">{{ users.length }}</span> user • 
            <span class="font-semibold text-blue-600 dark:text-blue-400">{{ adminCount }}</span> admin • 
            <span class="font-semibold text-green-600 dark:text-green-400">{{ editorCount }}</span> editor
          </div>
          <div class="flex flex-wrap gap-2">
            <!-- View Toggle Desktop -->
            <button 
              v-if="!isMobile"
              @click="toggleViewMode"
              class="hidden lg:flex items-center gap-1 px-2 py-1.5 text-xs bg-gray-100 dark:bg-gray-800 text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-200 dark:hover:bg-gray-700 transition-all duration-300 border border-gray-200 dark:border-gray-700"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path v-if="viewMode === 'grid'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 10h16M4 14h16M4 18h16" />
                <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
              </svg>
              <span>{{ viewMode === 'grid' ? 'List View' : 'Grid View' }}</span>
            </button>
            
            <button 
              @click="exportUsers"
              class="inline-flex items-center gap-1 px-2 py-1.5 text-xs bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400 rounded-md hover:bg-blue-100 dark:hover:bg-blue-900/30 transition-all duration-300 border border-blue-200 dark:border-blue-800"
            >
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"></path>
              </svg>
              <span>Export</span>
            </button>
            
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, onUnmounted } from 'vue'

const users = ref<any[]>([])
const loading = ref(true)
const searchQuery = ref('')
const roleFilter = ref('')
const sortBy = ref('newest')
const viewMode = ref('grid') // 'grid' atau 'list'
const isFilterOpen = ref(false)
const isMobile = ref(false)

// Deteksi ukuran layar
const checkScreenSize = () => {
  isMobile.value = window.innerWidth < 1024 // lg breakpoint
  if (isMobile.value) {
    viewMode.value = 'grid' // Default grid di mobile
  }
}

const toggleViewMode = () => {
  viewMode.value = viewMode.value === 'grid' ? 'list' : 'grid'
}

definePageMeta({ layout: 'admin' })

// Computed properties
const filteredUsers = computed(() => {
  let filtered = users.value

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(user => 
      user.name?.toLowerCase().includes(query) ||
      user.email?.toLowerCase().includes(query)
    )
  }

  // Filter by role
  if (roleFilter.value) {
    filtered = filtered.filter(user => 
      (user.role || 'user') === roleFilter.value
    )
  }

  // Sort results
  switch (sortBy.value) {
    case 'newest':
      filtered = [...filtered].sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
      break
    case 'oldest':
      filtered = [...filtered].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime())
      break
    case 'name':
      filtered = [...filtered].sort((a, b) => a.name?.localeCompare(b.name))
      break
  }

  return filtered
})

const adminCount = computed(() => {
  return users.value.filter(user => user.role === 'admin').length
})

const editorCount = computed(() => {
  return users.value.filter(user => user.role === 'editor').length
})

// Helper functions
const getUserAvatarColor = (id: number) => {
  const colors = [
    'bg-gradient-to-br from-purple-500 to-pink-500',
    'bg-gradient-to-br from-blue-500 to-cyan-500',
    'bg-gradient-to-br from-green-500 to-emerald-500',
    'bg-gradient-to-br from-orange-500 to-red-500',
    'bg-gradient-to-br from-indigo-500 to-purple-500',
    'bg-gradient-to-br from-teal-500 to-blue-500',
    'bg-gradient-to-br from-pink-500 to-rose-500',
    'bg-gradient-to-br from-yellow-500 to-orange-500'
  ]
  return colors[id % colors.length]
}

const getRoleBadgeClass = (role: string) => {
  const baseClass = 'border '
  switch (role) {
    case 'admin':
      return baseClass + 'bg-purple-50 dark:bg-purple-900/20 text-purple-700 dark:text-purple-300 border-purple-200 dark:border-purple-800'
    case 'editor':
      return baseClass + 'bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 border-blue-200 dark:border-blue-800'
    default:
      return baseClass + 'bg-gray-50 dark:bg-gray-900/20 text-gray-700 dark:text-gray-300 border-gray-200 dark:border-gray-700'
  }
}

const getRoleDotClass = (role: string) => {
  switch (role) {
    case 'admin':
      return 'bg-purple-500'
    case 'editor':
      return 'bg-blue-500'
    default:
      return 'bg-gray-500'
  }
}

const formatDate = (d: string | null | undefined) => {
  if (!d) return '-'
  try {
    return new Date(d).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' })
  } catch {
    return d
  }
}

const clearFilters = () => {
  searchQuery.value = ''
  roleFilter.value = ''
  sortBy.value = 'newest'
  isFilterOpen.value = false
}

const exportUsers = () => {
  const data = JSON.stringify(users.value, null, 2)
  const blob = new Blob([data], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = 'users-data.json'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}

// Load semua user
async function loadUsers() {
  loading.value = true
  try {
    const res: any = await $fetch('http://localhost:8080/api/users')
    users.value = Array.isArray(res) ? res : []
  } catch (e: any) {
    alert('Gagal memuat users: ' + (e?.message || e))
  } finally {
    loading.value = false
  }
}

// Hapus user
async function hapus(id: number | string) {
  if (!confirm('Yakin ingin menghapus user ini?')) return

  try {
    const res = await fetch(`http://localhost:8080/api/users/${id}`, { method: 'DELETE' })
    if (!res.ok) {
      const errJSON = await res.json().catch(() => null)
      throw new Error(errJSON?.error || errJSON?.message || `Hapus gagal (HTTP ${res.status})`)
    }
    users.value = users.value.filter(u => String(u.id) !== String(id))
    alert('User berhasil dihapus.')
  } catch (e: any) {
    alert('Gagal menghapus user: ' + (e?.message || e))
  }
}

onMounted(() => {
  checkScreenSize()
  window.addEventListener('resize', checkScreenSize)
  loadUsers()
})

onUnmounted(() => {
  window.removeEventListener('resize', checkScreenSize)
})
</script>

<style scoped>
.glass-card {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.dark .glass-card {
  background: rgba(17, 24, 39, 0.7);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.animation-delay-500 {
  animation-delay: 0.5s;
}

.hover-scale-102 {
  transform: scale(1);
  transition: transform 0.3s ease;
}

.hover-scale-102:hover {
  transform: scale(1.02);
}

/* Responsive touch targets */
@media (max-width: 640px) {
  button, 
  a[role="button"],
  .cursor-pointer {
    min-height: 44px;
    min-width: 44px;
  }
}
</style>