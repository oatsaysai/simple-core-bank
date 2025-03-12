package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
)

func LoggingMiddleware(app *app.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.OriginalURL() != "/api/health-check" {
			startTime := time.Now()
			appCtx := app.NewContext(c)
			logger := appCtx.Logger.WithFields(log.Fields{
				"package":   "http_api",
				"remote_ip": c.Context().RemoteIP().String(),
				"method":    c.Method(),
				"path":      c.OriginalURL(),
				"span_id":   GetSpanID(c),
				"trace_id":  GetTraceID(c),
			})

			duration := time.Since(startTime)
			statusCode := c.Response().StatusCode()
			logger = logger.WithFields(log.Fields{
				"duration":    duration.String(),
				"status_code": statusCode,
			})

			jsonReq := CompactJSON(c.Request().Body())
			maxLen := uint(len(string(jsonReq)))
			if maxLen > 2000 {
				maxLen = 2000
			}
			logger.Debugf("JSON Request: %.*s", maxLen, jsonReq)

			c.Next()

			logger.Debugf("JSON Response %s", c.Response().Body())

			if statusCode != http.StatusOK && statusCode != http.StatusCreated && statusCode != http.StatusAccepted {
				logger.Errorf("%s", c.Response().Body())
			}
			logger.Infof("%s %s", c.Method(), c.OriginalURL())
		} else {
			c.Next()
		}

		return nil
	}
}

func CompactJSON(src []byte) []byte {
	var dst bytes.Buffer
	if err := json.Compact(&dst, src); err != nil {
		return nil
	}
	return dst.Bytes()
}
