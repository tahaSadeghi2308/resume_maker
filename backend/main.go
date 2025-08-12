package main

import (
	"log"
	"net/http"

	"backend/handlers"
	"backend/middleware"
)

func main() {
	if err := middleware.LoadAPIKey(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/personal-info", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPersonalInfo(w, r)
		case http.MethodPost, http.MethodPut:
			middleware.RequireAPIKey(handlers.UpdatePersonalInfo)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/experience/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/api/skills/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
