package middleware

import (
	"libra-ry/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// CheckPermission adalah middleware untuk memeriksa apakah user memiliki izin
func CheckPermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token yang sudah didecode di middleware JWT
		userClaims := c.Locals("user")
		if userClaims == nil {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		// Konversi ke JWT MapClaims
		claims, ok := userClaims.(jwt.MapClaims)
		if !ok {
			return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token claims")
		}

		// Ambil permissions dari token
		permissions, ok := claims["permissions"].([]interface{})
		if !ok {
			return pkg.ErrorResponse(c, fiber.StatusForbidden, "No permissions found")
		}

		// Konversi permissions ke slice string
		var permissionList []string
		for _, perm := range permissions {
			if p, ok := perm.(string); ok {
				permissionList = append(permissionList, p)
			}
		}

		// Cek apakah requiredPermission ada di permissions
		if !contains(permissionList, requiredPermission) {
			return pkg.ErrorResponse(c, fiber.StatusForbidden, "Access denied")
		}

		return c.Next()
	}
}

// contains mengecek apakah slice memiliki elemen tertentu
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
