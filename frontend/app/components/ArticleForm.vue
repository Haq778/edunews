<template>
  <div class="min-h-screen bg-gradient-to-br from-indigo-50 via-white to-purple-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 py-4 sm:py-8">
    <div class="max-w-4xl mx-auto px-3 sm:px-4 lg:px-8">
      
      <!-- Header dengan gradient menarik - Responsif -->
      <div class="mb-8 sm:mb-12 relative">
        <div class="absolute -inset-3 sm:-inset-4 bg-gradient-to-r from-blue-500/10 to-purple-500/10 dark:from-blue-500/5 dark:to-purple-500/5 rounded-2xl blur-xl"></div>
        <div class="relative bg-gradient-to-r from-blue-500 to-purple-600 rounded-xl sm:rounded-2xl p-4 sm:p-6 shadow-lg">
          <div class="flex items-center gap-3">
            <button 
              @click="$emit('cancel')"
              class="p-2 bg-white/20 hover:bg-white/30 backdrop-blur-sm text-white 
                     rounded-lg sm:rounded-xl transition-all duration-300 hover:scale-105 flex-shrink-0"
              aria-label="Kembali"
            >
              <svg class="w-4 h-4 sm:w-5 sm:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
              </svg>
            </button>
            <div class="flex-1 min-w-0">
              <h1 class="text-lg sm:text-xl md:text-2xl lg:text-3xl font-bold text-white drop-shadow-md truncate">
                {{ submitText }} Berita
              </h1>
              <p class="text-blue-100 mt-1 sm:mt-2 text-xs sm:text-sm truncate">
                Isi formulir dengan informasi berita terbaru
              </p>
            </div>
            <div class="hidden sm:block p-2 sm:p-3 bg-white/20 backdrop-blur-sm rounded-lg sm:rounded-xl flex-shrink-0">
              <svg class="w-6 h-6 sm:w-8 sm:h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Form Card dengan desain modern - Responsif -->
      <div class="relative group">
        <!-- Background effect -->
        <div class="absolute -inset-0.5 sm:-inset-1 bg-gradient-to-r from-blue-400 to-purple-500 rounded-xl sm:rounded-2xl blur opacity-20 sm:opacity-25 group-hover:opacity-30 transition duration-1000"></div>
        
        <!-- Main Card -->
        <div class="relative bg-gradient-to-br from-white to-blue-50 dark:from-gray-800 dark:to-gray-900 
                    rounded-xl sm:rounded-2xl shadow-lg sm:shadow-xl border border-blue-100 dark:border-gray-700 overflow-hidden">
          
          <!-- Card Header dengan accent color - Responsif -->
          <div class="bg-gradient-to-r from-blue-500/10 to-purple-500/10 dark:from-blue-500/5 dark:to-purple-500/5 
                      px-4 sm:px-6 py-3 sm:py-5 border-b border-blue-100 dark:border-gray-700">
            <div class="flex items-center gap-2 sm:gap-3">
              <div class="p-1.5 sm:p-2 bg-gradient-to-r from-blue-500 to-purple-600 rounded-lg">
                <svg class="w-4 h-4 sm:w-5 sm:h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"/>
                </svg>
              </div>
              <div class="flex-1 min-w-0">
                <h2 class="text-base sm:text-lg lg:text-xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent truncate">
                  Formulir Berita
                </h2>
                <p class="text-xs sm:text-sm text-gray-600 dark:text-gray-400 mt-0.5 sm:mt-1 truncate">
                  Lengkapi semua informasi yang diperlukan
                </p>
              </div>
            </div>
          </div>

          <!-- Form Content - Responsif -->
          <form @submit.prevent="submitForm" class="p-4 sm:p-6 md:p-8 space-y-6 sm:space-y-8">

            <!-- Judul Section -->
            <div class="space-y-3 sm:space-y-4">
              <div class="flex items-center gap-2">
                <div class="w-1.5 h-4 sm:w-2 sm:h-6 bg-gradient-to-b from-blue-500 to-purple-500 rounded-full"></div>
                <label class="block text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300">
                  Judul Berita <span class="text-red-500">*</span>
                </label>
              </div>
              <input
                v-model="form.title"
                type="text"
                class="w-full px-4 py-3 sm:px-5 sm:py-4 rounded-lg sm:rounded-xl border-2 border-gray-200 dark:border-gray-700 
                       bg-white dark:bg-gray-800 text-gray-900 dark:text-white text-sm sm:text-base
                       focus:border-blue-500 focus:ring-2 sm:focus:ring-4 focus:ring-blue-500/20 focus:outline-none
                       transition-all duration-300 placeholder-gray-400 dark:placeholder-gray-500
                       hover:border-blue-300 dark:hover:border-blue-500"
                placeholder="Masukkan judul berita yang menarik..."
                required
              />
            </div>

            <!-- Kategori & Cover Section Grid - Responsif -->
            <div class="grid grid-cols-1 gap-6 sm:gap-8">
              <!-- Kategori -->
              <div class="space-y-3 sm:space-y-4">
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-4 sm:w-2 sm:h-6 bg-gradient-to-b from-green-500 to-emerald-500 rounded-full"></div>
                  <label class="block text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300">
                    Kategori <span class="text-red-500">*</span>
                  </label>
                </div>
                <div class="relative">
                  <select
                    v-model="form.category_id"
                    required
                    class="w-full px-4 py-3 sm:px-5 sm:py-4 rounded-lg sm:rounded-xl border-2 border-gray-200 dark:border-gray-700 
                           bg-white dark:bg-gray-800 text-gray-900 dark:text-white text-sm sm:text-base
                           focus:border-green-500 focus:ring-2 sm:focus:ring-4 focus:ring-green-500/20 focus:outline-none
                           transition-all duration-300 appearance-none
                           hover:border-green-300 dark:hover:border-green-500"
                  >
                    <option value="" disabled selected class="text-gray-400 dark:text-gray-500">Pilih kategori berita</option>
                    <option
                      v-for="kat in categories"
                      :key="kat.id"
                      :value="kat.id"
                      class="text-gray-900 dark:text-white py-2 text-sm sm:text-base"
                    >
                      {{ kat.name }}
                    </option>
                  </select>
                  <div class="absolute right-3 sm:right-4 top-1/2 transform -translate-y-1/2 pointer-events-none">
                    <svg class="w-4 h-4 sm:w-5 sm:h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Cover Upload - Responsif -->
              <div class="space-y-3 sm:space-y-4">
                <div class="flex items-center gap-2">
                  <div class="w-1.5 h-4 sm:w-2 sm:h-6 bg-gradient-to-b from-amber-500 to-orange-500 rounded-full"></div>
                  <label class="block text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300">
                    Cover Berita
                  </label>
                </div>
                
                <!-- File Upload Area - Responsif -->
                <div 
                  @click="triggerFileInput"
                  class="border-2 border-dashed border-gray-300 dark:border-gray-600 
                         rounded-lg sm:rounded-xl p-4 sm:p-6 text-center cursor-pointer
                         hover:border-blue-400 hover:bg-blue-50/50 dark:hover:border-blue-500 dark:hover:bg-blue-900/10
                         transition-all duration-300 group/upload"
                >
                  <div class="flex flex-col items-center gap-2 sm:gap-3">
                    <div class="p-2 sm:p-3 bg-gradient-to-br from-amber-100 to-orange-100 dark:from-amber-900/20 dark:to-orange-900/20 rounded-full">
                      <svg class="w-6 h-6 sm:w-8 sm:h-8 text-amber-600 dark:text-amber-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/>
                      </svg>
                    </div>
                    <div>
                      <p class="font-medium text-gray-700 dark:text-gray-300 text-sm sm:text-base">
                        Upload Cover Image
                      </p>
                      <p class="text-xs sm:text-sm text-gray-500 dark:text-gray-400 mt-0.5 sm:mt-1">
                        Klik untuk memilih file (max 5MB)
                      </p>
                    </div>
                  </div>
                  <input
                    ref="fileInput"
                    type="file"
                    accept="image/*"
                    @change="onFileChange"
                    class="hidden"
                  />
                </div>
              </div>
            </div>

            <!-- Image Previews - Responsif -->
            <div v-if="form.cover_url || preview" class="grid grid-cols-1 sm:grid-cols-2 gap-3 sm:gap-4">
              <!-- Current Cover -->
              <div v-if="form.cover_url && !preview" class="space-y-2 sm:space-y-3">
                <p class="text-xs sm:text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-1.5 sm:gap-2">
                  <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                  </svg>
                  Cover saat ini
                </p>
                <div class="relative group/image rounded-lg sm:rounded-xl overflow-hidden border-2 border-green-200 dark:border-green-800">
                  <img 
                    :src="form.cover_url" 
                    class="w-full h-36 sm:h-48 object-cover transition-transform duration-500 group-hover/image:scale-105"
                    alt="Cover saat ini"
                  />
                  <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent opacity-0 group-hover/image:opacity-100 transition-opacity duration-300"></div>
                  <button
                    type="button"
                    @click="clearImage"
                    class="absolute bottom-3 right-3 sm:bottom-4 sm:right-4 px-2 py-1 sm:px-4 sm:py-2 bg-red-500 hover:bg-red-600 text-white 
                           rounded-md sm:rounded-lg transition-all duration-300 transform translate-y-4 opacity-0 
                           group-hover/image:translate-y-0 group-hover/image:opacity-100 text-xs sm:text-sm"
                  >
                    Hapus
                  </button>
                </div>
              </div>

              <!-- New Preview -->
              <div v-if="preview" class="space-y-2 sm:space-y-3">
                <p class="text-xs sm:text-sm font-medium text-gray-700 dark:text-gray-300 flex items-center gap-1.5 sm:gap-2">
                  <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
                  </svg>
                  Preview baru
                </p>
                <div class="relative group/image rounded-lg sm:rounded-xl overflow-hidden border-2 border-blue-200 dark:border-blue-800">
                  <img 
                    :src="preview" 
                    class="w-full h-36 sm:h-48 object-cover transition-transform duration-500 group-hover/image:scale-105"
                    alt="Preview cover baru"
                  />
                  <div class="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent opacity-0 group-hover/image:opacity-100 transition-opacity duration-300"></div>
                  <button
                    type="button"
                    @click="clearImage"
                    class="absolute bottom-3 right-3 sm:bottom-4 sm:right-4 px-2 py-1 sm:px-4 sm:py-2 bg-gray-800 hover:bg-gray-900 text-white 
                           rounded-md sm:rounded-lg transition-all duration-300 transform translate-y-4 opacity-0 
                           group-hover/image:translate-y-0 group-hover/image:opacity-100 text-xs sm:text-sm"
                  >
                    Batal
                  </button>
                </div>
              </div>
            </div>

            <!-- Konten Section - Responsif -->
            <div class="space-y-3 sm:space-y-4">
              <div class="flex items-center gap-2">
                <div class="w-1.5 h-4 sm:w-2 sm:h-6 bg-gradient-to-b from-purple-500 to-pink-500 rounded-full"></div>
                <label class="block text-xs sm:text-sm font-semibold text-gray-700 dark:text-gray-300">
                  Konten Berita <span class="text-red-500">*</span>
                </label>
              </div>
              <div class="relative">
                <textarea
                  v-model="form.content"
                  rows="6"
                  class="w-full px-4 py-3 sm:px-5 sm:py-4 rounded-lg sm:rounded-xl border-2 border-gray-200 dark:border-gray-700 
                         bg-white dark:bg-gray-800 text-gray-900 dark:text-white text-sm sm:text-base
                         focus:border-purple-500 focus:ring-2 sm:focus:ring-4 focus:ring-purple-500/20 focus:outline-none
                         resize-vertical transition-all duration-300 placeholder-gray-400 dark:placeholder-gray-500
                         hover:border-purple-300 dark:hover:border-purple-500"
                  placeholder="Tulis konten berita yang informatif dan menarik..."
                  required
                ></textarea>
                <div class="absolute bottom-3 right-3 sm:bottom-4 sm:right-4 text-xs text-gray-400 dark:text-gray-500">
                  {{ form.content.length }}/5000
                </div>
              </div>
            </div>

            <!-- Action Buttons dengan desain menarik - Responsif -->
            <div class="flex flex-col sm:flex-row gap-3 sm:gap-4 pt-6 sm:pt-8 border-t border-gray-100 dark:border-gray-700">
              <button
                type="submit"
                :disabled="loading"
                class="flex-1 px-4 py-3 sm:px-6 sm:py-4 bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700
                       text-white font-semibold rounded-lg sm:rounded-xl transition-all duration-300 text-sm sm:text-base
                       focus:outline-none focus:ring-2 sm:focus:ring-4 focus:ring-purple-500/30
                       disabled:opacity-50 disabled:cursor-not-allowed
                       transform hover:-translate-y-0.5 active:translate-y-0
                       shadow-md sm:shadow-lg hover:shadow-lg sm:hover:shadow-xl"
              >
                <span class="flex items-center justify-center gap-2 sm:gap-3">
                  <svg v-if="loading" class="w-4 h-4 sm:w-5 sm:h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                  </svg>
                  <svg v-else class="w-4 h-4 sm:w-5 sm:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"/>
                  </svg>
                  <span class="truncate">{{ loading ? 'Menyimpan...' : submitText }}</span>
                </span>
              </button>
              
              <button
                type="button"
                @click="$emit('cancel')"
                class="flex-1 px-4 py-3 sm:px-6 sm:py-4 bg-gradient-to-r from-gray-100 to-gray-200 hover:from-gray-200 hover:to-gray-300 
                       dark:bg-gradient-to-r dark:from-gray-700 dark:to-gray-800 dark:hover:from-gray-600 dark:hover:to-gray-700
                       text-gray-700 dark:text-gray-300 font-semibold rounded-lg sm:rounded-xl transition-all duration-300 text-sm sm:text-base
                       focus:outline-none focus:ring-2 sm:focus:ring-4 focus:ring-gray-400/20
                       transform hover:-translate-y-0.5 active:translate-y-0"
              >
                <span class="flex items-center justify-center gap-2 sm:gap-3">
                  <svg class="w-4 h-4 sm:w-5 sm:h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
                  </svg>
                  <span class="truncate">Batal</span>
                </span>
              </button>
            </div>

          </form>
        </div>
      </div>

      <!-- Info Footer - Responsif -->
      <div class="mt-6 sm:mt-8 text-center">
        <p class="text-xs sm:text-sm text-gray-500 dark:text-gray-400">
          <span class="inline-flex items-center gap-1.5 sm:gap-2">
            <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            Pastikan semua informasi yang diisi sudah benar sebelum disimpan
          </span>
        </p>
      </div>

      <!-- Mobile Bottom Action Bar -->
      <div v-if="isMobile" class="fixed bottom-0 left-0 right-0 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 p-3 sm:hidden z-10">
        <div class="flex gap-2">
          <button
            type="submit"
            @click="submitForm"
            :disabled="loading"
            class="flex-1 px-4 py-3 bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700
                   text-white font-semibold rounded-lg transition-all duration-300 text-sm
                   disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span class="flex items-center justify-center gap-2">
              <svg v-if="loading" class="w-4 h-4 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
              </svg>
              <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"/>
              </svg>
              {{ loading ? 'Menyimpan...' : submitText }}
            </span>
          </button>
          
          <button
            type="button"
            @click="$emit('cancel')"
            class="flex-1 px-4 py-3 bg-gradient-to-r from-gray-100 to-gray-200 hover:from-gray-200 hover:to-gray-300 
                   dark:bg-gradient-to-r dark:from-gray-700 dark:to-gray-800 dark:hover:from-gray-600 dark:hover:to-gray-700
                   text-gray-700 dark:text-gray-300 font-semibold rounded-lg transition-all duration-300 text-sm"
          >
            <span class="flex items-center justify-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
              Batal
            </span>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch, onMounted, onUnmounted } from "vue"

const props = defineProps({
  modelValue: { type: Object, default: () => ({}) },
  categories: { type: Array, default: () => [] },
  submitText: { type: String, default: "Simpan" },
  loading: { type: Boolean, default: false }
})

const emit = defineEmits(["save", "cancel"])

const form = reactive({
  title: "",
  content: "",
  category_id: "",
  cover_url: ""
})

const preview = ref(null)
const newFile = ref(null)
const fileInput = ref(null)
const isMobile = ref(false)

// Deteksi ukuran layar
const checkScreenSize = () => {
  isMobile.value = window.innerWidth < 640 // sm breakpoint
}

// Prefill saat EDIT
watch(
  () => props.modelValue,
  (v) => {
    if (!v) return

    form.title = v.title || ""
    form.content = v.content || ""
    form.category_id = v.category_id ? Number(v.category_id) : ""
    form.cover_url = v.cover || v.cover_url || ""
    preview.value = null
    newFile.value = null
  },
  { immediate: true }
)

// Fungsi untuk trigger file input
function triggerFileInput() {
  if (fileInput.value) {
    fileInput.value.click()
  }
}

function onFileChange(e) {
  const file = e.target.files[0]
  if (!file) return

  if (!file.type.startsWith("image/")) {
    alert("File harus gambar (JPG, PNG, GIF)")
    return
  }
  
  if (file.size > 5 * 1024 * 1024) {
    alert("Ukuran file maksimal 5MB")
    return
  }

  newFile.value = file
  preview.value = URL.createObjectURL(file)
}

function clearImage() {
  newFile.value = null
  preview.value = null
  form.cover_url = ""
  
  // Reset file input juga
  if (fileInput.value) {
    fileInput.value.value = ""
  }
}

function submitForm() {
  emit("save", {
    data: {
      title: form.title,
      content: form.content,
      category_id: Number(form.category_id)
    },
    file: newFile.value
  })
}

onMounted(() => {
  checkScreenSize()
  window.addEventListener('resize', checkScreenSize)
})

onUnmounted(() => {
  window.removeEventListener('resize', checkScreenSize)
})
</script>

<style scoped>
/* Responsive touch targets */
@media (max-width: 640px) {
  button, 
  [role="button"],
  .cursor-pointer {
    min-height: 44px;
    min-width: 44px;
  }
  
  input,
  select,
  textarea {
    font-size: 16px; /* Prevents iOS zoom on focus */
  }
}

/* Adjust padding for mobile bottom bar */
.fixed {
  padding-bottom: env(safe-area-inset-bottom);
}
</style>