package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID int, username string, role string, permissions []string, expiry int, secret string) (string, error) {
	// Buat token JWT
	expiredTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     userID,
		"username":    username,
		"role":        role,
		"permissions": permissions,
		"exp":         expiredTime.Unix(), // Expiry 3 jam
	})

	// Sign token dengan secret key dari env
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
