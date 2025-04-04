package routes

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handler.UserHandler) {
	user := app.Group("/user", middleware.JWTMiddleware())

	user.Get("/", middleware.CheckPermission("user_read"), handler.GetAll)
	user.Get("/:id", middleware.CheckPermission("user_read"), handler.GetByID)
	user.Post("/", middleware.CheckPermission("user_write"), handler.Create)
	user.Put("/change-password", middleware.CheckPermission("user_write"), handler.ChangePassword)
	user.Put("/:id", middleware.CheckPermission("user_write"), handler.Update)
	user.Delete("/:id", middleware.CheckPermission("user_write"), handler.Delete)
}
