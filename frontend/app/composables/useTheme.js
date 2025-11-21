import { ref, onMounted } from '#imports'

export const useTheme = () => {
  const theme = ref('light')
  
  const setTheme = (newTheme) => {
    theme.value = newTheme
    localStorage.setItem('edunews-theme', newTheme)
    
    if (newTheme === 'dark') {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }
  
  const toggleTheme = () => {
    const newTheme = theme.value === 'light' ? 'dark' : 'light'
    setTheme(newTheme)
  }
  
  onMounted(() => {
    const savedTheme = localStorage.getItem('edunews-theme')
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
    
    if (savedTheme) {
      setTheme(savedTheme)
    } else if (prefersDark) {
      setTheme('dark')
    }
  })
  
  return {
    theme,        // <-- tidak pakai readonly
    setTheme,
    toggleTheme
  }
}
