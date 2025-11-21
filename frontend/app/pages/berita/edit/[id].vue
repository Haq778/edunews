<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Edit Berita</h1>
      <NuxtLink to="/berita" class="px-3 py-2 border rounded">Kembali</NuxtLink>
    </div>

    <div v-if="loading" class="text-center py-8">Loading...</div>

    <div v-else class="bg-white p-6 rounded shadow">
      <ArticleForm :modelValue="article" submitText="Simpan" @save="update" @cancel="goBack" />
      <div class="mt-4">
        <button @click="remove" class="px-3 py-2 rounded border text-red-600">Hapus Berita</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ArticleForm from '~/components/ArticleForm.vue'
import { useRoute, useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()
const article = ref<any>(null)
const loading = ref(true)

async function fetchOne() {
  loading.value = true
  const id = route.params.id as string
  try {
    article.value = await $fetch(`/api/berita/${id}`)
  } catch (e) {
    article.value = null
  } finally {
    loading.value = false
  }
}

async function update(payload: any) {
  const id = route.params.id as string
  // Build FormData if file provided
  if (payload.file) {
    const fd = new FormData()
    Object.entries(payload.data || {}).forEach(([k,v]) => {
      if (v !== undefined && v !== null) fd.append(k, String(v))
    })
    fd.append('cover', payload.file)
    await $fetch(`/api/berita/${id}`, { method: 'PUT', body: fd })
  } else {
    // send JSON
    await $fetch(`/api/berita/${id}`, { method: 'PUT', body: payload.data })
  }
  router.push('/berita')
}

async function remove() {
  if (!confirm('Yakin hapus berita ini?')) return
  const id = route.params.id as string
  await $fetch(`/api/berita/${id}`, { method: 'DELETE' })
  router.push('/berita')
}

function goBack(){ router.back() }

onMounted(fetchOne)
</script>
