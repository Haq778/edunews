<template>
  <div class="max-w-3xl mx-auto px-6 py-12 relative">

    <!-- Decorative floating dots -->
    <div class="absolute top-14 left-10 h-4 w-4 bg-purple-400/30 rounded-full blur-xl"></div>
    <div class="absolute bottom-20 right-20 h-5 w-5 bg-blue-400/30 rounded-full blur-xl"></div>

    <!-- Page Header -->
    <div class="mb-10 flex items-center gap-3">
      <div
        class="h-10 w-10 bg-gradient-to-br from-blue-500 to-purple-600 text-white grid place-content-center rounded-xl shadow-md"
      >
        🗂️
      </div>

      <div>
        <h1 class="text-3xl font-extrabold tracking-tight text-gray-900 dark:text-gray-100">
          Edit Kategori
        </h1>
        <p class="text-gray-500 dark:text-gray-400 mt-1 text-sm">
          Ubah nama dan slug kategori sesuai kebutuhan.
        </p>
      </div>
    </div>

    <!-- Loading -->
    <div
      v-if="loading"
      class="p-8 rounded-2xl bg-white/60 dark:bg-gray-800/50 backdrop-blur-xl border border-gray-200 dark:border-gray-700 shadow animate-pulse"
    >
      <div class="h-6 w-48 bg-gray-200 dark:bg-gray-700 rounded mb-5"></div>
      <div class="h-4 w-full bg-gray-200 dark:bg-gray-700 rounded mb-3"></div>
      <div class="h-4 w-3/4 bg-gray-200 dark:bg-gray-700 rounded"></div>
    </div>

    <!-- Form Card -->
    <div
      v-else
      class="rounded-2xl shadow-xl overflow-hidden border border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-900 transition-all"
    >
      <!-- Header -->
      <div
        class="p-6 bg-gradient-to-br from-blue-600 via-purple-600 to-blue-700 text-white shadow-inner"
      >
        <h2 class="text-xl font-semibold">Form Edit Kategori</h2>
        <p class="text-white/80 text-sm mt-1">Pastikan nama kategori ditulis dengan benar.</p>
      </div>

      <!-- Body -->
      <div class="p-8">

        <!-- Nama -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
            >Nama Kategori</label
          >
          <input
            v-model="form.name"
            type="text"
            class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-gray-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition"
          />
          <p v-if="errors.name" class="text-sm text-red-500 mt-1">
            {{ errors.name }}
          </p>
        </div>

        <!-- Slug -->
        <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Slug</label>
          <input
            v-model="form.slug"
            type="text"
            class="w-full p-3 rounded-xl border border-gray-300 dark:border-gray-700 bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-gray-200 focus:ring-2 focus:ring-purple-500 focus:border-purple-500 transition"
          />
        </div>

        <!-- Action Buttons -->
        <div class="flex gap-3 pt-4">
          <button
  :disabled="saving"
  @click="submit"
  class="px-5 py-2.5 rounded-xl font-semibold text-white bg-gradient-to-r from-yellow-500 to-yellow-600 shadow hover:brightness-110 disabled:opacity-50 transition"
>
  {{ saving ? "Menyimpan..." : "Update" }}
</button>


          <NuxtLink
            to="/admin/kategori"
            class="px-5 py-2.5 rounded-xl border border-gray-300 dark:border-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-800 transition"
          >
            Batal
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute, useRouter, useRuntimeConfig } from "#imports";

definePageMeta({ layout: "admin" });

const config = useRuntimeConfig();

const route = useRoute();
const router = useRouter();
const id = (route.params.id as string) || "";

const form = ref({ name: "", slug: "" });
const errors = ref<{ name?: string }>({});
const loading = ref(true);
const saving = ref(false);

function slugify(s = "") {
  return s
    .toLowerCase()
    .trim()
    .replace(/\s+/g, "-")
    .replace(/[^\w\-]+/g, "")
    .replace(/\-\-+/g, "-");
}

async function load() {
  loading.value = true;
  try {
    const data: any = await $fetch(`${config.public.apiBase}/api/categories/${id}`);
    form.value.name = data?.name || "";
    form.value.slug = data?.slug || "";
  } catch (e) {
    console.error("Gagal memuat kategori", e);
    alert("Gagal memuat kategori.");
    await router.push("/admin/kategori");
  } finally {
    loading.value = false;
  }
}

async function submit() {
  errors.value = {};

  if (!form.value.name || form.value.name.trim().length < 2) {
    errors.value.name = "Nama kategori minimal 2 karakter";
    return;
  }

  saving.value = true;

  try {
    const payload = {
      name: form.value.name.trim(),
      slug: form.value.slug?.trim() || slugify(form.value.name),
    };

    await $fetch(`${config.public.apiBase}/api/categories/${id}`, {
      method: "PUT",
      body: payload,
    });

    await router.push("/admin/kategori");
  } catch (e) {
    console.error("Gagal update kategori", e);
    alert("Gagal update kategori.");
  } finally {
    saving.value = false;
  }
}

onMounted(() => {
  if (!id) {
    router.push("/admin/kategori");
    return;
  }
  load();
});
</script>
