<template>
  <!-- Enhanced Background Layer -->
  <div class="fixed inset-0 overflow-hidden pointer-events-none transition-all duration-1000">
    <!-- Dynamic Gradient Background -->
    <div class="absolute inset-0 transition-all duration-2000" 
         :class="currentView === 'login' 
           ? 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-100 dark:from-gray-900 dark:via-blue-900/20 dark:to-indigo-900' 
           : 'bg-gradient-to-br from-emerald-50 via-green-50 to-teal-100 dark:from-gray-900 dark:via-emerald-900/20 dark:to-teal-900'">
    </div>
    
    <!-- Animated Floating Elements -->
    <div class="absolute -top-20 -left-20 w-60 h-60 bg-blue-400/20 rounded-full blur-3xl animate-orb-float"></div>
    <div class="absolute -bottom-20 -right-20 w-72 h-72 bg-purple-400/20 rounded-full blur-3xl animate-orb-float delay-1000"></div>
    <div class="absolute top-1/3 right-1/4 w-48 h-48 bg-cyan-400/15 rounded-full blur-2xl animate-orb-float delay-2000"></div>
    <div class="absolute bottom-1/3 left-1/4 w-52 h-52 bg-indigo-400/15 rounded-full blur-2xl animate-orb-float delay-1500"></div>
    
    <!-- Dynamic Grid Pattern -->
    <div class="absolute inset-0 opacity-20 dark:opacity-10 transition-all duration-1000">
      <div class="w-full h-full" :class="currentView === 'login' ? 'bg-grid-blue' : 'bg-grid-green'"></div>
    </div>

    <!-- Animated Connection Lines -->
    <div class="absolute inset-0 opacity-40">
      <div v-for="i in 8" :key="i" 
           class="absolute w-0.5 h-20 bg-gradient-to-b from-blue-500/30 to-transparent animate-connection-line"
           :style="{
             left: (i * 12.5) + '%',
             animationDelay: (i * 0.5) + 's',
             animationDuration: (3 + Math.random() * 2) + 's'
           }">
      </div>
    </div>

    <!-- Floating Particles with Enhanced Animation -->
    <div v-for="i in 20" :key="i" 
         class="absolute w-2 h-2 rounded-full animate-particle-float"
         :class="currentView === 'login' ? 'bg-blue-400/60' : 'bg-emerald-400/60'"
         :style="{
           top: Math.random() * 100 + '%',
           left: Math.random() * 100 + '%',
           animationDelay: Math.random() * 8 + 's',
           animationDuration: (6 + Math.random() * 6) + 's'
         }">
    </div>

    <!-- Pulse Waves -->
    <div class="absolute inset-0 flex items-center justify-center">
      <div v-for="i in 3" :key="i"
           class="absolute border-2 rounded-full animate-pulse-wave"
           :class="currentView === 'login' ? 'border-blue-400/20' : 'border-emerald-400/20'"
           :style="{
             width: 100 + i * 100 + 'px',
             height: 100 + i * 100 + 'px',
             animationDelay: i * 1 + 's'
           }">
      </div>
    </div>
  </div>

  <!-- Enhanced Content Layer -->
  <div class="min-h-screen flex items-center justify-center p-4 relative z-10">
    <div class="w-full max-w-md mx-auto">
      
      <!-- Enhanced Back to Home Button -->
      <div class="flex justify-start mb-8">
        <button
          @click="goToHome"
          class="group relative inline-flex items-center px-5 py-3 bg-white/90 dark:bg-gray-800/90 backdrop-blur-xl rounded-2xl border border-white/50 dark:border-gray-700/50 shadow-2xl hover:shadow-3xl transition-all duration-500 hover:scale-105 text-gray-700 dark:text-gray-300 hover:text-gray-900 dark:hover:text-white overflow-hidden"
        >
          <!-- Animated Background -->
          <div class="absolute inset-0 bg-gradient-to-r from-blue-500/0 via-cyan-500/10 to-blue-500/0 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-500 transform -translate-x-full group-hover:translate-x-full"></div>
          
          <!-- Icon and Text -->
          <svg class="w-5 h-5 mr-3 transform group-hover:-translate-x-1 transition-transform duration-300 relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
          </svg>
          <span class="font-semibold text-sm relative z-10">Back to Home</span>
          
          <!-- Corner Accents -->
          <div class="absolute top-0 left-0 w-3 h-3 border-t-2 border-l-2 border-blue-500/50 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
          <div class="absolute bottom-0 right-0 w-3 h-3 border-b-2 border-r-2 border-cyan-500/50 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
        </button>
      </div>

      <!-- Enhanced Main Card Container -->
      <div class="backdrop-blur-2xl bg-white/90 dark:bg-gray-800/90 rounded-3xl shadow-2xl border border-white/50 dark:border-gray-700/50 transition-all duration-700 overflow-hidden relative group/card"
           :class="{
             'hover:shadow-3xl hover:scale-105': !loading
           }">
        
        <!-- Card Background Effects -->
        <div class="absolute inset-0 opacity-0 group-hover/card:opacity-100 transition-opacity duration-1000">
          <div class="absolute inset-0 bg-gradient-to-br from-blue-500/5 to-cyan-500/5" 
               v-if="currentView === 'login'"></div>
          <div class="absolute inset-0 bg-gradient-to-br from-emerald-500/5 to-green-500/5" 
               v-else></div>
        </div>

        <!-- Animated Border -->
        <div class="absolute inset-0 rounded-3xl border-2 border-transparent opacity-0 group-hover/card:opacity-100 transition-all duration-1000"
             :class="currentView === 'login' 
               ? 'group-hover/card:border-blue-500/20 animate-border-glow-blue' 
               : 'group-hover/card:border-emerald-500/20 animate-border-glow-green'">
        </div>

        <!-- Enhanced Toggle Switch -->
        <div class="p-2 border-b border-white/30 dark:border-gray-700/30 bg-white/50 dark:bg-gray-800/50 backdrop-blur-xl relative">
          <div class="flex relative">
            <!-- Animated Background Slider -->
            <div class="absolute top-1 left-1 bottom-1 w-1/2 rounded-xl transition-all duration-700 shadow-lg transform"
                 :class="currentView === 'login' 
                   ? 'bg-gradient-to-r from-blue-500 to-cyan-500 translate-x-0 scale-105' 
                   : 'bg-gradient-to-r from-emerald-500 to-green-500 translate-x-full scale-105'">
              <!-- Shine Effect on Slider -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent transform -translate-x-full group-hover/card:translate-x-full transition-transform duration-1000"></div>
            </div>
            
            <button
              @click="switchToLogin"
              :class="[
                'flex-1 py-4 px-6 rounded-xl font-bold transition-all duration-700 relative z-10 text-sm',
                currentView === 'login' 
                  ? 'text-white drop-shadow-lg' 
                  : 'text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200'
              ]"
            >
              <div class="flex items-center justify-center space-x-3">
                <svg class="w-5 h-5 transition-transform duration-500" 
                     :class="{'scale-110 rotate-12': currentView === 'login'}" 
                     fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"/>
                </svg>
                <span class="transition-all duration-500"
                      :class="{'translate-x-1': currentView === 'login'}">
                  Login
                </span>
              </div>
            </button>
            
            <button
              @click="switchToRegister"
              :class="[
                'flex-1 py-4 px-6 rounded-xl font-bold transition-all duration-700 relative z-10 text-sm',
                currentView === 'register' 
                  ? 'text-white drop-shadow-lg' 
                  : 'text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200'
              ]"
            >
              <div class="flex items-center justify-center space-x-3">
                <svg class="w-5 h-5 transition-transform duration-500" 
                     :class="{'scale-110 rotate-12': currentView === 'register'}" 
                     fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
                </svg>
                <span class="transition-all duration-500"
                      :class="{'translate-x-1': currentView === 'register'}">
                  Register
                </span>
              </div>
            </button>
          </div>
        </div>

        <!-- Enhanced Form Content -->
        <div class="p-8 relative">
          <transition
            :name="currentView === 'login' ? 'form-rotate' : 'form-flip'"
            mode="out-in"
          >
            <!-- Login Form -->
            <div 
              v-if="currentView === 'login'" 
              key="login"
              class="space-y-6"
            >
              <!-- Enhanced Header -->
              <div class="text-center">
                <div class="relative inline-block mb-4 group/header">
                  <div class="w-16 h-16 bg-gradient-to-br from-blue-500 to-cyan-500 rounded-2xl flex items-center justify-center mx-auto shadow-2xl transform group-hover/header:scale-110 transition-all duration-500 relative overflow-hidden">
                    <!-- Animated Rings -->
                    <div class="absolute inset-0 border-2 border-white/30 rounded-2xl animate-ping-slow"></div>
                    <div class="absolute inset-2 border border-white/20 rounded-xl animate-ping-slower"></div>
                    
                    <!-- Shine Effect -->
                    <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/40 to-transparent transform -translate-x-full group-hover/header:translate-x-full transition-transform duration-1000"></div>
                    
                    <svg class="w-8 h-8 text-white relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                    </svg>
                  </div>
                </div>
                <h2 class="text-3xl font-black bg-gradient-to-r from-blue-600 to-cyan-600 bg-clip-text text-transparent mb-2">
                  Welcome Back
                </h2>
                <p class="text-gray-500 dark:text-cyan-300 text-sm">Sign in to continue your journey</p>
              </div>

              <form @submit.prevent="handleLogin" class="space-y-5">
                <!-- Enhanced Email Input -->
                <div class="space-y-3 group/input">
                  <label class="text-sm font-semibold text-gray-700 dark:text-cyan-200 transition-all duration-300 group-hover/input:translate-x-1">
                    Email Address
                  </label>
                  <div class="relative group/field">
                    <input 
                      v-model="loginEmail" 
                      type="email" 
                      required
                      placeholder="your@email.com"
                      class="w-full p-4 pl-12 rounded-xl border-2 border-gray-200 dark:border-gray-600 
                             bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                             focus:border-blue-500 dark:focus:border-cyan-500 transition-all duration-500
                             placeholder-gray-400 dark:placeholder-gray-500 text-sm shadow-lg
                             group-hover/field:shadow-xl group-hover/field:scale-105"
                    />
                    <div class="absolute left-4 top-1/2 transform -translate-y-1/2 transition-all duration-500 group-hover/field:scale-110">
                      <div class="w-6 h-6 bg-gradient-to-br from-blue-500 to-cyan-400 rounded-lg flex items-center justify-center shadow-md">
                        <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Enhanced Password Input -->
                <div class="space-y-3 group/input">
                  <label class="text-sm font-semibold text-gray-700 dark:text-cyan-200 transition-all duration-300 group-hover/input:translate-x-1">
                    Password
                  </label>
                  <div class="relative group/field">
                    <input 
                      v-model="loginPassword" 
                      :type="showLoginPassword ? 'text' : 'password'" 
                      required
                      placeholder="••••••••"
                      class="w-full p-4 pl-12 pr-12 rounded-xl border-2 border-gray-200 dark:border-gray-600 
                             bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                             focus:border-cyan-500 dark:focus:border-blue-500 transition-all duration-500
                             placeholder-gray-400 dark:placeholder-gray-500 text-sm shadow-lg
                             group-hover/field:shadow-xl group-hover/field:scale-105"
                    />
                    <div class="absolute left-4 top-1/2 transform -translate-y-1/2 transition-all duration-500 group-hover/field:scale-110">
                      <div class="w-6 h-6 bg-gradient-to-br from-cyan-500 to-blue-400 rounded-lg flex items-center justify-center shadow-md">
                        <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                        </svg>
                      </div>
                    </div>
                    <button 
                      type="button"
                      @click="showLoginPassword = !showLoginPassword"
                      class="absolute right-4 top-1/2 transform -translate-y-1/2 p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 transition-all duration-300 group/eye"
                    >
                      <svg class="w-4 h-4 text-gray-400 group-hover/eye:text-gray-600 dark:group-hover/eye:text-cyan-300 transition-colors duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path v-if="showLoginPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                        <path v-if="showLoginPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                        <path v-if="!showLoginPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L6.59 6.59m9.02 9.02l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <!-- Enhanced Submit Button -->
                <button 
                  type="submit" 
                  :disabled="loading"
                  class="w-full py-4 px-6 bg-gradient-to-r from-blue-500 to-cyan-500 text-white rounded-xl 
                         hover:from-blue-600 hover:to-cyan-600 transition-all duration-500 font-bold
                         shadow-2xl hover:shadow-3xl disabled:opacity-50 disabled:cursor-not-allowed
                         flex items-center justify-center space-x-3 text-sm relative overflow-hidden group/btn
                         transform hover:scale-105"
                >
                  <!-- Animated Background -->
                  <div class="absolute inset-0 bg-gradient-to-r from-blue-600 to-cyan-600 opacity-0 group-hover/btn:opacity-100 transition-opacity duration-500"></div>
                  
                  <!-- Shine Effect -->
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent transform -translate-x-full group-hover/btn:translate-x-full transition-transform duration-1000"></div>
                  
                  <span v-if="loading" class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin relative z-10"></span>
                  <span class="relative z-10 font-semibold">{{ loading ? 'Signing in...' : 'Sign In' }}</span>
                  <svg v-if="!loading" class="w-5 h-5 relative z-10 transform group-hover/btn:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
                  </svg>
                </button>
              </form>

              <div class="text-center pt-4">
                <p class="text-gray-500 dark:text-cyan-300 text-sm">
                  Don't have an account? 
                  <button @click="switchToRegister" class="text-blue-500 dark:text-cyan-400 font-bold hover:text-blue-600 dark:hover:text-cyan-300 transition-all duration-300 hover:underline hover:scale-105 transform inline-block">
                    Sign up now
                  </button>
                </p>
              </div>
            </div>

            <!-- Register Form -->
            <div 
              v-else
              key="register"
              class="space-y-6"
            >
              <!-- Enhanced Header -->
              <div class="text-center">
                <div class="relative inline-block mb-4 group/header">
                  <div class="w-16 h-16 bg-gradient-to-br from-emerald-500 to-green-500 rounded-2xl flex items-center justify-center mx-auto shadow-2xl transform group-hover/header:scale-110 transition-all duration-500 relative overflow-hidden">
                    <!-- Animated Rings -->
                    <div class="absolute inset-0 border-2 border-white/30 rounded-2xl animate-ping-slow"></div>
                    <div class="absolute inset-2 border border-white/20 rounded-xl animate-ping-slower"></div>
                    
                    <!-- Shine Effect -->
                    <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/40 to-transparent transform -translate-x-full group-hover/header:translate-x-full transition-transform duration-1000"></div>
                    
                    <svg class="w-8 h-8 text-white relative z-10" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
                    </svg>
                  </div>
                </div>
                <h2 class="text-3xl font-black bg-gradient-to-r from-emerald-600 to-green-600 bg-clip-text text-transparent mb-2">
                  Join Us
                </h2>
                <p class="text-gray-500 dark:text-emerald-300 text-sm">Start your amazing journey</p>
              </div>

              <form @submit.prevent="registerUser" class="space-y-5">
                <!-- Enhanced Name Input -->
                <div class="space-y-3 group/input">
                  <label class="text-sm font-semibold text-gray-700 dark:text-emerald-200 transition-all duration-300 group-hover/input:translate-x-1">
                    Full Name
                  </label>
                  <div class="relative group/field">
                    <input 
                      v-model="registerName" 
                      type="text" 
                      required
                      placeholder="Enter your full name"
                      class="w-full p-4 pl-12 rounded-xl border-2 border-gray-200 dark:border-gray-600 
                             bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                             focus:border-emerald-500 dark:focus:border-green-500 transition-all duration-500
                             placeholder-gray-400 dark:placeholder-gray-500 text-sm shadow-lg
                             group-hover/field:shadow-xl group-hover/field:scale-105"
                    />
                    <div class="absolute left-4 top-1/2 transform -translate-y-1/2 transition-all duration-500 group-hover/field:scale-110">
                      <div class="w-6 h-6 bg-gradient-to-br from-emerald-500 to-green-400 rounded-lg flex items-center justify-center shadow-md">
                        <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Enhanced Email Input -->
                <div class="space-y-3 group/input">
                  <label class="text-sm font-semibold text-gray-700 dark:text-emerald-200 transition-all duration-300 group-hover/input:translate-x-1">
                    Email Address
                  </label>
                  <div class="relative group/field">
                    <input 
                      v-model="registerEmail" 
                      type="email" 
                      required
                      placeholder="your@email.com"
                      class="w-full p-4 pl-12 rounded-xl border-2 border-gray-200 dark:border-gray-600 
                             bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                             focus:border-green-500 dark:focus:border-emerald-500 transition-all duration-500
                             placeholder-gray-400 dark:placeholder-gray-500 text-sm shadow-lg
                             group-hover/field:shadow-xl group-hover/field:scale-105"
                    />
                    <div class="absolute left-4 top-1/2 transform -translate-y-1/2 transition-all duration-500 group-hover/field:scale-110">
                      <div class="w-6 h-6 bg-gradient-to-br from-green-500 to-emerald-400 rounded-lg flex items-center justify-center shadow-md">
                        <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- Enhanced Password Input -->
                <div class="space-y-3 group/input">
                  <label class="text-sm font-semibold text-gray-700 dark:text-emerald-200 transition-all duration-300 group-hover/input:translate-x-1">
                    Password
                  </label>
                  <div class="relative group/field">
                    <input 
                      v-model="registerPassword" 
                      :type="showRegisterPassword ? 'text' : 'password'" 
                      required
                      placeholder="••••••••"
                      class="w-full p-4 pl-12 pr-12 rounded-xl border-2 border-gray-200 dark:border-gray-600 
                             bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none 
                             focus:border-emerald-500 dark:focus:border-green-500 transition-all duration-500
                             placeholder-gray-400 dark:placeholder-gray-500 text-sm shadow-lg
                             group-hover/field:shadow-xl group-hover/field:scale-105"
                    />
                    <div class="absolute left-4 top-1/2 transform -translate-y-1/2 transition-all duration-500 group-hover/field:scale-110">
                      <div class="w-6 h-6 bg-gradient-to-br from-green-500 to-emerald-400 rounded-lg flex items-center justify-center shadow-md">
                        <svg class="w-3 h-3 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                        </svg>
                      </div>
                    </div>
                    <button 
                      type="button"
                      @click="showRegisterPassword = !showRegisterPassword"
                      class="absolute right-4 top-1/2 transform -translate-y-1/2 p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 transition-all duration-300 group/eye"
                    >
                      <svg class="w-4 h-4 text-gray-400 group-hover/eye:text-gray-600 dark:group-hover/eye:text-emerald-300 transition-colors duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path v-if="showRegisterPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                        <path v-if="showRegisterPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                        <path v-if="!showRegisterPassword" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L6.59 6.59m9.02 9.02l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"/>
                      </svg>
                    </button>
                  </div>
                </div>

                <!-- Enhanced Submit Button -->
                <button 
                  type="submit" 
                  :disabled="loading"
                  class="w-full py-4 px-6 bg-gradient-to-r from-emerald-500 to-green-500 text-white rounded-xl 
                         hover:from-emerald-600 hover:to-green-600 transition-all duration-500 font-bold
                         shadow-2xl hover:shadow-3xl disabled:opacity-50 disabled:cursor-not-allowed
                         flex items-center justify-center space-x-3 text-sm relative overflow-hidden group/btn
                         transform hover:scale-105"
                >
                  <!-- Animated Background -->
                  <div class="absolute inset-0 bg-gradient-to-r from-emerald-600 to-green-600 opacity-0 group-hover/btn:opacity-100 transition-opacity duration-500"></div>
                  
                  <!-- Shine Effect -->
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/30 to-transparent transform -translate-x-full group-hover/btn:translate-x-full transition-transform duration-1000"></div>
                  
                  <span v-if="loading" class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin relative z-10"></span>
                  <span class="relative z-10 font-semibold">{{ loading ? 'Creating Account...' : 'Create Account' }}</span>
                  <svg v-if="!loading" class="w-5 h-5 relative z-10 transform group-hover/btn:translate-x-1 transition-transform duration-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"/>
                  </svg>
                </button>
              </form>

              <div class="text-center pt-4">
                <p class="text-gray-500 dark:text-emerald-300 text-sm">
                  Already have an account? 
                  <button @click="switchToLogin" class="text-emerald-500 dark:text-green-400 font-bold hover:text-emerald-600 dark:hover:text-green-300 transition-all duration-300 hover:underline hover:scale-105 transform inline-block">
                    Sign in here
                  </button>
                </p>
              </div>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { User } from '~/composables/useAuth'
import { useAuth } from '~/composables/useAuth'

// Reactive state untuk kedua form
const currentView = ref<'login' | 'register'>('login')
const loading = ref(false)

// State untuk form LOGIN
const loginEmail = ref('')
const loginPassword = ref('')
const showLoginPassword = ref(false)

// State untuk form REGISTER
const registerName = ref('')
const registerEmail = ref('')
const registerPassword = ref('')
const showRegisterPassword = ref(false)

const router = useRouter()
const auth = useAuth()

// Fungsi kembali ke beranda
function goToHome() {
  router.push('/')
}

// Enhanced switching functions
function switchToLogin() {
  if (currentView.value === 'login') return
  currentView.value = 'login'
}

function switchToRegister() {
  if (currentView.value === 'register') return
  currentView.value = 'register'
}

// Fungsi LOGIN
async function handleLogin() {
  if (loading.value) return
  
  loading.value = true
  try {
    console.log("📩 EMAIL TERKIRIM =", loginEmail.value)
    console.log("🔑 PASSWORD TERKIRIM =", loginPassword.value)

    const res = await fetch('http://localhost:8080/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        email: loginEmail.value, 
        password: loginPassword.value 
      })
    })

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}))
      throw new Error(errorData.message || `Login gagal (HTTP ${res.status})`)
    }

    const data = await res.json()

    if (!data.token || !data.user) {
      throw new Error('Login gagal: user atau token tidak valid')
    }

    // Simpan data login
    auth.login(data.token, data.user as User)

    console.log('✅ Login berhasil, user role:', data.user.role)

    // Redirect berdasarkan role
    if (data.user.role === 'admin') {
      router.push('/admin/berita')
    } else {
      router.push('/')
    }

  } catch (err: any) {
    console.error('❌ Login error:', err)
    alert(err?.message || 'Login gagal. Periksa email dan password Anda.')
  } finally {
    loading.value = false
  }
}

// Fungsi REGISTER
async function registerUser() {
  if (loading.value) return
  
  loading.value = true
  try {
    const res = await fetch('http://localhost:8080/api/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        name: registerName.value, 
        email: registerEmail.value, 
        password: registerPassword.value 
      })
    })

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}))
      throw new Error(errorData.message || `Registrasi gagal (HTTP ${res.status})`)
    }

    const data = await res.json()
    
    if (!data.token || !data.user) {
      throw new Error('Registrasi gagal: data user tidak valid')
    }

    console.log('✅ Registrasi berhasil')

    // Auto login setelah registrasi
    auth.login(data.token, data.user as User)
    
    // Redirect ke homepage
    router.push('/')
    
  } catch (err: any) {
    console.error('❌ Registrasi error:', err)
    alert(err?.message || 'Registrasi gagal. Coba lagi.')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* Enhanced Animation Styles */
.form-rotate-enter-active {
  transition: all 0.8s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.form-rotate-leave-active {
  transition: all 0.6s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.form-rotate-enter-from {
  opacity: 0;
  transform: rotateY(90deg) scale(0.8);
}

.form-rotate-leave-to {
  opacity: 0;
  transform: rotateY(-90deg) scale(0.8);
}

.form-flip-enter-active {
  transition: all 0.8s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.form-flip-leave-active {
  transition: all 0.6s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.form-flip-enter-from {
  opacity: 0;
  transform: rotateX(90deg) scale(0.8);
}

.form-flip-leave-to {
  opacity: 0;
  transform: rotateX(-90deg) scale(0.8);
}

/* Enhanced Custom Animations */
@keyframes orb-float {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg) scale(1);
  }
  33% {
    transform: translateY(-30px) translateX(20px) rotate(120deg) scale(1.1);
  }
  66% {
    transform: translateY(20px) translateX(-15px) rotate(240deg) scale(0.9);
  }
}

@keyframes particle-float {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
    opacity: 0.6;
  }
  25% {
    transform: translateY(-40px) translateX(15px) rotate(90deg);
    opacity: 0.8;
  }
  50% {
    transform: translateY(-60px) translateX(-10px) rotate(180deg);
    opacity: 1;
  }
  75% {
    transform: translateY(-40px) translateX(-15px) rotate(270deg);
    opacity: 0.8;
  }
}

@keyframes connection-line {
  0% {
    transform: translateY(-100px);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translateY(100vh);
    opacity: 0;
  }
}

@keyframes pulse-wave {
  0% {
    transform: scale(0.8);
    opacity: 1;
  }
  100% {
    transform: scale(2);
    opacity: 0;
  }
}

@keyframes border-glow-blue {
  0%, 100% {
    box-shadow: 0 0 20px rgba(59, 130, 246, 0.3);
  }
  50% {
    box-shadow: 0 0 40px rgba(59, 130, 246, 0.6);
  }
}

@keyframes border-glow-green {
  0%, 100% {
    box-shadow: 0 0 20px rgba(34, 197, 94, 0.3);
  }
  50% {
    box-shadow: 0 0 40px rgba(34, 197, 94, 0.6);
  }
}

@keyframes ping-slow {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.5;
  }
}

@keyframes ping-slower {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.3;
  }
}

.animate-orb-float {
  animation: orb-float 15s ease-in-out infinite;
}

.animate-particle-float {
  animation: particle-float 12s ease-in-out infinite;
}

.animate-connection-line {
  animation: connection-line 4s linear infinite;
}

.animate-pulse-wave {
  animation: pulse-wave 3s ease-out infinite;
}

.animate-border-glow-blue {
  animation: border-glow-blue 2s ease-in-out infinite;
}

.animate-border-glow-green {
  animation: border-glow-green 2s ease-in-out infinite;
}

.animate-ping-slow {
  animation: ping-slow 3s ease-in-out infinite;
}

.animate-ping-slower {
  animation: ping-slower 4s ease-in-out infinite;
}

/* Background Grid Patterns */
.bg-grid-blue {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
}

.bg-grid-green {
  background-image: 
    linear-gradient(rgba(34, 197, 94, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(34, 197, 94, 0.1) 1px, transparent 1px);
  background-size: 50px 50px;
}

/* Backdrop blur enhancement */
.backdrop-blur-2xl {
  backdrop-filter: blur(24px) saturate(200%);
  -webkit-backdrop-filter: blur(24px) saturate(200%);
}

/* Smooth transitions for all elements */
* {
  transition: 
    color 0.3s ease-in-out, 
    background-color 0.3s ease-in-out, 
    border-color 0.3s ease-in-out,
    transform 0.3s ease-in-out,
    box-shadow 0.3s ease-in-out;
}
</style>