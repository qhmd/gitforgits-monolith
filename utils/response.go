package utils

import "github.com/gofiber/fiber/v2"

type RegisUserResponse struct {
	AccessToken string `json:"token"`
	User        any    `json:"user"`
}

// Struktur umum respons API
type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// Fungsi helper untuk response sukses
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Fungsi helper untuk response gagal
func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errors any) error {
	return c.Status(statusCode).JSON(APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
