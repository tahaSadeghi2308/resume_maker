import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { PersonalInfo, Experience, Skill } from '@/services/api'
import { apiService } from '@/services/api'

export const useResumeStore = defineStore('resume', () => {
  // State
  const personalInfo = ref<PersonalInfo>({
    title: '',
    description: '',
    addr: '',
    email: '',
    phone_number: ''
  })
  
  const experiences = ref<Experience[]>([])
  const skills = ref<Skill[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const apiKey = ref('')

  // Computed
  const hasApiKey = computed(() => !!apiKey.value)
  const canEdit = computed(() => hasApiKey.value)

  // Actions
  const setApiKey = (key: string) => {
    apiKey.value = key
    apiService.setApiKey(key)
  }

  const clearApiKey = () => {
    apiKey.value = ''
    apiService.setApiKey('')
  }

  const setError = (message: string) => {
    error.value = message
    setTimeout(() => {
      error.value = null
    }, 5000)
  }

  // Personal Info
  const fetchPersonalInfo = async () => {
    try {
      loading.value = true
      personalInfo.value = await apiService.getPersonalInfo()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch personal info')
    } finally {
      loading.value = false
    }
  }

  const updatePersonalInfo = async (data: PersonalInfo) => {
    try {
      loading.value = true
      await apiService.updatePersonalInfo(data)
      personalInfo.value = data
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update personal info')
      throw err
    } finally {
      loading.value = false
    }
  }

  // Experience
  const fetchExperiences = async () => {
    try {
      loading.value = true
      experiences.value = await apiService.getExperience()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch experiences')
    } finally {
      loading.value = false
    }
  }

  const addExperience = async (data: Omit<Experience, 'id'>) => {
    try {
      loading.value = true
      const newExperience = await apiService.addExperience(data)
      experiences.value = await apiService.getExperience() // Refresh list
      return newExperience
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to add experience')
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateExperience = async (data: Experience) => {
    try {
      loading.value = true
      await apiService.updateExperience(data)
      const index = experiences.value.findIndex(exp => exp.id === data.id)
      if (index !== -1) {
        experiences.value[index] = data
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update experience')
      throw err
    } finally {
      loading.value = false
    }
  }

  const deleteExperience = async (id: number) => {
    try {
      loading.value = true
      experiences.value = await apiService.deleteExperience(id)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete experience')
      throw err
    } finally {
      loading.value = false
    }
  }

  // Skills
  const fetchSkills = async () => {
    try {
      loading.value = true
      skills.value = await apiService.getSkills()
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch skills')
    } finally {
      loading.value = false
    }
  }

  const addSkill = async (data: Omit<Skill, 'id'>) => {
    try {
      loading.value = true
      const newSkill = await apiService.addSkill(data)
      skills.value = await apiService.getSkills() // Refresh list
      return newSkill
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to add skill')
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateSkill = async (data: Skill) => {
    try {
      loading.value = true
      await apiService.updateSkill(data)
      const index = skills.value.findIndex(skill => skill.id === data.id)
      if (index !== -1) {
        skills.value[index] = data
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update skill')
      throw err
    } finally {
      loading.value = false
    }
  }

  const deleteSkill = async (id: number) => {
    try {
      loading.value = true
      skills.value = await apiService.deleteSkill(id)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete skill')
      throw err
    } finally {
      loading.value = false
    }
  }

  // Initialize data
  const initializeData = async () => {
    try {
      await Promise.all([
        fetchPersonalInfo(),
        fetchExperiences(),
        fetchSkills()
      ])
    } catch (error) {
      // Error is handled by the store's setError function
    }
  }

  return {
    // State
    personalInfo,
    experiences,
    skills,
    loading,
    error,
    apiKey,
    
    // Computed
    hasApiKey,
    canEdit,
    
    // Actions
    setApiKey,
    clearApiKey,
    initializeData,
    
    // Personal Info
    fetchPersonalInfo,
    updatePersonalInfo,
    
    // Experience
    fetchExperiences,
    addExperience,
    updateExperience,
    deleteExperience,
    
    // Skills
    fetchSkills,
    addSkill,
    updateSkill,
    deleteSkill
  }
}) 