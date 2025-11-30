<template>
  <div class="min-h-screen bg-gradient-to-br from-gray-50 to-blue-50/20 dark:from-gray-900 dark:to-gray-800 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">

      <!-- Header -->
      <div class="mb-8 flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
            {{ submitText }}
          </h1>
          <p class="text-gray-600 dark:text-gray-400">
            Perbarui informasi berita dengan data terbaru
          </p>
        </div>

        <button
          @click="$emit('cancel')"
          class="flex items-center space-x-2 px-4 py-2.5 text-gray-600 dark:text-gray-400
                 hover:text-gray-900 dark:hover:text-white transition-colors duration-300
                 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"/>
          </svg>
          <span>Tutup</span>
        </button>
      </div>

      <!-- Card -->
      <div class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg border border-gray-200 dark:border-gray-700 overflow-hidden">

        <!-- Card Header -->
        <div class="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-gray-700 dark:to-gray-800 px-6 py-4 border-b">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
            Formulir Berita
          </h2>
        </div>

        <!-- FORM -->
        <form @submit.prevent="submitForm" class="p-6 space-y-6">

          <!-- Judul -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
              Judul Berita <span class="text-red-500">*</span>
            </label>
            <input
              v-model="form.title"
              type="text"
              class="w-full px-4 py-3 rounded-xl border bg-white dark:bg-gray-700"
              placeholder="Masukkan judul..."
              required
            />
          </div>

          <!-- Kategori -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
              Kategori <span class="text-red-500">*</span>
            </label>

            <select
              v-model="form.category_id"
              required
              class="w-full px-4 py-3 rounded-xl border bg-white dark:bg-gray-700">
              <option value="" disabled>-- Pilih Kategori --</option>
              <option
                v-for="kat in categories"
                :key="kat.id"
                :value="kat.id">
                {{ kat.name }}
              </option>
            </select>
          </div>

          <!-- Konten -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
              Konten Berita <span class="text-red-500">*</span>
            </label>
            <textarea
              v-model="form.content"
              rows="7"
              class="w-full px-4 py-3 rounded-xl border bg-white dark:bg-gray-700 resize-vertical"
              placeholder="Tulis konten berita..."
              required
            ></textarea>
          </div>

          <!-- Cover -->
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
              Cover Berita
            </label>

            <!-- Cover Lama -->
            <div
              v-if="form.cover_url && !preview"
              class="flex items-center space-x-4 p-4 bg-gray-50 dark:bg-gray-700 rounded-xl mb-4">
              <img :src="form.cover_url" class="w-24 h-24 rounded-lg object-cover" />
              <p class="text-gray-500">Cover saat ini</p>
            </div>

            <!-- Preview Baru -->
            <div
              v-if="preview"
              class="flex items-center space-x-4 p-4 bg-blue-50 dark:bg-blue-900/20 border rounded-xl mb-4">
              <img :src="preview" class="w-24 h-24 rounded-lg object-cover" />
              <p class="text-blue-600 dark:text-blue-300">Preview cover baru</p>
            </div>

            <div class="flex items-center gap-4">
              <input
                type="file"
                accept="image/*"
                @change="onFileChange"
                class="block w-full text-sm"
              />

              <button
                v-if="preview || form.cover_url"
                type="button"
                @click="clearImage"
                class="px-3 py-2 text-sm text-red-500 hover:bg-red-50 rounded-lg">
                Hapus
              </button>
            </div>

            <p class="text-xs text-gray-500 mt-2">Max 5MB (JPG, PNG, GIF)</p>
          </div>

          <!-- Buttons -->
          <div class="flex flex-col sm:flex-row gap-4 pt-6 border-t">
            <button
              type="submit"
              class="flex-1 px-6 py-3 bg-blue-600 text-white rounded-xl">
              {{ submitText }}
            </button>

            <button
              type="button"
              @click="$emit('cancel')"
              class="flex-1 px-6 py-3 bg-gray-100 dark:bg-gray-700 rounded-xl">
              Batal
            </button>
          </div>

        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch } from "vue"

const props = defineProps({
  modelValue: { type: Object, default: () => ({}) },
  categories: { type: Array, default: () => [] },
  submitText: { type: String, default: "Simpan" },
  loading: { type: Boolean, default: false }
})

const emit = defineEmits(["save", "cancel"])

const form = reactive({
  title: "",
  content: "",
  category_id: "",
  cover_url: "" // harus string, tidak boleh null
})

const preview = ref(null)   // aman
const newFile = ref(null)   // aman

/* Prefill saat EDIT */
watch(
  () => props.modelValue,
  (v) => {
    if (!v) return

    form.title = v.title || ""
    form.content = v.content || ""
    form.category_id = v.category_id ? Number(v.category_id) : ""
    form.cover_url = v.cover || v.cover_url || ""   // fallback aman
    preview.value = null
    newFile.value = null
  },
  { immediate: true }
)

function onFileChange(e) {
  const file = e.target.files[0]
  if (!file) return

  if (!file.type.startsWith("image/")) return alert("File harus gambar")
  if (file.size > 5 * 1024 * 1024) return alert("Max 5MB")

  newFile.value = file
  preview.value = URL.createObjectURL(file)
}

function clearImage() {
  newFile.value = null
  preview.value = null
  form.cover_url = "" // hilangkan cover lama
}

function submitForm() {
  emit("save", {
    data: {
      title: form.title,
      content: form.content,
      category_id: Number(form.category_id)
    },
    file: newFile.value
  })
}
</script>
