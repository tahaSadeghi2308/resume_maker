package main

import (
	"log"
	"net/http"

	"backend/handlers"
	"backend/middleware"
)

// CORS middleware
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin (for development)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func main() {
	if err := middleware.LoadAPIKey(); err != nil {
		log.Fatal(err)
	}

	// Global OPTIONS handler for CORS preflight
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")
			w.WriteHeader(http.StatusOK)
			return
		}
		http.NotFound(w, r)
	})

	http.HandleFunc("/api/personal-info", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPersonalInfo(w, r)
		case http.MethodPost, http.MethodPut:
			middleware.RequireAPIKey(handlers.UpdatePersonalInfo)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/experience/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetExperienceInfo(w, r)
		case http.MethodPost, http.MethodPut:
			middleware.RequireAPIKey(handlers.UpdateExperienceInfo)(w, r)
		case http.MethodDelete:
			middleware.RequireAPIKey(handlers.DeleteExperienceById)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/skills/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetSkillsInfo(w, r)
		case http.MethodPost, http.MethodPut:
			middleware.RequireAPIKey(handlers.UpdateSkillsInfo)(w, r)
		case http.MethodDelete:
			middleware.RequireAPIKey(handlers.DeleteSkillById)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
