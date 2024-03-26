package Middleware

import (
	"errors"
	"fmt"
	"net/http"
	Models "server/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Handler[T error] func(c echo.Context, err T) *ErrorResponse
type ErrorResponse struct {
	Code   int
	Errors []string
}

func HandleError[T error](handler Handler[T]) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			var e T
			if errors.As(err, &e) {
				res := handler(c, e)
				return c.JSON(res.Code, res)
			}
			return err
		}
	}
}

func NotFoundErrorHandler(c echo.Context, err *Models.NotFoundError) *ErrorResponse {
	return &ErrorResponse{
		Code:   http.StatusNotFound,
		Errors: []string{err.Error()},
	}
}

func ValidationErrorHandler(c echo.Context, err *ValidationError) *ErrorResponse {
	errs := make([]string, 0, len(err.Errors()))
	for _, e := range err.Errors() {
		errs = append(errs, e.Error())
	}

	return &ErrorResponse{
		Code:   http.StatusBadRequest,
		Errors: errs,
	}
}

func HTTPErrorHandler(c echo.Context, err *echo.HTTPError) *ErrorResponse {
	return &ErrorResponse{
		Code:   err.Code,
		Errors: []string{fmt.Sprint(err.Message)},
	}
}

func DatabaseErrorHandler(c echo.Context, err *pq.Error) *ErrorResponse {
	// get the error code and determine how to handle the error based on the code
	fmt.Println(err.Code)
	switch err.Code {
	case "23505": // unique_violation
		return &ErrorResponse{
			Code:   http.StatusConflict,
			Errors: []string{"Resource already exists"},
		}
	// add more cases here
	default:
		return &ErrorResponse{
			Code:   http.StatusInternalServerError,
			Errors: []string{"Internal server error"},
		}
	}
}
