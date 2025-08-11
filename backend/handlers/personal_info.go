package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
)

func GetPersonalInfo(w http.ResponseWriter, r *http.Request) {
	models.PersonalMu.RLock()
	defer models.PersonalMu.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.PersonalData)
}

func UpdatePersonalInfo(w http.ResponseWriter, r *http.Request) {
	models.PersonalMu.Lock()
	defer models.PersonalMu.Unlock()
	var updated models.PersonalInfo
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	models.PersonalData = updated
	w.WriteHeader(http.StatusNoContent)
}
