package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto/book"
	"github.com/qhmd/gitforgits/utils"
)

func ValidateBook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req book.BookRequest

		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errors := make(map[string]string)
			for _, e := range validationErrors {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)

		}

		c.Locals("validateBook", req)
		return c.Next()
	}
}
