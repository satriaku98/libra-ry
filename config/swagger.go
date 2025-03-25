package config

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

// NewSwaggerConfig mengatur Swagger untuk dokumentasi API
func NewSwaggerConfig() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/",
		Path:     "docs",
		FilePath: "../docs/swagger.json",
	})
}
