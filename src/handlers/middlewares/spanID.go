package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/constant"
)

func CorrelationMiddleware(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {

		traceID := c.Get("X-Correlation-ID")
		spanID := uuid.NewString()

		c.Locals(constant.ContextKeyTraceID, traceID)
		c.Locals(constant.ContextKeySpanID, spanID)

		c.Context().SetUserValue(constant.ContextKeyTraceID, traceID)
		c.Context().SetUserValue(constant.ContextKeySpanID, spanID)

		return c.Next()
	}
}

func GetSpanID(c *fiber.Ctx) string {
	spanID := c.Locals(constant.ContextKeySpanID)
	if ret, ok := spanID.(string); ok {
		return ret
	}
	return ""
}

func GetTraceID(c *fiber.Ctx) string {
	traceID := c.Locals(constant.ContextKeyTraceID)
	if ret, ok := traceID.(string); ok {
		return ret
	}
	return ""
}
