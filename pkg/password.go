package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword menghasilkan hash dari password yang diberikan
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword membandingkan password yang dimasukkan dengan hash di database
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
