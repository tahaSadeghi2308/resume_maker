<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-lg w-full mx-4">
      <div class="px-6 py-4 border-b">
        <h3 class="text-lg font-semibold text-gray-800">
          {{ isEdit ? 'Edit Experience' : 'Add New Experience' }}
        </h3>
      </div>
      
      <form @submit.prevent="handleSubmit" class="px-6 py-4">
        <div class="space-y-4">
          <div>
            <label for="jobTitle" class="block text-sm font-medium text-gray-700 mb-1">
              Job Title *
            </label>
            <input
              id="jobTitle"
              v-model="form.job_title"
              type="text"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="e.g., Senior Software Engineer"
              required
            />
          </div>
          
          <div>
            <label for="description" class="block text-sm font-medium text-gray-700 mb-1">
              Job Description *
            </label>
            <textarea
              id="description"
              v-model="form.description"
              rows="4"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Describe your responsibilities and achievements"
              required
            ></textarea>
          </div>
          
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label for="startDate" class="block text-sm font-medium text-gray-700 mb-1">
                Start Date *
              </label>
              <input
                id="startDate"
                v-model="form.start_date"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
              />
            </div>
            
            <div>
              <label for="endDate" class="block text-sm font-medium text-gray-700 mb-1">
                End Date
              </label>
              <input
                id="endDate"
                v-model="form.end_date"
                type="date"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                :min="form.start_date"
              />
              <div class="flex items-center mt-2">
                <input
                  id="currentJob"
                  v-model="isCurrentJob"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                />
                <label for="currentJob" class="ml-2 text-sm text-gray-600">
                  Current Position
                </label>
              </div>
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
            {{ isEdit ? 'Update' : 'Add' }} Experience
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import type { Experience } from '@/services/api'

interface Props {
  experience?: Experience | null
  isEdit: boolean
}

interface Emits {
  (e: 'close'): void
  (e: 'save', data: Omit<Experience, 'id'> | Experience): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const form = ref<Omit<Experience, 'id'>>({
  job_title: '',
  description: '',
  start_date: '',
  end_date: ''
})

const isCurrentJob = ref(false)

onMounted(() => {
  if (props.experience) {
    form.value = {
      job_title: props.experience.job_title,
      description: props.experience.description,
      start_date: props.experience.start_date,
      end_date: props.experience.end_date
    }
    isCurrentJob.value = !props.experience.end_date
  } else {
    // Set default start date to current month
    const today = new Date()
    form.value.start_date = today.toISOString().split('T')[0]
  }
})

watch(isCurrentJob, (current) => {
  if (current) {
    form.value.end_date = ''
  }
})

const handleSubmit = () => {
  const data = { ...form.value }
  
  if (props.isEdit && props.experience) {
    emit('save', { ...data, id: props.experience.id } as Experience)
  } else {
    emit('save', data)
  }
}
</script> 