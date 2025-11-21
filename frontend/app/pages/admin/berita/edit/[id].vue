<template>
  <div class="max-w-3xl mx-auto py-10">
    <h1 class="text-2xl font-bold mb-6">Edit Berita</h1>

    <div v-if="loading" class="p-6 bg-white rounded shadow">Loading...</div>

    <div v-else class="bg-white p-6 rounded shadow">
      <ArticleForm
        :modelValue="article || {}"
        submitText="Simpan"
        @save="update"
        @cancel="goBack"
      />

      <div class="mt-4">
        <button @click="remove" class="px-3 py-2 rounded border text-red-600">
          Hapus Berita
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ArticleForm from '~/components/ArticleForm.vue'
import { useRoute, useRouter } from '#imports'
import { ref, watchEffect } from 'vue'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()
const article = ref<any>(null)
const loading = ref(true)

// helper: parse id safely
function parseId(raw: any) {
  const n = Number(raw)
  return Number.isFinite(n) && n > 0 ? n : null
}

// watch route.params.id dan baru fetch kalau valid
watchEffect(async (onInvalidate) => {
  loading.value = true
  article.value = null

  const id = parseId(route.params.id)
  if (!id) {
    // jika id belum tersedia / tidak valid, tunggu (don't call API)
    loading.value = false
    return
  }

  let aborted = false
  onInvalidate(() => { aborted = true })

  try {
    const res = await fetch(`http://localhost:8080/api/berita/${id}`)
    if (!res.ok) {
      const txt = await res.text().catch(()=> 'fetch error')
      throw new Error(txt || `HTTP ${res.status}`)
    }
    const data = await res.json()
    if (!aborted) article.value = data
  } catch (err) {
    console.error('Gagal memuat berita:', err)
    article.value = null
  } finally {
    if (!aborted) loading.value = false
  }
})

// update handler
async function update(payload: any) {
  const id = parseId(route.params.id)
  if (!id) return alert('ID tidak valid')

  try {
    if (payload.file) {
      const fd = new FormData()
      Object.entries(payload.data || {}).forEach(([k, v]) => {
        if (v !== undefined && v !== null) fd.append(k, String(v))
      })
      fd.append('cover', payload.file)
      const res = await fetch(`http://localhost:8080/api/berita/${id}`, {
        method: 'PUT',
        body: fd
      })
      if (!res.ok) throw new Error(await res.text())
    } else {
      const res = await fetch(`http://localhost:8080/api/berita/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload.data)
      })
      if (!res.ok) throw new Error(await res.text())
    }

    alert('Berita berhasil disimpan')
    router.push('/admin/berita')
  } catch (err) {
    console.error('Gagal update berita:', err)
    alert('Gagal update, cek console')
  }
}

// delete handler
async function remove() {
  if (!confirm('Yakin hapus berita ini?')) return
  const id = parseId(route.params.id)
  if (!id) return alert('ID tidak valid')

  try {
    const res = await fetch(`http://localhost:8080/api/berita/${id}`, { method: 'DELETE' })
    if (!res.ok) throw new Error(await res.text())
    alert('Berita berhasil dihapus')
    router.push('/admin/berita')
  } catch (err) {
    console.error('Gagal hapus berita:', err)
    alert('Gagal hapus, cek console')
  }
}

function goBack(){ router.back() }
</script>
