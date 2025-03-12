package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/custom_error"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/render"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

type TransferEndpoint interface {
	TransferIn(c *fiber.Ctx) error
	TransferOut(c *fiber.Ctx) error
	Transfer(c *fiber.Ctx) error
}

type transferEndpoint struct {
	App *app.App
}

func NewTransferEndpoint(app *app.App) TransferEndpoint {
	return &transferEndpoint{
		App: app,
	}
}

func (e *transferEndpoint) TransferIn(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.TransferInParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.TransferIn(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}

func (e *transferEndpoint) TransferOut(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.TransferOutParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.TransferOut(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}

func (e *transferEndpoint) Transfer(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.TransferParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.Transfer(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}
