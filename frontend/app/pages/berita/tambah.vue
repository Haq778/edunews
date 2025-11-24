<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Buat Berita</h1>
      <NuxtLink to="/berita" class="px-3 py-2 border rounded">Kembali</NuxtLink>
    </div>

    <div class="bg-white p-6 rounded shadow">
      <ArticleForm submitText="Buat" @save="store" @cancel="goBack" />
    </div>
  </div>
</template>

<script setup lang="ts">
import ArticleForm from '~/components/ArticleForm.vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const router = useRouter()

async function store(payload: any) {
  // payload = { data: {...}, file: File|null }
  const fd = new FormData()
  Object.entries(payload.data || {}).forEach(([k,v]) => {
    if (v !== undefined && v !== null) fd.append(k, String(v))
  })
  if (payload.file) {
    fd.append('cover', payload.file)
  }
  await $fetch('/api/berita', { method: 'POST', body: fd })
  router.push('/berita')
}

function goBack(){ router.back() }
</script>
