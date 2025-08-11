package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, msg string, code int) {
	http.Error(w, msg, code)
}
