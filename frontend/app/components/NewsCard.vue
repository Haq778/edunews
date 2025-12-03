<template>
  <article
    class="cursor-pointer group bg-white dark:bg-gray-800 rounded-lg overflow-hidden shadow-md hover:shadow-xl transition-all duration-300 border border-gray-100 dark:border-gray-700"
    @click="handleClick"
  >
    <!-- Cover dengan efek overlay -->
    <div class="relative h-40 sm:h-44 overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-t from-black/30 via-transparent to-transparent z-10"></div>
      
      <img
        v-if="validCover"
        :src="coverUrl"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
        :alt="berita?.title"
        @error="onImageError"
      />

      <div
        v-else
        class="w-full h-full bg-gradient-to-br from-cyan-600 to-blue-700 flex items-center justify-center text-white"
      >
        <svg class="w-10 h-10 opacity-50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
        </svg>
      </div>

      <!-- Category Badge - Warna diperbaiki -->
      <div class="absolute top-3 left-3 z-20">
        <span class="px-2.5 py-1 text-xs font-bold bg-gradient-to-r from-cyan-500 to-blue-600 text-white rounded-full shadow">
          {{ categoryText }}
        </span>
      </div>

      <!-- Breaking News Indicator - Warna diperbaiki -->
      <div v-if="isBreakingNews" class="absolute top-3 right-3 z-20">
        <span class="px-2 py-1 text-[10px] font-bold bg-gradient-to-r from-orange-500 to-red-500 text-white rounded animate-pulse">
          🔥 HOT
        </span>
      </div>
    </div>

    <!-- Content -->
    <div class="p-4">
      <!-- Meta Info -->
      <div class="flex items-center justify-between mb-2">
        <div class="flex items-center space-x-2">
          <svg class="w-3.5 h-3.5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span class="text-xs font-medium text-gray-500 dark:text-gray-400">
            {{ formattedDate }}
          </span>
        </div>
        <div class="flex items-center space-x-1">
          <svg class="w-3.5 h-3.5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
          <span class="text-xs font-medium text-gray-500 dark:text-gray-400">
            {{ viewsText }}
          </span>
        </div>
      </div>

      <!-- Title - Warna hover diperbaiki -->
      <h2
        class="text-base font-bold leading-tight text-gray-900 dark:text-white mb-2 group-hover:text-blue-600 dark:group-hover:text-cyan-400 transition-colors line-clamp-2"
      >
        {{ berita?.title }}
      </h2>

      <!-- Author (replacing excerpt) -->
      <div class="flex items-center text-gray-500 dark:text-gray-400 text-xs mb-3">
        <svg class="w-3.5 h-3.5 mr-1.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
        </svg>
        <span>{{ authorText }}</span>
      </div>

      <!-- Read More Button - Warna diperbaiki -->
      <div class="flex items-center justify-between pt-3 border-t border-gray-100 dark:border-gray-700">
        <div class="flex items-center text-blue-600 dark:text-cyan-400 font-semibold text-xs group-hover:text-blue-700 dark:group-hover:text-cyan-300 transition-colors">
          <span class="group-hover:underline">BACA SELENGKAPNYA</span>
          <svg class="w-3.5 h-3.5 ml-1.5 transform group-hover:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
          </svg>
        </div>
        
        <!-- Share Icon - Warna diperbaiki -->
        <button 
          class="p-1.5 text-gray-400 hover:text-blue-600 dark:hover:text-cyan-400 transition-colors rounded-full hover:bg-gray-100 dark:hover:bg-gray-700" 
          @click.stop="shareNews"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
          </svg>
        </button>
      </div>
    </div>
  </article>
</template>

<script setup>
import { computed, ref } from "vue";
import { navigateTo } from "#app";

const props = defineProps({
  berita: { type: Object, required: true },
});

const imageError = ref(false);

// Cover URL
const coverUrl = computed(() => {
  if (!props.berita?.cover) return null;
  return `http://localhost:8080/uploads/berita/${props.berita.cover}`;
});

const validCover = computed(() => props.berita?.cover && !imageError.value);

// Category
const categoryText = computed(() => {
  return props.berita?.category?.name || props.berita?.category || "NEWS";
});

// Views
const viewsText = computed(() => {
  const views = props.berita?.views || 0;
  if (views >= 1000) {
    return `${(views / 1000).toFixed(1)}k`;
  }
  return `${views}`;
});

// Author
const authorText = computed(() => {
  return props.berita?.author || "Reporter";
});

// Check if news is breaking (less than 1 hour old)
const isBreakingNews = computed(() => {
  const newsDate = props.berita?.date || props.berita?.created_at;
  if (!newsDate) return false;
  
  const newsTime = new Date(newsDate).getTime();
  const currentTime = new Date().getTime();
  const oneHour = 60 * 60 * 1000;
  
  return (currentTime - newsTime) < oneHour;
});

// Date
const formattedDate = computed(() => {
  const raw = props.berita?.date || props.berita?.created_at;
  if (!raw) return "Baru saja";

  const now = new Date();
  const newsDate = new Date(raw);
  const diffInHours = Math.floor((now - newsDate) / (1000 * 60 * 60));

  if (diffInHours < 1) {
    const diffInMinutes = Math.floor((now - newsDate) / (1000 * 60));
    if (diffInMinutes < 1) return "Baru saja";
    return `${diffInMinutes} mnt lalu`;
  } else if (diffInHours < 24) {
    return `${diffInHours} jam lalu`;
  }

  return newsDate.toLocaleDateString("id-ID", {
    month: "short",
    day: "numeric",
  });
});

function handleClick() {
  navigateTo(`/berita/${props.berita.id}`);
}

function onImageError() {
  imageError.value = true;
}

function shareNews() {
  if (navigator.share) {
    navigator.share({
      title: props.berita?.title,
      text: props.berita?.excerpt,
      url: window.location.origin + `/berita/${props.berita.id}`
    });
  }
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

article {
  animation: fadeIn 0.3s ease-out;
}
</style>