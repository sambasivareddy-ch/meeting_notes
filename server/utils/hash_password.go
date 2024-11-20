package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenerateSessionId(access_token string) string {
	hasher := sha256.New()

	hasher.Write([]byte(access_token))

	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash
}
