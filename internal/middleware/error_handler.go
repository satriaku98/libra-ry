package middleware

import (
	"libra-ry/pkg"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	return pkg.ErrorResponse(c, 500, err.Error())
}
