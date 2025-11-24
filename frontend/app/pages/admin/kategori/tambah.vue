<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Tambah Kategori</h1>

    <form @submit.prevent="submit">
      <div class="mb-4">
        <label class="block text-sm font-medium mb-1">Nama Kategori</label>
        <input v-model="form.name" type="text" class="w-full p-3 border rounded" />
        <p v-if="errors.name" class="text-sm text-red-500 mt-1">{{ errors.name }}</p>
      </div>

      <div class="mb-6">
        <label class="block text-sm font-medium mb-1">Slug (opsional)</label>
        <input v-model="form.slug" type="text" class="w-full p-3 border rounded" />
        <p class="text-xs text-gray-500 mt-1">Jika kosong, akan dibuat otomatis dari nama.</p>
      </div>

      <div class="flex gap-2">
        <button :disabled="saving" class="px-4 py-2 bg-green-600 text-white rounded">
          {{ saving ? 'Menyimpan...' : 'Simpan' }}
        </button>
        <NuxtLink to="/admin/kategori" class="px-4 py-2 border rounded">Batal</NuxtLink>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const router = useRouter()
const form = ref({ name: '', slug: '' })
const errors = ref<{ name?: string }>({})
const saving = ref(false)

function slugify(s = '') {
  return s
    .toString()
    .trim()
    .toLowerCase()
    .replace(/\s+/g, '-')
    .replace(/[^\w\-]+/g, '')
    .replace(/\-\-+/g, '-')
}

async function submit() {
  errors.value = {}
  if (!form.value.name || form.value.name.trim().length < 2) {
    errors.value.name = 'Nama kategori minimal 2 karakter'
    return
  }

  saving.value = true
  try {
    const payload = {
      name: form.value.name.trim(),
      slug: form.value.slug?.trim() || slugify(form.value.name)
    }

    // Sesuaikan endpoint jika berbeda
    await $fetch('/api/kategori', {
      method: 'POST',
      body: payload
    })

    // kembali ke daftar
    await router.push('/admin/kategori')
  } catch (e: any) {
    console.error('Gagal menyimpan kategori', e)
    alert('Gagal menyimpan kategori. Cek console.')
  } finally {
    saving.value = false
  }
}
</script>
