package main

import (
	"libra-ry/config"
	authorization "libra-ry/internal/auth"
	"libra-ry/internal/handler"
	"libra-ry/internal/middleware"
	"libra-ry/internal/repository"
	"libra-ry/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables and initialize logger and database connection
	config.LoadEnv()
	// Initialize logger and database connection
	logger := config.NewLogger()
	db := config.NewDB(logger)
	// Initalize Fiber app
	app := fiber.New()

	// Middleware global
	app.Use(middleware.LoggerMiddleware(logger))
	app.Use(middleware.RateLimiter())
	app.Use(middleware.CORSMiddleware())
	app.Use(config.NewSwaggerConfig())

	// Initialize repository, use case, and handler for Buku
	repo := repository.NewBukuRepository(db)
	useCase := usecase.NewBukuUseCase(repo)
	handler := handler.NewBukuHandler(useCase)

	// Route untuk Swagger UI
	// Group buku dengan middleware JWT
	buku := app.Group("/buku", middleware.JWTMiddleware())
	// GET hanya bisa diakses oleh user dengan izin "buku_read"
	buku.Get("/", middleware.CheckPermission("buku_read"), handler.GetBuku)
	// POST, PUT, DELETE hanya bisa diakses oleh user dengan izin "buku_write"
	buku.Post("/", middleware.CheckPermission("buku_write"), handler.CreateBuku)
	buku.Put("/:id", middleware.CheckPermission("buku_write"), handler.UpdateBuku)
	buku.Delete("/:id", middleware.CheckPermission("buku_write"), handler.DeleteBuku)

	// Route untuk Swagger UI
	// Group auth tanpa middleware JWT
	auth := app.Group("/auth")
	auth.Post("/login", authorization.Login)

	port := config.GetEnv("APP_PORT", "3000")
	logger.Info("Server is running on port " + port)
	app.Listen(":" + port)
}
