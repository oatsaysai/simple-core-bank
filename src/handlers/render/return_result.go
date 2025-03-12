package render

import (
	"net/http"

	"repo.blockfint.com/sakkarin/go-http-server-template/src/model"
)

type Result struct {
	ResponseCode    int               `json:"response_code"`
	ResponseMessage string            `json:"response_message"`
	Pagination      *model.Pagination `json:"pagination,omitempty"`
	Data            any               `json:"data,omitempty"`
}

// Error error description
func (rs Result) Error() string {
	return rs.ResponseMessage
}

// ErrorCode error code
func (rs Result) ErrorCode() int {
	return rs.ResponseCode
}

// HTTPStatusCode http status code
func (rs Result) HTTPStatusCode() int {
	switch rs.ResponseCode {
	case 0, 200: // success
		return http.StatusOK
	case 400: // bad request
		return http.StatusBadRequest
	case 404: // connection_error
		return http.StatusNotFound
	case 401: // unauthorized
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

// NewResultWithMessage new result with message
func NewResultWithMessage(code int, message string) Result {
	return Result{
		ResponseCode:    code,
		ResponseMessage: message,
	}
}
