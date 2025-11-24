<template>
  <div class="py-8">
    <Navbar />

    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700 text-white overflow-hidden">
      <div class="absolute inset-0 bg-black opacity-10 pointer-events-none"></div>
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12 pt-8">
          <div class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-br from-orange-500 to-red-500 rounded-2xl shadow-lg mb-6 animate-float">
            <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"/>
            </svg>
          </div>

          <h1 class="text-4xl md:text-5xl font-bold text-gray-900 dark:text-white mb-4">
            Berita <span class="text-transparent bg-clip-text bg-gradient-to-r from-orange-500 to-red-500">Populer</span>
          </h1>

          <p class="text-lg text-gray-600 dark:text-gray-300 max-w-2xl mx-auto">
            Artikel paling banyak dibaca dan trending di komunitas EduNews
          </p>

          <!-- Filter Tabs -->
          <div class="flex flex-wrap justify-center gap-4 mt-8">
            <button
              v-for="tab in tabs"
              :key="tab.id"
              @click="activeTab = tab.id"
              :aria-pressed="activeTab === tab.id"
              :class="[
                'px-5 py-2 rounded-xl font-semibold transition-all duration-200 transform hover:scale-105 focus:outline-none',
                activeTab === tab.id
                  ? 'bg-gradient-to-r from-orange-500 to-red-500 text-white shadow-lg'
                  : 'bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 border border-gray-200 dark:border-gray-700'
              ]"
            >
              {{ tab.label }}
            </button>
          </div>
        </div>

        <!-- Stats Overview -->
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          <div
            v-for="stat in stats"
            :key="stat.id"
            class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg p-6 text-center hover:shadow-xl transition-transform duration-300"
          >
            <div class="text-2xl md:text-3xl font-bold text-gray-900 dark:text-white mb-2">{{ stat.value }}{{ stat.suffix }}</div>
            <div class="text-sm text-gray-500 dark:text-gray-400">{{ stat.label }}</div>
          </div>
        </div>

        <!-- Popular Articles Grid -->
        <div class="grid lg:grid-cols-2 xl:grid-cols-3 gap-8 mb-16">
          <article
            v-for="(article, index) in filteredArticles"
            :key="article.id"
            class="bg-white dark:bg-gray-800 rounded-3xl shadow-lg overflow-hidden hover:shadow-2xl transition-transform duration-300 group"
          >
            <!-- Article Image / Hero -->
            <div class="relative h-48 bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
              <span class="text-white text-6xl font-bold opacity-20 select-none">{{ article.emoji }}</span>

              <div class="absolute top-4 right-4">
                <span class="bg-red-500 text-white px-3 py-1 rounded-full text-sm font-semibold flex items-center">
                  <svg class="w-4 h-4 mr-1" fill="currentColor" viewBox="0 0 20 20"><path d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z"/></svg>
                  <span>{{ article.views }}</span>
                </span>
              </div>

              <div class="absolute bottom-4 left-4">
                <span class="bg-white/90 dark:bg-gray-800/90 backdrop-blur-sm text-gray-700 dark:text-gray-300 px-3 py-1 rounded-full text-sm font-medium">{{ article.category }}</span>
              </div>
            </div>

            <!-- Article Content -->
            <div class="p-6">
              <div class="flex items-center justify-between mb-3 text-sm text-gray-500 dark:text-gray-400">
                <span>{{ article.date }}</span>
                <div class="flex items-center space-x-1 text-yellow-500">
                  <svg
                    v-for="n in 5"
                    :key="n"
                    class="w-4 h-4"
                    :class="n <= article.rating ? 'text-yellow-400' : 'text-gray-300 dark:text-gray-600'"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                  >
                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                  </svg>
                </div>
              </div>

              <h3 class="text-xl font-bold text-gray-900 dark:text-white mb-3 line-clamp-2 group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors duration-300">{{ article.title }}</h3>

              <p class="text-gray-600 dark:text-gray-300 mb-4 line-clamp-3">{{ article.excerpt }}</p>

              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 bg-gradient-to-br from-green-500 to-blue-500 rounded-full flex items-center justify-center text-white text-sm font-bold">{{ article.author.initials }}</div>
                  <span class="text-sm text-gray-600 dark:text-gray-400">{{ article.author.name }}</span>
                </div>

                <button
                  @click="showArticle(article)"
                  class="flex items-center space-x-2 text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 font-semibold transition-colors duration-200"
                >
                  <span>Baca</span>
                  <svg class="w-4 h-4 transform transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                  </svg>
                </button>
              </div>
            </div>
          </article>
        </div>

        <!-- Coming Soon Features -->
        <div class="bg-gradient-to-r from-orange-500 to-red-500 rounded-3xl shadow-2xl p-8 lg:p-12 text-white text-center mb-16">
          <div class="max-w-4xl mx-auto">
            <h2 class="text-2xl md:text-3xl font-bold mb-4">Fitur Menarik Segera Hadir! 🚀</h2>
            <p class="text-base opacity-90 mb-8">Kami sedang mengembangkan fitur analitik real-time, rekomendasi personalisasi, dan leaderboard pembaca aktif untuk pengalaman yang lebih baik.</p>

            <div class="grid md:grid-cols-3 gap-6 mb-8">
              <div v-for="feature in upcomingFeatures" :key="feature.id" class="bg-white/10 backdrop-blur-sm rounded-2xl p-6 border border-white/20">
                <div class="text-3xl mb-3">{{ feature.emoji }}</div>
                <h3 class="font-semibold text-lg mb-2">{{ feature.title }}</h3>
                <p class="text-white/80 text-sm">{{ feature.description }}</p>
              </div>
            </div>

            <div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
              <NuxtLink to="/" class="bg-white text-orange-600 px-6 py-3 rounded-xl font-semibold hover:shadow-2xl transition duration-200 inline-flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/></svg>
                <span>Kembali ke Beranda</span>
              </NuxtLink>

              <button @click="showNotification = true" class="bg-transparent border-2 border-white text-white px-6 py-3 rounded-xl font-semibold hover:bg-white/10 transition duration-200 inline-flex items-center space-x-2">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM21 12a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
                <span>Notify Me</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Notification Toast -->
        <div v-if="showNotification" class="fixed bottom-4 right-4 bg-green-500 text-white px-6 py-4 rounded-2xl shadow-2xl max-w-sm z-50">
          <div class="flex items-center space-x-3">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/></svg>
            <div>
              <p class="font-semibold">Terima kasih!</p>
              <p class="text-sm opacity-90">Kami akan beri tahu ketika fitur baru tersedia.</p>
            </div>
            <button @click="showNotification = false" class="text-white/80 hover:text-white">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>
            </button>
          </div>
        </div>
      </div>
    </section>
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

const tabs = ref([
  { id: 'all', label: 'Semua' },
  { id: 'weekly', label: 'Minggu Ini' },
  { id: 'monthly', label: 'Bulan Ini' },
  { id: 'yearly', label: 'Tahun Ini' }
])

const stats = ref([
  { id: 1, value: 156, suffix: 'K', label: 'Total Views' },
  { id: 2, value: 892, suffix: '',  label: 'Artikel Populer' },
  { id: 3, value: 45, suffix: 'K', label: 'Pembaca Aktif' },
  { id: 4, value: 98, suffix: '%', label: 'Kepuasan' }
])

const articles = ref([
  { id: 1, title: "Revolusi Pendidikan Digital: Transformasi Pembelajaran di Era 4.0", excerpt: "Bagaimana teknologi mengubah landscape pendidikan Indonesia dan peluang yang terbuka...", category: "Teknologi", views: "15.2K", rating: 5, date: "2 hari lalu", emoji: "📚", slug: "revolusi-pendidikan-digital", author: { name: "Dr. Sari Dewi", initials: "SD" } },
  { id: 2, title: "Strategi Sukses Masuk PTN Favorit 2024", excerpt: "Tips dan trik terbaru untuk menghadapi seleksi masuk perguruan tinggi negeri...", category: "Kampus", views: "12.8K", rating: 4, date: "1 minggu lalu", emoji: "🎓", slug: "strategi-sukses-ptn", author: { name: "Prof. Ahmad", initials: "PA" } },
  { id: 3, title: "Inovasi Kurikulum Merdeka: Apa yang Perlu Diketahui?", excerpt: "Analisis mendalam tentang implementasi kurikulum merdeka...", category: "Kebijakan", views: "10.5K", rating: 5, date: "3 hari lalu", emoji: "💡", slug: "inovasi-kurikulum-merdeka", author: { name: "Budi Santoso", initials: "BS" } },
  { id: 4, title: "Beasiswa Luar Negeri 2024: Peluang Emas untuk Mahasiswa", excerpt: "Daftar lengkap beasiswa internasional dan panduan aplikasi sukses...", category: "Beasiswa", views: "9.7K", rating: 4, date: "5 hari lalu", emoji: "✈️", slug: "beasiswa-luar-negeri", author: { name: "Maria Silva", initials: "MS" } },
  { id: 5, title: "Mental Health di Kalangan Pelajar", excerpt: "Mengenal tanda-tanda stress akademik dan cara menjaga kesehatan mental...", category: "Kesehatan", views: "8.9K", rating: 5, date: "1 minggu lalu", emoji: "❤️", slug: "mental-health-pelajar", author: { name: "Dr. Lisa Wang", initials: "LW" } },
  { id: 6, title: "Skill yang Paling Dicari di Dunia Kerja 2024", excerpt: "Analisis tren keterampilan yang dibutuhkan industri...", category: "Karir", views: "7.4K", rating: 4, date: "4 hari lalu", emoji: "🧠", slug: "skill-dunia-kerja", author: { name: "Andi Prasetyo", initials: "AP" } }
])

const upcomingFeatures = ref([
  { id: 1, emoji: '📊', title: 'Real-time Analytics', description: 'Statistik pembaca langsung diperbarui.' },
  { id: 2, emoji: '✨', title: 'Rekomendasi Cerdas', description: 'Artikel disesuaikan dengan minat Anda.' },
  { id: 3, emoji: '🏆', title: 'Top Reader Leaderboard', description: 'Peringkat pembaca paling aktif.' }
])

const filteredArticles = computed(() => {
  if (activeTab.value === 'all') return articles.value
  if (activeTab.value === 'weekly') return articles.value.filter(a => /hari|minggu/i.test(a.date))
  if (activeTab.value === 'monthly') return articles.value.slice(0, 4)
  if (activeTab.value === 'yearly') return articles.value.slice(0, 2)
  return articles.value
})

const showArticle = (article) => {
  // jika rute detail belum siap, cukup arahkan ke slug
  router.push(`/berita/${article.slug}`)
}

// auto-hide toast
onMounted(() => {
  if (showNotification.value) {
    setTimeout(() => showNotification.value = false, 5000)
  }
})
</script>

<style>
/* Animations (dipusatkan di sini agar dapat bekerja dengan dark mode) */
@keyframes float { 0%,100%{transform:translateY(0)}50%{transform:translateY(-8px)} }
.animate-float{animation:float 6s ease-in-out infinite}

@keyframes fadeInUp{0%{opacity:0;transform:translateY(20px)}100%{opacity:1;transform:translateY(0)}}
.fade-in-up{opacity:0;animation:fadeInUp 0.8s ease-out forwards}

/* Line clamp (fallback jika plugin belum aktif) */
.line-clamp-2{display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden}
.line-clamp-3{display:-webkit-box;-webkit-line-clamp:3;-webkit-box-orient:vertical;overflow:hidden}

/* Card group hover effect */
.group:hover .group-hover\\:text-blue-600{color:var(--tw-text-opacity)}
</style>