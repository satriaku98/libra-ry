package routes

import (
	"libra-ry/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, handler *handler.AuthHandler) {
	authGroup := app.Group("/auth")
	authGroup.Post("/login", handler.Login)
}
