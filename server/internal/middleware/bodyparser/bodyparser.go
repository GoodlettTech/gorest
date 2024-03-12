package BodyParser

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Parse[T interface{}](name string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var data T
			if err := c.Bind(&data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			if err := c.Validate(&data); err != nil {
				return err
			}

			c.Set(name, data)
			return next(c)
		}
	}
}
