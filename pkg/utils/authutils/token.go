package authutils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomHex(length int) string {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // Handle error appropriately in production
	}

	return hex.EncodeToString(randomBytes)
}
