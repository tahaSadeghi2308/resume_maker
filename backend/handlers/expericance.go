package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
)

func GetExperienceInfo(w http.ResponseWriter, r *http.Request) {
	models.ExperienceMu.RLock()
	defer models.ExperienceMu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ExperienceData)
}

func UpdateExperienceInfo(w http.ResponseWriter, r *http.Request) {
    models.ExperienceMu.Lock()
    defer models.ExperienceMu.Unlock()

    var updated models.Experience
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    updated.ID = uint8(len(models.ExperienceData) + 1)
    models.ExperienceData = append(models.ExperienceData, updated)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(updated)
}
