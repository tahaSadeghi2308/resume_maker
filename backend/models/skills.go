package models

import "sync"

type Skill struct {
	ID         uint8  `json:"id"`
	Title      string `json:"skill_title"`
	Percentage uint8  `json:"percentage"`
}

var (
	SkillsData []Skill
	SkillsMu   sync.RWMutex
)
