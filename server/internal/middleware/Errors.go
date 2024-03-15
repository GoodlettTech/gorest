package Middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler[T error] (func(c echo.Context, err T) error)
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
				return handler(c, e)
			}
			return err
		}
	}
}

func ValidationErrorHandler(c echo.Context, err *ValidationError) error {
	errs := make([]string, 0, len(err.Errors()))
	for _, e := range err.Errors() {
		errs = append(errs, e.Error())
	}

	return c.JSON(http.StatusBadRequest, ErrorResponse{
		Code:   http.StatusBadRequest,
		Errors: errs,
	})
}

func HTTPErrorHandler(c echo.Context, err *echo.HTTPError) error {
	return c.JSON(err.Code, ErrorResponse{
		Code:   err.Code,
		Errors: []string{fmt.Sprint(err.Message)},
	})
}
