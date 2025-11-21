<template>
  <div class="flex min-h-screen bg-gray-100 dark:bg-gray-900">

    <!-- SIDEBAR -->
    <aside class="w-64 bg-white dark:bg-gray-800 shadow-md px-4 py-6">
      <h2 class="text-xl font-bold mb-6 dark:text-white">Edunews Admin</h2>

      <nav class="space-y-2">
        <NuxtLink
          to="/admin/berita"
          class="block px-3 py-2 rounded-md hover:bg-gray-200 dark:hover:bg-gray-700 dark:text-white"
        >
          📄 Manajemen Berita
        </NuxtLink>

        <NuxtLink
          to="/admin/kategori"
          class="block px-3 py-2 rounded-md hover:bg-gray-200 dark:hover:bg-gray-700 dark:text-white"
        >
          🗂️ Manajemen Kategori
        </NuxtLink>

        <NuxtLink
          to="/admin/users"
          class="block px-3 py-2 rounded-md hover:bg-gray-200 dark:hover:bg-gray-700 dark:text-white"
        >
          👤 Manajemen User
        </NuxtLink>
      </nav>
    </aside>

    <!-- MAIN CONTENT -->
    <main class="flex-1 p-8">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-2xl font-bold dark:text-white">
          <slot name="title"></slot>
        </h1>

        <!-- Theme Toggle -->
        <button
          @click="toggleTheme"
          class="px-3 py-1 rounded-lg border dark:text-white"
        >
          Toggle Theme
        </button>
      </div>

      <div class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow">
        <slot />
      </div>
    </main>

  </div>
</template>

<script setup>
import { useTheme } from "~/composables/useTheme"
const { toggleTheme } = useTheme()
</script>
