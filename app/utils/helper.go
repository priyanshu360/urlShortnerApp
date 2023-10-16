package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/priyanshu360/urlShortnerApp.git/models"
)

func GenerateRandomHash() string {
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Failed to generate random bytes")
	}

	return base64.RawURLEncoding.EncodeToString(randomBytes)[:6]
}

type apiFunc func(http.ResponseWriter, *http.Request) models.APIResult

func Handle(f apiFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		result := f(rw, r)
		resultJSON, err := json.Marshal(result.Body)
		if err != nil {
			http.Error(rw, "Failed to marshal JSON response", http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(result.Status)
		rw.Write(resultJSON)
	}
}