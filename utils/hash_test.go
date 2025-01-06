package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	Password := "my_secure_password"
	hashed := HashPassword(Password)

	if hashed == "" {
		t.Errorf("HashPassword returned an empty string")
	}

	if len(hashed) != 64 { // SHA-256 хэш всегда 64 символа
		t.Errorf("Expected hashed password length of 64, got %d", len(hashed))
	}
}
