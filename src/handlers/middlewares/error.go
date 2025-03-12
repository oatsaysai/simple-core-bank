package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/render"
)

// WrapError wrap error
func WrapError() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			return render.Error(c, err)
		}
		return nil
	}
}
