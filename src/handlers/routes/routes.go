package routes

import (
	ctx "context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/middlewares"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/routes/endpoint"
	log "repo.blockfint.com/sakkarin/go-http-server-template/src/logger"
)

func NewRouter(config *Config, logger log.Logger, app *app.App) {

	fiberApp := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // this is the default limit of 4MB
	})

	fiberApp.Use(
		cors.New(),
		middlewares.SetContentTypeJSON(),
		middlewares.CorrelationMiddleware(app),
		middlewares.LoggingMiddleware(app),
		middlewares.WrapError(),
	)

	// Initialize endpoints
	healthCheckEndpoint := endpoint.NewHealthCheckEndpoint(app)
	accountsEndpoint := endpoint.NewAccountsEndpoint(app)
	transferEndpoint := endpoint.NewTransferEndpoint(app)

	api := fiberApp.Group("/api")
	api.Get("/health-check", healthCheckEndpoint.HealthCheck)

	apiV1 := api.Group("/v1")

	// Accounts routes
	accountsGroup := apiV1.Group("")
	accountsGroup.Post("/pre-generate-account-numbers", accountsEndpoint.PreGenerateAccountNumbers)
	accountsGroup.Post("/create-account", accountsEndpoint.CreateAccount)
	accountsGroup.Post("/get-account", accountsEndpoint.GetAccount)
	accountsGroup.Post("/get-transaction-by-account-no", accountsEndpoint.GetTransactionByAccountNo)

	// Transfer routes
	transferGroup := apiV1.Group("")
	transferGroup.Post("/transfer-in", transferEndpoint.TransferIn)
	transferGroup.Post("/transfer-out", transferEndpoint.TransferOut)
	transferGroup.Post("/transfer", transferEndpoint.Transfer)

	// Graceful shutdown
	// Waiting os signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
		defer cancel()

		logger.Infof("Gracefully shutting down...")
		_ = fiberApp.Shutdown()
	}()

	logger.Infof("Serving HTTP API at http://127.0.0.1:%d", config.Port)
	err := fiberApp.Listen(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Panicf(err.Error())
	}
}
