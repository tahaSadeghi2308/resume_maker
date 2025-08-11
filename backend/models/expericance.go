package models

import "sync"

type Experience struct {
    ID          uint8  `json:"id"`
    JobTitle    string `json:"job_title"`
    Description string `json:"description"`
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
}

var (
	ExperienceData []Experience
	ExperienceMu   sync.RWMutex
)
