package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"backend/models"
)

func GetSkillsInfo(w http.ResponseWriter, r *http.Request) {
	models.SkillsMu.RLock()
	defer models.SkillsMu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ExperienceData)
}

func UpdateSkillsInfo(w http.ResponseWriter, r *http.Request) {
	models.SkillsMu.Lock()
	defer models.SkillsMu.Unlock()

	var updated models.Skill
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if updated.Percentage > 100 {
		http.Error(w, "Percentage out of range", http.StatusBadRequest)
		return
	}

	updated.ID = uint8(len(models.SkillsData) + 1)
	models.SkillsData = append(models.SkillsData, updated)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updated)
}

func DeleteSkillById(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	skillStringID := urlParts[3]
	skillIntID, err := strconv.Atoi(skillStringID)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// deleting logic but not very good !!!
	if skillIntID-1 < 0 || skillIntID-1 >= len(models.SkillsData) {
		http.Error(w, "Out of range ID", http.StatusBadRequest)
		return
	}
	models.SkillsData = append(models.SkillsData[:skillIntID-1], models.SkillsData[skillIntID:]...)

	for idx, skill := range models.SkillsData {
		skill.ID = uint8(idx + 1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SkillsData)
}