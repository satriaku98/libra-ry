package pkg

import (
	"encoding/json"
	"libra-ry/internal/domain"

	"github.com/gofiber/fiber/v2"
)

// Response sukses standar
func SuccessResponse(c *fiber.Ctx, message string, data interface{}, totaldata int) error {
	return c.JSON(fiber.Map{
		"success":   true,
		"message":   message,
		"data":      data,
		"totaldata": totaldata,
	})
}

// Response error standar
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

func MapUserToResponse(user domain.User) domain.UserResponse {
	var permissions []string
	_ = json.Unmarshal(user.Permissions, &permissions)

	return domain.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Role:        user.Role,
		Permissions: permissions,
	}
}

func MapUsersToResponse(users []domain.User) []domain.UserResponse {
	var list []domain.UserResponse
	for _, u := range users {
		list = append(list, MapUserToResponse(u))
	}
	return list
}
