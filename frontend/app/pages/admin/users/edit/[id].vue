<template>
  <div class="min-h-screen bg-gradient-to-br from-indigo-50 via-white to-purple-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 py-8">
    <div class="max-w-xl mx-auto px-4 sm:px-6 lg:px-8">
      
      <!-- Header dengan gradient menarik -->
      <div class="mb-10 relative">
        <div class="absolute -inset-4 bg-gradient-to-r from-indigo-500/10 to-purple-500/10 dark:from-indigo-500/5 dark:to-purple-500/5 rounded-2xl blur-xl"></div>
        <div class="relative bg-gradient-to-r from-indigo-600 to-purple-600 rounded-2xl p-6 shadow-lg">
          <div class="flex items-center gap-4">
            <button 
              @click="goBack"
              class="p-2 bg-white/20 hover:bg-white/30 backdrop-blur-sm text-white 
                     rounded-xl transition-all duration-300 hover:scale-105"
              aria-label="Kembali"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
              </svg>
            </button>
            <div class="flex-1">
              <h1 class="text-2xl md:text-3xl font-bold text-white drop-shadow-md flex items-center gap-3">
                <div class="p-2 bg-white/20 backdrop-blur-sm rounded-lg">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                  </svg>
                </div>
                Edit User
              </h1>
              <p class="text-indigo-100 mt-2">
                Perbarui informasi pengguna dengan data terbaru
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Form Card -->
      <div class="relative group">
        <!-- Background effect -->
        <div class="absolute -inset-1 bg-gradient-to-r from-indigo-400 to-purple-500 rounded-2xl blur opacity-20 group-hover:opacity-30 transition duration-1000"></div>
        
        <!-- Main Card -->
        <div class="relative bg-gradient-to-br from-white to-indigo-50 dark:from-gray-800 dark:to-gray-900 
                    rounded-2xl shadow-xl border border-indigo-100 dark:border-gray-700 overflow-hidden">
          
          <!-- Card Header -->
          <div class="bg-gradient-to-r from-indigo-500/10 to-purple-500/10 dark:from-indigo-500/5 dark:to-purple-500/5 
                      px-6 py-5 border-b border-indigo-100 dark:border-gray-700">
            <div class="flex items-center gap-3">
              <div class="p-2 bg-gradient-to-r from-indigo-500 to-purple-600 rounded-lg">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                </svg>
              </div>
              <div>
                <h2 class="text-lg font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
                  Formulir Pengguna
                </h2>
                <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">
                  Isi semua informasi yang diperlukan
                </p>
              </div>
            </div>
          </div>

          <!-- Form Content -->
          <div class="p-6 md:p-8">
            <!-- Error Alert -->
            <div v-if="error" class="mb-6 p-4 bg-gradient-to-r from-red-50 to-orange-50 dark:from-red-900/20 dark:to-orange-900/20 
                         border-l-4 border-red-500 rounded-r-lg">
              <div class="flex items-start gap-3">
                <svg class="w-5 h-5 text-red-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                <p class="text-sm text-red-700 dark:text-red-300">{{ error }}</p>
              </div>
            </div>

            <form @submit.prevent="save" class="space-y-6">
              <!-- Nama -->
              <div class="space-y-2">
                <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 flex items-center gap-2">
                  <div class="w-2 h-5 bg-gradient-to-b from-indigo-500 to-purple-500 rounded-full"></div>
                  Nama Lengkap <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="name" 
                  type="text" 
                  class="w-full px-4 py-3 rounded-lg border-2 border-gray-200 dark:border-gray-700 
                         bg-white dark:bg-gray-800 text-gray-900 dark:text-white
                         focus:border-indigo-500 focus:ring-4 focus:ring-indigo-500/20 focus:outline-none
                         transition-all duration-300 placeholder-gray-400 dark:placeholder-gray-500
                         hover:border-indigo-300 dark:hover:border-indigo-500"
                  placeholder="Masukkan nama lengkap pengguna"
                  required
                />
              </div>

              <!-- Email -->
              <div class="space-y-2">
                <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 flex items-center gap-2">
                  <div class="w-2 h-5 bg-gradient-to-b from-blue-500 to-cyan-500 rounded-full"></div>
                  Alamat Email <span class="text-red-500">*</span>
                </label>
                <input 
                  v-model="email" 
                  type="email" 
                  class="w-full px-4 py-3 rounded-lg border-2 border-gray-200 dark:border-gray-700 
                         bg-white dark:bg-gray-800 text-gray-900 dark:text-white
                         focus:border-blue-500 focus:ring-4 focus:ring-blue-500/20 focus:outline-none
                         transition-all duration-300 placeholder-gray-400 dark:placeholder-gray-500
                         hover:border-blue-300 dark:hover:border-blue-500"
                  placeholder="Masukkan alamat email aktif"
                  required
                />
              </div>

              <!-- Role -->
              <div class="space-y-2">
                <label class="block text-sm font-semibold text-gray-700 dark:text-gray-300 flex items-center gap-2">
                  <div class="w-2 h-5 bg-gradient-to-b from-emerald-500 to-teal-500 rounded-full"></div>
                  Role Pengguna <span class="text-red-500">*</span>
                </label>
                <div class="relative">
                  <select 
                    v-model="role" 
                    class="w-full px-4 py-3 rounded-lg border-2 border-gray-200 dark:border-gray-700 
                           bg-white dark:bg-gray-800 text-gray-900 dark:text-white
                           focus:border-emerald-500 focus:ring-4 focus:ring-emerald-500/20 focus:outline-none
                           transition-all duration-300 appearance-none
                           hover:border-emerald-300 dark:hover:border-emerald-500"
                  >
                    <option value="user" class="text-gray-900 dark:text-white">User</option>
                    <option value="admin" class="text-gray-900 dark:text-white">Admin</option>
                  </select>
                  <div class="absolute right-4 top-1/2 transform -translate-y-1/2 pointer-events-none">
                    <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
                    </svg>
                  </div>
                </div>
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-2">
                  Pilih peran untuk menentukan hak akses pengguna
                </p>
              </div>

              <!-- Action Buttons -->
              <div class="flex flex-col sm:flex-row gap-4 pt-8 border-t border-gray-100 dark:border-gray-700">
                <button
                  type="button"
                  @click="goBack"
                  class="flex-1 px-6 py-3.5 bg-gradient-to-r from-gray-100 to-gray-200 hover:from-gray-200 hover:to-gray-300 
                         dark:bg-gradient-to-r dark:from-gray-700 dark:to-gray-800 dark:hover:from-gray-600 dark:hover:to-gray-700
                         text-gray-700 dark:text-gray-300 font-semibold rounded-lg transition-all duration-300
                         focus:outline-none focus:ring-4 focus:ring-gray-400/20
                         transform hover:-translate-y-0.5 active:translate-y-0"
                >
                  <span class="flex items-center justify-center gap-3">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
                    </svg>
                    Kembali
                  </span>
                </button>
                
                <button
                  type="submit"
                  :disabled="loading"
                  class="flex-1 px-6 py-3.5 bg-gradient-to-r from-indigo-500 to-purple-600 hover:from-indigo-600 hover:to-purple-700
                         text-white font-semibold rounded-lg transition-all duration-300
                         focus:outline-none focus:ring-4 focus:ring-purple-500/30
                         disabled:opacity-50 disabled:cursor-not-allowed
                         transform hover:-translate-y-0.5 active:translate-y-0
                         shadow-lg hover:shadow-xl"
                >
                  <span class="flex items-center justify-center gap-3">
                    <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                    </svg>
                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                    </svg>
                    {{ loading ? 'Menyimpan...' : 'Simpan Perubahan' }}
                  </span>
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <!-- Info Footer -->
      <div class="mt-8 text-center">
        <p class="text-sm text-gray-500 dark:text-gray-400">
          <span class="inline-flex items-center gap-2">
            <svg class="w-4 h-4 text-emerald-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            Pastikan informasi pengguna sudah benar sebelum disimpan
          </span>
        </p>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()
const id = route.params.id as string

const name = ref('')
const email = ref('')
const role = ref('user')
const loading = ref(false)
const error = ref('')

interface User {
  name: string
  email: string
  role: string
}

// Fetch user data saat halaman mount
onMounted(async () => {
  try {
    const data = await $fetch<User>(`http://localhost:8080/api/users/${id}`)
    name.value = data.name
    email.value = data.email
    role.value = data.role
  } catch (err) {
    console.error('Gagal fetch user:', err)
    error.value = 'Gagal mengambil data user.'
  }
})

// Fungsi simpan perubahan
const save = async () => {
  loading.value = true
  error.value = ''
  try {
    await $fetch(`http://localhost:8080/api/users/${id}`, {
      method: 'PUT',
      body: {
        name: name.value,
        email: email.value,
        role: role.value
      }
    })

    alert('Berhasil disimpan!')
    router.push('/admin/users') // redirect ke halaman list users

  } catch (err) {
    console.error('Gagal menyimpan user:', err)
    error.value = 'Gagal menyimpan user. Coba lagi.'
  } finally {
    loading.value = false
  }
}

// Fungsi kembali
const goBack = () => {
  router.back()
}
</script>

<style scoped>
/* Animation untuk focus state */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fadeIn 0.3s ease-out;
}
</style>