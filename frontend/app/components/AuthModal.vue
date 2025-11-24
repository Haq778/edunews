<template>
  <ClientOnly>
    <div class="fixed inset-0 bg-black/50 flex justify-center items-center z-50">
      <div class="relative bg-white dark:bg-gray-800 rounded-2xl p-8 w-full max-w-md">
        <!-- Tombol Close -->
        <button
          type="button"
          class="absolute top-4 right-4 text-gray-500 hover:text-gray-700 dark:hover:text-white"
          @click="close"
        >
          ✕
        </button>

        <h2 class="text-2xl font-bold mb-4">{{ tab === 'login' ? 'Login' : 'Register' }}</h2>

        <form @submit.prevent="tab === 'login' ? login() : register()">
          <div v-if="tab === 'register'" class="mb-4">
            <input v-model="registerForm.name" type="text" placeholder="Nama" class="input" />
          </div>
          <div class="mb-4">
            <input v-model="form.email" type="email" placeholder="Email" class="input" />
          </div>
          <div class="mb-4">
            <input v-model="form.password" type="password" placeholder="Password" class="input" />
          </div>
          <p v-if="errorMessage" class="text-red-500 mb-2">{{ errorMessage }}</p>
          <button type="submit" class="btn w-full">{{ tab === 'login' ? 'Login' : 'Register' }}</button>
        </form>

        <p class="mt-4 text-center text-gray-500">
          <span v-if="tab === 'login'">
            Belum punya akun? 
            <button type="button" @click="tab='register'" class="text-blue-600">Daftar</button>
          </span>
          <span v-else>
            Sudah punya akun? 
            <button type="button" @click="tab='login'" class="text-blue-600">Login</button>
          </span>
        </p>
      </div>
    </div>
  </ClientOnly>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface User { id: number; name: string; email: string }

const emit = defineEmits<{
  (e: 'login-success', user: User): void
  (e: 'auth-modal-close'): void
}>()

const tab = ref<'login' | 'register'>('login')
const form = ref({ email: '', password: '' })
const registerForm = ref({ name: '', email: '', password: '' })
const loading = ref(false)
const errorMessage = ref('')

// Close modal
const close = () => {
  errorMessage.value = ''
  tab.value = 'login'
  emit('auth-modal-close')
}

// Login
const login = async () => {
  loading.value = true
  errorMessage.value = ''
  try {
    // @ts-ignore
    const res = await $api.post('/login', form.value)
    const user: User = res.data.user
    const token: string = res.data.token
    localStorage.setItem('token', token)
    emit('login-success', user)
    close()
  } catch (err: unknown) {
    errorMessage.value = (err as any)?.response?.data?.message || 'Gagal login'
  } finally {
    loading.value = false
  }
}

// Register
const register = async () => {
  loading.value = true
  errorMessage.value = ''
  try {
    // @ts-ignore
    const res = await $api.post('/register', registerForm.value)
    const user: User = res.data.user
    const token: string = res.data.token
    localStorage.setItem('token', token)
    emit('login-success', user)
    close()
  } catch (err: unknown) {
    errorMessage.value = (err as any)?.response?.data?.message || 'Gagal register'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ccc;
  border-radius: 0.75rem;
  background: white;
  color: black;
}
.btn {
  padding: 0.75rem;
  border-radius: 0.75rem;
  background-color: #3b82f6;
  color: white;
  font-weight: bold;
}
</style>
