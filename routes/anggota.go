package routes

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AnggotaRoutes(app *fiber.App, h *handler.AnggotaHandler) {
	anggota := app.Group("/anggota", middleware.JWTMiddleware())

	anggota.Get("/", h.Get)
	anggota.Post("/", h.Create)
	anggota.Put("/", h.Update)
	anggota.Delete("/", h.Delete)
}
