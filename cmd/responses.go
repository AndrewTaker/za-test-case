package main

import "github.com/gofiber/fiber/v2"

func ResponseGenericError(c *fiber.Ctx, status int, message string) error {
	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
		Success: false,
		Message: message,
	})
}
