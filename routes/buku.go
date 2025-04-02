package routes

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BukuRoutes(app *fiber.App, bukuHandler *handler.BukuHandler) {
	buku := app.Group("/buku", middleware.JWTMiddleware())

	// GET hanya untuk user dengan izin "buku_read"
	buku.Get("/", middleware.CheckPermission("buku_read"), bukuHandler.GetBuku)

	// POST, PUT, DELETE hanya untuk user dengan izin "buku_write"
	buku.Post("/", middleware.CheckPermission("buku_write"), bukuHandler.CreateBuku)
	buku.Put("/:id", middleware.CheckPermission("buku_write"), bukuHandler.UpdateBuku)
	buku.Delete("/:id", middleware.CheckPermission("buku_write"), bukuHandler.DeleteBuku)
}
