export default defineNuxtConfig({
  compatibilityDate: '2025-11-21',

  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode'
  ],

  tailwindcss: {
    viewer: true
  },

  colorMode: {
    classSuffix: '', // penting agar class = "dark"
    preference: 'system', 
    fallback: 'light'
  },

  css: ['~/assets/css/tailwind.css'],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  app: {
    head: {
      title: 'EduNews - Platform Berita Pendidikan',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Platform berita pendidikan terdepan di Indonesia' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },

  runtimeConfig: {
  public: {
    apiBase: "http://localhost:8080" // backend golang kamu
  }
},

  vite: { css: { preprocessorOptions: { } } }

  
})
