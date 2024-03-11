package AuthMiddleware

import (
	AuthService "server/server/internal/services/auth"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get jwt string from auth header
		authHeader := c.Request().Header.Get("Auth")

		// parse the jwt string and handle errors
		token, err := AuthService.ValidateToken(authHeader)
		if err != nil {
			return echo.NewHTTPError(401, "must be authenticated to use this route")
		}

		c.Set("token", token)
		return next(c)
	}
}
