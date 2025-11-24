<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-8">
    <h1 class="text-3xl font-bold text-gray-800 dark:text-white mb-6">Admin Dashboard</h1>

    <p class="text-gray-700 dark:text-gray-300 mb-4">
      Selamat datang, <span class="font-semibold">{{ userName }}</span>!
    </p>

    <div class="bg-white dark:bg-gray-800 p-6 rounded-2xl shadow-md">
      <h2 class="text-xl font-bold mb-4 text-gray-800 dark:text-white">Fitur Admin</h2>
      <ul class="list-disc pl-5 space-y-2 text-gray-700 dark:text-gray-300">
        <li>Kelola Berita</li>
        <li>Kelola User</li>
        <li>Statistik & Laporan</li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useAuth } from '~/composables/useAuth'

const auth = useAuth()

const userName = computed(() => auth.user.value?.name || 'Admin')
</script>
