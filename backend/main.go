package main

import (
	"log"
	"net/http"

	"backend/handlers"
)

func main() {
	http.HandleFunc("/api/personal-info", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPersonalInfo(w, r)
		case http.MethodPost, http.MethodPut:
			handlers.UpdatePersonalInfo(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
