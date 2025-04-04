package routes

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, handler *handler.UserHandler) {
	user := app.Group("/user")

	user.Get("/", middleware.JWTMiddleware(), middleware.CheckPermission("user_read"), handler.GetAll)
	user.Get("/:id", middleware.JWTMiddleware(), middleware.CheckPermission("user_read"), handler.GetByID)
	user.Post("/", handler.Create) // Create user without authentication
	user.Put("/change-password", middleware.JWTMiddleware(), middleware.CheckPermission("user_write"), handler.ChangePassword)
	user.Put("/:id", middleware.JWTMiddleware(), middleware.CheckPermission("user_write"), handler.Update)
	user.Delete("/:id", middleware.JWTMiddleware(), middleware.CheckPermission("user_write"), handler.Delete)
}
