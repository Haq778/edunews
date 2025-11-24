<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Edit Kategori</h1>

    <div v-if="loading" class="text-gray-500">Memuat...</div>
    <div v-else>
      <form @submit.prevent="submit">
        <div class="mb-4">
          <label class="block text-sm font-medium mb-1">Nama Kategori</label>
          <input v-model="form.name" type="text" class="w-full p-3 border rounded" />
          <p v-if="errors.name" class="text-sm text-red-500 mt-1">{{ errors.name }}</p>
        </div>

        <div class="mb-6">
          <label class="block text-sm font-medium mb-1">Slug</label>
          <input v-model="form.slug" type="text" class="w-full p-3 border rounded" />
        </div>

        <div class="flex gap-2">
          <button :disabled="saving" class="px-4 py-2 bg-yellow-600 text-white rounded">
            {{ saving ? 'Menyimpan...' : 'Update' }}
          </button>
          <NuxtLink to="/admin/kategori" class="px-4 py-2 border rounded">Batal</NuxtLink>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()
const id = (route.params.id as string) || ''

const form = ref({ name: '', slug: '' })
const errors = ref<{ name?: string }>({})
const loading = ref(true)
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

async function load() {
  loading.value = true
  try {
    const data: any = await $fetch(`/api/kategori/${id}`)
    form.value.name = data?.name || ''
    form.value.slug = data?.slug || ''
  } catch (e) {
    console.error('Gagal memuat kategori', e)
    alert('Gagal memuat kategori. Cek console.')
    await router.push('/admin/kategori')
  } finally {
    loading.value = false
  }
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

    await $fetch(`/api/kategori/${id}`, {
      method: 'PUT',
      body: payload
    })

    await router.push('/admin/kategori')
  } catch (e) {
    console.error('Gagal update kategori', e)
    alert('Gagal update kategori. Cek console.')
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  if (!id) {
    router.push('/admin/kategori')
    return
  }
  load()
})
</script>
