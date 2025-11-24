<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900">
    <div class="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
      <h2 class="text-2xl font-bold mb-6 text-gray-800 dark:text-white text-center">Login</h2>

      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block mb-2 text-gray-700 dark:text-gray-300">Email</label>
          <input v-model="email" type="email" required
            class="w-full p-3 rounded-xl border border-gray-200 dark:border-gray-600 
                   bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                   focus:ring-2 focus:ring-blue-500"/>
        </div>

        <div>
          <label class="block mb-2 text-gray-700 dark:text-gray-300">Password</label>
          <input v-model="password" type="password" required
            class="w-full p-3 rounded-xl border border-gray-200 dark:border-gray-600 
                   bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                   focus:ring-2 focus:ring-blue-500"/>
        </div>

        <button type="submit" 
                class="w-full py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-500 transition font-semibold">
          Login
        </button>
      </form>

      <p class="mt-4 text-gray-600 dark:text-gray-300 text-center">
        Belum punya akun? 
        <NuxtLink to="/register" class="text-blue-600 dark:text-blue-400 font-semibold hover:underline">Daftar</NuxtLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { User } from '~/composables/useAuth'
import { useAuth } from '~/composables/useAuth'

const email = ref('')
const password = ref('')
const router = useRouter()
const auth = useAuth()

async function handleLogin() {
  try {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: email.value, password: password.value })
    })

    if (!res.ok) throw new Error(`Login gagal (HTTP ${res.status})`)

    const data = await res.json()
    if (!data.token || !data.user) throw new Error('Login gagal: user atau token tidak valid')

    // simpan di state reactive (tanpa localStorage)
    auth.login(data.token, data.user as User)

    // redirect berdasarkan role
    // simpan di state reactive (tanpa localStorage)
    // tunggu jika auth.login mengembalikan Promise
    const loginResult: void | Promise<void> = auth.login(data.token, data.user as User)
    if (typeof loginResult === "object" && typeof (loginResult as Promise<void>)?.then === "function") await loginResult

    // debug: pastikan role sesuai
    // (hapus atau ubah jadi logger yang sesuai di production)
    // eslint-disable-next-line no-console
    console.log('logged user role:', data.user?.role)

    // redirect berdasarkan role
    if (data.user?.role === 'admin') {
      await router.push('/admin/berita')
    } else {
      await router.push('/')
    }
    } catch (err: any) {
    alert(err?.message || 'Login gagal')
    }
}
</script>
