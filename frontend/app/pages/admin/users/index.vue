<template>
  <div class="p-6 max-w-4xl mx-auto">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-2xl font-bold">Manajemen Users</h1>
      <!-- Optional: tombol tambah user jika ada endpoint -->
      <!-- <NuxtLink to="/admin/users/tambah" class="px-4 py-2 rounded bg-green-600 text-white">Tambah User</NuxtLink> -->
    </div>

    <div v-if="loading" class="text-gray-500">Memuat users...</div>

    <div v-else>
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-gray-50">
            <th class="p-2">ID</th>
            <th class="p-2">Nama</th>
            <th class="p-2">Email</th>
            <th class="p-2">Role</th>
            <th class="p-2">Dibuat</th>
            <th class="p-2 w-40">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in users" :key="u.id" class="border-t hover:bg-gray-50">
            <td class="p-2">{{ u.id }}</td>
            <td class="p-2">{{ u.name }}</td>
            <td class="p-2">{{ u.email }}</td>
            <td class="p-2">{{ u.role }}</td>
            <td class="p-2 text-sm text-gray-500">{{ formatDate(u.created_at) }}</td>
            <td class="p-2">
              <div class="flex gap-2">
                <NuxtLink :to="`/admin/users/edit/${u.id}`" class="px-2 py-1 border rounded">Edit</NuxtLink>
                <button @click="hapus(u.id)" class="px-2 py-1 border rounded text-red-600">Delete</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="users.length === 0" class="text-gray-500 mt-4">Belum ada user.</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from '#imports'

definePageMeta({ layout: 'admin' })

const users = ref<any[]>([])
const loading = ref(true)
const router = useRouter()

function formatDate(d: any) {
  if (!d) return '-'
  try { return new Date(d).toLocaleString() } catch { return d }
}

async function loadUsers() {
  loading.value = true
  try {
    const res: any = await $fetch('/api/users')
    users.value = Array.isArray(res) ? res : []
  } catch (e) {
    console.error('Gagal memuat users', e)
    alert('Gagal memuat users. Cek console.')
  } finally {
    loading.value = false
  }
}

async function hapus(id: number | string) {
  if (!confirm('Yakin ingin menghapus user ini?')) return
  try {
    await $fetch(`/api/users/${id}`, { method: 'DELETE' })
    await loadUsers()
  } catch (e) {
    console.error('Gagal menghapus user', e)
    alert('Gagal menghapus user. Cek console.')
  }
}

onMounted(loadUsers)
</script>