package pkg

import "github.com/gofiber/fiber/v2"

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
