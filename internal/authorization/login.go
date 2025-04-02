package authorization

import (
	"libra-ry/config"
	"libra-ry/pkg"
	"strconv"
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
		return pkg.ErrorResponse(c, 401, "Invalid credentials")
	}

	// Contoh hardcoded user (sebaiknya dari DB)
	if req.Username != "admin" || req.Password != "admin" {
		return pkg.ErrorResponse(c, 401, "Invalid credentials")
	}

	// Buat token JWT
	expiry, _ := strconv.Atoi(config.GetEnv("JWT_EXPIRY", "3"))
	expiredTime := time.Now().Add(time.Duration(expiry) * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    req.Username,
		"role":        "admin",
		"permissions": []string{"buku_read", "buku_write"},
		"exp":         expiredTime.Unix(), // Expiry 3 jam
	})

	// Sign token dengan secret key dari env
	secret := config.GetEnv("JWT_SECRET", "default_secret")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return pkg.ErrorResponse(c, 500, "Could not generate token")
	}

	return pkg.SuccessResponse(c, "Login successful", fiber.Map{"token": tokenString, "expired_at": expiredTime.Format(time.RFC3339)}, 1)
}
