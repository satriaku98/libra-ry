package middleware

import (
	"libra-ry/config"
	"libra-ry/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header Authorization
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Missing token")
		}

		// Hapus "Bearer " jika ada
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Ambil secret key dari environment
		secret := config.GetEnv("JWT_SECRET", "default_secret")

		// Parse dan verifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
		}

		// Simpan token yang sudah di-decode ke context agar bisa digunakan di handler
		c.Locals("user", token.Claims)

		return c.Next()
	}
}
