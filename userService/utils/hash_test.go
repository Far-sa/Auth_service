package utils

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
		return
	}
	if len(hashedPassword) == 0 {
		t.Errorf("HashPassword() = %v, want a hashed password", hashedPassword)
	}
}
