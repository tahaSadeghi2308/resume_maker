package middleware

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var apiKey string

func LoadAPIKey() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		return ErrNoAPIKey
	}
	return nil
}

var ErrNoAPIKey = &APIError{"API_KEY not set in .env"}

type APIError struct {
	Msg string
}

func (e *APIError) Error() string {
	return e.Msg
}

func RequireAPIKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientKey := r.Header.Get("X-API-Key")
		if clientKey != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

