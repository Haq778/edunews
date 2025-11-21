<template>
  <Transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="transform opacity-0 scale-95"
    enter-to-class="transform opacity-100 scale-100"
    leave-active-class="transition-all duration-200 ease-in"
    leave-from-class="transform opacity-100 scale-100"
    leave-to-class="transform opacity-0 scale-95"
  >
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <!-- Backdrop -->
      <div 
        class="fixed inset-0 bg-black/60 backdrop-blur-sm transition-opacity"
        @click="close"
      ></div>

      <!-- Modal Container -->
      <div class="flex min-h-full items-center justify-center p-4">
        <!-- Modal Content -->
        <div class="relative bg-white dark:bg-gray-800 rounded-2xl shadow-2xl w-full max-w-4xl max-h-[90vh] overflow-hidden transform transition-all">
          
          <!-- Header -->
          <div class="sticky top-0 bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 px-6 py-4 z-10">
            <div class="flex justify-between items-center">
              <div>
                <span class="inline-block px-3 py-1 text-xs font-semibold bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full">
                  {{ article?.category || 'Pendidikan' }}
                </span>
              </div>
              <button
                @click="close"
                class="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-full transition-colors duration-200"
              >
                <svg class="w-6 h-6 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>

          <!-- Scrollable Content -->
          <div class="overflow-y-auto max-h-[calc(90vh-80px)]">
            <!-- Article Image -->
            <div class="relative h-64 sm:h-80 md:h-96 bg-gray-200 dark:bg-gray-700 overflow-hidden">
              <img
                v-if="article?.image"
                :src="getImageUrl(article.image)"
                :alt="article?.title"
                class="w-full h-full object-cover transition-transform duration-700 hover:scale-105"
                @error="handleImageError"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-r from-blue-400 to-purple-500">
                <svg class="w-16 h-16 text-white opacity-70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </div>
              
              <!-- Image Loading State -->
              <div v-if="imageLoading" class="absolute inset-0 flex items-center justify-center bg-gray-100 dark:bg-gray-800">
                <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
              </div>
            </div>

            <!-- Article Content -->
            <div class="px-6 py-8">
              <!-- Article Meta -->
              <div class="flex flex-wrap items-center gap-4 text-sm text-gray-500 dark:text-gray-400 mb-6">
                <div class="flex items-center gap-1">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <span>{{ formatDate(article?.date) }}</span>
                </div>
                <div class="flex items-center gap-1">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <span>{{ article?.views || 0 }} views</span>
                </div>
              </div>

              <!-- Article Title -->
              <h1 class="text-3xl lg:text-4xl font-bold text-gray-900 dark:text-white mb-6 leading-tight">
                {{ article?.title }}
              </h1>

              <!-- Article Excerpt -->
              <p class="text-lg text-gray-600 dark:text-gray-300 mb-8 leading-relaxed font-medium">
                {{ article?.excerpt }}
              </p>

              <!-- Article Body -->
              <div class="prose prose-lg dark:prose-invert max-w-none">
                <div class="text-gray-700 dark:text-gray-300 leading-8 space-y-4">
                  <p v-for="(paragraph, index) in getContentParagraphs()" :key="index" class="text-justify">
                    {{ paragraph }}
                  </p>
                </div>

                <!-- Sample Content (if no content provided) -->
                <div v-if="!article?.content" class="text-gray-700 dark:text-gray-300 leading-8 space-y-4">
                  <p>
                    Transformasi digital dalam dunia pendidikan telah membawa perubahan signifikan 
                    dalam cara belajar dan mengajar. Di era new normal ini, teknologi menjadi 
                    tulang punggung sistem pendidikan yang lebih inklusif dan adaptif.
                  </p>
                  <p>
                    Dengan adopsi berbagai platform pembelajaran digital, siswa dan guru dapat 
                    berinteraksi secara lebih efektif meskipun berada di lokasi yang berbeda. 
                    Hal ini tidak hanya meningkatkan aksesibilitas pendidikan, tetapi juga 
                    membuka peluang baru untuk pengembangan metode pembelajaran yang lebih kreatif.
                  </p>
                  <p>
                    Namun, tantangan tetap ada. Kesetaraan akses teknologi dan kemampuan digital 
                    yang beragam di antara peserta didik menjadi perhatian utama yang perlu 
                    diatasi untuk memastikan tidak ada yang tertinggal dalam proses transformasi ini.
                  </p>
                </div>
              </div>

              <!-- Tags -->
              <div class="mt-8 pt-6 border-t border-gray-200 dark:border-gray-700">
                <div class="flex flex-wrap gap-2">
                  <span 
                    v-for="tag in getTags()" 
                    :key="tag"
                    class="px-3 py-1 bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 text-sm rounded-full hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors cursor-pointer"
                  >
                    #{{ tag }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="sticky bottom-0 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 px-6 py-4">
            <div class="flex justify-between items-center">
              <button
                @click="close"
                class="px-6 py-2 text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200 transition-colors duration-200 font-medium"
              >
                Tutup
              </button>
              <div class="flex gap-2">
                <button
                  @click="shareArticle"
                  class="flex items-center gap-2 px-4 py-2 text-blue-600 dark:text-blue-400 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-lg transition-colors duration-200"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                  </svg>
                  Bagikan
                </button>
                <button
                  @click="bookmarkArticle"
                  class="flex items-center gap-2 px-4 py-2 text-blue-600 dark:text-blue-400 hover:bg-blue-50 dark:hover:bg-blue-900/30 rounded-lg transition-colors duration-200"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                  </svg>
                  Simpan
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  isOpen: Boolean,
  article: Object
})

const emit = defineEmits(['close'])

const imageLoading = ref(true)
const imageError = ref(false)

// Watch for article changes to reset image states
watch(() => props.article, () => {
  imageLoading.value = true
  imageError.value = false
})

function close() {
  emit('close')
}

function getImageUrl(imagePath) {
  // Handle different image path formats
  if (!imagePath) return ''
  
  // If it's already a full URL, return as is
  if (imagePath.startsWith('http')) return imagePath
  
  // If it's a relative path, construct the full URL
  if (imagePath.startsWith('/')) {
    return `${window.location.origin}${imagePath}`
  }
  
  // For other cases, assume it's a relative path from assets
  return `/images/${imagePath}`
}

function handleImageError(event) {
  console.error('Error loading image:', props.article?.image)
  imageError.value = true
  imageLoading.value = false
  
  // You can set a fallback image here
  event.target.style.display = 'none'
}

function handleImageLoad() {
  imageLoading.value = false
  imageError.value = false
}

function formatDate(dateString) {
  if (!dateString) return 'Tanggal tidak tersedia'
  
  const options = { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric',
    timeZone: 'UTC'
  }
  return new Date(dateString).toLocaleDateString('id-ID', options)
}

function getContentParagraphs() {
  if (!props.article?.content) return []
  return props.article.content.split('\n').filter(p => p.trim())
}

function getTags() {
  const baseTags = ['pendidikan', 'berita', 'edunews']
  if (props.article?.category) {
    baseTags.unshift(props.article.category.toLowerCase())
  }
  return [...new Set(baseTags)] // Remove duplicates
}

function shareArticle() {
  if (navigator.share) {
    navigator.share({
      title: props.article?.title,
      text: props.article?.excerpt,
      url: window.location.href,
    })
    .catch(error => console.log('Error sharing:', error))
  } else {
    // Fallback: copy to clipboard
    navigator.clipboard.writeText(window.location.href)
    alert('Link berita telah disalin ke clipboard!')
  }
}

function bookmarkArticle() {
  // Implement bookmark functionality
  const bookmarks = JSON.parse(localStorage.getItem('edunews-bookmarks') || '[]')
  const articleId = props.article?.id
  
  if (!bookmarks.find(b => b.id === articleId)) {
    bookmarks.push(props.article)
    localStorage.setItem('edunews-bookmarks', JSON.stringify(bookmarks))
    alert('Berita disimpan!')
  } else {
    alert('Berita sudah disimpan sebelumnya!')
  }
}

// Add keyboard event listener for ESC key
const handleKeydown = (event) => {
  if (event.key === 'Escape' && props.isOpen) {
    close()
  }
}

// Add/remove event listener when component mounts/unmounts or when modal opens/closes
watch(() => props.isOpen, (isOpen) => {
  if (isOpen) {
    document.addEventListener('keydown', handleKeydown)
    document.body.style.overflow = 'hidden'
  } else {
    document.removeEventListener('keydown', handleKeydown)
    document.body.style.overflow = 'auto'
  }
})
</script>

<style scoped>
/* Custom scrollbar for modal */
.scrollable-content::-webkit-scrollbar {
  width: 6px;
}

.scrollable-content::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.scrollable-content::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.scrollable-content::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

.dark .scrollable-content::-webkit-scrollbar-track {
  background: #374151;
}

.dark .scrollable-content::-webkit-scrollbar-thumb {
  background: #6b7280;
}

.dark .scrollable-content::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>