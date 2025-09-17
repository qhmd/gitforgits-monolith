package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/utils"
)

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("userRole").(string)
		if role != "admin" {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "cannot access this is request", nil)
		}
		return c.Next()
	}
}
