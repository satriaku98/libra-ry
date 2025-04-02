package main

import (
	"libra-ry/config"
	"libra-ry/internal/middleware"
	"libra-ry/routes"

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

	// Initialize Dependencies
	deps := config.InitDependencies(db, logger)

	// Register Routes
	routes.BukuRoutes(app, deps.BukuHandler)
	routes.AuthRoutes(app)

	port := config.GetEnv("APP_PORT", "3000")
	logger.Info("Server is running on port " + port)
	app.Listen(":" + port)
}
