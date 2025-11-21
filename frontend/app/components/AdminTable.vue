<template>
  <div class="bg-white border rounded">
    <table class="w-full text-left">
      <thead class="bg-gray-50">
        <tr>
          <th class="p-3">Cover</th>
          <th class="p-3">Title</th>
          <th class="p-3">Category</th>
          <th class="p-3">Created</th>
          <th class="p-3 w-48">Actions</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="item in props.items"
          :key="item.id"
          class="border-t"
        >
          <td class="p-3">
            <img
              v-if="item.cover"
              :src="item.cover"
              class="h-16 w-24 object-cover rounded"
            />
            <div v-else class="h-16 w-24 bg-gray-100 rounded flex items-center justify-center text-sm text-gray-500">No</div>
          </td>

          <td class="p-3">
            <div class="font-semibold">{{ item.title }}</div>
            <div class="text-sm text-gray-500 truncate">{{ item.slug }}</div>
          </td>

          <td class="p-3">{{ item.category }}</td>

          <td class="p-3 text-sm text-gray-500">{{ formatDate(item.created_at) }}</td>

          <td class="p-3">
            <div class="flex gap-2">
              <NuxtLink :to="`/berita/${item.id}`" class="px-2 py-1 border rounded">View</NuxtLink>
              <NuxtLink :to="`/admin/berita/edit/${item.id}`" class="px-2 py-1 border rounded">Edit</NuxtLink>
              <button @click="emitDelete(item.id)" class="px-2 py-1 border rounded text-red-600">Delete</button>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
interface BeritaItem {
  id: number | string
  title: string
  slug: string
  category: string
  cover?: string
  created_at?: string | Date
}

const props = defineProps<{
  items: BeritaItem[]
}>()

const emits = defineEmits<{
  (e: 'delete', id: number | string): void
}>()

function emitDelete(id: number | string) {
  emits('delete', id)
}

function formatDate(d: any) {
  try {
    return new Date(d).toLocaleString()
  } catch {
    return d
  }
}
</script>
