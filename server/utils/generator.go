package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecretKey() string {
	bytes := make([]byte, 32) // 256-bit key
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
