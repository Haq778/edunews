<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50/30 to-indigo-50/20 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 transition-colors duration-500 overflow-x-hidden">
    <Navbar />
    
    <!-- Enhanced Hero Section -->
    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-800 text-white overflow-hidden py-16 lg:py-20">
      <!-- Background Animations -->
      <div class="absolute inset-0 overflow-hidden">
        <!-- Floating Shapes -->
        <div class="absolute top-1/4 left-1/4 w-32 h-32 bg-white/10 rounded-full blur-3xl animate-float-slow"></div>
        <div class="absolute top-1/3 right-1/4 w-24 h-24 bg-purple-400/20 rounded-full blur-2xl animate-float-medium delay-1000"></div>
        <div class="absolute bottom-1/4 left-1/3 w-28 h-28 bg-indigo-400/15 rounded-full blur-3xl animate-float-slow delay-2000"></div>
        <div class="absolute top-3/4 right-1/3 w-20 h-20 bg-blue-400/25 rounded-full blur-2xl animate-float-fast delay-1500"></div>

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
          🎓 Tentang EduNews
        </div>
        
        <h1 class="text-4xl md:text-5xl font-bold text-white mb-4 animate-text-slide-up">
          Tentang 
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-yellow-300 via-amber-300 to-orange-300">
            EduNews
          </span>
        </h1>
        
        <p class="text-lg text-white/90 mb-8 max-w-2xl mx-auto animate-text-slide-up delay-200 leading-relaxed">
          Menghadirkan informasi pendidikan terpercaya dengan cara yang sederhana dan bermakna untuk seluruh komunitas pendidikan Indonesia.
        </p>

        <!-- Stats -->
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

    <!-- Visi Misi Enhanced -->
    <section class="py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid lg:grid-cols-2 gap-8">
          <!-- Visi -->
          <div class="bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg hover:shadow-xl transition-all duration-500 transform hover:scale-105 slide-in-left border border-gray-200 dark:border-gray-700">
            <div class="w-16 h-16 bg-gradient-to-br from-blue-500 to-purple-600 rounded-2xl text-white mb-6 flex items-center justify-center shadow-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-gray-800 dark:text-white mb-4">Visi Kami</h2>
            <p class="text-gray-600 dark:text-gray-300 leading-relaxed">
              Menjadi platform berita pendidikan terpercaya nomor satu di Indonesia yang memberikan 
              akses informasi berkualitas untuk mendukung transformasi pendidikan nasional.
            </p>
          </div>

          <!-- Misi -->
          <div class="bg-white dark:bg-gray-800 rounded-2xl p-8 shadow-lg hover:shadow-xl transition-all duration-500 transform hover:scale-105 slide-in-right border border-gray-200 dark:border-gray-700">
            <div class="w-16 h-16 bg-gradient-to-br from-purple-500 to-indigo-600 rounded-2xl text-white mb-6 flex items-center justify-center shadow-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-gray-800 dark:text-white mb-4">Misi Kami</h2>
            <div class="space-y-4">
              <div v-for="(mission, index) in missions" :key="index" 
                   class="flex items-start space-x-4 p-4 bg-gray-50 dark:bg-gray-700/50 rounded-xl mission-item transition-all duration-300 group hover:bg-gray-100 dark:hover:bg-gray-700"
                   :style="`animation-delay: ${index * 100}ms`">
                <div class="flex-shrink-0 w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center text-white text-sm group-hover:scale-110 transition-transform duration-300">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                  </svg>
                </div>
                <div>
                  <h3 class="font-semibold text-gray-800 dark:text-white mb-1 text-sm">{{ mission.title }}</h3>
                  <p class="text-gray-600 dark:text-gray-300 text-xs">{{ mission.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Nilai Inti Enhanced -->
    <section class="py-16 bg-gradient-to-br from-blue-600 to-indigo-700 text-white">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold mb-4 fade-in-up">Nilai Kami</h2>
          <p class="text-blue-100 fade-in-up delay-200">Prinsip yang kami pegang teguh dalam setiap langkah</p>
        </div>
        
        <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="(value, index) in values" :key="value.id" 
               class="bg-white/10 backdrop-blur-sm rounded-2xl p-6 text-center value-card border border-white/20 hover:bg-white/15 transition-all duration-500 transform hover:scale-105 group"
               :style="`animation-delay: ${index * 100}ms`">
            <div class="w-12 h-12 bg-white/20 rounded-2xl text-white mb-4 mx-auto flex items-center justify-center group-hover:scale-110 transition-transform duration-300">
              <component :is="value.icon" class="w-6 h-6" />
            </div>
            <h3 class="font-bold mb-2 text-sm">{{ value.title }}</h3>
            <p class="text-blue-100 text-xs leading-relaxed">{{ value.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Tim Inti Enhanced -->
    <section class="py-16 bg-white dark:bg-gray-800">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold text-gray-800 dark:text-white mb-4 fade-in-up">Tim Inti</h2>
          <p class="text-gray-600 dark:text-gray-300 fade-in-up delay-200">Orang-orang berdedikasi di balik EduNews</p>
        </div>
        
        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="(member, index) in team" :key="member.id" 
               class="text-center team-member group"
               :style="`animation-delay: ${index * 150}ms`">
            <div class="w-24 h-24 mx-auto mb-6 rounded-2xl bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-xl font-bold shadow-lg group-hover:scale-110 transition-transform duration-500">
              {{ member.initials }}
            </div>
            <h3 class="text-lg font-bold text-gray-800 dark:text-white mb-2">{{ member.name }}</h3>
            <p class="text-purple-600 dark:text-purple-400 font-semibold mb-3 text-sm">{{ member.role }}</p>
            <p class="text-gray-600 dark:text-gray-300 text-sm leading-relaxed">{{ member.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Pencapaian Enhanced -->
    <section class="py-16 bg-gradient-to-br from-gray-50 to-blue-50/30 dark:from-gray-800 dark:to-gray-700/50">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl font-bold text-gray-800 dark:text-white mb-12 fade-in-up">Pencapaian Kami</h2>
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="(stat, index) in achievements" :key="stat.id" 
               class="achievement-item group"
               :style="`animation-delay: ${index * 100}ms`">
            <div class="text-3xl font-bold text-gray-800 dark:text-white mb-2 group-hover:scale-110 transition-transform duration-300">{{ stat.value }}{{ stat.suffix }}</div>
            <div class="text-gray-600 dark:text-gray-300 text-sm">{{ stat.label }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Kontak Enhanced -->
    <section class="py-16 bg-white dark:bg-gray-800">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold text-gray-800 dark:text-white mb-4 fade-in-up">Hubungi Kami</h2>
          <p class="text-gray-600 dark:text-gray-300 fade-in-up delay-200">Sampaikan pertanyaan, masukan, atau kolaborasi</p>
        </div>
        
        <div class="grid lg:grid-cols-2 gap-8">
          <!-- Info Kontak -->
          <div class="space-y-4 slide-in-left">
            <div v-for="(contact, index) in contacts" :key="contact.type"
                 class="flex items-center space-x-4 p-4 bg-gray-50 dark:bg-gray-700 rounded-2xl contact-item transition-all duration-300 transform hover:scale-105 group"
                 :style="`animation-delay: ${index * 100}ms`">
              <div class="flex-shrink-0 w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-xl flex items-center justify-center text-white group-hover:scale-110 transition-transform duration-300">
                <component :is="contact.icon" class="w-6 h-6" />
              </div>
              <div>
                <h3 class="font-semibold text-gray-800 dark:text-white text-sm">{{ contact.type }}</h3>
                <p class="text-gray-600 dark:text-gray-300 text-xs">{{ contact.value }}</p>
              </div>
            </div>
          </div>

          <!-- Form Kontak -->
          <div class="bg-gray-50 dark:bg-gray-700 rounded-2xl p-6 slide-in-right transition-all duration-300">
            <form @submit.prevent="sendMessage" class="space-y-4">
              <input 
                v-model="contactForm.name"
                type="text" 
                placeholder="Nama Lengkap" 
                class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white dark:bg-gray-800 text-gray-900 dark:text-white transition-all duration-300 text-sm"
                required
              >
              <input 
                v-model="contactForm.email"
                type="email" 
                placeholder="Email" 
                class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white dark:bg-gray-800 text-gray-900 dark:text-white transition-all duration-300 text-sm"
                required
              >
              <textarea 
                v-model="contactForm.message"
                placeholder="Pesan Anda..." 
                rows="4"
                class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white dark:bg-gray-800 text-gray-900 dark:text-white resize-none transition-all duration-300 text-sm"
                required
              ></textarea>
              <button 
                type="submit"
                class="w-full bg-gradient-to-r from-blue-500 to-purple-600 text-white py-3 px-6 rounded-xl font-medium hover:shadow-lg transition-all duration-300 transform hover:scale-105 text-sm"
                :disabled="isSubmitting"
                :class="{'opacity-50 cursor-not-allowed': isSubmitting}"
              >
                {{ isSubmitting ? 'Mengirim...' : 'Kirim Pesan' }}
              </button>
            </form>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Navbar from '~/components/Navbar.vue'

// Data
const team = ref([
  {
    id: 1,
    name: "Ahmad Wijaya",
    role: "Editor Utama",
    description: "Memastikan setiap konten berkualitas dan akurat dengan pengalaman 10+ tahun di jurnalistik",
    initials: "AW"
  },
  {
    id: 2,
    name: "Sari Dewi", 
    role: "Manajer Konten",
    description: "Mengelola dan mengembangkan materi pendidikan dengan fokus pada kurikulum terbaru",
    initials: "SD"
  },
  {
    id: 3,
    name: "Budi Santoso",
    role: "Lead Developer",
    description: "Menjaga platform tetap optimal dan mudah digunakan dengan teknologi terkini",
    initials: "BS"
  }
])

const missions = ref([
  {
    title: "Akurasi & Kredibilitas",
    description: "Menyajikan berita pendidikan yang akurat dan terverifikasi"
  },
  {
    title: "Inovasi Konten", 
    description: "Mengembangkan format konten kreatif dan engaging"
  },
  {
    title: "Kolaborasi Edukatif",
    description: "Bermitra dengan institusi pendidikan dan pakar"
  },
  {
    title: "Aksesibilitas Universal", 
    description: "Informasi dapat diakses semua lapisan masyarakat"
  }
])

const values = ref([
  {
    id: 1,
    title: "Integritas",
    description: "Kejujuran dan transparansi dalam setiap konten",
    icon: 'ShieldIcon'
  },
  {
    id: 2,
    title: "Inovasi",
    description: "Terus berinovasi dalam konten pendidikan",
    icon: 'LightIcon'
  },
  {
    id: 3, 
    title: "Kolaborasi",
    description: "Bekerja sama untuk kemajuan pendidikan",
    icon: 'TeamIcon'
  },
  {
    id: 4,
    title: "Ekselen", 
    description: "Memberikan yang terbaik dalam setiap layanan",
    icon: 'StarIcon'
  }
])

const achievements = ref([
  { id: 1, value: "500", suffix: "+", label: "Artikel" },
  { id: 2, value: "25", suffix: "K+", label: "Pembaca" },
  { id: 3, value: "50", suffix: "+", label: "Sekolah" },
  { id: 4, value: "3", suffix: "+", label: "Tahun" }
])

const stats = ref([
  { id: 1, value: 500, suffix: "+", label: "Artikel" },
  { id: 2, value: 25, suffix: "K+", label: "Pembaca" },
  { id: 3, value: 50, suffix: "+", label: "Sekolah" },
  { id: 4, value: 3, suffix: "+", label: "Tahun" }
])

const contacts = ref([
  {
    type: 'Email',
    value: 'hello@edunews.id',
    icon: 'EmailIcon'
  },
  {
    type: 'Telepon',
    value: '+62 21 1234 5678', 
    icon: 'PhoneIcon'
  },
  {
    type: 'Alamat',
    value: 'Jakarta, Indonesia',
    icon: 'LocationIcon'
  }
])

const contactForm = ref({
  name: '',
  email: '', 
  message: ''
})

const isSubmitting = ref(false)

// Icons
const ShieldIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/></svg>`
}

const LightIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/></svg>`
}

const TeamIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"/></svg>`
}

const StarIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/></svg>`
}

const EmailIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/></svg>`
}

const PhoneIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"/></svg>`
}

const LocationIcon = {
  template: `<svg fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/></svg>`
}

const sendMessage = async () => {
  isSubmitting.value = true
  await new Promise(resolve => setTimeout(resolve, 2000))
  alert('Terima kasih! Pesan Anda telah kami terima.')
  contactForm.value = { name: '', email: '', message: '' }
  isSubmitting.value = false
}

// Animations
onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-in')
      }
    })
  }, { threshold: 0.1 })

  document.querySelectorAll('.fade-in-up, .slide-in-left, .slide-in-right, .mission-item, .value-card, .team-member, .achievement-item, .contact-item').forEach(el => {
    observer.observe(el)
  })
})
</script>

<style scoped>
/* Enhanced Animation System */
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

@keyframes slideInLeft {
  0% {
    opacity: 0;
    transform: translateX(-30px);
  }
  100% {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInRight {
  0% {
    opacity: 0;
    transform: translateX(30px);
  }
  100% {
    opacity: 1;
    transform: translateX(0);
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

.slide-in-left.animate-in {
  animation: slideInLeft 0.6s ease-out forwards;
}

.slide-in-right.animate-in {
  animation: slideInRight 0.6s ease-out forwards;
}

/* Grid Pattern */
.bg-grid-white {
  background-image: 
    linear-gradient(rgba(255, 255, 255, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.1) 1px, transparent 1px);
  background-size: 30px 30px;
}

/* Animation Delays */
.delay-200 {
  animation-delay: 200ms;
}

/* Initial States */
.fade-in-up,
.slide-in-left, 
.slide-in-right,
.mission-item,
.value-card,
.team-member,
.achievement-item,
.contact-item {
  opacity: 0;
}

.fade-in-up.animate-in {
  animation: text-slide-up 0.6s ease-out forwards;
}

.mission-item.animate-in,
.value-card.animate-in,
.team-member.animate-in,
.achievement-item.animate-in,
.contact-item.animate-in {
  animation: text-slide-up 0.6s ease-out forwards;
}

/* Enhanced backdrop blur */
.backdrop-blur-sm {
  backdrop-filter: blur(8px) saturate(180%);
  -webkit-backdrop-filter: blur(8px) saturate(180%);
}

/* Smooth scrolling */
html {
  scroll-behavior: smooth;
}
</style>