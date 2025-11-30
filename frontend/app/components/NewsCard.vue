<template>
  <div
    class="bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-xl transition-all duration-300 overflow-hidden cursor-pointer group"
    @click="handleClick"
  >
    <!-- COVER -->
    <div class="relative h-48 bg-gray-200 dark:bg-gray-700 overflow-hidden">

      <img
        v-if="validCover"
        :src="coverUrl"
        :alt="berita?.title || 'Gambar Berita'"
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        @error="onImageError"
      />

      <!-- Placeholder kalau tidak ada cover -->
      <div
        v-else
        class="w-full h-full flex items-center justify-center bg-gradient-to-r from-blue-400 to-purple-500"
      >
        <svg class="w-12 h-12 text-white opacity-70" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2" />
        </svg>
      </div>

      <!-- Kategori -->
      <div class="absolute top-3 left-3">
        <span class="px-3 py-1 text-xs font-semibold bg-blue-600 text-white rounded-full">
          {{ categoryText }}
        </span>
      </div>
    </div>

    <!-- KONTEN -->
    <div class="p-6">

      <!-- Tanggal dan Views -->
      <div class="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400 mb-3">
        <span>{{ formattedDate }}</span>
        <span class="flex items-center gap-1">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          {{ berita?.views || 0 }}
        </span>
      </div>

      <!-- Judul -->
      <h3 class="text-lg font-bold text-gray-900 dark:text-white mb-3 line-clamp-2 group-hover:text-blue-600 dark:group-hover:text-blue-400 transition-colors">
        {{ berita?.title || "Judul tidak tersedia" }}
      </h3>

      <!-- Excerpt -->
      <p class="text-gray-600 dark:text-gray-300 text-sm line-clamp-3 mb-4">
        {{ berita?.excerpt || "Ringkasan tidak tersedia." }}
      </p>

      <div class="flex items-center justify-between">
        <span class="text-blue-600 dark:text-blue-400 font-medium text-sm group-hover:underline">
          Baca selengkapnya →
        </span>
        <span class="text-xs text-gray-400">{{ index + 1 }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";

const props = defineProps({
  berita: { type: Object, required: true },
  index: { type: Number, default: 0 },
});

const emit = defineEmits(["article-click"]);

const imageError = ref(false);

// COVER URL fix
const coverUrl = computed(() => {
  if (!props.berita?.cover) return null;
  return `http://localhost:8080/uploads/berita/${props.berita.cover}`;
});

const validCover = computed(() => {
  return props.berita?.cover && !imageError.value;
});

// fallback kategori → ambil relasi jika ada
const categoryText = computed(() => {
  return props.berita?.category?.name || props.berita?.category || "Umum";
});

// fix tanggal fleksibel (date / created_at)
const formattedDate = computed(() => {
  try {
    const raw = props.berita?.date || props.berita?.created_at;
    if (!raw) return "Tidak ada tanggal";

    return new Date(raw).toLocaleDateString("id-ID", {
      year: "numeric",
      month: "short",
      day: "numeric"
    });
  } catch {
    return "Tanggal tidak valid";
  }
});

function handleClick() {
  emit("article-click", props.berita);
}

function onImageError() {
  imageError.value = true;
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
  