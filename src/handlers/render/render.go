package render

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/custom_error"
	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

// JSON render json to client
func JSON(c *fiber.Ctx, response any) error {
	return c.
		Status(http.StatusOK).
		JSON(Result{
			ResponseCode:    0,
			ResponseMessage: "Success",
			Data:            response,
		})
}

func JSONWithOutCode(c *fiber.Ctx, response any) error {
	return c.
		Status(http.StatusOK).
		JSON(response)
}

// JSON with pagination render json to client
func JSONWithPagination(c *fiber.Ctx, response any, pagination *model.Pagination) error {
	return c.
		Status(http.StatusOK).
		JSON(Result{
			ResponseCode:    0,
			ResponseMessage: "Success",
			Pagination:      pagination,
			Data:            response,
		})
}

// Byte render byte to client
func Byte(c *fiber.Ctx, bytes []byte) error {
	_, err := c.Status(http.StatusOK).
		Write(bytes)

	return err
}

// Error render error to client
func Error(c *fiber.Ctx, err error) error {

	if locErr, ok := err.(Result); ok {
		return c.
			Status(locErr.HTTPStatusCode()).
			JSON(locErr)
	}

	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.
			Status(fiberErr.Code).
			JSON(NewResultWithMessage(fiberErr.Code, fiberErr.Error()))
	}

	if customErr, ok := err.(*custom_error.ValidationError); ok {
		return c.
			Status(http.StatusBadRequest).
			JSON(customErr)
	}

	if customErr, ok := err.(*custom_error.UserError); ok {
		return c.
			Status(http.StatusBadRequest).
			JSON(customErr)
	}

	defaultErr := Result{
		ResponseCode:    http.StatusInternalServerError,
		ResponseMessage: err.Error(),
	}
	return c.
		Status(defaultErr.HTTPStatusCode()).
		JSON(defaultErr)
}
