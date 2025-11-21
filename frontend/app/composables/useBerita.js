import { ref } from 'vue'

const API_URL = 'http://localhost:8080/api/berita'

export function useBerita() {
  const beritaList = ref([])
  const loading = ref(false)
  const error = ref(null)

  const getBerita = async () => {
    loading.value = true
    try {
      const res = await fetch(API_URL)
      beritaList.value = await res.json()
    } catch (err) {
      error.value = err
    } finally {
      loading.value = false
    }
  }

  const createBerita = async (data) => {
    try {
      const res = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
      })
      return await res.json()
    } catch (err) {
      error.value = err
    }
  }

  const deleteBerita = async (id) => {
    try {
      const res = await fetch(`${API_URL}/${id}`, { method: 'DELETE' })
      return await res.json()
    } catch (err) {
      error.value = err
    }
  }

  return { beritaList, loading, error, getBerita, createBerita, deleteBerita }
}
