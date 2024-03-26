package Middleware

import (
	AuthService "server/internal/services/auth"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			return echo.NewHTTPError(401, "must be authenticated to use this route")
		}

		tokenString := cookie.Value

		// parse the jwt string and handle errors
		token, err := AuthService.ValidateToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(401, "must be authenticated to use this route")
		}

		c.Set("token", token)
		return next(c)
	}
}
