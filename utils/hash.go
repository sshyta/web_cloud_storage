package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword принимает пароль в виде строки и возвращает его хэш SHA-256.
func HashPassword(Password string) string {
	hash := sha256.Sum256([]byte(Password))
	return hex.EncodeToString(hash[:])
}
