<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900">
    <div class="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
      <h1 class="text-2xl font-bold mb-6 text-center text-gray-800 dark:text-white">Register</h1>

      <form @submit.prevent="registerUser" class="space-y-4">
        <div>
          <label class="block text-sm text-gray-700 dark:text-gray-300">Nama</label>
          <input v-model="name" type="text" required
            class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 
                   text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-green-500"/>
        </div>

        <div>
          <label class="block text-sm text-gray-700 dark:text-gray-300">Email</label>
          <input v-model="email" type="email" required
            class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 
                   text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-green-500"/>
        </div>

        <div>
          <label class="block text-sm text-gray-700 dark:text-gray-300">Password</label>
          <input v-model="password" type="password" required
            class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 
                   text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-green-500"/>
        </div>

        <button type="submit" 
                class="w-full py-3 bg-green-600 text-white rounded-xl hover:bg-green-500 transition font-semibold">
          Register
        </button>
      </form>

      <p class="mt-4 text-sm text-gray-600 dark:text-gray-400 text-center">
        Sudah punya akun? 
        <NuxtLink to="/login" class="text-blue-600 dark:text-blue-400 hover:underline">Login</NuxtLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const name = ref('')
const email = ref('')
const password = ref('')
const router = useRouter()

async function registerUser() {
  try {
    const res = await fetch('http://localhost:8080/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: name.value, email: email.value, password: password.value })
    })
    if (!res.ok) throw new Error('Registrasi gagal')

    const data = await res.json()
    localStorage.setItem('token', data.token)
    localStorage.setItem('user', JSON.stringify(data.user))

    router.push('/') // redirect ke homepage
  } catch (err: any) {
    alert(err.message)
  }
}
</script>
