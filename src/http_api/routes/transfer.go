package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/oatsaysai/simple-core-bank/src/app"
	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	"github.com/oatsaysai/simple-core-bank/src/model"
)

func TransferRouter(fiberApp fiber.Router) {
	fiberApp.Post("/transfer-in", TransferIn())
	fiberApp.Post("/transfer-out", TransferOut())
	fiberApp.Post("/transfer", Transfer())
	fiberApp.Post("/transfer-in-for-load-test", TransferInForLoadTest())
	fiberApp.Post("/transfer-out-for-load-test", TransferOutForLoadTest())
	fiberApp.Post("/transfer-for-load-test", TransferForLoadTest())
}

func TransferIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferInParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.TransferIn(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func TransferOut() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferOutParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.TransferOut(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func Transfer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.Transfer(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func TransferInForLoadTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferForLoadTestParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.TransferInForLoadTest(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func TransferOutForLoadTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferForLoadTestParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.TransferOutForLoadTest(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}

func TransferForLoadTest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params model.TransferForLoadTestParams

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&custom_error.UserError{
				Code:           custom_error.InvalidJSONString,
				Message:        "Invalid JSON string",
				HTTPStatusCode: http.StatusBadRequest,
			})
		}

		appCtx := c.Locals(APP_CTX_KEY).(app.Context)
		result, err := appCtx.TransferForLoadTest(params)
		if err != nil {
			return ReturnCustomError(c, err)
		}

		return c.JSON(&Response{
			Code:    0,
			Message: "",
			Data:    result,
		})
	}
}
