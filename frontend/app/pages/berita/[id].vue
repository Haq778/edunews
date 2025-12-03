<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Animated Background Elements -->
    <div class="absolute inset-0 z-0 overflow-hidden">
      <!-- Floating orbs/particles -->
      <div class="absolute top-1/4 left-10 w-72 h-72 bg-gradient-to-r from-cyan-500/10 to-blue-500/10 rounded-full blur-3xl animate-float-slow"></div>
      <div class="absolute bottom-1/4 right-10 w-96 h-96 bg-gradient-to-r from-purple-500/10 to-pink-500/10 rounded-full blur-3xl animate-float-medium"></div>
      <div class="absolute top-3/4 left-1/4 w-64 h-64 bg-gradient-to-r from-orange-500/10 to-yellow-500/10 rounded-full blur-3xl animate-float-fast"></div>
      
      <!-- Animated grid pattern -->
      <div class="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,0.05)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.05)_1px,transparent_1px)] bg-[size:40px_40px] [mask-image:radular-gradient(ellipse_at_center,black_30%,transparent_70%)]"></div>
    </div>

    <!-- Navigation Bar -->
    <div class="relative z-50">
      <div class="max-w-7xl mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <!-- Back to Home Button -->
          <button 
            @click="goToHome"
            class="group flex items-center space-x-3 px-5 py-3 bg-gradient-to-r from-white/90 to-white/80 dark:from-gray-800/90 dark:to-gray-900/80 backdrop-blur-xl rounded-2xl shadow-xl hover:shadow-2xl border border-white/20 dark:border-gray-700/20 transition-all duration-300 transform hover:-translate-y-0.5"
          >
            <svg class="w-5 h-5 text-gray-700 dark:text-gray-300 transform group-hover:-translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
            <span class="font-bold text-gray-700 dark:text-gray-300 group-hover:text-blue-600 dark:group-hover:text-cyan-400 transition-colors">
              Kembali ke Beranda
            </span>
          </button>

          <!-- Site Logo/Name -->
          <div class="text-center">
            <div class="text-2xl font-black bg-gradient-to-r from-blue-600 via-cyan-500 to-teal-500 bg-clip-text text-transparent">
              NEWS<span class="text-blue-600">.</span>ID
            </div>
            <div class="text-xs text-gray-500 dark:text-gray-400 mt-1">Portal Berita Terkini</div>
          </div>

          <!-- Reading Time -->
          <div class="flex items-center space-x-2 px-4 py-3 bg-gradient-to-r from-cyan-600 to-blue-600 rounded-2xl shadow-lg">
            <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="text-sm font-semibold text-white">{{ readingTime }} menit baca</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Breaking News Banner -->
    <div v-if="isBreakingNews" class="relative z-10 bg-gradient-to-r from-orange-500 via-orange-400 to-yellow-500 text-white py-3 px-4 animate-pulse">
      <div class="max-w-6xl mx-auto flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <div class="relative">
            <div class="w-3 h-3 bg-white rounded-full animate-ping"></div>
            <div class="w-3 h-3 bg-white rounded-full absolute top-0"></div>
          </div>
          <span class="font-bold uppercase text-sm tracking-wider animate-bounce">🚨 BERITA TERKINI</span>
        </div>
        <div class="text-sm font-medium flex items-center space-x-2">
          <span class="animate-pulse">•</span>
          <span>LIVE UPDATES</span>
          <span class="animate-pulse">•</span>
        </div>
      </div>
    </div>

    <!-- HERO SECTION -->
    <div class="relative z-10">
      <!-- Title Card -->
      <div class="max-w-6xl mx-auto px-4 py-8">
        <div class="bg-gradient-to-br from-white/95 to-white/90 dark:from-gray-800/95 dark:to-gray-900/95 backdrop-blur-xl rounded-3xl shadow-2xl p-8 md:p-12 border border-white/20 dark:border-gray-700/20 transform transition-all duration-300 hover:shadow-3xl">
          <div class="flex flex-col md:flex-row gap-8">
            <!-- Image Card -->
            <div class="md:w-2/5">
              <div class="relative group">
                <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500 via-blue-500 to-purple-500 rounded-2xl blur opacity-30 group-hover:opacity-70 transition duration-1000 group-hover:duration-200 animate-gradient-x"></div>
                <div class="relative bg-white dark:bg-gray-800 rounded-2xl overflow-hidden shadow-xl">
                  <img
                    v-if="berita?.cover"
                    :src="coverUrl"
                    class="w-full h-[300px] md:h-[350px] object-cover transition-transform duration-500 group-hover:scale-105"
                    alt="Cover Berita"
                  />
                  <div v-else class="w-full h-[300px] md:h-[350px] bg-gradient-to-br from-blue-900 to-purple-900 flex items-center justify-center">
                    <svg class="w-20 h-20 text-white/50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <!-- Image Badge -->
                  <div class="absolute top-4 left-4">
                    <span class="px-4 py-2 text-sm font-bold bg-gradient-to-r from-blue-600 to-cyan-500 text-white rounded-full shadow-lg">
                      {{ berita?.category?.name?.toUpperCase() || "UTAMA" }}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Content -->
            <div class="md:w-3/5">
              <div class="space-y-6">
                <div>
                  <h1 class="text-3xl md:text-4xl lg:text-5xl font-black text-gray-900 dark:text-white leading-tight mb-4">
                    {{ berita?.title }}
                  </h1>
                  
                  <div class="flex flex-wrap items-center gap-4 mb-2">
                    <div class="flex items-center space-x-2 bg-gray-100 dark:bg-gray-800 px-3 py-1.5 rounded-full">
                      <svg class="w-4 h-4 text-cyan-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                      </svg>
                      <span class="font-medium text-sm text-gray-700 dark:text-gray-300">{{ formattedDate }}</span>
                    </div>
                    <div class="flex items-center space-x-2 bg-gray-100 dark:bg-gray-800 px-3 py-1.5 rounded-full">
                      <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                      <span class="font-medium text-sm text-gray-700 dark:text-gray-300">{{ berita?.views }}x dibaca</span>
                    </div>
                    <div class="flex items-center space-x-2 bg-gray-100 dark:bg-gray-800 px-3 py-1.5 rounded-full">
                      <svg class="w-4 h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                      </svg>
                      <span class="font-medium text-sm text-gray-700 dark:text-gray-300">Reporter: {{ berita?.author || "Admin" }}</span>
                    </div>
                  </div>
                </div>

                <!-- Lead Paragraph -->
                <div class="bg-gradient-to-r from-gray-50 to-white dark:from-gray-800 dark:to-gray-900 p-6 rounded-xl border-l-4 border-blue-500">
                  <p class="text-lg md:text-xl text-gray-700 dark:text-gray-200 font-medium leading-relaxed">
                    {{ berita?.excerpt || "Ringkasan berita akan ditampilkan di sini." }}
                  </p>
                </div>

                <!-- Quick Actions -->
                <div class="flex flex-wrap gap-3">
                  <button @click="copyLink" class="flex items-center space-x-2 px-5 py-3 bg-gradient-to-r from-blue-600 to-cyan-500 text-white rounded-xl hover:from-blue-700 hover:to-cyan-600 transition-all duration-300 transform hover:-translate-y-0.5 shadow-lg hover:shadow-xl">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                    </svg>
                    <span class="font-medium">Bagikan</span>
                  </button>
                  
                  <button class="flex items-center space-x-2 px-5 py-3 bg-gradient-to-r from-gray-800 to-gray-700 dark:from-gray-700 dark:to-gray-600 text-white rounded-xl hover:from-gray-900 hover:to-gray-800 transition-all duration-300 transform hover:-translate-y-0.5 shadow-lg hover:shadow-xl">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                    </svg>
                    <span class="font-medium">Simpan</span>
                  </button>
                  
                  <button class="flex items-center space-x-2 px-5 py-3 bg-gradient-to-r from-green-600 to-emerald-500 text-white rounded-xl hover:from-green-700 hover:to-emerald-600 transition-all duration-300 transform hover:-translate-y-0.5 shadow-lg hover:shadow-xl">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h4.764a2 2 0 011.789 2.894l-3.5 7A2 2 0 0115.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 00-2-2h-.095c-.5 0-.905.405-.905.905 0 .714-.211 1.412-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 01-2-2v-6a2 2 0 012-2h2.5" />
                    </svg>
                    <span class="font-medium">Suka</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- CONTENT SECTION -->
    <div class="relative z-10 max-w-4xl mx-auto px-4 py-8">
      <!-- Floating Social Share (Only Twitter) -->
      <div class="fixed left-4 top-1/2 transform -translate-y-1/2 hidden lg:block z-40">
        <div class="bg-gradient-to-b from-white/90 to-white/80 dark:from-gray-800/90 dark:to-gray-900/80 backdrop-blur-xl rounded-2xl shadow-2xl p-4 space-y-4 border border-white/20 dark:border-gray-700/20">
          <!-- Twitter Share -->
          <button @click="shareTwitter" class="p-3 hover:bg-blue-100/50 dark:hover:bg-gray-700/50 rounded-xl transition-all duration-300 transform hover:-translate-x-1 group">
            <div class="flex items-center justify-center w-8 h-8 bg-gradient-to-br from-cyan-500 to-blue-400 rounded-lg group-hover:scale-110 transition-transform">
              <span class="text-white font-bold text-sm">𝕏</span>
            </div>
          </button>
          
          <!-- Copy Link -->
          <button @click="copyLink" class="p-3 hover:bg-purple-100/50 dark:hover:bg-gray-700/50 rounded-xl transition-all duration-300 transform hover:-translate-x-1 group">
            <div class="flex items-center justify-center w-8 h-8 bg-gradient-to-br from-blue-600 to-cyan-500 rounded-lg group-hover:scale-110 transition-transform">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
              </svg>
            </div>
          </button>
          
          <!-- Back to Top -->
          <button @click="scrollToTop" class="p-3 hover:bg-red-100/50 dark:hover:bg-gray-700/50 rounded-xl transition-all duration-300 transform hover:-translate-x-1 group">
            <div class="flex items-center justify-center w-8 h-8 bg-gradient-to-br from-blue-600 to-cyan-500 rounded-lg group-hover:scale-110 transition-transform">
              <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18" />
              </svg>
            </div>
          </button>
        </div>
      </div>

      <!-- Main Content Card -->
      <div class="relative group">
        <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500/20 via-blue-500/20 to-purple-500/20 rounded-3xl blur-xl opacity-70 group-hover:opacity-100 transition duration-1000"></div>
        <div class="relative bg-white/95 dark:bg-gray-800/95 backdrop-blur-xl rounded-3xl shadow-2xl overflow-hidden border border-white/20 dark:border-gray-700/20">
          <!-- Content Header -->
          <div class="px-8 py-6 bg-gradient-to-r from-gray-50 to-white dark:from-gray-900 dark:to-gray-800 border-b border-gray-100 dark:border-gray-700/50">
            <div class="flex items-center justify-between">
              <div class="text-lg font-bold text-gray-900 dark:text-white">
                📝 Isi Berita
              </div>
              <div class="flex items-center space-x-4">
                <div class="flex items-center space-x-2 text-sm text-gray-600 dark:text-gray-400">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <span>{{ berita?.views || 0 }} pembaca</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Article Content -->
          <div class="px-8 py-10">
            <!-- Custom prose styling untuk dark mode -->
            <article
              class="article-content max-w-none leading-relaxed text-gray-800 dark:text-gray-200"
              v-html="berita?.content"
            ></article>

            <!-- Tags -->
            <div class="mt-10 pt-8 border-t border-gray-200 dark:border-gray-700/50">
              <div class="flex flex-wrap gap-3">
                <span class="px-4 py-2 bg-gradient-to-r from-blue-600 to-cyan-500 text-white rounded-full text-sm font-medium shadow-lg">
                  #{{ berita?.category?.name || "berita" }}
                </span>
                <span v-for="(tag, index) in relatedTags" :key="index" 
                      class="px-4 py-2 bg-gradient-to-r from-blue-600/10 to-cyan-600/10 dark:from-blue-500/20 dark:to-cyan-500/20 text-blue-700 dark:text-cyan-300 rounded-full text-sm font-medium border border-blue-200 dark:border-cyan-500/30">
                  #{{ tag }}
                </span>
              </div>
            </div>

            <!-- Author Box -->
            <div class="mt-10 relative group/author">
              <div class="absolute -inset-1 bg-gradient-to-r from-blue-500/10 to-cyan-500/10 rounded-2xl blur opacity-50 group-hover/author:opacity-100 transition duration-500"></div>
              <div class="relative bg-gradient-to-r from-gray-50 to-white dark:from-gray-900 dark:to-gray-800 rounded-2xl p-6">
                <div class="flex items-center space-x-4">
                  <div class="relative">
                    <div class="absolute -inset-1 bg-gradient-to-r from-blue-500 to-cyan-600 rounded-full blur opacity-60 animate-pulse"></div>
                    <div class="relative w-14 h-14 bg-gradient-to-br from-blue-500 to-cyan-600 rounded-full flex items-center justify-center text-white font-bold text-xl shadow-lg">
                      {{ berita?.author?.charAt(0) || "A" }}
                    </div>
                  </div>
                  <div>
                    <h4 class="font-bold text-gray-900 dark:text-white text-lg">{{ berita?.author || "Admin" }}</h4>
                    <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Reporter Berita</p>
                    <p class="text-sm text-gray-600 dark:text-gray-300 mt-2">Penulis artikel berita dengan pengalaman lebih dari 5 tahun di industri jurnalistik.</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- RELATED NEWS SECTION -->
      <div class="mt-16 mb-20">
        <div class="flex items-center justify-between mb-8">
          <div class="relative">
            <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500/20 to-blue-500/20 rounded-lg blur"></div>
            <h2 class="relative text-2xl md:text-3xl font-black text-gray-900 dark:text-white bg-white dark:bg-gray-900 px-4 py-2 rounded-lg">
              📰 BERITA TERKAIT
            </h2>
          </div>
          <button 
            @click="goToHome"
            class="group flex items-center space-x-2 px-5 py-3 bg-gradient-to-r from-blue-600 to-cyan-500 text-white rounded-xl hover:from-blue-700 hover:to-cyan-600 transition-all duration-300 shadow-lg hover:shadow-xl transform hover:-translate-y-0.5"
          >
            <span class="font-semibold">Lihat Semua Berita</span>
            <svg class="w-4 h-4 transform group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
            </svg>
          </button>
        </div>

        <div class="grid md:grid-cols-3 gap-6">
          <div
            v-for="item in related"
            :key="item.id"
            class="group relative"
            @click="navigateTo(`/berita/${item.id}`)"
          >
            <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500/20 via-blue-500/20 to-purple-500/20 rounded-2xl blur opacity-0 group-hover:opacity-100 transition duration-500"></div>
            <div class="relative bg-white dark:bg-gray-800 rounded-2xl shadow-xl hover:shadow-2xl cursor-pointer overflow-hidden transition-all duration-500 border border-gray-100 dark:border-gray-700 group-hover:-translate-y-2">
              <div class="relative h-48 overflow-hidden">
                <img
                  :src="getRelatedCover(item)"
                  class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110"
                />
                <div class="absolute inset-0 bg-gradient-to-t from-black/50 via-transparent to-transparent"></div>
                <div class="absolute top-4 left-4">
                  <span class="px-3 py-1 bg-gradient-to-r from-blue-600 to-cyan-500 text-white text-xs font-bold rounded-full shadow-lg">
                    TERKAIT
                  </span>
                </div>
                <div class="absolute bottom-4 left-4 right-4">
                  <div class="text-xs text-white/90 font-medium bg-black/30 backdrop-blur-sm px-3 py-1.5 rounded-full inline-block">
                    {{ formatRelatedDate(item.created_at) }}
                  </div>
                </div>
              </div>
              <div class="p-6">
                <h3 class="font-bold text-gray-900 dark:text-white text-lg leading-tight mb-3 group-hover:text-blue-600 dark:group-hover:text-cyan-400 transition-colors line-clamp-2">
                  {{ item.title }}
                </h3>
                <p class="text-gray-600 dark:text-gray-300 text-sm line-clamp-2 mb-4">
                  {{ item.excerpt || "" }}
                </p>
                <div class="flex items-center justify-between">
                  <div class="text-blue-600 dark:text-cyan-400 text-sm font-semibold flex items-center group-hover:translate-x-1 transition-transform">
                    Baca Selengkapnya
                    <svg class="w-4 h-4 ml-1 transform group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
                    </svg>
                  </div>
                  <div class="text-xs text-gray-500 dark:text-gray-400">
                    {{ Math.floor(Math.random() * 1000) + 100 }} views
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Newsletter Subscription -->
    <div class="relative z-10">
      <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500 via-blue-500 to-purple-500 rounded-3xl blur-xl opacity-20 animate-pulse"></div>
      <div class="relative bg-gradient-to-r from-gray-900 via-blue-900 to-purple-900 text-white py-16 px-4 overflow-hidden">
        <!-- Animated dots in background -->
        <div class="absolute inset-0">
          <div class="absolute top-10 left-10 w-2 h-2 bg-white/30 rounded-full animate-ping"></div>
          <div class="absolute top-20 right-20 w-3 h-3 bg-white/30 rounded-full animate-ping" style="animation-delay: 0.2s"></div>
          <div class="absolute bottom-20 left-1/4 w-2 h-2 bg-white/30 rounded-full animate-ping" style="animation-delay: 0.4s"></div>
        </div>
        
        <div class="relative max-w-4xl mx-auto text-center">
          <h3 class="text-3xl font-bold mb-4 animate-bounce-subtle">🚀 Dapatkan Update Berita Terbaru</h3>
          <p class="text-cyan-200/90 mb-8 text-lg">Berlangganan newsletter kami untuk mendapatkan berita terkini langsung ke email Anda</p>
          <div class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto">
            <input 
              type="email" 
              placeholder="Masukkan email Anda" 
              class="flex-grow px-6 py-4 rounded-xl text-gray-900 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-lg transform transition-transform hover:scale-[1.02]"
            >
            <button class="bg-gradient-to-r from-blue-600 to-cyan-500 hover:from-blue-700 hover:to-cyan-600 px-8 py-4 rounded-xl font-bold transition-all duration-300 transform hover:-translate-y-0.5 shadow-lg hover:shadow-xl animate-pulse-subtle">
              BERLANGGANAN
            </button>
          </div>
          <p class="text-gray-400 text-sm mt-4">Kami tidak akan mengirim spam. Berhenti berlangganan kapan saja.</p>
        </div>
      </div>
    </div>

    <!-- Toast Notification -->
    <transition name="slide-up">
      <div v-if="showShareToast" class="fixed bottom-4 right-4 z-50">
        <div class="bg-gradient-to-r from-green-600 to-emerald-500 text-white px-6 py-3 rounded-xl shadow-2xl flex items-center space-x-3 animate-bounce-in">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span class="font-semibold">Link berhasil disalin!</span>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, navigateTo } from "#app";

const route = useRoute();
const berita = ref(null);
const related = ref([]);
const showShareToast = ref(false);

async function fetchBerita() {
  const res = await $fetch(`http://localhost:8080/api/berita/${route.params.id}`);
  berita.value = res;
}

async function fetchRelated() {
  const res = await $fetch(`http://localhost:8080/api/berita`);
  related.value = res
    .filter((x) => x.id !== berita.value.id)
    .slice(0, 3);
}

onMounted(async () => {
  await fetchBerita();
  await fetchRelated();
});

const coverUrl = computed(() => {
  if (!berita.value?.cover) return "";
  return `http://localhost:8080/uploads/berita/${berita.value.cover}`;
});

const formattedDate = computed(() => {
  if (!berita.value?.created_at) return "-";
  return new Date(berita.value.created_at).toLocaleDateString("id-ID", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
});

const isBreakingNews = computed(() => {
  if (!berita.value?.created_at) return false;
  const newsDate = new Date(berita.value.created_at);
  const now = new Date();
  const diffInHours = Math.floor((now - newsDate) / (1000 * 60 * 60));
  return diffInHours < 24;
});

const readingTime = computed(() => {
  if (!berita.value?.content) return 3;
  const words = berita.value.content.split(' ').length;
  const minutes = Math.ceil(words / 200);
  return minutes || 3;
});

const relatedTags = computed(() => {
  if (!berita.value?.title) return [];
  const words = berita.value.title.split(' ').slice(0, 3);
  return words.map(word => word.toLowerCase().replace(/[^\w]/g, ''));
});

function getRelatedCover(item) {
  return item.cover
    ? `http://localhost:8080/uploads/berita/${item.cover}`
    : "https://images.unsplash.com/photo-1588681664899-f142ff2dc9b1?ixlib=rb-4.0.3&auto=format&fit=crop&w=400&q=80";
}

function formatRelatedDate(dateString) {
  if (!dateString) return "";
  const date = new Date(dateString);
  const now = new Date();
  const diffInHours = Math.floor((now - date) / (1000 * 60 * 60));
  
  if (diffInHours < 1) return "Baru saja";
  if (diffInHours < 24) return `${diffInHours} jam lalu`;
  
  return date.toLocaleDateString("id-ID", {
    month: "short",
    day: "numeric"
  });
}

function copyLink() {
  const url = window.location.href;
  navigator.clipboard.writeText(url);
  showShareToast.value = true;
  setTimeout(() => showShareToast.value = false, 2000);
}

function shareTwitter() {
  window.open(`https://twitter.com/intent/tweet?url=${encodeURIComponent(window.location.href)}&text=${encodeURIComponent(berita.value?.title)}`, '_blank');
}

function goToHome() {
  navigateTo('/');
}

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}
</script>

<style>
/* Custom styling untuk artikel content dengan dark mode support */
.article-content {
  font-family: 'Georgia', 'Times New Roman', serif;
  font-size: 1.125rem; /* text-lg */
  line-height: 1.75;
  color: #1f2937; /* gray-800 */
}

.dark .article-content {
  color: #e5e7eb; /* gray-200 */
}

.article-content h2 {
  font-size: 1.875rem; /* text-3xl */
  font-weight: bold;
  margin-top: 2rem; /* mt-8 */
  margin-bottom: 1rem; /* mb-4 */
  color: #111827; /* gray-900 */
}

.dark .article-content h2 {
  color: #f9fafb; /* gray-50 */
}

.article-content h3 {
  font-size: 1.5rem; /* text-2xl */
  font-weight: bold;
  margin-top: 1.5rem; /* mt-6 */
  margin-bottom: 0.75rem; /* mb-3 */
  color: #1f2937; /* gray-800 */
}

.dark .article-content h3 {
  color: #f3f4f6; /* gray-100 */
}

.article-content p {
  margin-bottom: 1rem; /* mb-4 */
  color: #374151; /* gray-700 */
}

.dark .article-content p {
  color: #d1d5db; /* gray-300 */
}

.article-content img {
  border-radius: 0.75rem; /* rounded-xl */
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
  margin-top: 1.5rem; /* my-6 */
  margin-bottom: 1.5rem; /* my-6 */
  max-width: 100%;
  height: auto;
}

.article-content blockquote {
  border-left: 4px solid #2563eb; /* border-blue-600 */
  padding-left: 1.5rem; /* pl-6 */
  padding-top: 0.5rem; /* py-2 */
  padding-bottom: 0.5rem; /* py-2 */
  margin-top: 1.5rem; /* my-6 */
  margin-bottom: 1.5rem; /* my-6 */
  font-style: italic;
  color: #4b5563; /* gray-600 */
}

.dark .article-content blockquote {
  color: #d1d5db; /* gray-300 */
  border-left-color: #06b6d4; /* border-cyan-500 */
}

.article-content ul {
  list-style-type: disc;
  padding-left: 1.25rem; /* pl-5 */
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
  color: #374151; /* gray-700 */
}

.dark .article-content ul {
  color: #d1d5db; /* gray-300 */
}

.article-content ol {
  list-style-type: decimal;
  padding-left: 1.25rem; /* pl-5 */
  margin-top: 0.5rem;
  margin-bottom: 0.5rem;
  color: #374151; /* gray-700 */
}

.dark .article-content ol {
  color: #d1d5db; /* gray-300 */
}

.article-content li {
  margin-bottom: 0.25rem;
}

.article-content strong {
  font-weight: bold;
  color: #111827; /* gray-900 */
}

.dark .article-content strong {
  color: #f9fafb; /* gray-50 */
}

.article-content a {
  color: #2563eb; /* blue-600 */
  text-decoration: none;
}

.dark .article-content a {
  color: #06b6d4; /* cyan-500 */
}

.article-content a:hover {
  text-decoration: underline;
}

/* Custom animations */
@keyframes float-slow {
  0%, 100% { transform: translateY(0) translateX(0); }
  50% { transform: translateY(-20px) translateX(10px); }
}

@keyframes float-medium {
  0%, 100% { transform: translateY(0) translateX(0); }
  50% { transform: translateY(-15px) translateX(-15px); }
}

@keyframes float-fast {
  0%, 100% { transform: translateY(0) translateX(0); }
  50% { transform: translateY(-10px) translateX(5px); }
}

@keyframes gradient-x {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

@keyframes bounce-subtle {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-5px); }
}

@keyframes pulse-subtle {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.8; }
}

@keyframes bounce-in {
  0% {
    opacity: 0;
    transform: translateY(20px) scale(0.9);
  }
  50% {
    transform: translateY(-10px) scale(1.05);
  }
  70% {
    transform: translateY(5px) scale(1);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes slide-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-float-slow {
  animation: float-slow 15s ease-in-out infinite;
}

.animate-float-medium {
  animation: float-medium 12s ease-in-out infinite;
}

.animate-float-fast {
  animation: float-fast 10s ease-in-out infinite;
}

.animate-gradient-x {
  animation: gradient-x 3s ease infinite;
  background-size: 200% 200%;
}

.animate-bounce-subtle {
  animation: bounce-subtle 2s ease-in-out infinite;
}

.animate-pulse-subtle {
  animation: pulse-subtle 2s ease-in-out infinite;
}

.animate-bounce-in {
  animation: bounce-in 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s ease;
}

.slide-up-enter-from,
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

/* Enhanced scrollbar */
::-webkit-scrollbar {
  width: 12px;
}

::-webkit-scrollbar-track {
  background-color: #f3f4f6; /* gray-100 */
}

.dark ::-webkit-scrollbar-track {
  background-color: #111827; /* gray-900 */
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #06b6d4, #3b82f6); /* cyan-500 to blue-500 */
  border-radius: 9999px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #0891b2, #2563eb); /* cyan-600 to blue-600 */
}

/* Smooth content animation */
.article-content {
  animation: slide-up 0.6s ease-out;
}

/* Card hover effects */
.group:hover .group-hover\:shadow-3xl {
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}

/* Text selection color */
::selection {
  background-color: rgba(6, 182, 212, 0.3); /* cyan-500/30 */
  color: #111827; /* gray-900 */
}

.dark ::selection {
  color: #f9fafb; /* white */
}

/* Improve dark mode text visibility */
.dark .text-gray-600 {
  color: #d1d5db; /* gray-300 */
}

.dark .text-gray-700 {
  color: #d1d5db; /* gray-300 */
}

.dark .text-gray-800 {
  color: #e5e7eb; /* gray-200 */
}

.dark .text-gray-900 {
  color: #f9fafb; /* white */
}
</style>