<template>
  <div class="container">
    <div class="card">
      <div class="card-header">
        <h2>
          <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
          </svg>
          Edit User
        </h2>
      </div>
      
      <div class="card-body">
        <div v-if="error" class="alert alert-error">
          {{ error }}
        </div>
        
        <form @submit.prevent="save">
          <div class="form-group">
            <label class="form-label">Nama</label>
            <input 
              v-model="name" 
              type="text" 
              class="form-control" 
              placeholder="Masukkan nama lengkap"
              required
            />
          </div>
          
          <div class="form-group">
            <label class="form-label">Email</label>
            <input 
              v-model="email" 
              type="email" 
              class="form-control" 
              placeholder="Masukkan alamat email"
              required
            />
          </div>
          
          <div class="form-group">
            <label class="form-label">Role</label>
            <select v-model="role" class="form-select">
              <option value="admin">Admin</option>
              <option value="user">User</option>
            </select>
          </div>
          
          <div class="action-buttons">
            <button 
              type="button" 
              class="btn btn-secondary"
              @click="goBack"
            >
              <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              Kembali
            </button>
            
            <button
              type="submit"
              :disabled="loading"
              class="btn btn-primary"
            >
              <span v-if="loading" class="loading-spinner"></span>
              <svg v-else class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              {{ loading ? 'Menyimpan...' : 'Simpan Perubahan' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

definePageMeta({ layout: 'admin' })

const route = useRoute()
const router = useRouter()
const id = route.params.id as string

const name = ref('')
const email = ref('')
const role = ref('user')
const loading = ref(false)
const error = ref('')

interface User {
  name: string
  email: string
  role: string
}

// Fetch user data saat halaman mount
onMounted(async () => {
  try {
    const data = await $fetch<User>(`http://localhost:8080/api/users/${id}`)
    name.value = data.name
    email.value = data.email
    role.value = data.role
  } catch (err) {
    console.error('Gagal fetch user:', err)
    error.value = 'Gagal mengambil data user.'
  }
})

// Fungsi simpan perubahan
const save = async () => {
  loading.value = true
  error.value = ''
  try {
    await $fetch(`http://localhost:8080/api/users/${id}`, {
      method: 'PUT',
      body: {
        name: name.value,
        email: email.value,
        role: role.value
      }
    })

    alert('Berhasil disimpan!')
    router.push('/admin/users') // redirect ke halaman list users

  } catch (err) {
    console.error('Gagal menyimpan user:', err)
    error.value = 'Gagal menyimpan user. Coba lagi.'
  } finally {
    loading.value = false
  }
}

// Fungsi kembali
const goBack = () => {
  router.back()
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.container {
  max-width: 500px;
  margin: 2rem auto;
  padding: 0 1rem;
}

.card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.08);
  overflow: hidden;
}

.card-header {
  padding: 1.5rem;
  border-bottom: 1px solid #eaeaea;
  background: linear-gradient(135deg, #4f46e5, #7c3aed);
  color: white;
}

.card-header h2 {
  font-size: 1.5rem;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.card-body {
  padding: 1.5rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #4b5563;
}

.form-control {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.2s ease;
}

.form-control:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.form-select {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 1px solid #d1d5db;
  border-radius: 8px;
  font-size: 1rem;
  background-color: white;
  cursor: pointer;
  transition: all 0.2s ease;
}

.form-select:focus {
  outline: none;
  border-color: #4f46e5;
  box-shadow: 0 0 0 3px rgba(79, 70, 229, 0.1);
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  font-weight: 500;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background-color: #4f46e5;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background-color: #4338ca;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(79, 70, 229, 0.3);
}

.btn-primary:disabled {
  background-color: #9ca3af;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.btn-secondary {
  background-color: #6b7280;
  color: white;
  margin-right: 0.75rem;
}

.btn-secondary:hover {
  background-color: #4b5563;
}

.action-buttons {
  display: flex;
  justify-content: space-between;
  margin-top: 2rem;
}

.loading-spinner {
  display: inline-block;
  width: 1rem;
  height: 1rem;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.alert {
  padding: 0.75rem 1rem;
  border-radius: 8px;
  margin-bottom: 1.5rem;
  font-size: 0.9rem;
}

.alert-error {
  background-color: #fee2e2;
  color: #dc2626;
  border-left: 4px solid #dc2626;
}

.icon {
  width: 1.25rem;
  height: 1.25rem;
}
</style>