package UserMiddleware

import (
	"net/http"
	UserModel "server/server/internal/models/users"

	"github.com/labstack/echo/v4"
)

func TakesUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// parse body json into user object
		var user UserModel.User

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&user); err != nil {
			return err
		}

		c.Set("user", user)
		return next(c)
	}
}

func TakesCredentials(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// parse body json into credentials object
		var credentials UserModel.Credentials

		err := c.Bind(&credentials)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&credentials); err != nil {
			return err
		}

		c.Set("credentials", credentials)

		return next(c)
	}
}
