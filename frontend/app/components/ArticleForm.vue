<template>
  <form @submit.prevent="submitForm" class="space-y-4">

    <!-- Judul -->
    <div>
      <label class="block font-medium mb-1">Judul</label>
      <input
        v-model="form.title"
        type="text"
        class="w-full border px-3 py-2 rounded"
        required
      />
    </div>

    <!-- Kategori -->
    <div>
      <label class="block font-medium mb-1">Kategori</label>
      <select
        v-model="form.category"
        class="w-full border px-3 py-2 rounded"
        required
      >
        <option disabled value="">-- pilih kategori --</option>
        <option value="pendidikan">Pendidikan</option>
        <option value="informasi">Informasi</option>
        <option value="pengumuman">Pengumuman</option>
      </select>
    </div>

    <!-- Isi -->
    <div>
      <label class="block font-medium mb-1">Isi</label>
      <textarea
        v-model="form.content"
        rows="6"
        class="w-full border px-3 py-2 rounded"
        required
      ></textarea>
    </div>

    <!-- Cover -->
    <div>
      <label class="block font-medium mb-1">Cover</label>

      <!-- Preview cover lama -->
      <div v-if="form.cover && !newFile" class="mb-2">
        <img
          :src="form.cover"
          alt="cover lama"
          class="w-48 rounded shadow"
        />
      </div>

      <!-- Preview cover baru -->
      <div v-if="preview" class="mb-2">
        <img
          :src="preview"
          class="w-48 rounded shadow"
        />
      </div>

      <input type="file" @change="onFileChange" />
    </div>

    <!-- Tombol -->
    <div class="flex gap-3">
      <button
        type="submit"
        class="px-4 py-2 bg-blue-600 text-white rounded"
      >
        {{ submitText }}
      </button>

      <button
        type="button"
        @click="$emit('cancel')"
        class="px-4 py-2 bg-gray-300 rounded"
      >
        Batal
      </button>
    </div>
  </form>
</template>


<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Object, default: () => ({}) },
  submitText: { type: String, default: 'Simpan' }
})

const emit = defineEmits(['save', 'cancel'])

const form = ref({
  title: '',
  category: '',
  content: '',
  cover: ''  // URL cover lama
})

const newFile = ref(null)
const preview = ref(null)

watch(
  () => props.modelValue,
  (val) => {
    // isi form berdasarkan data edit
    form.value = {
      title: val.title || '',
      category: val.category || '',
      content: val.content || '',
      cover: val.cover || ''
    }
  },
  { immediate: true }
)

function onFileChange(e) {
  const file = e.target.files[0]
  if (!file) return

  newFile.value = file
  preview.value = URL.createObjectURL(file)
}

function submitForm() {
  emit('save', {
    data: {
      title: form.value.title,
      category: form.value.category,
      content: form.value.content
    },
    file: newFile.value
  })
}
</script>
