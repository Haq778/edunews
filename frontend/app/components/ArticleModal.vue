<template>
  <Transition name="fade">
    <div
      v-if="isOpen"
      class="fixed inset-0 z-[9999] flex items-center justify-center bg-black/70 backdrop-blur-sm p-4"
    >
      <div
        class="relative bg-white dark:bg-gray-900 rounded-2xl shadow-2xl max-w-4xl w-full max-h-[90vh] overflow-hidden flex flex-col animate-fadeIn"
      >
        <!-- Close Button -->
        <button
          @click="$emit('close')"
          class="absolute top-4 right-4 z-20 bg-white/30 backdrop-blur-sm text-white dark:text-gray-300 p-2 rounded-full hover:bg-white/50 transition"
        >
          ✕
        </button>

        <!-- Image Section -->
        <div class="relative h-64 md:h-80 w-full overflow-hidden">
          
          <!-- Render hanya jika url valid -->
          <img
            v-if="coverUrl"
            :src="coverUrl"
            :alt="article?.title"
            class="w-full h-full object-cover transition-transform duration-700 hover:scale-105"
            @load="handleImageLoad"
            @error="handleImageError"
          />

          <!-- Placeholder aman (bukan route) -->
          <div
            v-else
            class="w-full h-full bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center text-white opacity-80"
          >
            <span class="text-lg font-semibold">Tidak ada gambar</span>
          </div>

          <!-- Loading skeleton -->
          <div
            v-if="imageLoading"
            class="absolute inset-0 animate-pulse bg-gray-300 dark:bg-gray-700"
          ></div>
        </div>

        <!-- Content -->
        <div class="p-6 overflow-y-auto">
          <h2 class="text-3xl font-bold mb-4 text-gray-800 dark:text-gray-100">
            {{ article?.title }}
          </h2>

          <p class="text-gray-500 dark:text-gray-400 mb-4 text-sm">
            {{ formattedDate }} • 👁️ {{ article?.views || 0 }} views
          </p>

          <p class="text-gray-700 dark:text-gray-300 leading-relaxed whitespace-pre-line">
            {{ article?.content }}
          </p>
        </div>

      </div>
    </div>
  </Transition>
</template>

<script setup>
import { computed, ref } from 'vue'



const props = defineProps({
  article: Object,
  isOpen: Boolean,
})

const imageLoading = ref(true)

// URL cover backend — tidak pakai default-cover.jpg
const coverUrl = computed(() => {
  if (!props.article?.cover) return null
  return `http://localhost:8080/uploads/berita/${props.article.cover}`
})

const formattedDate = computed(() => {
  if (!props.article?.date) return ''
  return new Date(props.article.date).toLocaleString('id-ID', {
    dateStyle: 'long',
    timeStyle: 'short'
  })
})

function handleImageLoad() {
  imageLoading.value = false
}

function handleImageError(event) {
  console.warn("Gambar gagal dimuat, memakai placeholder")
  event.target.style.display = "none" // hide broken image
  imageLoading.value = false
}
</script>
