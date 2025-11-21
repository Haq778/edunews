<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Daftar Berita</h1>
      <NuxtLink to="/berita/create" class="px-4 py-2 rounded bg-green-600 text-white">
        Buat Berita
      </NuxtLink>
    </div>

    <div v-if="loading" class="text-center py-8">Loading...</div>

    <div v-else>
      <AdminTable :items="articles" @delete="handleDelete" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import AdminTable from '~/components/AdminTable.vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const articles = ref<any[]>([])
const loading = ref(true)
const router = useRouter()

async function fetchAll() {
  loading.value = true
  try {
    const res = await $fetch('/api/berita')
    articles.value = res || []
  } catch (e) {
    console.error(e)
    articles.value = []
  } finally {
    loading.value = false
  }
}

async function handleDelete(id: string | number) {
  if (!confirm('Yakin ingin menghapus berita ini?')) return
  await $fetch(`/api/berita/${id}`, { method: 'DELETE' })
  await fetchAll()
}

onMounted(fetchAll)
</script>
