<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Hero Section -->
    <Navbar />
    <section class="relative bg-gradient-to-br from-blue-600 via-purple-600 to-indigo-700 text-white overflow-hidden">
      <div class="absolute inset-0 bg-black opacity-10"></div>
      <div class="absolute inset-0">
        <div class="absolute inset-0 bg-gradient-to-t from-blue-900/20 to-transparent"></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24 lg:py-32">
        <div class="text-center">
          <div class="inline-flex items-center justify-center w-20 h-20 bg-white/10 rounded-2xl backdrop-blur-sm border border-white/20 mb-8 floating">
            <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9m0 0v12m0 0h4m-4 0h4"/>
            </svg>
          </div>
          <h1 class="text-5xl md:text-7xl font-bold mb-6 leading-tight fade-in-up">
            Tentang 
            <span class="block text-transparent bg-clip-text bg-gradient-to-r from-yellow-300 to-orange-300 mt-2">
              EduNews
            </span>
          </h1>
          <p class="text-xl md:text-2xl mb-8 max-w-3xl mx-auto opacity-90 leading-relaxed fade-in-up delay-200">
            Menjadi jembatan informasi pendidikan terdepan yang menghubungkan komunitas 
            pendidikan dengan perkembangan terkini di Indonesia
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center items-center fade-in-up delay-400">
            <div class="flex items-center space-x-6 text-white/80">
              <div class="text-center">
                <div class="text-2xl font-bold counter">{{ stats.articles }}+</div>
                <div class="text-sm">Artikel</div>
              </div>
              <div class="w-px h-8 bg-white/30"></div>
              <div class="text-center">
                <div class="text-2xl font-bold counter">{{ stats.readers }}+</div>
                <div class="text-sm">Pembaca</div>
              </div>
              <div class="w-px h-8 bg-white/30"></div>
              <div class="text-center">
                <div class="text-2xl font-bold counter">{{ stats.years }}+</div>
                <div class="text-sm">Tahun</div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Wave Divider -->
      <div class="absolute bottom-0 left-0 right-0">
        <svg viewBox="0 0 1200 120" preserveAspectRatio="none" class="w-full h-12">
          <path d="M0,0V46.29c47.79,22.2,103.59,32.17,158,28,70.36-5.37,136.33-33.31,206.8-37.5C438.64,32.43,512.34,53.67,583,72.05c69.27,18,138.3,24.88,209.4,13.08,36.15-6,69.85-17.84,104.45-29.34C989.49,25,1113-14.29,1200,52.47V0Z" opacity=".25" class="fill-current text-gray-50"></path>
          <path d="M0,0V15.81C13,36.92,27.64,56.86,47.69,72.05,99.41,111.27,165,111,224.58,91.58c31.15-10.15,60.09-26.07,89.67-39.8,40.92-19,84.73-46,130.83-49.67,36.26-2.85,70.9,9.42,98.6,31.56,31.77,25.39,62.32,62,103.63,73,40.44,10.79,81.35-6.69,119.13-24.28s75.16-39,116.92-43.05c59.73-5.85,113.28,22.88,168.9,38.84,30.2,8.66,59,6.17,87.09-7.5,22.43-10.89,48-26.93,60.65-49.24V0Z" opacity=".5" class="fill-current text-gray-50"></path>
          <path d="M0,0V5.63C149.93,59,314.09,71.32,475.83,42.57c43-7.64,84.23-20.12,127.61-26.46,59-8.63,112.48,12.24,165.56,35.4C827.93,77.22,886,95.24,951.2,90c86.53-7,172.46-45.71,248.8-84.81V0Z" class="fill-current text-gray-50"></path>
        </svg>
      </div>
    </section>

    <!-- Visi Misi -->
    <section class="py-20 lg:py-28">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid lg:grid-cols-2 gap-16 items-center">
          <!-- Visi -->
          <div class="relative">
            <div class="absolute -top-4 -left-4 w-24 h-24 bg-blue-100 rounded-2xl opacity-50 floating"></div>
            <div class="absolute -bottom-4 -right-4 w-20 h-20 bg-purple-100 rounded-2xl opacity-50 floating delay-2000"></div>
            <div class="relative bg-white rounded-3xl shadow-xl p-8 lg:p-10 border border-gray-100 slide-in-left">
              <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-blue-500 to-blue-600 rounded-2xl text-white mb-6">
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                </svg>
              </div>
              <h2 class="text-3xl lg:text-4xl font-bold text-gray-800 mb-6">Visi Kami</h2>
              <p class="text-lg lg:text-xl text-gray-600 leading-relaxed">
                Menjadi platform berita pendidikan terpercaya nomor satu di Indonesia yang memberikan 
                akses informasi berkualitas, terkini, dan inspiratif untuk mendukung transformasi 
                pendidikan nasional.
              </p>
            </div>
          </div>

          <!-- Misi -->
          <div class="space-y-6 slide-in-right">
            <h2 class="text-3xl lg:text-4xl font-bold text-gray-800 mb-8">Misi Kami</h2>
            <div class="space-y-6">
              <div v-for="(mission, index) in missions" :key="index" 
                   class="flex items-start space-x-4 p-6 bg-white rounded-2xl shadow-lg border border-gray-100 hover:shadow-xl transition-all duration-300 group mission-item"
                   :style="`animation-delay: ${index * 200}ms`">
                <div class="flex-shrink-0 w-12 h-12 bg-gradient-to-br from-green-500 to-emerald-600 rounded-xl flex items-center justify-center text-white transform group-hover:scale-110 transition duration-300">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                </div>
                <div class="flex-1">
                  <h3 class="text-lg font-semibold text-gray-800 mb-2">{{ mission.title }}</h3>
                  <p class="text-gray-600 leading-relaxed">{{ mission.description }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Nilai Perusahaan -->
    <section class="py-20 bg-gradient-to-br from-gray-50 to-blue-50/30">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl lg:text-5xl font-bold text-gray-800 mb-4 fade-in-up">Nilai Kami</h2>
          <p class="text-xl text-gray-600 max-w-3xl mx-auto fade-in-up delay-200">
            Prinsip-prinsip yang menjadi fondasi dalam setiap langkah kami
          </p>
        </div>
        <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
          <div v-for="(value, index) in values" :key="value.id" 
               class="bg-white rounded-2xl shadow-lg p-8 text-center hover:shadow-2xl transition-all duration-500 transform hover:-translate-y-2 group value-card"
               :style="`animation-delay: ${index * 150}ms`">
            <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-blue-500 to-purple-600 rounded-2xl text-white mb-6 group-hover:scale-110 transition duration-300">
              <svg v-if="value.id === 1" class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
              <svg v-else-if="value.id === 2" class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
              </svg>
              <svg v-else-if="value.id === 3" class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"/>
              </svg>
              <svg v-else class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"/>
              </svg>
            </div>
            <h3 class="text-xl font-bold text-gray-800 mb-4">{{ value.title }}</h3>
            <p class="text-gray-600 leading-relaxed">{{ value.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Tim -->
    <section class="py-20 lg:py-28 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl lg:text-5xl font-bold text-gray-800 mb-4 fade-in-up">Tim Kami</h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto fade-in-up delay-200">
            Para profesional berpengalaman yang berdedikasi menghadirkan konten pendidikan terbaik
          </p>
        </div>
        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8 lg:gap-12">
          <div v-for="(member, index) in team" :key="member.id" 
               class="text-center group team-member"
               :style="`animation-delay: ${index * 200}ms`">
            <div class="relative mb-6">
              <div class="w-40 h-40 mx-auto rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-5xl font-bold transform group-hover:scale-105 transition duration-500 shadow-2xl">
                {{ member.initials }}
              </div>
              <div class="absolute inset-0 rounded-full bg-gradient-to-br from-blue-500 to-purple-600 opacity-0 group-hover:opacity-20 transition duration-500 transform group-hover:scale-110"></div>
            </div>
            <h3 class="text-2xl font-bold text-gray-800 mb-2">{{ member.name }}</h3>
            <p class="text-blue-600 font-semibold text-lg mb-3">{{ member.role }}</p>
            <p class="text-gray-600 leading-relaxed">{{ member.description }}</p>
            <div class="flex justify-center space-x-4 mt-4">
              <span v-for="(skill, skillIndex) in member.skills" :key="skillIndex" 
                    class="px-3 py-1 bg-blue-100 text-blue-700 rounded-full text-sm font-medium skill-tag"
                    :style="`animation-delay: ${skillIndex * 100 + 600}ms`">
                {{ skill }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Pencapaian -->
    <section class="py-20 bg-gradient-to-r from-blue-600 to-purple-700 text-white relative overflow-hidden">
      <div class="absolute inset-0 bg-black opacity-10"></div>
      <div class="absolute top-0 right-0 w-64 h-64 bg-white/5 rounded-full -translate-y-32 translate-x-32 floating"></div>
      <div class="absolute bottom-0 left-0 w-80 h-80 bg-white/5 rounded-full translate-y-40 -translate-x-40 floating delay-3000"></div>
      
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-4xl lg:text-5xl font-bold mb-4 fade-in-up">Pencapaian Kami</h2>
          <p class="text-xl opacity-90 max-w-2xl mx-auto fade-in-up delay-200">
            Perjalanan kami dalam menghadirkan konten pendidikan berkualitas
          </p>
        </div>
        <div class="grid grid-cols-2 lg:grid-cols-4 gap-8">
          <div v-for="(stat, index) in achievements" :key="stat.id" 
               class="text-center transform hover:scale-105 transition duration-300 achievement-item"
               :style="`animation-delay: ${index * 150}ms`">
            <div class="text-4xl lg:text-5xl font-bold mb-2 bg-white/10 rounded-2xl p-6 backdrop-blur-sm border border-white/20 counter">
              {{ stat.value }}{{ stat.suffix }}
            </div>
            <div class="text-blue-100 font-semibold">{{ stat.label }}</div>
            <div class="text-blue-200 text-sm mt-1">{{ stat.description }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Kontak -->
    <section class="py-20 lg:py-28 bg-gray-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid lg:grid-cols-2 gap-16 items-start">
          <!-- Info Kontak -->
          <div class="slide-in-left">
            <h2 class="text-4xl lg:text-5xl font-bold text-gray-800 mb-6">Mari Berkolaborasi</h2>
            <p class="text-xl text-gray-600 mb-8 leading-relaxed">
              Tertarik untuk berkolaborasi atau memiliki pertanyaan? 
              Kami senang mendengar dari Anda dan siap membantu.
            </p>
            
            <div class="space-y-6">
              <div v-for="(contact, index) in contacts" :key="contact.type"
                   class="flex items-center space-x-4 p-6 bg-white rounded-2xl shadow-lg border border-gray-100 hover:shadow-xl transition-all duration-300 group contact-item"
                   :style="`animation-delay: ${index * 100}ms`">
                <div class="flex-shrink-0 w-12 h-12 bg-gradient-to-br from-blue-500 to-blue-600 rounded-xl flex items-center justify-center text-white transform group-hover:scale-110 transition duration-300">
                  <component :is="contact.icon" class="w-6 h-6" />
                </div>
                <div>
                  <h3 class="font-semibold text-gray-800 mb-1">{{ contact.type }}</h3>
                  <p class="text-gray-600">{{ contact.value }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Form Kontak -->
          <div class="bg-white rounded-3xl shadow-2xl p-8 lg:p-10 border border-gray-100 slide-in-right">
            <h3 class="text-2xl lg:text-3xl font-bold text-gray-800 mb-6">Kirim Pesan</h3>
            <form @submit.prevent="sendMessage" class="space-y-6">
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">Nama Lengkap</label>
                <input 
                  v-model="contactForm.name"
                  type="text" 
                  placeholder="Masukkan nama Anda" 
                  class="w-full p-4 rounded-xl border border-gray-300 focus:outline-none focus:ring-3 focus:ring-blue-500 focus:border-transparent transition duration-300"
                  required
                >
              </div>
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">Email</label>
                <input 
                  v-model="contactForm.email"
                  type="email" 
                  placeholder="email@contoh.com" 
                  class="w-full p-4 rounded-xl border border-gray-300 focus:outline-none focus:ring-3 focus:ring-blue-500 focus:border-transparent transition duration-300"
                  required
                >
              </div>
              <div>
                <label class="block text-sm font-semibold text-gray-700 mb-2">Pesan</label>
                <textarea 
                  v-model="contactForm.message"
                  placeholder="Tulis pesan Anda di sini..." 
                  rows="5"
                  class="w-full p-4 rounded-xl border border-gray-300 focus:outline-none focus:ring-3 focus:ring-blue-500 focus:border-transparent transition duration-300 resize-none"
                  required
                ></textarea>
              </div>
              <button 
                type="submit"
                class="w-full bg-gradient-to-r from-blue-600 to-purple-600 text-white py-4 px-8 rounded-xl font-semibold text-lg hover:shadow-2xl transform hover:scale-105 transition duration-300 flex items-center justify-center space-x-2"
                :disabled="isSubmitting"
                :class="{'opacity-50 cursor-not-allowed': isSubmitting}"
              >
                <span v-if="!isSubmitting">Kirim Pesan</span>
                <span v-else>Mengirim...</span>
                <svg v-if="!isSubmitting" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"/>
                </svg>
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

// Team data dengan skills
const team = ref([
  {
    id: 1,
    name: "Ahmad Wijaya",
    role: "Editor in Chief",
    description: "Memiliki pengalaman 10 tahun di dunia jurnalistik pendidikan dengan fokus pada konten edukatif",
    initials: "AW",
    skills: ["Jurnalistik", "Editing", "Research"]
  },
  {
    id: 2,
    name: "Sari Dewi",
    role: "Content Manager",
    description: "Spesialis dalam pengembangan konten pendidikan dan strategi kurikulum digital",
    initials: "SD",
    skills: ["Content Strategy", "Curriculum", "SEO"]
  },
  {
    id: 3,
    name: "Budi Santoso",
    role: "Tech Lead",
    description: "Ahli dalam pengembangan platform digital dan optimasi teknologi pendidikan",
    initials: "BS",
    skills: ["Web Development", "UX/UI", "DevOps"]
  }
])

// Misi perusahaan
const missions = ref([
  {
    title: "Akurasi & Kredibilitas",
    description: "Menyajikan berita pendidikan yang akurat, terverifikasi, dan dapat dipertanggungjawabkan"
  },
  {
    title: "Inovasi Konten",
    description: "Mengembangkan format konten yang kreatif dan engaging untuk berbagai kalangan pembaca"
  },
  {
    title: "Kolaborasi Edukatif",
    description: "Membangun kemitraan dengan institusi pendidikan dan pakar di bidangnya"
  },
  {
    title: "Aksesibilitas Universal",
    description: "Memastikan informasi pendidikan dapat diakses oleh semua lapisan masyarakat"
  }
])

// Nilai perusahaan
const values = ref([
  {
    id: 1,
    title: "Integritas",
    description: "Selalu menjunjung tinggi kejujuran dan transparansi dalam setiap konten"
  },
  {
    id: 2,
    title: "Inovasi",
    description: "Terus berinovasi dalam menyajikan konten pendidikan yang relevan"
  },
  {
    id: 3,
    title: "Kolaborasi",
    description: "Bekerja sama dengan berbagai pihak untuk kemajuan pendidikan"
  },
  {
    id: 4,
    title: "Ekselen",
    description: "Berusaha memberikan yang terbaik dalam setiap konten dan layanan"
  }
])

// Pencapaian
const achievements = ref([
  {
    id: 1,
    value: "500",
    suffix: "+",
    label: "Artikel",
    description: "Konten Berkualitas"
  },
  {
    id: 2,
    value: "25",
    suffix: "K+",
    label: "Pembaca",
    description: "Per Bulan"
  },
  {
    id: 3,
    value: "15",
    suffix: "+",
    label: "Kategori",
    description: "Topik Pendidikan"
  },
  {
    id: 4,
    value: "3",
    suffix: " Tahun",
    label: "Pengalaman",
    description: "Dedikasi"
  }
])

// Stats untuk hero section
const stats = ref({
  articles: 500,
  readers: 25000,
  years: 3
})

// Form kontak
const contactForm = ref({
  name: '',
  email: '',
  message: ''
})

const isSubmitting = ref(false)

// Contact info dengan icons
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
  },
  {
    type: 'Jam Operasional',
    value: 'Senin - Jumat, 09:00 - 17:00',
    icon: 'ClockIcon'
  }
])

// Icons untuk contact section
const EmailIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
    </svg>
  `
}

const PhoneIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"/>
    </svg>
  `
}

const LocationIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"/>
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"/>
    </svg>
  `
}

const ClockIcon = {
  template: `
    <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
    </svg>
  `
}

const sendMessage = async () => {
  isSubmitting.value = true
  
  // Simulasi pengiriman pesan
  await new Promise(resolve => setTimeout(resolve, 2000))
  
  alert('Pesan berhasil dikirim! Terima kasih telah menghubungi kami. Kami akan membalas dalam 1x24 jam.')
  contactForm.value = { name: '', email: '', message: '' }
  isSubmitting.value = false
}

// Animasi on scroll
onMounted(() => {
  const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
  }

  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-in')
      }
    })
  }, observerOptions)

  // Observe semua elemen dengan class animasi
  document.querySelectorAll('.fade-in-up, .slide-in-left, .slide-in-right, .mission-item, .value-card, .team-member, .achievement-item, .contact-item').forEach(el => {
    observer.observe(el)
  })
})
</script>

<style scoped>
/* Custom animations */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-50px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(50px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes countUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Animation classes */
.floating {
  animation: float 6s ease-in-out infinite;
}

.delay-200 {
  animation-delay: 200ms;
}

.delay-400 {
  animation-delay: 400ms;
}

.delay-2000 {
  animation-delay: 2000ms;
}

.delay-3000 {
  animation-delay: 3000ms;
}

/* Initial state for animated elements */
.fade-in-up,
.slide-in-left,
.slide-in-right,
.mission-item,
.value-card,
.team-member,
.achievement-item,
.contact-item,
.skill-tag {
  opacity: 0;
}

/* Animation when in view */
.fade-in-up.animate-in {
  animation: fadeInUp 0.8s ease-out forwards;
}

.slide-in-left.animate-in {
  animation: slideInLeft 0.8s ease-out forwards;
}

.slide-in-right.animate-in {
  animation: slideInRight 0.8s ease-out forwards;
}

.mission-item.animate-in,
.value-card.animate-in,
.team-member.animate-in,
.achievement-item.animate-in,
.contact-item.animate-in {
  animation: fadeInUp 0.6s ease-out forwards;
}

.skill-tag.animate-in {
  animation: countUp 0.4s ease-out forwards;
}

/* Counter animation */
.counter {
  animation: countUp 1s ease-out forwards;
}

/* Smooth transitions */
* {
  transition: all 0.3s ease;
}

/* Hover effects enhancement */
.group:hover .group-hover\:scale-110 {
  transform: scale(1.1);
}

/* Custom scroll behavior */
html {
  scroll-behavior: smooth;
}

/* Loading state for form */
button:disabled {
  cursor: not-allowed;
  opacity: 0.7;
}
</style>