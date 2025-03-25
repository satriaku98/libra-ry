package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        20,              // Maksimum request per IP
		Expiration: 1 * time.Minute, // Reset setiap 1 menit
	})
}
