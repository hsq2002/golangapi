package api

import "github.com/gofiber/fiber/v2"

func apiKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("Authorization")
	if apiKey != "Your_Secret_API_Key" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Success": false, "message": "Unauthorized"})
	}
	return c.Next()
}
