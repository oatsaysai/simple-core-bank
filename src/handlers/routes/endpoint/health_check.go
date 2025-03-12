package endpoint

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/render"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/version"
)

type HealthCheckEndpoint interface {
	HealthCheck(c *fiber.Ctx) error
}

type healthCheckEndpoint struct {
	App       *app.App
	startTime time.Time
}

func NewHealthCheckEndpoint(app *app.App) HealthCheckEndpoint {
	return &healthCheckEndpoint{
		App:       app,
		startTime: time.Now(),
	}
}

type HealthCheckServiceDetail struct {
	ServiceName string `json:"service_name"`
	StartTime   string `json:"start_time,omitempty"`
	UpTime      string `json:"up_time,omitempty"`
	Version     string `json:"version,omitempty"`
	Commit      string `json:"commit,omitempty"`
	Data        any    `json:"data,omitempty"`
}

// It work only when run just one instance
// If you run multiple instance, you should use other way to save start time
// such as save in database or redis
func (ep *healthCheckEndpoint) HealthCheck(c *fiber.Ctx) error {
	return render.JSON(
		c,
		HealthCheckServiceDetail{
			ServiceName: "go-http-server-template",
			StartTime:   ep.startTime.Format(time.RFC3339),
			UpTime:      time.Since(ep.startTime).String(),
			Version:     version.AppSemVer,
			Commit:      version.GitCommit,
		})
}
