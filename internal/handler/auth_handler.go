package handler

import (
	"encoding/json"
	"libra-ry/internal/usecase"
	"libra-ry/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUseCase    *usecase.AuthUseCase
	tokenExpiry    int
	tokenSecretKey string
}

// NewAuthHandler membuat instance baru dari AuthHandler
func NewAuthHandler(authUseCase *usecase.AuthUseCase, tokenExpiry int, tokenSecretKey string) *AuthHandler {
	return &AuthHandler{
		authUseCase:    authUseCase,
		tokenExpiry:    tokenExpiry,
		tokenSecretKey: tokenSecretKey,
	}
}

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
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var LoginRequest LoginRequest
	if err := c.BodyParser(&LoginRequest); err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid request")
	}

	user, err := h.authUseCase.Login(LoginRequest.Username, LoginRequest.Password)
	if err != nil {
		return pkg.ErrorResponse(c, 401, "Invalid credentials")
	}

	// Permissions dalam bentuk JSON array
	var permissions []string
	err = json.Unmarshal([]byte(user.Permissions), &permissions)
	if err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to parse permissions")
	}

	token, err := pkg.GenerateJWT(user.Username, user.Role, permissions, h.tokenExpiry, h.tokenSecretKey)
	if err != nil {
		return pkg.ErrorResponse(c, 500, "Failed to generate token")
	}

	return pkg.SuccessResponse(c, "Successfully logged in", fiber.Map{"token": token}, 1)
}
