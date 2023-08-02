package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomHash() string {
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("Failed to generate random bytes")
	}

	return base64.RawURLEncoding.EncodeToString(randomBytes)[:6]
}