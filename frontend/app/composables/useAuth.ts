import { ref } from 'vue'

// Interface User
export interface User {
  id: number
  name: string
  email: string
  role: 'admin' | 'user' // untuk redirect
}

// Reactive state
export const user = ref<User | null>(null)
export const userToken = ref<string | null>(null)
export const isLoggedIn = ref(false)

export function useAuth() {
  // login: simpan token & user di state saja
  const login = (token: string, u?: User) => {
    userToken.value = token
    if (u) user.value = u
    isLoggedIn.value = true
  }

  // logout: hapus token & user dari state
  const logout = () => {
    userToken.value = null
    user.value = null
    isLoggedIn.value = false
  }

  // checkToken di sini tidak perlu karena tidak ada localStorage

  return { login, logout, user, userToken, isLoggedIn }
}
