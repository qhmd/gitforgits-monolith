package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qhmd/gitforgits/utils"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(400).JSON(fiber.Map{"error": "missing authorization header"})
		}
		tokenStr := authHeader[len("Bearer "):]
		token, err := utils.VerifyAccessToken(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token claims"})
		}
		c.Locals("userID", claims["id"])
		c.Locals("userEmail", claims["email"])
		c.Locals("userName", claims["name"])
		c.Locals("userRole", claims["role"])

		return c.Next()
	}
}
