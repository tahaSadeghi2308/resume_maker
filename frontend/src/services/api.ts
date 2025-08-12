import { config } from '@/config'

const API_BASE_URL = config.apiBaseUrl

export interface PersonalInfo {
  title: string
  description: string
  addr: string
  email: string
  phone_number: string
}

export interface Experience {
  id: number
  job_title: string
  description: string
  start_date: string
  end_date: string
}

export interface Skill {
  id: number
  skill_title: string
  percentage: number
}

class ApiService {
  private apiKey: string = ''

  setApiKey(key: string) {
    this.apiKey = key
  }

  private getHeaders(): HeadersInit {
    const headers: HeadersInit = {
      'Content-Type': 'application/json'
    }
    
    if (this.apiKey) {
      headers['X-API-Key'] = this.apiKey
    }
    
    return headers
  }

  // Personal Info
  async getPersonalInfo(): Promise<PersonalInfo> {
    const response = await fetch(`${API_BASE_URL}/personal-info`)
    if (!response.ok) throw new Error('Failed to fetch personal info')
    return response.json()
  }

  async updatePersonalInfo(data: PersonalInfo): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/personal-info`, {
      method: 'PUT',
      headers: this.getHeaders(),
      body: JSON.stringify(data)
    })
    if (!response.ok) throw new Error('Failed to update personal info')
  }

  // Experience
  async getExperience(): Promise<Experience[]> {
    const response = await fetch(`${API_BASE_URL}/experience/`)
    if (!response.ok) throw new Error('Failed to fetch experience')
    return response.json()
  }

  async addExperience(data: Omit<Experience, 'id'>): Promise<Experience> {
    const response = await fetch(`${API_BASE_URL}/experience/`, {
      method: 'POST',
      headers: this.getHeaders(),
      body: JSON.stringify(data)
    })
    if (!response.ok) throw new Error('Failed to add experience')
    return response.json()
  }

  async updateExperience(data: Experience): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/experience/`, {
      method: 'PUT',
      headers: this.getHeaders(),
      body: JSON.stringify(data)
    })
    if (!response.ok) throw new Error('Failed to update experience')
  }

  async deleteExperience(id: number): Promise<Experience[]> {
    const response = await fetch(`${API_BASE_URL}/experience/${id}`, {
      method: 'DELETE',
      headers: this.getHeaders()
    })
    if (!response.ok) throw new Error('Failed to delete experience')
    return response.json()
  }

  // Skills
  async getSkills(): Promise<Skill[]> {
    const response = await fetch(`${API_BASE_URL}/skills/`)
    if (!response.ok) throw new Error('Failed to fetch skills')
    return response.json()
  }

  async addSkill(data: Omit<Skill, 'id'>): Promise<Skill> {
    const response = await fetch(`${API_BASE_URL}/skills/`, {
      method: 'POST',
      headers: this.getHeaders(),
      body: JSON.stringify(data)
    })
    if (!response.ok) throw new Error('Failed to add skill')
    return response.json()
  }

  async updateSkill(data: Skill): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/skills/`, {
      method: 'PUT',
      headers: this.getHeaders(),
      body: JSON.stringify(data)
    })
    if (!response.ok) throw new Error('Failed to update skill')
  }

  async deleteSkill(id: number): Promise<Skill[]> {
    const response = await fetch(`${API_BASE_URL}/skills/${id}`, {
      method: 'DELETE',
      headers: this.getHeaders()
    })
    if (!response.ok) throw new Error('Failed to delete skill')
    return response.json()
  }
}

export const apiService = new ApiService() 