<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8">
    <div class="max-w-2xl mx-auto px-4 sm:px-6 lg:px-8">
      <!-- Header Section -->
      <div class="mb-8">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
              Tambah Kategori Baru
            </h1>
            <p class="text-gray-600 dark:text-gray-400">
              Buat kategori baru untuk mengorganisir berita pendidikan
            </p>
          </div>
          
          <NuxtLink 
            to="/admin/kategori"
            class="flex items-center space-x-2 px-4 py-2 text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200 transition-colors duration-200 bg-white dark:bg-gray-800 rounded-xl border border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
            </svg>
            <span>Kembali</span>
          </NuxtLink>
        </div>
      </div>

      <!-- Form Card -->
      <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border border-gray-200 dark:border-gray-700 overflow-hidden">
        <!-- Card Header -->
        <div class="bg-gradient-to-r from-green-600 to-emerald-700 px-6 py-4">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center">
              <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
              </svg>
            </div>
            <div>
              <h2 class="text-xl font-semibold text-white">Form Kategori Baru</h2>
              <p class="text-green-100 text-sm">Isi informasi kategori di bawah ini</p>
            </div>
          </div>
        </div>

        <!-- Form Content -->
        <form @submit.prevent="submit" class="p-6 space-y-6">
          <!-- Nama Kategori Field -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3 flex items-center">
              <svg class="w-5 h-5 mr-2 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z"/>
              </svg>
              Nama Kategori *
            </label>
            <input 
              v-model="form.name" 
              type="text" 
              placeholder="Masukkan nama kategori"
              class="w-full p-4 rounded-xl border border-gray-300 dark:border-gray-600 focus:outline-none focus:ring-3 focus:ring-green-500/20 focus:border-green-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white transition-all duration-300"
              :class="{ 'border-red-500 dark:border-red-400': errors.name }"
            />
            <p v-if="errors.name" class="text-sm text-red-500 dark:text-red-400 mt-2 flex items-center">
              <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L4.082 16.5c-.77.833.192 2.5 1.732 2.5z"/>
              </svg>
              {{ errors.name }}
            </p>
            <p class="text-xs text-gray-500 dark:text-gray-400 mt-2">
              Nama kategori harus unik dan deskriptif
            </p>
          </div>

          <!-- Slug Field -->
          <div>
            <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 mb-3 flex items-center">
              <svg class="w-5 h-5 mr-2 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"/>
              </svg>
              Slug (URL)
            </label>
            <input 
              v-model="form.slug" 
              type="text" 
              placeholder="Slug akan dibuat otomatis"
              class="w-full p-4 rounded-xl border border-gray-300 dark:border-gray-600 focus:outline-none focus:ring-3 focus:ring-blue-500/20 focus:border-blue-500 bg-white dark:bg-gray-700 text-gray-900 dark:text-white transition-all duration-300"
            />
            <p class="text-xs text-gray-500 dark:text-gray-400 mt-2">
              Slug digunakan untuk URL. Jika kosong, akan dibuat otomatis dari nama kategori.
            </p>
            
            <!-- Auto-generated Slug Preview -->
            <div v-if="form.name && !form.slug" class="mt-3 p-3 bg-blue-50 dark:bg-blue-900/20 rounded-xl border border-blue-200 dark:border-blue-800">
              <p class="text-sm text-blue-700 dark:text-blue-300 flex items-center">
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <span class="font-medium">Slug otomatis:</span>
                <code class="ml-2 bg-white dark:bg-blue-800 px-2 py-1 rounded text-blue-800 dark:text-blue-200 font-mono">
                  {{ slugify(form.name) }}
                </code>
              </p>
            </div>
          </div>

          <!-- Action Buttons -->
          <div class="flex flex-col sm:flex-row gap-4 pt-6 border-t border-gray-200 dark:border-gray-700">
            <button 
              type="submit"
              :disabled="saving"
              class="flex-1 bg-gradient-to-r from-green-600 to-emerald-700 text-white py-4 px-6 rounded-xl font-semibold hover:shadow-lg transition-all duration-300 transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none flex items-center justify-center space-x-2"
            >
              <span v-if="!saving">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
              </span>
              <span v-else class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
              <span>{{ saving ? 'Menyimpan...' : 'Simpan Kategori' }}</span>
            </button>
            
            <NuxtLink 
              to="/admin/kategori"
              class="flex-1 py-4 px-6 text-center border border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 rounded-xl font-semibold hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors duration-300 flex items-center justify-center space-x-2"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
              <span>Batal</span>
            </NuxtLink>
          </div>
        </form>
      </div>

      <!-- Success Notification -->
      <div 
        v-if="showSuccess" 
        class="fixed top-4 right-4 bg-green-500 text-white px-6 py-4 rounded-2xl shadow-2xl z-50 transform transition-all duration-500 flex items-center space-x-3"
      >
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <div>
          <p class="font-semibold">Sukses!</p>
          <p class="text-sm opacity-90">Kategori berhasil ditambahkan</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const router = useRouter()
const form = ref({ name: '', slug: '' })
const errors = ref<{ name?: string }>({})
const saving = ref(false)
const showSuccess = ref(false)

// Auto-generate slug ketika nama berubah
watch(() => form.value.name, (newName) => {
  if (newName && !form.value.slug) {
    // Hanya generate slug jika user belum mengisi slug manual
    form.value.slug = slugify(newName)
  }
})

function slugify(text: string) {
  return text
    .toString()
    .trim()
    .toLowerCase()
    .replace(/\s+/g, '-')
    .replace(/[^\w\-]+/g, '')
    .replace(/\-\-+/g, '-')
    .replace(/^-+/, '')
    .replace(/-+$/, '')
}

async function submit() {
  // Reset errors
  errors.value = {}
  
  // Validation
  if (!form.value.name || form.value.name.trim().length < 2) {
    errors.value.name = 'Nama kategori minimal 2 karakter'
    return
  }

  if (form.value.name.trim().length > 50) {
    errors.value.name = 'Nama kategori maksimal 50 karakter'
    return
  }

  saving.value = true

  try {
    const payload = {
      name: form.value.name.trim(),
      slug: form.value.slug?.trim() || slugify(form.value.name)
    }

    const response = await $fetch('http://localhost:8080/api/categories', {
      method: 'POST',
      body: payload,
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      }
    })

    // Show success notification
    showSuccess.value = true
    
    // Redirect setelah delay
    setTimeout(() => {
      router.push('/admin/kategori')
    }, 1500)

  } catch (error: any) {
    console.error('❌ Gagal menyimpan kategori:', error)
    
    // Handle specific error cases
    if (error?.status === 409) {
      errors.value.name = 'Kategori dengan nama ini sudah ada'
    } else if (error?.status === 400) {
      errors.value.name = 'Data yang dikirim tidak valid'
    } else {
      errors.value.name = 'Terjadi kesalahan saat menyimpan kategori'
    }
  } finally {
    saving.value = false
  }
}
</script>

<style scoped>
/* Custom animations */
.fixed {
  transform: translateX(100%);
  opacity: 0;
}

.fixed.transform {
  transform: translateX(0);
  opacity: 1;
}

/* Smooth transitions */
* {
  transition: all 0.3s ease;
}

/* Custom focus styles */
input:focus {
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

/* Dark mode adjustments */
.dark input:focus {
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.2);
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
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

.dark ::-webkit-scrollbar-track {
  background: #374151;
}

.dark ::-webkit-scrollbar-thumb {
  background: #6b7280;
}

.dark ::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}
</style>