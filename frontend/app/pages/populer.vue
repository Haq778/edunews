<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50/30 to-indigo-50/20 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 transition-colors duration-500">
    <Navbar />

    <!-- Compact Hero Section -->
    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-800 text-white overflow-hidden py-16 lg:py-20">
      <!-- Animated Background -->
      <div class="absolute inset-0 overflow-hidden">
        <!-- Floating Shapes -->
        <div class="absolute top-10 left-10 w-24 h-24 bg-white/10 rounded-full blur-2xl animate-float-slow"></div>
        <div class="absolute top-20 right-20 w-16 h-16 bg-purple-400/20 rounded-full blur-xl animate-float-medium delay-1000"></div>
        <div class="absolute bottom-10 left-20 w-20 h-20 bg-indigo-400/15 rounded-full blur-2xl animate-float-slow delay-2000"></div>
        <div class="absolute bottom-20 right-10 w-12 h-12 bg-blue-400/25 rounded-full blur-xl animate-float-fast delay-1500"></div>

        <!-- Animated Grid -->
        <div class="absolute inset-0 opacity-10">
          <div class="w-full h-full bg-grid-white animate-grid-drift"></div>
        </div>

        <!-- Floating Particles -->
        <div 
          v-for="i in 8" 
          :key="i"
          class="absolute w-1.5 h-1.5 bg-white/50 rounded-full animate-particle-float"
          :style="{
            top: Math.random() * 100 + '%',
            left: Math.random() * 100 + '%',
            animationDelay: Math.random() * 6 + 's',
            animationDuration: (8 + Math.random() * 8) + 's'
          }"
        ></div>
      </div>

      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <!-- Animated Badge -->
        <div class="inline-flex items-center px-4 py-2 bg-white/20 backdrop-blur-sm rounded-xl text-sm font-medium mb-6 border border-white/30 animate-badge-float">
          <span class="w-2 h-2 bg-yellow-400 rounded-full mr-2 animate-pulse"></span>
          🏆 Konten Paling Populer
        </div>

        <!-- Compact Hero Text -->
        <div class="space-y-4 animate-text-slide-up">
          <h1 class="text-3xl md:text-4xl font-bold mb-4 leading-tight">
            Berita <span class="text-transparent bg-clip-text bg-gradient-to-r from-yellow-300 via-amber-300 to-orange-300">Populer</span>
          </h1>
          
          <p class="text-lg text-white/90 max-w-2xl mx-auto leading-relaxed">
            Artikel paling banyak dibaca dan trending di komunitas EduNews
          </p>
        </div>

        <!-- Enhanced Filter Tabs -->
        <div class="flex flex-wrap justify-center gap-2 mt-8">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'px-4 py-2.5 rounded-xl font-medium transition-all duration-300 transform hover:scale-105 backdrop-blur-sm border text-sm',
              activeTab === tab.id
                ? 'bg-white text-purple-700 shadow-lg border-white'
                : 'bg-white/10 text-white border-white/20 hover:bg-white/20'
            ]"
          >
            <span class="flex items-center space-x-2">
              <span>{{ tab.emoji }}</span>
              <span>{{ tab.label }}</span>
            </span>
          </button>
        </div>

        <!-- Compact Stats -->
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-8 max-w-md mx-auto">
          <div
            v-for="stat in stats"
            :key="stat.id"
            class="bg-white/10 backdrop-blur-sm rounded-2xl p-4 text-center hover:shadow-lg transition-all duration-300 transform hover:scale-105 border border-white/20"
          >
            <div class="text-xl font-bold text-white mb-1">{{ stat.value }}{{ stat.suffix }}</div>
            <div class="text-white/80 text-xs">{{ stat.label }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Main Content -->
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <!-- Articles Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-3 gap-6 mb-12">
        <article
          v-for="(article, index) in filteredArticles"
          :key="article.id"
          class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-500 transform hover:scale-105 group animate-card-enter border border-gray-200 dark:border-gray-700 overflow-hidden"
          :style="{
            animationDelay: `${index * 100}ms`
          }"
        >
          <!-- Article Header -->
          <div class="relative h-40 bg-gradient-to-br from-blue-500 via-purple-500 to-indigo-600 flex items-center justify-center overflow-hidden">
            <div class="absolute inset-0 bg-black/20 group-hover:bg-black/30 transition-colors duration-300"></div>
            <span class="text-white text-4xl font-bold opacity-30 select-none z-10">{{ article.emoji }}</span>
            
            <!-- Trending Badge -->
            <div class="absolute top-3 right-3 z-20">
              <span class="bg-amber-500 text-white px-2 py-1 rounded-lg text-xs font-semibold flex items-center shadow-md">
                <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z"/>
                </svg>
                <span>{{ formatViews(article.views) }}</span>
              </span>
            </div>

            <!-- Category Badge -->
            <div class="absolute bottom-3 left-3 z-20">
              <span class="bg-white/90 dark:bg-gray-800/90 backdrop-blur-sm text-gray-700 dark:text-gray-300 px-2 py-1 rounded-lg text-xs font-medium border border-white/50">
                {{ article.category }}
              </span>
            </div>
          </div>

          <!-- Article Content -->
          <div class="p-5">
            <div class="flex items-center justify-between mb-3 text-xs text-gray-500 dark:text-gray-400">
              <span class="flex items-center space-x-1">
                <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                </svg>
                <span>{{ formatDate(article.date) }}</span>
              </span>
              <div class="flex items-center space-x-0.5">
                <svg
                  v-for="n in 5"
                  :key="n"
                  class="w-3 h-3"
                  :class="n <= article.rating ? 'text-amber-400' : 'text-gray-300 dark:text-gray-600'"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                </svg>
              </div>
            </div>

            <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-2 line-clamp-2 group-hover:text-purple-600 dark:group-hover:text-purple-400 transition-colors duration-300 leading-tight">
              {{ article.title }}
            </h3>

            <p class="text-gray-600 dark:text-gray-300 text-sm mb-4 line-clamp-2 leading-relaxed">
              {{ article.excerpt }}
            </p>

            <div class="flex items-center justify-between pt-3 border-t border-gray-200 dark:border-gray-700">
              <div class="flex items-center space-x-2">
                <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-500 rounded-full flex items-center justify-center text-white text-xs font-bold shadow-md">
                  {{ article.author.initials }}
                </div>
                <div>
                  <div class="text-xs font-semibold text-gray-900 dark:text-white">{{ article.author.name }}</div>
                  <div class="text-xs text-gray-500 dark:text-gray-400">Penulis</div>
                </div>
              </div>

              <button
                @click="showArticle(article)"
                class="flex items-center space-x-1 text-purple-600 dark:text-purple-400 hover:text-purple-700 dark:hover:text-purple-300 font-medium text-sm transition-colors duration-200 group/btn"
              >
                <span>Baca</span>
                <svg class="w-3 h-3 transform group-hover/btn:translate-x-0.5 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                </svg>
              </button>
            </div>
          </div>
        </article>
      </div>

      <!-- Load More Section -->
      <div class="text-center">
        <button
          @click="loadMore"
          class="inline-flex items-center space-x-2 bg-gradient-to-r from-blue-500 to-purple-600 text-white px-6 py-3 rounded-xl hover:shadow-lg transition-all duration-300 transform hover:scale-105 font-medium"
        >
          <span>Muat Lebih Banyak</span>
          <svg class="w-4 h-4 transform group-hover:translate-y-0.5 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
          </svg>
        </button>
      </div>

      <!-- Features Section -->
      <div class="mt-16 bg-gradient-to-r from-blue-500 to-purple-600 rounded-2xl shadow-xl p-8 text-white relative overflow-hidden">
        <!-- Background Pattern -->
        <div class="absolute inset-0 opacity-10">
          <div class="absolute top-0 right-0 w-32 h-32 bg-white rounded-full -translate-y-16 translate-x-16"></div>
          <div class="absolute bottom-0 left-0 w-24 h-24 bg-white rounded-full translate-y-12 -translate-x-12"></div>
        </div>
        
        <div class="relative text-center">
          <h2 class="text-2xl font-bold mb-4">Fitur Menarik Segera Hadir! 🚀</h2>
          <p class="text-white/90 mb-8 max-w-md mx-auto text-sm leading-relaxed">
            Kami sedang mengembangkan fitur analitik real-time dan rekomendasi personalisasi untuk pengalaman yang lebih baik.
          </p>

          <div class="grid md:grid-cols-2 gap-6 mb-8">
            <div 
              v-for="feature in upcomingFeatures" 
              :key="feature.id" 
              class="bg-white/10 backdrop-blur-sm rounded-xl p-4 border border-white/20 hover:bg-white/20 transition-all duration-300 transform hover:scale-105"
            >
              <div class="text-2xl mb-2">{{ feature.emoji }}</div>
              <h3 class="font-semibold text-sm mb-1">{{ feature.title }}</h3>
              <p class="text-white/80 text-xs leading-relaxed">{{ feature.description }}</p>
            </div>
          </div>

          <div class="flex flex-col sm:flex-row gap-3 justify-center items-center">
            <NuxtLink 
              to="/" 
              class="bg-white text-purple-600 px-6 py-2.5 rounded-xl font-medium hover:shadow-lg transition-all duration-300 transform hover:scale-105 inline-flex items-center space-x-2 text-sm"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
              </svg>
              <span>Kembali ke Beranda</span>
            </NuxtLink>

            <button 
              @click="showNotification = true" 
              class="bg-transparent border border-white text-white px-6 py-2.5 rounded-xl font-medium hover:bg-white/10 transition-all duration-300 transform hover:scale-105 inline-flex items-center space-x-2 text-sm"
            >
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <span>Notify Me</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Notification Toast -->
    <Transition
      enter-active-class="transition-all duration-500 ease-out"
      enter-from-class="transform opacity-0 translate-x-full"
      enter-to-class="transform opacity-100 translate-x-0"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="transform opacity-100 translate-x-0"
      leave-to-class="transform opacity-0 translate-x-full"
    >
      <div 
        v-if="showNotification" 
        class="fixed bottom-4 right-4 bg-purple-500 text-white px-4 py-3 rounded-xl shadow-lg max-w-xs z-50 backdrop-blur-sm bg-purple-500/95 border border-purple-400"
      >
        <div class="flex items-center space-x-2">
          <div class="w-6 h-6 bg-white/20 rounded-full flex items-center justify-center">
            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
          </div>
          <div class="flex-1">
            <p class="font-medium text-sm">Terima kasih!</p>
            <p class="text-xs opacity-90 mt-0.5">Kami akan beri tahu ketika fitur baru tersedia.</p>
          </div>
          <button 
            @click="showNotification = false" 
            class="text-white/80 hover:text-white transition-colors duration-200 p-0.5 rounded-full hover:bg-white/10"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from '~/components/Navbar.vue'

const router = useRouter()

// State
const activeTab = ref('all')
const showNotification = ref(false)
const visibleCount = ref(6)

const tabs = ref([
  { id: 'all', label: 'Semua', emoji: '🔥' },
  { id: 'weekly', label: 'Minggu Ini', emoji: '📅' },
  { id: 'monthly', label: 'Bulan Ini', emoji: '📊' },
  { id: 'yearly', label: 'Tahun Ini', emoji: '🏆' }
])

const stats = ref([
  { id: 1, value: 156, suffix: 'K', label: 'Total Views' },
  { id: 2, value: 892, suffix: '',  label: 'Artikel Populer' },
  { id: 3, value: 45, suffix: 'K', label: 'Pembaca Aktif' },
  { id: 4, value: 98, suffix: '%', label: 'Kepuasan' }
])

const articles = ref([
  { 
    id: 1, 
    title: "Revolusi Pendidikan Digital: Transformasi Pembelajaran di Era 4.0", 
    excerpt: "Bagaimana teknologi mengubah landscape pendidikan Indonesia dan peluang yang terbuka untuk generasi digital native...", 
    category: "Teknologi", 
    views: "15200", 
    rating: 5, 
    date: "2024-01-15", 
    emoji: "📚", 
    slug: "revolusi-pendidikan-digital", 
    author: { name: "Dr. Sari Dewi", initials: "SD" } 
  },
  { 
    id: 2, 
    title: "Strategi Sukses Masuk PTN Favorit 2024", 
    excerpt: "Tips dan trik terbaru untuk menghadapi seleksi masuk perguruan tinggi negeri dengan persiapan yang matang...", 
    category: "Kampus", 
    views: "12800", 
    rating: 4, 
    date: "2024-01-10", 
    emoji: "🎓", 
    slug: "strategi-sukses-ptn", 
    author: { name: "Prof. Ahmad", initials: "PA" } 
  },
  { 
    id: 3, 
    title: "Inovasi Kurikulum Merdeka: Apa yang Perlu Diketahui?", 
    excerpt: "Analisis mendalam tentang implementasi kurikulum merdeka dan dampaknya terhadap kualitas pendidikan nasional...", 
    category: "Kebijakan", 
    views: "10500", 
    rating: 5, 
    date: "2024-01-12", 
    emoji: "💡", 
    slug: "inovasi-kurikulum-merdeka", 
    author: { name: "Budi Santoso", initials: "BS" } 
  },
  { 
    id: 4, 
    title: "Beasiswa Luar Negeri 2024: Peluang Emas untuk Mahasiswa", 
    excerpt: "Daftar lengkap beasiswa internasional dan panduan aplikasi sukses untuk meraih kesempatan studi...", 
    category: "Beasiswa", 
    views: "9700", 
    rating: 4, 
    date: "2024-01-08", 
    emoji: "✈️", 
    slug: "beasiswa-luar-negeri", 
    author: { name: "Maria Silva", initials: "MS" } 
  },
  { 
    id: 5, 
    title: "Mental Health di Kalangan Pelajar", 
    excerpt: "Mengenal tanda-tanda stress akademik dan cara menjaga kesehatan mental dalam menghadapi tekanan...", 
    category: "Kesehatan", 
    views: "8900", 
    rating: 5, 
    date: "2024-01-05", 
    emoji: "❤️", 
    slug: "mental-health-pelajar", 
    author: { name: "Dr. Lisa Wang", initials: "LW" } 
  },
  { 
    id: 6, 
    title: "Skill yang Paling Dicari di Dunia Kerja 2024", 
    excerpt: "Analisis tren keterampilan yang dibutuhkan industri modern dan bagaimana mempersiapkan diri...", 
    category: "Karir", 
    views: "7400", 
    rating: 4, 
    date: "2024-01-14", 
    emoji: "🧠", 
    slug: "skill-dunia-kerja", 
    author: { name: "Andi Prasetyo", initials: "AP" } 
  },
  { 
    id: 7, 
    title: "Teknologi AI dalam Pendidikan: Masa Depan atau Ancaman?", 
    excerpt: "Eksplorasi dampak kecerdasan buatan terhadap sistem pendidikan dan peran guru di era digital...", 
    category: "Teknologi", 
    views: "6800", 
    rating: 5, 
    date: "2024-01-11", 
    emoji: "🤖", 
    slug: "ai-pendidikan", 
    author: { name: "Dr. Tech", initials: "DT" } 
  },
  { 
    id: 8, 
    title: "Metode Pembelajaran Efektif untuk Generasi Z", 
    excerpt: "Strategi pengajaran yang sesuai dengan karakteristik dan kebutuhan belajar generasi digital native...", 
    category: "Metode", 
    views: "6200", 
    rating: 4, 
    date: "2024-01-09", 
    emoji: "🌟", 
    slug: "metode-pembelajaran-z", 
    author: { name: "Guru Inovatif", initials: "GI" } 
  }
])

const upcomingFeatures = ref([
  { id: 1, emoji: '📊', title: 'Real-time Analytics', description: 'Statistik pembaca langsung dengan visualisasi data interaktif' },
  { id: 2, emoji: '✨', title: 'Rekomendasi Cerdas', description: 'Artikel disesuaikan dengan minat dan perilaku membaca Anda' }
])

const filteredArticles = computed(() => {
  let filtered = articles.value
  
  if (activeTab.value === 'weekly') {
    filtered = filtered.filter(a => new Date(a.date) > new Date(Date.now() - 7 * 24 * 60 * 60 * 1000))
  } else if (activeTab.value === 'monthly') {
    filtered = filtered.filter(a => new Date(a.date) > new Date(Date.now() - 30 * 24 * 60 * 60 * 1000))
  } else if (activeTab.value === 'yearly') {
    filtered = filtered.filter(a => new Date(a.date) > new Date(Date.now() - 365 * 24 * 60 * 60 * 1000))
  }
  
  return filtered.slice(0, visibleCount.value)
})

const showArticle = (article) => {
  router.push(`/berita/${article.slug}`)
}

const loadMore = () => {
  visibleCount.value += 3
}

const formatViews = (views) => {
  const num = parseInt(views)
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffTime = Math.abs(now - date)
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 1) return 'Kemarin'
  if (diffDays < 7) return `${diffDays} hari lalu`
  if (diffDays < 30) return `${Math.ceil(diffDays / 7)} minggu lalu`
  return date.toLocaleDateString('id-ID', { day: 'numeric', month: 'short' })
}

// Auto-hide toast
onMounted(() => {
  if (showNotification.value) {
    setTimeout(() => showNotification.value = false, 5000)
  }
})
</script>

<style scoped>
/* Enhanced Animations */
@keyframes float-slow {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-8px) translateX(4px) rotate(120deg);
  }
  66% {
    transform: translateY(4px) translateX(-2px) rotate(240deg);
  }
}

@keyframes float-medium {
  0%, 100% {
    transform: translateY(0px) translateX(0px) scale(1);
  }
  50% {
    transform: translateY(-6px) translateX(3px) scale(1.02);
  }
}

@keyframes float-fast {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-4px) rotate(180deg);
  }
}

@keyframes particle-float {
  0%, 100% {
    transform: translate(0, 0) rotate(0deg);
    opacity: 0.6;
  }
  25% {
    transform: translate(15px, -20px) rotate(90deg);
    opacity: 0.8;
  }
  50% {
    transform: translate(-10px, -30px) rotate(180deg);
    opacity: 1;
  }
  75% {
    transform: translate(-15px, -20px) rotate(270deg);
    opacity: 0.8;
  }
}

@keyframes grid-drift {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 50px 50px;
  }
}

@keyframes badge-float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-3px);
  }
}

@keyframes text-slide-up {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes card-enter {
  0% {
    opacity: 0;
    transform: translateY(15px) scale(0.98);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* Animation Classes */
.animate-float-slow {
  animation: float-slow 15s ease-in-out infinite;
}

.animate-float-medium {
  animation: float-medium 12s ease-in-out infinite;
}

.animate-float-fast {
  animation: float-fast 8s ease-in-out infinite;
}

.animate-particle-float {
  animation: particle-float 20s ease-in-out infinite;
}

.animate-grid-drift {
  animation: grid-drift 30s linear infinite;
}

.animate-badge-float {
  animation: badge-float 4s ease-in-out infinite;
}

.animate-text-slide-up {
  animation: text-slide-up 0.8s ease-out forwards;
}

.animate-card-enter {
  animation: card-enter 0.6s ease-out forwards;
}

/* Grid Pattern */
.bg-grid-white {
  background-image: 
    linear-gradient(rgba(255, 255, 255, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 30px 30px;
}

/* Line clamp utilities */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Smooth scrolling */
html {
  scroll-behavior: smooth;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #3b82f6, #8b5cf6);
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #2563eb, #7c3aed);
}

.dark ::-webkit-scrollbar-track {
  background: #374151;
}

.dark ::-webkit-scrollbar-thumb {
  background: linear-gradient(to bottom, #60a5fa, #a78bfa);
}

.dark ::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(to bottom, #3b82f6, #8b5cf6);
}

/* Enhanced backdrop blur */
.backdrop-blur-sm {
  backdrop-filter: blur(8px) saturate(180%);
  -webkit-backdrop-filter: blur(8px) saturate(180%);
}
</style>