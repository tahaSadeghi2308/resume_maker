<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useResumeStore } from '@/stores/resume'

const resumeStore = useResumeStore()
const apiKeyInput = ref('')

const { hasApiKey, loading, error, setApiKey: setStoreApiKey, clearApiKey, initializeData } = resumeStore

const setApiKey = () => {
  if (apiKeyInput.value.trim()) {
    setStoreApiKey(apiKeyInput.value.trim())
    apiKeyInput.value = ''
  }
}

onMounted(() => {
  initializeData()
})
</script>

<template>
  <div id="app">
    <header class="bg-gradient-to-r from-blue-600 to-purple-600 text-white shadow-lg">
      <div class="container mx-auto px-4 py-6">
        <div class="flex justify-between items-center">
          <h1 class="text-3xl font-bold">Resume Maker</h1>
          <div class="flex items-center space-x-4">
            <div v-if="!hasApiKey" class="flex items-center space-x-2">
              <input
                v-model="apiKeyInput"
                type="password"
                placeholder="Enter API Key"
                class="px-3 py-2 rounded text-gray-800 text-sm"
                @keyup.enter="setApiKey"
              />
              <button
                @click="setApiKey"
                class="bg-white text-blue-600 px-4 py-2 rounded font-medium hover:bg-gray-100 transition-colors"
              >
                Set Key
              </button>
            </div>
            <div v-else class="flex items-center space-x-2">
              <span class="text-sm opacity-90">API Key Set</span>
              <button
                @click="clearApiKey"
                class="bg-red-500 text-white px-3 py-2 rounded text-sm hover:bg-red-600 transition-colors"
              >
                Clear
              </button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <main class="container mx-auto px-4 py-8">
      <!-- Error Display -->
      <div v-if="error" class="mb-6 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded">
        {{ error }}
      </div>

      <!-- Loading Indicator -->
      <div v-if="loading" class="flex justify-center items-center py-8">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Main Content -->
      <div v-else>
        <router-view />
      </div>
    </main>

    <footer class="bg-gray-100 border-t mt-12">
      <div class="container mx-auto px-4 py-6 text-center text-gray-600">
        <p>&copy; 2024 Resume Maker. Built with Vue.js and Go.</p>
      </div>
    </footer>
  </div>
</template>

<style>
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

main {
  flex: 1;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style>
