<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useResumeStore } from '@/stores/resume'
import type { PersonalInfo, Experience, Skill } from '@/services/api'
import PersonalInfoModal from '@/components/PersonalInfoModal.vue'
import ExperienceModal from '@/components/ExperienceModal.vue'
import SkillModal from '@/components/SkillModal.vue'

const resumeStore = useResumeStore()

// Use storeToRefs to maintain reactivity when destructuring
const {
  personalInfo,
  experiences,
  skills,
  canEdit
} = storeToRefs(resumeStore)

// Destructure methods (these don't need to be reactive)
const {
  updatePersonalInfo,
  addExperience,
  updateExperience,
  deleteExperience,
  addSkill,
  updateSkill,
  deleteSkill
} = resumeStore

// Modal states
const editPersonalInfo = ref(false)
const showAddExperience = ref(false)
const editingExperience = ref<Experience | null>(null)
const showAddSkill = ref(false)
const editingSkill = ref<Skill | null>(null)

// Handlers
const handlePersonalInfoSave = async (data: PersonalInfo) => {
  try {
    await updatePersonalInfo(data)
    editPersonalInfo.value = false
  } catch (error) {
    // Error is handled by the store
  }
}

const editExperience = (experience: Experience) => {
  editingExperience.value = experience
}

const closeExperienceModal = () => {
  showAddExperience.value = false
  editingExperience.value = null
}

const handleExperienceSave = async (data: Omit<Experience, 'id'> | Experience) => {
  try {
    if (editingExperience.value) {
      await updateExperience(data as Experience)
    } else {
      await addExperience(data as Omit<Experience, 'id'>)
    }
    closeExperienceModal()
  } catch (error) {
    // Error is handled by the store
  }
}

const editSkill = (skill: Skill) => {
  editingSkill.value = skill
}

const closeSkillModal = () => {
  showAddSkill.value = false
  editingSkill.value = null
}

const handleSkillSave = async (data: Omit<Skill, 'id'> | Skill) => {
  try {
    if (editingSkill.value) {
      await updateSkill(data as Skill)
    } else {
      await addSkill(data as Omit<Skill, 'id'>)
    }
    closeSkillModal()
  } catch (error) {
    // Error is handled by the store
  }
}
</script>

<template>
  <div class="space-y-8">
    <!-- Personal Information Section -->
    <section class="bg-white rounded-lg shadow-md p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold text-gray-800">Personal Information</h2>
        <button
          v-if="canEdit"
          @click="editPersonalInfo = true"
          class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors"
        >
          Edit
        </button>
      </div>
      
      <div class="grid md:grid-cols-2 gap-6">
        <div>
          <h3 class="text-xl font-semibold text-gray-700 mb-2">
            {{ personalInfo?.title || 'Your Name' }}
          </h3>
          <p class="text-gray-600 mb-4">{{ personalInfo?.description || 'No description available' }}</p>
        </div>
        
        <div class="space-y-3">
          <div class="flex items-center space-x-3">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            <span class="text-gray-700">{{ personalInfo?.email || 'email@example.com' }}</span>
          </div>
          
          <div class="flex items-center space-x-3">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
            </svg>
            <span class="text-gray-700">{{ personalInfo?.phone_number || '+1 234 567 8900' }}</span>
          </div>
          
          <div v-if="personalInfo?.addr" class="flex items-center space-x-3">
            <svg class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            <span class="text-gray-700">{{ personalInfo.addr }}</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Experience Section -->
    <section class="bg-white rounded-lg shadow-md p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold text-gray-800">Work Experience</h2>
        <button
          v-if="canEdit"
          @click="showAddExperience = true"
          class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition-colors"
        >
          Add Experience
        </button>
      </div>
      
      <div v-if="!experiences || experiences.length === 0" class="text-center py-8 text-gray-500">
        <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2-2v2m8 0V6a2 2 0 012 2v6a2 2 0 01-2 2H8a2 2 0 01-2-2V8a2 2 0 012-2z" />
        </svg>
        <p>No work experience added yet.</p>
      </div>
      
      <div v-else class="space-y-6">
        <div
          v-for="experience in experiences"
          :key="experience.id"
          class="border-l-4 border-blue-500 pl-6 py-4"
        >
          <div class="flex justify-between items-start">
            <div class="flex-1">
              <h3 class="text-xl font-semibold text-gray-800 mb-2">{{ experience.job_title }}</h3>
              <p class="text-gray-600 mb-3">{{ experience.description }}</p>
              <div class="flex items-center space-x-4 text-sm text-gray-500">
                <span>{{ experience.start_date }}</span>
                <span>-</span>
                <span>{{ experience.end_date || 'Present' }}</span>
              </div>
            </div>
            
            <div v-if="canEdit" class="flex space-x-2">
              <button
                @click="editExperience(experience)"
                class="text-blue-600 hover:text-blue-800 transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button
                @click="deleteExperience(experience.id)"
                class="text-red-600 hover:text-red-800 transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Skills Section -->
    <section class="bg-white rounded-lg shadow-md p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-2xl font-bold text-gray-800">Skills</h2>
        <button
          v-if="canEdit"
          @click="showAddSkill = true"
          class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition-colors"
        >
          Add Skill
        </button>
      </div>
      
      <div v-if="!skills || skills.length === 0" class="text-center py-8 text-gray-500">
        <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
        </svg>
        <p>No skills added yet.</p>
      </div>
      
      <div v-else class="grid md:grid-cols-2 gap-6">
        <div
          v-for="skill in skills"
          :key="skill.id"
          class="bg-gray-50 rounded-lg p-4"
        >
          <div class="flex justify-between items-center mb-3">
            <h3 class="font-semibold text-gray-800">{{ skill.skill_title }}</h3>
            
            <div v-if="canEdit" class="flex space-x-2">
              <button
                @click="editSkill(skill)"
                class="text-blue-600 hover:text-blue-800 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                </svg>
              </button>
              <button
                @click="deleteSkill(skill.id)"
                class="text-red-600 hover:text-red-800 transition-colors"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>
          
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div
              class="bg-gradient-to-r from-blue-500 to-purple-600 h-2.5 rounded-full transition-all duration-300"
              :style="{ width: skill.percentage + '%' }"
            ></div>
          </div>
          <div class="text-right mt-2">
            <span class="text-sm font-medium text-gray-600">{{ skill.percentage }}%</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Modals -->
    <PersonalInfoModal
      v-if="editPersonalInfo"
      :personal-info="personalInfo"
      @close="editPersonalInfo = false"
      @save="handlePersonalInfoSave"
    />
    
    <ExperienceModal
      v-if="showAddExperience || editingExperience"
      :experience="editingExperience"
      :is-edit="!!editingExperience"
      @close="closeExperienceModal"
      @save="handleExperienceSave"
    />
    
    <SkillModal
      v-if="showAddSkill || editingSkill"
      :skill="editingSkill"
      :is-edit="!!editingSkill"
      @close="closeSkillModal"
      @save="handleSkillSave"
    />
  </div>
</template>
