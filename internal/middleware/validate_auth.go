package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/internal/dto/auth"
	"github.com/qhmd/gitforgits/utils"
)

func ValidateAuthLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req auth.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)

		}
		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}

func ValidateAuthReg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req auth.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body", err)
		}

		if err := utils.Validate.Struct(req); err != nil {
			validationError := err.(validator.ValidationErrors)
			errors := make(map[string]string)
			for _, e := range validationError {
				errors[e.Field()] = utils.MsgForTag(e)
			}
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "validation error", errors)
		}
		c.Locals("validateAuth", req)
		return c.Next()
	}
}
