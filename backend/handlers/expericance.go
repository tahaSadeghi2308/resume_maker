package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

func DeleteExperienceById(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	experStringID := urlParts[3]
	experIntID, err := strconv.Atoi(experStringID)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// deleting logic but not very good !!!
	if experIntID-1 < 0 || experIntID-1 >= len(models.ExperienceData) {
		http.Error(w, "Out of range ID", http.StatusBadRequest)
		return
	}
	models.ExperienceData = append(models.ExperienceData[:experIntID-1], models.ExperienceData[experIntID:]...)

	for idx := range models.ExperienceData {
		models.ExperienceData[idx].ID = uint8(idx + 1)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.ExperienceData)
}
