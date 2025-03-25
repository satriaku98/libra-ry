package main

import (
	"libra-ry/config"
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"
	"libra-ry/internal/repository"
	"libra-ry/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := config.NewLogger()
	db := config.NewDB(logger)

	app := fiber.New()

	// Middleware global
	app.Use(middleware.LoggerMiddleware(logger))
	app.Use(middleware.RateLimiter())
	app.Use(middleware.CORSMiddleware())
	app.Use(config.NewSwaggerConfig())

	repo := repository.NewBukuRepository(db)
	useCase := usecase.NewBukuUseCase(repo)
	handler := handler.NewBukuHandler(useCase)

	app.Get("/buku", handler.GetBuku)
	app.Get("/buku/:id", handler.GetBukuByID)
	app.Post("/buku", handler.CreateBuku)
	app.Put("/buku/:id", handler.UpdateBuku)
	app.Delete("/buku/:id", handler.DeleteBuku)

	logger.Info("Server is running on port 3000")
	app.Listen(":3000")
}
