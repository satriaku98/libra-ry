package routes

import (
	"libra-ry/internal/authorization"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	authGroup := app.Group("/auth")
	authGroup.Post("/login", authorization.Login)
}
