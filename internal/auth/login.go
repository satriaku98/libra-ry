package authorization

import (
	"libra-ry/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Login and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Username and Password"
// @Success 200 {object} map[string]string "JWT Token"
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Contoh hardcoded user (sebaiknya dari DB)
	if req.Username != "admin" || req.Password != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    req.Username,
		"role":        "admin",
		"permissions": []string{"buku_read", "buku_write"},
		"exp":         time.Now().Add(3 * time.Hour).Unix(), // Expiry 3 jam
	})

	// Sign token dengan secret key dari env
	secret := config.GetEnv("JWT_SECRET", "default_secret")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}
