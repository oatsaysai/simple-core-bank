package endpoint

import (
	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/app"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/custom_error"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/handlers/render"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

type AccountsEndpoint interface {
	PreGenerateAccountNumbers(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	GetAccount(c *fiber.Ctx) error
	GetTransactionByAccountNo(c *fiber.Ctx) error
}

type accountsEndpoint struct {
	App *app.App
}

func NewAccountsEndpoint(app *app.App) AccountsEndpoint {
	return &accountsEndpoint{
		App: app,
	}
}

func (e *accountsEndpoint) PreGenerateAccountNumbers(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.PreGenerateAccountNoParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	_, err := ctx.PreGenerateAccountNumbers(params)
	if err != nil {
		return err
	}

	return render.JSON(c, nil)
}

func (e *accountsEndpoint) CreateAccount(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.CreateAccountParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.CreateAccount(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}

func (e *accountsEndpoint) GetAccount(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.GetAccountParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.GetAccount(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}

func (e *accountsEndpoint) GetTransactionByAccountNo(c *fiber.Ctx) error {
	ctx := e.App.NewContext(c)

	var params *model.GetTransactionParams
	if err := c.BodyParser(&params); err != nil {
		return &custom_error.ValidationError{
			Code:    custom_error.InvalidJSONString,
			Message: "Invalid JSON string",
		}
	}

	res, err := ctx.GetTransactionByAccountNo(params)
	if err != nil {
		return err
	}

	return render.JSON(c, res)
}
