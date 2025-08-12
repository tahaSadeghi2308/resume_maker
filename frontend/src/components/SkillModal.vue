<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-md w-full mx-4">
      <div class="px-6 py-4 border-b">
        <h3 class="text-lg font-semibold text-gray-800">
          {{ isEdit ? 'Edit Skill' : 'Add New Skill' }}
        </h3>
      </div>
      
      <form @submit.prevent="handleSubmit" class="px-6 py-4">
        <div class="space-y-4">
          <div>
            <label for="skillTitle" class="block text-sm font-medium text-gray-700 mb-1">
              Skill Name *
            </label>
            <input
              id="skillTitle"
              v-model="form.skill_title"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="e.g., JavaScript, Project Management"
              required
            />
          </div>
          
          <div>
            <label for="percentage" class="block text-sm font-medium text-gray-700 mb-1">
              Proficiency Level *
            </label>
            <div class="flex items-center space-x-4">
              <input
                id="percentage"
                v-model.number="form.percentage"
                type="range"
                min="0"
                max="100"
                step="5"
                class="flex-1 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer slider"
                required
              />
              <span class="text-lg font-semibold text-blue-600 w-16 text-center">
                {{ form.percentage }}%
              </span>
            </div>
            <div class="flex justify-between text-xs text-gray-500 mt-1">
              <span>Beginner</span>
              <span>Intermediate</span>
              <span>Expert</span>
            </div>
          </div>
          
          <div class="bg-gray-50 rounded-lg p-3">
            <div class="w-full bg-gray-200 rounded-full h-2">
              <div
                class="bg-gradient-to-r from-blue-500 to-purple-600 h-2 rounded-full transition-all duration-300"
                :style="{ width: form.percentage + '%' }"
              ></div>
            </div>
          </div>
        </div>
        
        <div class="flex justify-end space-x-3 mt-6">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 text-gray-600 border border-gray-300 rounded-md hover:bg-gray-50 transition-colors"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
          >
            {{ isEdit ? 'Update' : 'Add' }} Skill
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Skill } from '@/services/api'

interface Props {
  skill?: Skill | null
  isEdit: boolean
}

interface Emits {
  (e: 'close'): void
  (e: 'save', data: Omit<Skill, 'id'> | Skill): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const form = ref<Omit<Skill, 'id'>>({
  skill_title: '',
  percentage: 50
})

onMounted(() => {
  if (props.skill) {
    form.value = {
      skill_title: props.skill.skill_title,
      percentage: props.skill.percentage
    }
  }
})

const handleSubmit = () => {
  const data = { ...form.value }
  
  if (props.isEdit && props.skill) {
    emit('save', { ...data, id: props.skill.id } as Skill)
  } else {
    emit('save', data)
  }
}
</script>

<style scoped>
.slider::-webkit-slider-thumb {
  appearance: none;
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.slider::-moz-range-thumb {
  height: 20px;
  width: 20px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
  border: none;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
</style> 