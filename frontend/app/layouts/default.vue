<template>
  <div :class="colorMode.value" class="min-h-screen transition-colors duration-300">
    <main class="bg-gray-50 dark:bg-gray-900 min-h-screen">
      <slot />
    </main>
  </div>
</template>

<script setup>
const colorMode = useColorMode()
</script>
