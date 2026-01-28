package middleware

import "github.com/gofiber/fiber/v2"

func JWT(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}
	return c.Next()
}
