package config

import (
	"strconv"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

// NewSwaggerConfig mengatur Swagger untuk dokumentasi API
func NewSwaggerConfig() fiber.Handler {
	// Baca dari env dan convert ke int
	cacheAgeStr := GetEnv("SWAGGER_CACHE_AGE", "3600")
	cacheAge, err := strconv.Atoi(cacheAgeStr)
	if err != nil {
		cacheAge = 3600 // fallback jika parsing gagal
	}
	return swagger.New(swagger.Config{
		BasePath: "/",
		Path:     "docs",
		FilePath: "../docs/swagger.json",
		CacheAge: cacheAge,
	})
}
