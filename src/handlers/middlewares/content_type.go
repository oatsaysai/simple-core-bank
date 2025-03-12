package middlewares

import "github.com/gofiber/fiber/v2"

func SetContentTypeJSON() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Request().Header.Set("Content-Type", "application/json")
		return c.Next()
	}
}
