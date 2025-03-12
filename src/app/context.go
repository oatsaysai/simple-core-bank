package app

import (
	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/constant"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/db"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
)

type Context struct {
	Logger    log.Logger
	Config    *Config
	DB        db.DB
	RequestID string
	FiberCtx  *fiber.Ctx
}

func (app *App) NewContext(c *fiber.Ctx) *Context {

	newCtx := &Context{
		Logger:   app.Logger,
		Config:   app.Config,
		DB:       app.DB,
		FiberCtx: c,
	}

	// Add package to log
	newCtx.Logger = newCtx.Logger.WithFields(log.Fields{"package": "app"})

	if c != nil {
		if traceID, ok := c.Locals(constant.ContextKeyTraceID).(string); ok {
			newCtx.Logger = newCtx.Logger.WithFields(log.Fields{"trace_id": traceID})
		}
		if spanID, ok := c.Locals(constant.ContextKeySpanID).(string); ok {
			newCtx.Logger = newCtx.Logger.WithFields(log.Fields{"span_id": spanID})
		}
	}

	return newCtx
}

// func (ctx *Context) WithLogger(logger log.Logger) *Context {
// 	ret := *ctx
// 	ret.Logger = logger
// 	return &ret
// }
