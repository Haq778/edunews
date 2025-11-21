<template>
  <div class="max-w-3xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Tambah Berita</h1>

    <div class="bg-white p-6 rounded shadow">
      <ArticleForm
        :modelValue="{}"
        submitText="Tambah"
        @save="save"
        @cancel="goBack"
      />
    </div>
  </div>
</template>


<script setup lang="ts">
import ArticleForm from '~/components/ArticleForm.vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const router = useRouter()

async function save(payload: any) {
  try {
    if (payload.file) {
      const fd = new FormData()
      Object.entries(payload.data || {}).forEach(([k, v]) => {
        if (v !== undefined && v !== null) fd.append(k, String(v))
      })
      fd.append('cover', payload.file)

      const res = await fetch('http://localhost:8080/api/berita', {
        method: 'POST',
        body: fd,
      })
      if (!res.ok) throw new Error(await res.text())
    } else {
      const res = await fetch('http://localhost:8080/api/berita', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload.data),
      })
      if (!res.ok) throw new Error(await res.text())
    }

    router.push('/admin/berita')
  } catch (err) {
    console.error(err)
    alert('Gagal menyimpan berita')
  }
}

function goBack() {
  router.back()
}
</script>


